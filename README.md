# Learning go


# Ten point roadmap:
Very good. Now expand the list so that it contains 10 items instead.


## 1. Go Syntax & Basic Language Constructs

Variables, constants, types, control flow, functions, parameters, returning values, named returns, and pointers.


### **Exercise Set 1: Strings**

```go
// TODO:
// 1. Create a string with your full name.
// 2. Print its length using len().
// 3. Access the first and last character (as bytes).
// 4. Convert the string to uppercase using strings.ToUpper.
// 5. Count how many vowels it contains.
// 6. Check if it contains the substring "a" or "A".
// 7. Extract the substring from index 2 to 5.
// 8. Split the name into words using strings.Split.
// 9. Join the words back using a hyphen.
// 10. Replace all vowel characters with '*'.
// 11. Iterate over the string as runes and print each rune + its index.
// 12. Reverse the string (properly handling runes).
```


### **Exercise Set 2: Variables, Types & Type Conversion**

```go
// TODO:
// 1. Declare variables of types: int, float32, float64, string, bool.
// 2. Print their types using fmt.Printf and %T.
// 3. Use short declaration (:=) to create inferred types.
// 4. Declare multiple variables at once.
// 5. Convert an int to float64.
// 6. Convert a float64 to int (and observe rounding behavior).
// 7. Convert a string containing a number ("42") to int using strconv.Atoi.
// 8. Declare a zero-value variable of each basic type.
// 9. Declare an untyped constant and print its type after assignment.
// 10. Use iota to create a simple enum-like constant list.
// 11. Swap values of two integer variables without a temp variable.
// 12. Show the difference between := and var in shadowing behavior.
```



### **Exercise Set 3: Constants**

```go
// TODO:
// 1. Create a constant for Pi (approx).
// 2. Create a typed constant with type float64.
// 3. Create an untyped constant and assign it to different variable types.
// 4. Use iota to define constants representing weekdays.
// 5. Create a bitmask using iota (1 << iota).
// 6. Try modifying a constant and observe compile-time error.
// 7. Create a large untyped constant and assign it to both int64 and float64.
```



### **Exercise Set 4: Conditionals (if, else-if, switch)**

```go
// TODO:
// 1. Read an integer from the user.
// 2. Check if the number is positive, negative, or zero.
// 3. Check if the number is even or odd.
// 4. Use if with a short declaration: if x := ...; x > 10.
// 5. Use logical operators (&&, ||, !).
// 6. Write a switch that categorizes numbers: 0, 1–5, 6–10, >10.
// 7. Write a switch on strings for detecting: "go", "java", "python".
// 8. Write a switch with no condition (like chained ifs).
// 9. Use fallthrough intentionally and observe behavior.
// 10. Nest conditionals to test multiple properties.
```



### **Exercise Set 5: Loops (for)**

```go
// TODO:
// 1. Print numbers 1 to 10.
// 2. Print even numbers until 20.
// 3. Sum numbers in a slice: []int{3, 7, 2, 9, 4}.
// 4. Find the largest number in a slice.
// 5. Use range to print index + value.
// 6. Use continue to skip odd numbers.
// 7. Use break to stop at a given value.
// 8. Create an infinite loop that stops when counter reaches 5.
// 9. Loop through a string and print each rune.
// 10. Nested loops: print multiplication table 1–5.
// 11. Use labeled loops to break out of outer loop.
// 12. Implement a loop to find the first divisible-by-7 number ≥ 1,000.
```



### **Exercise Set 6: Arrays & Slices**

```go
// TODO:
// 1. Create an array of 5 integers.
// 2. Create a slice from the array.
// 3. Append elements to the slice.
// 4. Check slice length and capacity.
// 5. Copy one slice to another using copy().
// 6. Create a slice literal []int{1, 2, 3, 4}.
// 7. Slice the slice: first 2 elements, last 2 elements.
// 8. Modify the array and observe changes in slices pointing to it.
// 9. Implement manual dynamic expansion (append in loop).
// 10. Remove an element from index 2.
// 11. Create a 2D slice and print it.
// 12. Preallocate capacity with make() and compare performance.
```



### **Exercise Set 7: Maps**

```go
// TODO:
// 1. Create a map[string]int.
// 2. Add three key–value pairs.
// 3. Lookup a key using the “comma ok” idiom.
// 4. Delete a key.
// 5. Loop through all key–value pairs.
// 6. Count how many values are above a threshold.
// 7. Find the largest value and its key.
// 8. Initialize a map literal with values.
// 9. Check if map is nil before using it.
// 10. Create a map[string][]string (map to slice).
// 11. Merge two maps.
// 12. Invert a map[int]string → map[string]int.
```



