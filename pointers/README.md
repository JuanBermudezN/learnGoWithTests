 Go, when you call a function or a method the arguments are copied.

When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.

Without getting too computer-sciency, when you create a value - like a wallet, it is stored somewhere in memory. You can find out what the address of that bit of memory with &myVal.

and seemingly addressed the object directly. In fact, the code above using (*w) is absolutely valid. However, the makers of Go deemed this notation cumbersome, so the language permits us to write w.balance, without an explicit dereference. These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine. However, by convention you should keep your method receiver types the same for consistency.

so you can also return (*w).balance

We said we were making a Bitcoin wallet but we have not mentioned them so far. We've been using int because they're a good type for counting things!
It seems a bit overkill to create a struct for this. int is fine in terms of the way it works but it's not descriptive.

Go lets you create new types from existing ones.
The syntax is type MyName OriginalType

To make Bitcoin you just use the syntax Bitcoin(999).
By doing this we're making a new type and we can declare methods on them. This can be very useful when you want to add some domain specific functionality on top of existing types.

Stringer

This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints.

What should happen if you try to Withdraw more than is left in the account? For now, our requirement is to assume there is not an overdraft facility.
How do we signal a problem when using Withdraw?
In Go, if you want to indicate an error it is idiomatic for your function to return an err for the caller to check and act on.

We want Withdraw to return an error if you try to take out more than you have and the balance should stay the same.

We then check an error has returned by failing the test if it is nil.
nil is synonymous with null from other programming languages. Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

Like null if you try to access a value that is nil it will throw a runtime panic. This is bad! You should make sure that you check for nils.

And now the test is easier to follow too.
I have moved the helpers out of the main test function just so when someone opens up a file they can start reading our assertions first, rather than some helpers.
Another useful property of tests is that they help us understand the real usage of our code so we can make sympathetic code. We can see here that a developer can simply 