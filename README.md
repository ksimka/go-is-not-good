# What's this

This repository is a list of articles that complain about **golang**'s imperfection.

## Motivation

Seems like complaining about **go**'s flaws is becoming a trend. Any newbie must have a chance to read all the **go**-is-bad arguments before they go too far. So here it is.

## What it's for

This repo is not aimed to offend or insult someone (at least not more than each author does it in its article), especially **golang** itself, its authors and the community. It is for educational purpose only. Any contributor can have an absolutely different point of view.

I don't think anyone would deny that **go** has weaknesses: it certainly has. But how do you know, *is it really a language design flaw or is it just you, doing something completely wrong*? This list here to help you quickly answer the question.

## How to use it

You're writing some code. And suddenly you understand you need something that language can't give you. You go here and check if you're the one with that issue or not. If it's a common issue, it'll be here. Then you decide what to do: choose another tool for your task or go find a better solution or a workaround.

## See also

[How to complain about Go](https://divan.github.io/posts/go_complain_howto/)

# The List

+ http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh, 2009)
  - slower than C++
  - CSP implementation is only partial
  - standard library is limited
  - no GUI support in standard library
  - no database support in standard library
  - statically linked
  - no exceptions
  - syntax is too different from C, C++, and Java
+ https://cowlark.com/2009-11-15-go/ (David Given, 2009)
  - c-style
  - poor design
  - no constructors
  - `new` and `make` instead of one
  - stuck in '70s
  - no OOP
  - types go after identifiers, not before
  - return types after function name, not before
  - no while loops
  - can't range over user-defined types
  - constants can only be numbers or strings
  - pointers are a mess
  - can't name nested functions
  - single-pass compiler
  - designers did not consult the literature
  - too simple / lack of syntactic sugar
+ http://www.lessonsoffailure.com/software/google-go-good-for-nothing/ (Dave Rodenbaugh, 2009)
  - slower than C++
  - memory footprint is worse than C++
  - syntax is too different from C, C++, and Java
+ https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis, 2010)
  - no language interoperability (only C)
  - no versioning model
  - no OOP
  - no GUI support
  - `new` and `make` instead of one
  - internationalization is limited to Unicode support
  - has pointers
  - no semicolons at line endings
  - no `this`
  - no function/operator overloading
  - not real-world oriented
  - `defer` lets you be lazy
  - ignores C#
  - no exceptions
+ http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish, 2012)
  - c-style
  - unused imports are compile errors
  - unused variables are compile errors
  - bad Unicode support
  - types go after identifiers, not before
  - return types after function name, not before
  - no ternary operator
  - too verbose
  - semicolons are injected by the lexer
  - hard to pass function pointers to C
  - no incremental or parallel compilation
  - closures capture variables by reference, not by value
  - no assertions
+ https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield, 2012)
  - error handling is mandatory
  - stuck in '70s
  - no exceptions
+ http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King, 2012)
  - un-Googlable name
  - stuck in '70s
  - no exceptions
  - is compiled
  - no read-eval-print loop
  - `fork()` is either wrong or impossible
  - string manipulation is a library package
  - too young
+ http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey, 2013)
  - error handling is mandatory
  - no exceptions
  - inconvenient `range`
  - can't range over user-defined types
  - tab width too wide
  - gofmt's style is no one's favourite
  - no generics
+ http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi, 2013)
  - too opinionated
  - unused imports are compile errors
  - unused variables are compile errors
+ https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson, 2013)
  - too simple / lack of syntactic sugar
  - no generics
  - string manipulation is a library package
  - too young
  - error handling is not mandatory
  - compiler warnings are errors
  - no function/operator overloading
  - no keyword arguments
  - no shebang lines
+ http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man, 2013)
  - no OOP
  - no assertions
  - has pointers
  - solves nobody's problem
  - no inheritance
  - no function/operator overloading
  - designers did not consult the literature
  - weird mascot (gopher)
+ https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman, 2014)
  - no decent IDE
  - GOPATH is a mess
  - pointers are a mess
  - case-defined scoping is bad
  - hidden types
  - immature GC
  - hard to extend
  - no exceptions
  - no generics
  - channels are overrated
  - duck typing of local variables
  - dependency management sucks