### **Exercise Set 8: Pointers (Basic Pointer Concepts Only)**

```go
// TODO:
// 1. Declare an int variable and a pointer to it.
// 2. Print the variable, pointer, and dereferenced value.
// 3. Change the value via the pointer.
// 4. Pass a pointer to a function that increments the value.
// 5. Create a function returning a pointer to a local variable (is it allowed?).
// 6. Create a struct and modify its fields via pointer.
// 7. Create a pointer to a slice and modify slice contents via the pointer.
// 8. Compare pointers for equality.
// 9. Create a nil pointer and test it.
// 10. Create a pointer to a pointer and print its value.
```



### **Exercise Set 9: Functions**

```go
// TODO:
// 1. Write a function that adds two ints.
// 2. Write a function that returns two values (sum, difference).
// 3. Write a function with named return values.
// 4. Write a function that checks if a number is prime.
// 5. Write a function with a variadic parameter (...int).
// 6. Write a function that takes a function as a parameter.
// 7. Write a function that returns a function (closure).
// 8. Write a function that modifies a variable via pointer.
// 9. Write a recursive function (factorial or Fibonacci).
// 10. Write a function with an anonymous function inside it.
// 11. Write a function that panics and one that recovers from panic.
// 12. Document your function with a proper Go doc comment.
```



### **Exercise Set 10: I/O Basics**

```go
// TODO:
// 1. Read a string from the user.
// 2. Read two integers and print their sum.
// 3. Read a whole line using bufio.NewReader.
// 4. Convert input strings to numbers using strconv.
// 5. Implement simple menu input with switch.
// 6. Read until the user enters "quit".
// 7. Implement a small REPL that echoes input lines.
```



### **Exercise Set 11: Basic Error Handling**

```go
// TODO:
// 1. Read a number from the user and handle invalid input.
// 2. Create a function that returns (int, error).
// 3. Return a custom error using errors.New.
// 4. Wrap errors using fmt.Errorf(... %w ...).
// 5. Check an error with errors.Is.
// 6. Check an error with errors.As.
// 7. Use panic() intentionally and recover gracefully in defer.
// 8. Write a function that validates emails (simple rules) and returns errors.
```





## 2. Composite Types & Interfaces

Structs, slices, maps, method sets, embedding, interface design, and idiomatic composition over inheritance.


### **Exercise Set 1: Struct Basics**

```go
// TODO:
// 1. Define a struct `Person` with fields: Name string, Age int.
// 2. Create a Person value and initialize all fields.
// 3. Create another Person using named field syntax.
// 4. Modify a field of the struct.
// 5. Write a function that takes a Person and prints their details.
// 6. Write a function that takes *Person and increments the Age.
// 7. Compare two Person structs for equality.
// 8. Create an anonymous struct with 3 fields and print it.
// 9. Create a slice of Person and initialize it with 3 people.
// 10. Loop through the slice and print everyone's name.
```



### **Exercise Set 2: Methods on Structs**

```go
// TODO:
// 1. Add a method Greet() to Person that prints "Hello, my name is X".
// 2. Add a method IsAdult() that returns true if Age >= 18.
// 3. Add a pointer receiver method HaveBirthday() that increments Age.
// 4. Add a method ChangeName(newName string) using a pointer receiver.
// 5. Create several Person values and call each method on them.
// 6. Create a method on a non-struct type (e.g., type Score int).
```



### **Exercise Set 3: Struct Embedding (Composition)**

```go
// TODO:
// 1. Create a struct Address with fields City and Country.
// 2. Embed Address inside Person (anonymous field).
// 3. Create a Person and set the embedded fields.
// 4. Access embedded fields directly (e.g., p.City).
// 5. Add a method to Address and call it via Person.
// 6. Create two structs that each embed the same type.
// 7. Show method promotion in embedded types.
// 8. Create a struct that embeds multiple types (multi-embedding).
```



### **Exercise Set 4: Slices & Maps of Structs**

```go
// TODO:
// 1. Create a slice of Person.
// 2. Append new entries.
// 3. Filter the slice to get only adults.
// 4. Count how many people are under 30.
// 5. Sort the slice by Age (using sort.Slice).
// 6. Create a map[string]Person keyed by name.
// 7. Lookup a Person in the map using comma-ok.
// 8. Modify a Person’s age inside the map.
// 9. Create a map[string][]Person grouping people by city.
// 10. Build an index: map[int][]Person grouped by age.
```



### **Exercise Set 5: Interfaces — Basics**

