# Maps

## Introduction

Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the [].
The key type is special. It can only be a comparable type because without the ability to tell if 2 keys are equal, we have no way to ensure that we are getting the correct value.

The basic search was very easy to implement, but what will happen if we supply a word that's not in our dictionary?

We actually get nothing back. This is good because the program can continue to run, but there is a better approach. The function can report that the word is not in the dictionary. This way, the user isn't left wondering if the word doesn't exist or if there is just no definition (this might not seem very useful for a dictionary. However, it's a scenario that could be key in other usecases).

The way to handle this scenario in Go is to return a second argument which is an Error type.
Errors can be converted to a string with the .Error() method, which we do when passing it to the assertion. We are also protecting assertStrings with if to ensure we don't call .Error() on nil.

In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition.

## Mpas aand pointers

An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.

A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic. You can read more about maps here.
Therefore, you should never initialize an empty map variable:

var m map[string]string

Instead, you can initialize an empty map like we were doing above, or use the make keyword to create a map for you:
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic.

We made the errors constant; this required us to create our own DictionaryErr type which implements the error interface. You can read more about the details in this excellent article by Dave Cheney. Simply put, it makes the errors more reusable and immutable.
Next, let's create a function to Update the definition of a word.

Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.
The delete function returns nothing, and we based our Delete method on the same notion. Since deleting a value that's not there has no effect, unlike our Update and Add methods, we don't need to complicate the API with errors.