+ http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin, 2014)
  - confusing/stupid syntax
  - immature toolchain
  - immature GC
  - no generics
  - multiple return values have no type checking
  - cumbersome interface
  - goroutines are not original
  - `defer` is abused
  - stuck in Unix thinking
  - types go after identifiers, not before
  - no decent IDE
  - dependency management sucks
  - unwieldy to code new collections
  - designers did not consult the literature
  - error handling is mandatory
  - summary: not elegant as Python, not strong as Java
+ http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski, 2014)
  - the worst compiler toolchain ever
  - Rob Pike and other core developers are arrogant
  - stuck in '70s
  - Fisher-Price assembly language
  - stuck in Plan9 thinking
  - technical hubris
  - assembler uses Unicode mid-dot character
  - confusing and undebuggable
+ http://yager.io/programming/go.html (Will Yager, 2014)
  - no generics
  - no function/operator overloading
  - nil as a failure marker
  - type inference is too simple
  - no immutables
  - not a good language
  - unwieldy to code new collections
  - can't range over user-defined types
  - no ternary operator
  - dynamic memory allocation
  - garbage collector requires implicit (hidden) trade-offs
  - not real-time oriented
  - no unsafe code isolation
  - stuck in '70s
  - no pattern matching
+ https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan, 2014)
  - no decent IDE
  - error handling is mandatory
  - unexpected variable shadowing
  - no generics
  - no `map`/`reduce`/`filter`
  - no immutables
  - not-so-obvious slices behaviour
+ http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda, 2014)
  - no built-in set type
  - no built-in BitSet/BitArray type
  - maps are slow
  - slower than Java
  - Rob Pike and other core developers are arrogant
  - no generics
+ http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel, 2015)
  - not-so-obvious slices behaviour
  - error handling is mandatory
  - hard to test, hard to mock
+ https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat, 2015)
  - no generics
  - interfaces are not contracts
  - no function/operator overloading
  - no exceptions
  - can't range over user-defined types
  - sending to a closed channel panics
  - channels are not as efficient as plain mutexes
  - `defer` is slow
  - interface indirection is expensive
  - can't peek into channels
  - can't receive multiple values from channels
  - goroutines make it easy to leak things
  - dependency management sucks
  - immature toolchain
  - Rob Pike and other core developers are arrogant
  - can't add methods to types from other packages
+ http://blog.goodstuff.im/golang (David Pollak, 2015)
  - no immutables
  - no ternary operator
  - inconvenient `range`
  - := is weird
  - no exceptions
  - unused imports are compile errors
  - unused variables are compile errors
  - no read-eval-print loop
  - too simple / lack of syntactic sugar
  - confusing/stupid syntax
  - too opinionated
  - error handling is mandatory
  - error handling is not mandatory
  - designers did not consult the literature
  - stuck in '70s
  - no generics
  - case-defined scoping is bad
  - no `map`/`reduce`/`filter`
+ http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov, 2015)
  - designed for stupid people
  - too verbose
  - no ternary operator
  - no macros or templates
  - can't change hash function in maps
  - unwieldy to code new collections
  - inconvenient `range`
  - no function/operator overloading
  - can't declare/validate implements interface
  - no virtual functions
  - case-defined scoping is bad
+ https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell, 2015)
  - regex engine is slow
  - no generics
  - serialization is inefficient
  - is garbage-collected
  - no good integration test suite
  - no dynamic linking / plugins
  - no decent XML parser
+ http://www.evanmiller.org/four-days-of-go.html (Evan Miller, 2015)
  - unused imports are compile errors
  - too opinionated
  - poor std math lib
  - := is weird
  - no while loops
  - 'rune' is a weird name
  - is garbage-collected
  - unused variables are compile errors
  - pseudointellectual arrogance of Rob Pike and everything he stands for
  - weird mascot (gopher)
+ http://byrd.im/go-is-poor/ (Ian Byrd, 2015)
  - not-so-obvious slices behaviour
  - nil interfaces are not entirely nil
  - unexpected variable shadowing
  - interfaces are not first-class values
  - questionable compiler rigidity
  - go generate is a quirk
+ https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi, 2015)
  - no generics
  - slow json parsing
  - dependency management sucks
  - too verbose
  - GOPATH is a mess
  - no subpackages
+ http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg, 2015)
  - un-Googlable name
  - confusing/stupid syntax
  - polymorphism is broken
  - compilation rules are too confining
  - designed for stupid people
  - can't declare/validate implements interface
  - dependency management sucks