```go
// TODO:
// 1. Create an interface Speaker with method Speak() string.
// 2. Make Person implement Speak().
// 3. Create another type Robot that also implements Speak().
// 4. Write a function PrintSpeech(s Speaker) that prints s.Speak().
// 5. Pass both Person and Robot to PrintSpeech.
// 6. Create a slice []Speaker and store both types.
// 7. Loop over the slice and call Speak() on each.
```



### **Exercise Set 6: Interfaces — Type Assertions & Type Switches**

```go
// TODO:
// 1. Create an interface{} value containing an int.
// 2. Use type assertion to extract the int.
// 3. Try a failing type assertion and handle the boolean result.
// 4. Write a function Describe(i interface{}) that prints type/value.
// 5. Use a type switch inside Describe() to handle int, string, bool.
// 6. Create a slice of interface{} values and handle each with Describe.
// 7. Write a function that accepts interface{} and only processes structs.
```



### **Exercise Set 7: Interfaces — Behavioral Composition**

```go
// TODO:
// 1. Create two interfaces: Walker and Talker.
//    - Walker has Walk() string
//    - Talker has Talk() string
// 2. Create a type Human that implements both.
// 3. Create a type Dog that implements only Walk().
// 4. Write a function UseWalker(w Walker).
// 5. Write a function UseTalker(t Talker).
// 6. Put both Human and Dog in a []Walker and call Walk().
// 7. Show that Dog cannot be used as a Talker.
// 8. Create an interface Animal combining Walker and Talker.
// 9. Show that only Human satisfies Animal.
```



### **Exercise Set 8: Interface Design Patterns**

```go
// TODO:
// 1. Create an interface Stringer with method String() string.
// 2. Make a struct implement fmt.Stringer (String() string).
// 3. Use fmt.Println to print the custom string representation.
// 4. Create a Reader-like interface with method Read() string.
// 5. Implement it for two different types.
// 6. Write a function that accepts the Reader interface.
// 7. Demonstrate interface-based dependency injection:
//    - Define an interface Logger with Log().
//    - Define two implementations: ConsoleLogger, NullLogger.
//    - Write a function Process(logger Logger).
//    - Call Process with both loggers.
```



### **Exercise Set 9: Empty Interface & Generics (Optional But Useful)**

```go
// TODO:
// 1. Store different types in a []interface{} slice.
// 2. Print each value using reflection (reflect.TypeOf).
// 3. Write a generic function Max[T constraints.Ordered](a, b T) T.
// 4. Write a generic function Map[T any, U any](...) that transforms slices.
// 5. Write a generic Stack type using type parameters.
// 6. Compare generic vs. interface{}-based collection performance.
```



### **Exercise Set 10: Practical Composition Example**

```go
// TODO:
// 1. Create struct Engine with field Horsepower.
// 2. Create struct Wheels with field Count.
// 3. Create struct Car that embeds both Engine and Wheels.
// 4. Add method Start() to Engine and call it from Car.
// 5. Add method Drive() to Car that uses embedded fields.
// 6. Implement interface Vehicle with Drive() and Fuel() methods.
// 7. Make Car implement Vehicle.
// 8. Create struct Bicycle that embeds Wheels and implements Vehicle.
// 9. Store Car and Bicycle in a []Vehicle and call Drive() on each.
// 10. Write a function DescribeVehicle(Vehicle) using type switches.
// 11. Use Promote methods to show how embedding passes methods up.
```



### **Exercise Set 11: Deep Interface Integration**

```go
// TODO:
// 1. Create a File-like interface with Open(), Read(), Close().
// 2. Create a MemoryFile type implementing the interface.
// 3. Create a Logger interface and wrap MemoryFile in a LoggedFile decorator.
// 4. Write a function ProcessFile(f File) that:
//    - Opens it
//    - Reads it
//    - Closes it
// 5. Show how the decorator adds behavior without modifying MemoryFile.
// 6. Add a CountingFile decorator to count number of reads.
// 7. Chain decorators: LoggedFile → CountingFile → MemoryFile.
```


## 3. Concurrency in Go

Goroutines, channels, select, sync primitives, race conditions, and common concurrency patterns (pipelines, worker pools).


### **Exercise Set 1: Goroutines**

```go
// TODO:
// 1. Write a function say(msg string) that prints the message.
// 2. Call say("normal call") normally.
// 3. Call say("goroutine call") inside a goroutine.
// 4. Launch 5 goroutines in a loop, each printing its index.
// 5. Observe interleaving of output.
// 6. Use time.Sleep to keep main alive long enough to see all output.
// 7. Replace Sleep with sync.WaitGroup so goroutines finish cleanly.
```



