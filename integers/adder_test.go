package integers

import "testing"

func TestAdder(t *testing.T)) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', but got '%d", expected, sum)
	}
}

// You will notice that we're using %d as our format strings rather than %q. That's because we want it to print an integer rather than a string.
// Also note that we are no longer using the main package, instead we've defined a package named integers, as the name suggests this will group 
// functions for working with integers such as Add.