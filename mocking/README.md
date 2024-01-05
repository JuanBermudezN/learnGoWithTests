# Mocking

/mock_test.go:11:12: too many arguments in call to Countdown
	have (*bytes.Buffer)
	want ()

The compiler is telling you what your function signature could be, so update it.

`func Countdown(out *bytes.Buffer) {}`

We're using fmt.Fprint which takes an io.Writer (like *bytes.Buffer) and sends a string to it. The test should pass.

The tests still pass and the software works as intended but we have some problems:

* Our tests take 3 seconds to run.

* Every forward-thinking post about software development emphasises the importance of quick feedback loops.

* Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 3s added to the test run for every new test of Countdown?

* We have not tested an important property of our function.

We have a dependency on Sleeping which we need to extract so we can then control it in our tests.

If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep and then we can spy on the calls to make assertions on them.

Our latest change only asserts that it has slept 3 times, but those sleeps could occur out of sequence.

When writing tests if you're not confident that your tests are giving you sufficient confidence, just break it! (make sure you have committed your changes to source control first though). Change the code to the following

Our SpyCountdownOperations implements both io.Writer and Sleeper, recording every call into one slice. In this test we're only concerned about the order of operations, so just recording them as list of named operations is sufficient.

A nice feature would be for the Sleeper to be configurable. This means that we can adjust the sleep time in our main program.

we are using duration to configure the time slept and sleep as a way to pass in a sleep function. The signature of sleep is the same as for time.