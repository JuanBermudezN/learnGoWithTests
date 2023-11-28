package main

import "testing"

// Rules

// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file

// Let's start by capturing these requirements in a test. This is basic test driven development and allows us to make sure our test is actually
// testing what we want. When you retrospectively write tests there is the risk that your test may continue to pass even if the code doesn't work as intended.

// Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

}

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
// so you can call helper functions from a test, or a benchmark (don't worry if words like "interface" mean nothing to you right now,
//it will be covered later).e that *testing.T and *testing.B both satisfy, so you can call helper functions from a test, package main

// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be
// in our function call rather than inside our test helper.
//This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output.

// Let's go over the cycle again
// Write a test
// Make the compiler pass
// Run the test, see that it fails and check the error message is meaningful
// Write enough code to make the test pass
// Refactor

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