+ https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot, 2016)
  - case-defined scoping is bad
  - interfaces are not first-class values
  - no exceptions
  - difficult to generate code automatically
  - no ternary operator
  - sort.Interface approach is clumsy
  - no versioning model
  - import-based vendoring is terrible
  - no generics
+ http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds, 2016)
  - channels are slow
  - channels are an anti-pattern
  - channels are not as efficient as plain mutexes
  - channel buffers are a fixed size
  - sending to a closed channel panics
  - sending to a nil channel does not panic
  - channel API is inconsistent
+ https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker, 2016)
  - dependency management sucks
  - no generics
  - no sum types
  - no separation of pure code from effectful code
  - godoc doesn't support even basic formatting
  - godoc doesn't use source code ordering
  - profiler doesn't graph allocations over time
  - profiler doesn't track threads or garbage collection
  - zero values are a bad idea
  - too verbose
  - error handling is mandatory
  - weak typing
+ https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus, 2016)
  - not stable
  - no `map`/`reduce`/`filter`
  - no OOP
  - dependency management sucks
+ http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal, 2016)
  - stuck in Plan9 thinking
  - no assertions
  - un-Googlable name
  - tracebacks limited to 100 stack frames
  - confusing and undebuggable
  - no frame pointers
  - too opinionated
+ https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn, 2016)
  - misleading marketing around garbage collector
  - garbage collector is nothing new; concurrent mark/sweep from the '70s
  - garbage collector requires implicit (hidden) trade-offs
  - stuck in '70s
  - garbage collector optimized for pause times at the cost of other desirable GC features
+ https://faiface.github.io/post/context-should-go-away-go2/ (Michal Štrba, 2017)
  - context spreads like a virus
  - context is an inefficient linked list
  - context is not an elegant solution to the cancellation problem
  - ctx.Value is a map of meaningless objects to meaningless objects
+ http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett, 2017)
  - no non-nullable types
  - error handling is mandatory
  - error handling is not mandatory
  - no generics
  - no sum types
  - zero values are a bad idea
  - too verbose
  - no macros or templates
  - tuples are not first-class values
  - no `map`/`reduce`/`filter`
  - can't range over user-defined types
  - can't `make` user-defined types
  - unwieldy to code new collections
  - designers did not consult the literature
  - concurrency and parallelism are mixed
+ https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson, 2017)
  - too opinionated
  - error handling is mandatory
  - pointers are a mess
+ https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze, 2017)
  - no sum types
  - no algebraic data types
  - tuples are not first-class values
  - can't pass multiple return values to a function
  - multiple return parameters are overrated
+ https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol, 2017)
  - shortage of experienced developers
  - import-based vendoring is terrible
  - compiler warnings are errors
  - circular dependencies are errors
  - unwieldy to code new collections
  - no generics
  - too young
  - debugging support is poor
  - profiling support is poor
  - public library packages are half-baked or abandoned
  - dependency management sucks
+ https://zapier.com/engineering/go-no-go/ (Jordan Sherer, 2017)
  - unwieldy to code new collections
  - dependency management sucks
  - no inheritance
  - no classes
  - no generics
  - no exceptions
+ https://grimoire.ca/dev/go (Owen Jacobson, 2018)
  - hostile to developer ergonomics
  - too opinionated
  - dependency management sucks
  - GOPATH is a mess
  - Rob Pike and other core developers are arrogant
  - no sum types
  - too verbose
  - error handling is mandatory
+ https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez, 2018)
  - no generics
  - sort.Interface approach is clumsy
  - designers did not consult the literature
  - Fisher-Price assembly language
  - no functional programming
  - is garbage-collected
  - interfaces are not contracts
  - can't declare/validate implements interface
  - no enums
  - unexpected variable shadowing
  - zero values are a bad idea
  - error handling is mandatory
  - writing to a zero map panics
  - dependency management sucks
  - GOPATH is a mess
  - not-so-obvious slices behaviour
  - no immutables
  - race conditions are possible
  - io.EOF is an error
  - not Rust
  - interface values that hold nil concrete values are non-nil
  - struct field tags are free-form strings
  - no constructors
  - unwieldy to code new collections
  - can't range over user-defined types
  - no macros or templates
  - go generate is a quirk
  - no exceptions


# Reverse complaints index

