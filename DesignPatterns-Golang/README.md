# When to use each of these patterns?

## Singleton
Use when a class(in this case a struct) in your program should have just a single instance available to all clients.

## Facade
Use when a system is very complex or difficult to understand because the system has many interdependent classes(in this case a structs) or because its source code is unavailable. This pattern hides the complexities of the larger system and provides a simpler interface to the client.

## Observer
Use in any situation where you have several objects which are dependent on another object and are required to perform an action when the state of that object changes.

## Strategy
Use when you want to use different variants of an algorithm within an object and be able to switch from one algorithm to another during runtime.

## Factory
Use when you want an alternative to constructors, when you need a complicated process for constructing the object, when the construction need a dependency that you do not want for the actual class(in this case a struct).

## Bridge
Use when both the class(in this case a struct) and what it does vary often.

## Adapter
Use to make existing classes(in this case a structs) work with others without modifying their source code.

## Composite
Use when you need to treat a group of objects in similar way as a single object. Composite pattern composes objects in term of a tree structure to represent part as well as whole hierarchy.

## Decorator
Use when you want a solution to modify an object in run time, applying independent features or behaviors which can be combined and accumulated in any order.

## Proxy
Use when you need to create a wrapper to cover the main object's complexity from the client.
