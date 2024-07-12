---
author: "Frank Murillo"
title: "Gophercon 2024 Notes Dump"
description: "Gophercon 2024 Notes Dump"
date: "2024-07-12"
tags: ["golang", "gophercon"]
---
# GopherCon 2024 Notes 
# Talks
---
## Day 1
---
### Charm CLi (Habbit Tracker) - Donia Chaiehloudj
---
- Charm is a teminal UI written in go
- uses ELM archetecure - Model, View, Update
- grpC backend? (To store habbits etc)
- check out book `Learning Go with pocket sized projects`
- TLDR;Pretty cool, have used before not very succesfully, if I work with the Architecture of the project, I think it will be a lot easier to work with.

### Adding "four" loop to Go compiler - Riley Thompson
---
- Add "four" and "unless" statements
- First Stage of Compiler
    - Lexical Analysis and Parsing
    - source code Scaned and tokenized
    - Recusrive Decent Tree
- Add the new tokens as struct, embed the stmt struct
- Type Checking
- IR (Intermediate Representation) Contruction (noding)
    - Converting to AST (Abstract Syntax Tree)
- Escape analysis (what is this)
- Walking through the AST
    - "Order" step
- AST is converted to an IR in Static Single Assignment form
- IR is converted to SSA form
- This can be implemented in different stages in the compiler
- TLDR; Good talk, reminded me alot of the book `Making an interpreter in go` seems to be very applicable 

### Advanced Generics Patterns - Alex Wagner
---
- Generics in Go are a hot topic (Are they event used? I have never used them)
- They are objectively limited in Go, also a culutural limitation
- Generics allow us to define types with Type arguments 
```go
    type Slice[E Any] []E
    func (s Slice[E]) Append(e E) Slice[E] 
```
- Have to intantiate a type explicitly and fully when using it
- constraints hard to call methods on certain types
- Can use union types, call methods, or have the function needed as a function argument (go has union types?)
- Unmarshal seems like a good case for generics, however it is not a good fit for the Go implementation of generics, like `json.Unmarshal`
- TLDR; Experiment with Generics and see how far you can get with them, they have lots of contraints and you need to work around these contstraints.

###  Building a Deterministic Interpreter in Go: Readability vs Performance - Jae Kwon 
---
- a Smart Contract is a program that is run on the blockchain, it is usually untrusted code
- Devoped GnoVM (Go Virtual Machine) to solve issues with smart contracts and WASM
- Variables declared as interface always creates a pointer
- Scope != Allocation (for loop vars)
- Go reflection is limited

### Lightning Talks 
---
- Standardizing Errors in Go: A Practical Guide with Dapr | Cassie Coyle, Diagrid
    - Richer Error Model: code, message, details 
    - Errors messages should provide as much information as possible, evenn resources to help fix the problem
    - Dapr provides an error package that can be used to create errors with rich details
- Chan Chan Chan T: A Generic Tale | Branden J. Brown
    - finding a match in communication, talking to other goroutines
    - communication deduplicates state
- Go's use in Competitive Programming
    - CP is answering questions via code usually taken in STDIN and STDOUT
    - Pros: ...? Its fast? 
    - Cons: Compiler support, problems arent suitable to parallelizing, stdlib lacks in DSA, verbose
    - CPP is anti-software engineering (its throw-away code), it is not a good fit for Go
- 7 Surprising Features of the Go Playground | Matt Dale, MongoDB
    - support for pulling in modules
    - allows to create multiple files (go.mod) `---filepath---`
    - can write tests in the playground
    - clear the terminal by prinitng `fmt.Println("\x1b[2J\x1b[H")`
    - display images  
    - goplay.tool goplay.space

## Saving African Savannah Wildlife with Go-MQTT - Marvin Hosea
---
- Decided to use Go for the backend, because it is a good language for IoT
- MQTT is a lightweight pub/sub protocol, used for IoT -> use the Mosquitto broker (free!)
- This Go MQTT client library was not very good, hand bugs and not very well tested, speaker decided to work on it himself
- The end...
- TLDR; Short speach, but a good talk for a even better cause. Cool how Go can be used to solve real world problems, IOT is a big one.