It's a reverse complaints index, generated by https://github.com/ksimka/go-is-not-good/blob/master/generator.go (thanks to @capoferro)

+ 'rune' is a weird name
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ := is weird
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ CSP implementation is only partial
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
+ Fisher-Price assembly language
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ GOPATH is a mess
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ Rob Pike and other core developers are arrogant
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ `defer` is abused
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ `defer` is slow
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ `defer` lets you be lazy
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ `fork()` is either wrong or impossible
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ `new` and `make` instead of one
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ assembler uses Unicode mid-dot character
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ bad Unicode support
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ c-style
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ can't `make` user-defined types
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ can't add methods to types from other packages
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ can't change hash function in maps
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ can't declare/validate implements interface
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ can't name nested functions
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ can't pass multiple return values to a function
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ can't peek into channels
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ can't range over user-defined types
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ can't receive multiple values from channels
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ case-defined scoping is bad
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ channel API is inconsistent
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channel buffers are a fixed size
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channels are an anti-pattern
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channels are not as efficient as plain mutexes
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channels are overrated
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ channels are slow
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ circular dependencies are errors
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ closures capture variables by reference, not by value
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ compilation rules are too confining
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ compiler warnings are errors
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ concurrency and parallelism are mixed
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ confusing and undebuggable
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ confusing/stupid syntax
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ constants can only be numbers or strings
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ context is an inefficient linked list
  - https://faiface.github.io/post/context-should-go-away-go2/ (Michal Štrba 2017)
+ context is not an elegant solution to the cancellation problem
  - https://faiface.github.io/post/context-should-go-away-go2/ (Michal Štrba 2017)
+ context spreads like a virus
  - https://faiface.github.io/post/context-should-go-away-go2/ (Michal Štrba 2017)
+ ctx.Value is a map of meaningless objects to meaningless objects
  - https://faiface.github.io/post/context-should-go-away-go2/ (Michal Štrba 2017)
+ cumbersome interface
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ debugging support is poor
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ dependency management sucks
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ designed for stupid people
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ designers did not consult the literature
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ difficult to generate code automatically
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ duck typing of local variables
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ dynamic memory allocation
  - http://yager.io/programming/go.html (Will Yager 2014)
+ error handling is mandatory
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ error handling is not mandatory
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ garbage collector is nothing new; concurrent mark/sweep from the '70s
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector optimized for pause times at the cost of other desirable GC features
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector requires implicit (hidden) trade-offs
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ go generate is a quirk
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ godoc doesn't support even basic formatting
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ godoc doesn't use source code ordering
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ gofmt's style is no one's favourite
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ goroutines are not original
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ goroutines make it easy to leak things
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ hard to extend
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ hard to pass function pointers to C
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ hard to test, hard to mock
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
+ has pointers
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ hidden types
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ hostile to developer ergonomics
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ ignores C#
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ immature GC
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ immature toolchain
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ import-based vendoring is terrible
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ inconvenient `range`
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ interface indirection is expensive
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
+ interface values that hold nil concrete values are non-nil
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ interfaces are not contracts
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ interfaces are not first-class values
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ internationalization is limited to Unicode support
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ io.EOF is an error
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ is compiled
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ is garbage-collected
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ maps are slow
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
+ memory footprint is worse than C++
  - http://www.lessonsoffailure.com/software/google-go-good-for-nothing/ (Dave Rodenbaugh 2009)
+ misleading marketing around garbage collector
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ multiple return parameters are overrated
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ multiple return values have no type checking
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ nil as a failure marker
  - http://yager.io/programming/go.html (Will Yager 2014)
+ nil interfaces are not entirely nil
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ no GUI support
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no GUI support in standard library
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
+ no OOP
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ no `map`/`reduce`/`filter`
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no `this`
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no algebraic data types
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ no assertions
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ no built-in BitSet/BitArray type
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
+ no built-in set type
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
+ no classes
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
+ no constructors
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no database support in standard library
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
+ no decent IDE
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
+ no decent XML parser
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
+ no dynamic linking / plugins
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
+ no enums
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no exceptions
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no frame pointers
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ no function/operator overloading
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no functional programming
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no generics
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no good integration test suite
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
+ no immutables
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no incremental or parallel compilation
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ no inheritance
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
+ no keyword arguments
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
+ no language interoperability (only C)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no macros or templates
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ no non-nullable types
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no pattern matching
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no read-eval-print loop
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ no semicolons at line endings
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no separation of pure code from effectful code
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ no shebang lines
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
+ no subpackages
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ no sum types
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ no ternary operator
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no unsafe code isolation
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no versioning model
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no virtual functions
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no while loops
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ not Rust
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ not a good language
  - http://yager.io/programming/go.html (Will Yager 2014)
