# Dependency injection

Just to recap, here is what that function could look like

But how can we test this? Calling fmt.Printf prints to stdout, which is pretty hard for us to capture using the testing framework.
What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing.

f we do that, we can then change the implementation to print to something we control so that we can test it. In "real life" you would inject in something that writes to stdout.

The Buffer type from the bytes package implements the Writer interface, because it has the method Write(p []byte) (n int, err error).
So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet

Use the writer to send the greeting to the buffer in our test. Remember fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to, whereas fmt.Printf defaults to stdout.


Earlier the compiler told us to pass in a pointer to a bytes.Buffer. This is technically correct but not very useful.
To demonstrate this, try wiring up the Greet function into a Go application where we want it to print to stdout.

./di.go:14:7: cannot use os.Stdout (type *os.File) as type *bytes.Buffer in argument to Greet

As discussed earlier fmt.Fprintf allows you to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.

If we change our code to use the more general purpose interface we can now use it in both tests and in our application.

When you write an HTTP handler, you are given an http.ResponseWriter and the http.Request that was used to make the request. When you implement your server you write your response using the writer.

You can probably guess that http.ResponseWriter also implements io.Writer so this is why we could re-use our Greet function inside our handler.