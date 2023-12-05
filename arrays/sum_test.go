package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal,
// the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD,
// it's quite likely you'll have close to 100% coverage anyway.

// Go does not let you use equality operators with slices. You could write a function to iterate over each got and want slice and check their values but
// for convenience sake, we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.

// 's important to note that reflect.DeepEqual is not "type safe" - the code will compile even if you did something a bit silly.
// To see this in action, temporarily change the test to:

// What we have done here is try to compare a slice with a string. This makes no sense, but the test compiles! So while
// using reflect.DeepEqual is a convenient way of comparing slices (and other things) you must be careful when using it

// By defining this function inside the test, it cannot be used by other functions in this package.
//  Hiding variables and functions that don't need to be exported is an important design consideration.
// A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly adds a new test with checkSums(t, got, "dave")
//  the compiler will stop them in their tracks.
