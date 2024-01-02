# Basics

You will notice that we're using %d as our format strings rather than %q. That's because we want it to print an integer rather than a string.

Also note that we are no longer using the main package, instead we've defined a package named integers, as the name suggests this will group
functions for working with integers such as Add.

Go examples are executed just like tests so you can be confident examples reflect what the code actually does.
// Examples are compiled (and optionally executed) as part of a package's test suite.

As with typical tests, examples are functions that reside in a package's _test.go files. Add the following ExampleAdd function to the adder_test.go file.

Please note that the example function will not be executed if you remove the comment // Output: 6. Although the function will be compiled, it won't be executed.

By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.
