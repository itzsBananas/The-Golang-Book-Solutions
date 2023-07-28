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
-   use type assertions (custom interface or not) to check if a object has a certain method or is a certain type

Goroutines (Concurrency)

-   Do not communicate by sharing memory; instead, share memory by communicating
-   Where possible, confine variables to a single goroutine; for all other variables, use mutual exclusion.\
-   two main ways of building concurrency: shared variables and locks, or communicating sequential processes;
    shared: allows multiple writers, but need locks; seq: one writer, multiple readers (monitor goroutine)

-   Goroutines are lightweight processes that may run concurrently with other goroutines. Concurrent means we cannot determine what operation gets called between two goroutines
-   Channels allow communication between these goroutines (besides sharing memory); useful for sync'ing operations and limiting shared memory - Buffered channels are good if sync is not as important as it is harder to block coroutines with buffered channels
-   Channels can be useful to parallelize workloads or run processes on the side while responding to requests or timing
    -   also useful to limit # of processes or resources (counting / binary semaphore)
-   Multiplexing makes processing multiple recieving channels easy (fan-in) w/o running additional goroutines; fan-out is vice versa
-   Waitgroups are useful to indicate how many coroutines are running in the background and block if not zero
-   closed channels are useful for cancelling goroutines / nil channels are useful to block certain cases of select to happen
-   patterns: generators, multiplexing (monitor goroutine), cancellation thru done channel - monitor goroutine: goroutine that brokers access to a confined variable using channel requests
-   bounded parallelism by using counting semaphore or limiting the amount of coroutines
