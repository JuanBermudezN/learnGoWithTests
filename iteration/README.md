# Iterations

Benchmarking
Writing benchmarks in Go is another first-class feature of the language and it is very similar to writing tests.

The amount of times the code is run shouldn't matter to you, the framework will determine what is a "good" value for that to let you have some decent results.
To run the benchmarks do go test -bench=. (or if you're in Windows Powershell go test -bench=".")

What 136 ns/op means is our function takes on average 136 nanoseconds to run (on my computer). Which is pretty ok! To test this it ran it 10000000 times.
