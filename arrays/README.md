# Arrays and Slices


There's a new way to create a slice. make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.

As mentioned, slices have a capacity. If you have a slice with a capacity of 2 and try to do mySlice[10] = 1 you will get a runtime error.
However, you can use the append function which takes a slice and a new value, then returns a new slice with all the items in it.

Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides of the : it captures everything to that side of it.
In our case, we are saying "take from 1 to the end" with numbers[1:]. You may wish to spend some time writing other tests around slices and experiment with
the slice operator to get more familiar with it.

## Testing

Go's built-in testing toolkit features a coverage tool. Whilst striving for 100% coverage should not be your end goal,
the coverage tool can help identify areas of your code not covered by tests. If you have been strict with TDD,
it's quite likely you'll have close to 100% coverage anyway.

Go does not let you use equality operators with slices. You could write a function to iterate over each got and want slice and check their values but
for convenience sake, we can use reflect.DeepEqual which is useful for seeing if any two variables are the same.

It's important to note that reflect.DeepEqual is not "type safe" - the code will compile even if you did something a bit silly.
To see this in action, temporarily change the test to:

What we have done here is try to compare a slice with a string. This makes no sense, but the test compiles! So while
using reflect.DeepEqual is a convenient way of comparing slices (and other things) you must be careful when using it

By defining this function inside the test, it cannot be used by other functions in this package.
Hiding variables and functions that don't need to be exported is an important design consideration.
A handy side-effect of this is this adds a little type-safety to our code. If a developer mistakenly adds a new test with checkSums(t, got, "dave")
the compiler will stop them in their tracks.