## Building a promgrammable firewall with Go - Mason Johnson
---
- Firewalls are a very important part of any network, they are used to prevent unwanted traffic
- Using go packages (`netlink, nftables, firewall_toolkit`)
- Get netlink connection using netlink package
- Add a table using the netlink connection (`need to call .Flush()`)
- Add a chain to the table and Flush
- create a new rule using firewall_toolkit lib
- Can create a set to store the mutilple IP
- Can connect via Kafa (sends blocked IP's) and send the blocked IP's to the firewall updating the blocked IP's set
- TLDR; Go is great for Network programing, can program a firewall with an ergonomic API

## Automating Efficiency Improvement by Profile Guided Optimization - Jin Lee
---
- Uber has over 26000 Go microservices
- PGO is a compiler optimization technique that can be used to improve performance of Go programs
- Looks like a pretty complex process...
- Go-json PGO inlining saw improvments of 12% in performance
- PGO originally exptended compile times by 4x, until PGO pre-processing was introduced
- TLDR; PGO is a great tool to improve performance of Go programs, but it is a complex process.

### Prototyping a Go, GraphQL, gRPC Microservice - Sadie Freeman
---
- Should always start with a prototype, articulate the problem what stack, why, etc
    - Documents the process (pros-cons)
    - Use existing libs/frameworks
    - Use useful logs (!)
    - Don't overcomplicate, make assumptions
- Archetcture - inlcudes two services one for creating Alpacha and another for analytics each with their own DB
- Why GraphQL?
    - More flexible querying, front-end friendly
    - Fewer server requets 
    - Powerful pagination 
- Define proto file schema and generate code (resolver)
- gRPC (service to service communication)
    - Fast & Efficent
    - HTTP/2
    - Protocol Buffers (Binary serialization)
    - Clear Proto schmea
- Can quickly get up and running with this stack
- TLDR; Cool talk, very practical and useful, would like looking for into in with GraphQL, gRPC, and Buf 

### Closing Talk (Day 1): A Decade of Debugging - Derek Parker
---
- Go debuggers were immature, did not track goroutines properly
- in 2014 Delve was created to address this issue by Derek
- Q? Who uses this? Neovimers? ...No everyone! It is used behind the scenes for all these debuggers
- TLDR; mostly a talk about the delve debugger and change log throught out the 10 years...kinda cool debuggers are interesting

---
## Day 2
---
### Who Tests the Tests? - Daniela Petruzalek
---
- A way to validate tests - Rotation Testing
- Why do we write tests? 
    - Tests are our insurence policy
    - The code does what its supposed to do, happy path :)
    - The code doesnt do what its not supposed to do, sad path :(
- Tests are also design tools
    - First clients of our code
    - Strong correlation between code that is easy to test and code that is easy to read/maintain
- Measauring tests
    - Code coverage
    - Test coverage 
- Goodharts Law - A metric that becomes a goal stops being a good metric
- Test coverage is easily faked
- You can test tests via mutations, esentially mutate the code you are testing and it should fail. This is mutant tests coverage.
- In Go you can use an access the AST to program these mutations to happen during runtime
- Write tests for the right reasons, choose your test inputs wisely, mutation testing in go is not mature yet
- TLDR; very interesting talk, first time hearing of Rotation/Mutation testing, however seems like alot of work to get running

###  Interface Internals - Keith Randall
---
- Multiply simply asks the machine to multiply - nothing needed by Go compiler
- For interfaces there needs to be work...find the contained type..get list of methods..etc
- Can be done in 3 machine code instructions
- Interface variables bridge the gap between static and dynamic worlds, it can be reassigned 
- Interfaces in Go are just 2 word structs with a pointer to the interface type and a pointer to the interface data
- Utilize an interface tab on interface data pointer to store addiotnal data ->  all methods of an interface (sorted order) `fun[0]`
- TLDR; Very interesting talk, very useful to know how interfaces are implemented in Go and how the comiler converts to machine code

### Building a High-Performance Concurrent Map in Go - YunHao Zhang
---
- Go's built in map is not concurrent safe 
- `sync.Map` is provided by stdlib 
- using built in map with a `sync.RWMutex`  (together in a struct?)
- `sync.Map` is optimized for cahce contention - only fast to reads
- For RWMap there is a bottle neck for writes (blocks all reads)
- We can create a shared map = multiple read-write mutex maps
- Sounds very similar to databse sharding but for maps, writes will only block the map it is writing too
- Hash function to know which map to write/read to
- In all cases this implementation is faster, however takes more memory
- Can create a sub-map each new key-value pair creates a new node and all nodes are connected via pointers (LinkedList) 
- SkipList = multiple linked lists
- SkipMap is 20x faster  in typical cases however it is O(logn) on read/writes, `sync.Map` can be faster in low-currency cases
- TLDR; cool talk about data structures and system design, kina got lost towards the end but very interesting

### The Go Cryptography State of the Union - Fillipo Valsorda
---
- Post Quantum Cryptography is about the future, these computers are powerful and can crack these cryptographic algorithms
- Hyribs of new and old implementations to solve both problems
- New lib `crypto/internal/mlkem768` 300 lines made for readability
- fast and low memory footprint
- `cryoto/tls` updated defaults in `1.24`
- encrypted client hello in `1.23`
- `math/rand` should not be used for security before version `1.22`
- `ChaCha8Rand` is now uses in `rand` packages to make secure
- still should use `crypto/rand` these apis shold no longer return erros as they are now deterministic
- `ssh` package is catching up
- cool new API `secret.Do` that takes a closure and code is cleared from stack after done
- TLDR; cryptography is hard, Go maintainers are up for the challenge and the libs are battle tested and improving

### Advanced Code Instrumentation Techniques for High-Performance Trading Systems - Holly Casaletto
---
- "You can observe a system, but be prepared for it to act like it knows your watchinig!"
- A trading needs to be highly durable and performant (high throughout - low latency - concurrent)
- Merics API where we need to handle locking, etc
- TLDR; was kinda tired for this one, was very verbose and a bit too much info for me

### Digital Audio from Scratch - Patrick Stephen
---
- Finally some music
- using a lib `daw` we are sending bytes into speaker -> higher byte = higher volume, pattern = frequency
- speaker is going through different waves/numbers/and math to get a cool sound
- TLDR; Music is cool and interesting (also complicated), you can do it in Go if you want..

### Lightning Talks
---
- How to Mock Your Coworkers - Dylan Bourque from CrowdStrike
    - Mocking is test doubling - you can mock out the dependencies of a function
    - Don't export your mocks...write(or generate) your mocks internal to your project
    - ...dont actaully mock your coworkers lol
- Lightning-Fast Database Tests -  Robin Hallabro-Kokko
    - DB tests are usually the slowest tests
    - Fast tests lower the barier for creating new tests (CI/CD, faster development)
    - `pgtestdb` creates emphemal postgres databases quickly and handles lifecycle
    - it is very quick and seems easy to use
- Built it! With or WIthout Tools - Matthew Sanabria
    - Speaker compares refactoring production code to building his new kitchen 
    - Right Tools, Wrong Knowledge - Learn! Read the docs
    - Wrong Tools, Write Tests - It happends, verify tools work by writing tests
    - Know when to seek an expert, and observe their work
    - Greatest tool in your toolbox is you
- You can store that in a container registery? - Siva Manivannan
    - Yes, you can store things that are not containers in a container registry
    - Container Registry is metadata and a blob
    - If you implment these ~20 API endpoints you have a container registry 
- Would You Like a GUI With That? - Andy Joseph, ProntoGUI
    - ProntoGUI is a lib that allows to build GUI's in Go
    - Events from frontend to backend are via gRPC 
    - Kinda cool? Don't see how its better/worse than other GUI libs
- Embracing the Replace: Our Journey to Fixing the Plugin System in HashiCorp Packer - Wilken Rivera, HashiCorp 
    - Company had a breaking change bc they were using a experimental package that eventaully was dropped
    - Options were limited, decided to fork the package and fix it
    - added a CLI command to fix the issue for their users
    - when having a breaking change, stop the bleed, and do whats best for users even if its the hardest

### A 10-Year Retrospective on Building an Infrastructure Start-up in Go - Alan Shreve
---
- ngrok originall use was to have public url for local project running on a port
- when started in 2014 go was still very immature in problems related to ngrok
- sqlc for database stuff
- Use lots of gRPC over ~100 services
- Interceptors (like middleware but for gRPC) can be used on client and server side
- use RichErrors (custom Error type struct) and have a client side interceptor that catches this 
- Replicate all the things
- Integration Testing - matrix of thosands of tests run in production 24/7 in a loop (wow)
- Define your errors to help users
- ngrok's success is acredited to Go and its proerties
- dont get attached to code
    - software is meant to be rewritten

### Processing Millions of Events Per Second Reliably Using Generics - Alan Braithwaite
---
- Stream processing processes a continouts stream of distint events consuming from some source external 
to the system and delivering to some desination external to the system
-  Goals for the project, be fast, plugins are easy to create, support any serializable type, batteries included, at least one delivery
- Okay this is where generics comes in, using the Go lib [kawa](https://github.com/runreveal/kawa) the types are used as generics
- Speaker argues against using channels due to their overhead and unseen complexity 
- TLDR; stream proccessing is very interesting, i want to look into this. Project seems cool

### Building a Self-Regulating Pressure Relief Mechanism in Go - Ellen Gao
---
- it is importnat for systems to have healthy utlization (cpu?) 
- Not always best option to just increase resources, (cost, complexity)
- A presure valave can detect a soft limit and perform action
- slow down processing of non-critical processes, these messges are processed later
- Program uses a messague broker pub/sub system to mimic opearations and priority
- Basic throttler implemented with `time.Sleep` program now completes
- Adds 3 different throttling strategies (block, fixed delay, and backoff) for each priority via interfaces 
- Then implements recovery, most was already implemented so not much changes needed
- TLDR; interesting talk, good speaker, learned some about throttling and pressure relief, liked the the flow from talk to code. 

### Go in the Smallest of Places - Patricio Whittingslow]
---
- First major software was the Apollo spaceship and Margarette Lewis was first SWE, written in assembly
- C was created a higher level language (comparatievly) for micro controllers
- TinyGo makes Go usuable for MCU development
- Working with MCU's is verbose, you needd to shift a bunch of bits in memory (registers)
- TLDR; TinyGo is awesome and I need to use it with the programabale badge I have

### Range Over Function Types - Ian Lance Taylor 
---
- New feature in 1.23 release
- Ways to loop over all of an elements in a Set in Go
    - Extend the for range, or add iterator type, how about both
- `.All()` method should be created for container types which returns an iterator 
- TLDR; finally iterators in go, altough don't love the implementation

## End of Day 2 
