Good to know

-   If statements can have multiple statements
-   Function recievers transform reciving argument to correct type (but type cannot be underlying pointer or interface)
-   Lots of different ways to simplify # of words (to declare, initialize variables, structs, arrays for sure).
-   Slices and Structs are the only pointer types (structs and arrays are not)
-   No subtypes (is-relationship), but embedded variables promote methods to the holding struct (has-relationship)
-   it’s possible and sometimes useful for unnamed struct types to have methods too.
-   We cannot call methods on a non-addressable receiver, because there’s no way to obtain the address of a temporary value.
-   A type satisfies an interface if it possesses all the methods the interface requires.
-   pointer types are by no means the only types that satisfy interfaces, and even interfaces
    with mutator methods may be satisfied by one of Go’s other reference types.
-   Conceptually, a value of an interface type, or interface value, has two components, a concrete
    type and a value of that type. These are called the interface’s dynamic type and dynamic value.
-   it’s possible and sometimes useful for unnamed struct types to have
    methods too.groups together the two related variables in a single package-level variable