+ not real-time oriented
  - http://yager.io/programming/go.html (Will Yager 2014)
+ not real-world oriented
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ not stable
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ not-so-obvious slices behaviour
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ pointers are a mess
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
+ polymorphism is broken
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ poor design
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ poor std math lib
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ profiler doesn't graph allocations over time
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ profiler doesn't track threads or garbage collection
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ profiling support is poor
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ pseudointellectual arrogance of Rob Pike and everything he stands for
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ public library packages are half-baked or abandoned
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ questionable compiler rigidity
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ race conditions are possible
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ regex engine is slow
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
+ return types after function name, not before
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ semicolons are injected by the lexer
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ sending to a closed channel panics
  - https://bravenewgeek.com/go-is-unapologetically-flawed-heres-why-we-use-it/ (Tyler Treat 2015)
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ sending to a nil channel does not panic
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ serialization is inefficient
  - https://thenewstack.io/switching-to-go-and-learning-what-features-its-missing-along-the-way/ (Matthew Campbell 2015)
+ shortage of experienced developers
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ single-pass compiler
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ slow json parsing
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ slower than C++
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
  - http://www.lessonsoffailure.com/software/google-go-good-for-nothing/ (Dave Rodenbaugh 2009)
+ slower than Java
  - http://oneofmanyworlds.blogspot.co.uk/2014/01/another-go-at-go-failed.html (Srinivas Jonnalagadda 2014)
+ solves nobody's problem
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ sort.Interface approach is clumsy
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ standard library is limited
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
+ statically linked
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
+ string manipulation is a library package
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
+ struct field tags are free-form strings
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ stuck in '70s
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ stuck in Plan9 thinking
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ stuck in Unix thinking
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ summary: not elegant as Python, not strong as Java
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ syntax is too different from C, C++, and Java
  - http://www.lessonsoffailure.com/software/googles-go-not-getting-us-anywhere-part-2/ (Dave Rodenbaugh 2009)
  - http://www.lessonsoffailure.com/software/google-go-good-for-nothing/ (Dave Rodenbaugh 2009)
+ tab width too wide
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ technical hubris
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ the worst compiler toolchain ever
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ too opinionated
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ too simple / lack of syntactic sugar
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ too verbose
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ too young
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - https://blog.carlmjohnson.net/post/google-go-the-good-the-bad-and-the-meh/ (Carl Johnson 2013)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
+ tracebacks limited to 100 stack frames
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ tuples are not first-class values
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ type inference is too simple
  - http://yager.io/programming/go.html (Will Yager 2014)
+ types go after identifiers, not before
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ un-Googlable name
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - http://dtrace.org/blogs/ahl/2016/08/02/i-love-go-i-hate-go/ (Adam Leventhal 2016)
+ unexpected variable shadowing
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ unused imports are compile errors
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ unused variables are compile errors
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ unwieldy to code new collections
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bugfender.com/blog/go-pros-cons-using-go-programming-language/ (Aleix Ventayol 2017)
  - https://zapier.com/engineering/go-no-go/ (Jordan Sherer 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ weak typing
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
+ weird mascot (gopher)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ writing to a zero map panics
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)
+ zero values are a bad idea
  - https://memo.barrucadu.co.uk/three-months-of-go.html (Michael Walker 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://bluxte.net/musings/2018/04/10/go-good-bad-ugly/ (Sylvain Wallez 2018)

# Get involved

Feel free to add a PR with a new or old article you found on the internet. The structure is simple, just look at existing entries. Run `make` and check in the resulting `README.md` along with your updated `entries.json`.

```json
{
	"URL": "https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/",
	"Author": "Kaushal Subedi",
	"Year":  2015,
	"Complaints":[
		"no generics",
		"slow json parsing",
		"bad dependency management",
		"no subpackages"
	]
}
```

## TODO

+ merge complaints with the same ideas under the same names (make the complaints list smaller)
+ sort reverse index by the number of articles which have that complain (popular complains to the top)
