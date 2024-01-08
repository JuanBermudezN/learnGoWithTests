# Concurrency

Here's the setup: a colleague has written a function, CheckWebsites, that checks the status of a list of URLs.

t returns a map of each URL checked to a boolean value: true for a good response; false for a bad response.
You also have to pass in a WebsiteChecker which takes a single URL and returns a boolean. This is used by the function to check all the websites.
Using dependency injection has allowed them to test the function without making real HTTP calls, making it reliable and fast.

he benchmark tests CheckWebsites using a slice of one hundred urls and uses a new fake implementation of WebsiteChecker. slowStubWebsiteChecker is deliberately slow. It uses time.Sleep to wait exactly twenty milliseconds and then it returns true. We use b.ResetTimer() in this test to reset the time of our test before it actually runs
When we run the benchmark using go test -bench=. (or if you're in Windows Powershell go test -bench="."):


Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.

To tell Go to start a new goroutine we turn a function call into a go statement by putting the keyword go in front of it: go doSomething().

Because the only way to start a goroutine is to put go in front of a function call, we often use anonymous functions when we want to start a goroutine. An anonymous function literal looks just the same as a normal function declaration, but without a name (unsurprisingly). You can see one above in the body of the for loop.

Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing. Secondly they maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

The body of the anonymous function above is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function). Each goroutine will add its result to the results map

We are caught by the original test CheckWebsites, it's now returning an empty map. What went wrong?

None of the goroutines that our for loop started had enough time to add their result to the results map; the WebsiteChecker function is too fast for them, and it returns the still empty map
.
To fix this we can just wait while all the goroutines do their work, and then return. Two seconds ought to do it, right?

This isn't great - why only one result? We might try and fix this by increasing the time we wait - try it if you like. It won't work. The problem here is that the variable url is reused for each iteration of the for loop - it takes a new value from urls each time. But each of our goroutines have a reference to the url variable - they don't have their own independent copy. So they're all writing the value that url has at the end of the iteration - the last url. Which is why the one result we have is the last url.

By giving each anonymous function a parameter for the url - u - and then callinThis is long and scary, but all we need to do is take a breath and read the stacktrace: fatal error: concurrent map writes. Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time. Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.

This is long and scary, but all we need to do is take a breath and read the stacktrace: fatal error: concurrent map writes. Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time. Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.

This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.
Go can help us to spot race conditions with its built in race detector. To enable this feature, run the tests with the race flag: go test -race.

## Channels

We can solve this data race by coordinating our goroutines using channels. Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.
In this case we want to think about the communication between the parent process and each of the goroutines that it makes to do the work of running the WebsiteChecker function with the url.

The next for loop iterates once for each of the urls. Inside we're using a receive expression, which assigns a value received from a channel to a variable. This also uses the <- operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left:

We then use the result received to update the map.
By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time. Although each of the calls of wc, and each send to the result channel, is happening in parallel inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.

We have parallelized the part of the code that we wanted to make faster, while making sure that the part that cannot happen in parallel still happens linearly. And we have communicated across the multiple processes involved by using channels.