### **Exercise Set 2: Channels (Unbuffered & Buffered)**

```go
// TODO:
// 1. Create an unbuffered channel of type int.
// 2. Launch a goroutine that sends a number on the channel.
// 3. Receive it in main and print it.
// 4. Create a buffered channel of size 3.
// 5. Send 3 values without blocking.
// 6. Receive and print all values.
// 7. Write a sendString(ch chan string) function and call it as a goroutine.
// 8. Use a channel to return a computed result from a goroutine.
```



### **Exercise Set 3: Select & Time-Based Operations**

```go
// TODO:
// 1. Create two channels: ch1 and ch2.
// 2. Launch a goroutine that sends to ch1 after 200ms.
// 3. Launch a goroutine that sends to ch2 after 500ms.
// 4. Use select to print whichever value arrives first.
// 5. Add a default case that prints "waiting..." when no channel is ready.
// 6. Use time.After to trigger a timeout if no value arrives within 300ms.
// 7. Create a time.Ticker that prints "tick" every 100ms for 1 second.
// 8. Use select to handle ticker.C and a quit channel that stops the ticker.
```



### **Exercise Set 4: Sync Primitives (WaitGroup, Mutex, RWMutex, Atomic)**

```go
// TODO:
// 1. Use WaitGroup to run 5 goroutines printing their ID.
// 2. Wait for all goroutines to finish.
// 3. Create a shared counter.
// 4. Increment it with 100 goroutines, 100 times each → expect 10,000.
// 5. Protect the counter using sync.Mutex.
// 6. Rewrite the increment using sync/atomic.AddInt64.
// 7. Create a map shared between goroutines.
// 8. Protect it with sync.RWMutex during reads/writes.
// 9. Measure atomic vs mutex performance using time.Since.
```



### **Exercise Set 5: Race Conditions**

```go
// TODO:
// 1. Create a shared variable x := 0.
// 2. Start 50 goroutines that increment x without synchronization.
// 3. Print final x and observe it is NOT 50.
// 4. Run using: go run -race main.go
// 5. Observe race detector output.
// 6. Fix the race using a Mutex.
// 7. Fix the race using an Atomic integer.
// 8. Explain in comments why races happen even with append() on slices.
```



### **Exercise Set 6: Concurrency Patterns — Pipelines**

```go
// TODO:
//
// 1. Build a 3-stage pipeline:
//      Stage 1: generate numbers 1–5 into a channel.
//      Stage 2: square each number.
//      Stage 3: convert each number to a string.
// 2. Connect stages using channels.
// 3. Print the final output strings.
// 4. Modify pipeline to handle errors (return a struct {value, err}).
// 5. Add cancellation using context.WithCancel.
// 6. Make Stage 2 parallel by launching 2 worker goroutines.
// 7. Ensure Stage 3 collects results in correct order.
```



### **Exercise Set 7: Concurrency Patterns — Worker Pools**

```go
// TODO:
//
// 1. Write a worker function:
//      worker(id int, jobs <-chan int, results chan<- int)
//    It should square the input.
//
// 2. Create channels: jobs and results.
//
// 3. Launch 3 workers as goroutines.
//
// 4. Send 10 jobs (numbers 1–10), then close jobs.
//
// 5. Collect 10 results and print them.
//
// 6. Modify workers to simulate variable processing time using Sleep.
//
// 7. Add context cancellation to stop all workers early.
//
// 8. Add error handling by returning structs {jobID, result, err}.
```




## 4. Error Handling & Panic/Recover

Idiomatic error handling, wrapping/annotating errors, sentinel errors, panic, recover, and structured error patterns.

## 5. Package Management & Tooling (Environment)

Go modules (go mod init, go mod tidy, etc.), dependency caching, go fmt, go vet, go test, linters, static analysis, and workspace structure.

## 6. Testing, Benchmarking & Profiling

Unit tests, table-driven tests, subtests, fuzz testing, benchmarks, CPU and memory profiling, pprof, and tracing.

## 7. Building & Deploying Go Applications

Static binaries, build tags, cross-compilation, environment-based config, module vendoring, and deployment strategies.

## 8. Standard Library Mastery

Key packages such as fmt, io, bufio, strings, bytes, time, context, net/http, encoding/json, and common idioms around them.

## 9. Networking & Web Development

HTTP servers and clients, middleware patterns, request handling, JSON APIs, Go’s net package, and concurrency in networked applications.

## 10. Memory Management & Performance

Escape analysis, stack vs. heap allocation, garbage collection basics, profiling tools, and writing memory-efficient Go code.

If you want, I can also create beginner → advanced learning paths, weekly study plans, or example exercises for each topic.