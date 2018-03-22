# What's this

This repository is a list of articles that complain about **golang**'s imperfection.

## Motivation

Seems like complaining about **go**'s flaws is becoming a trend. Any newbie must have a chance to read all the **go**-is-bad arguments before they go too far. So here it is.

## What it's for

This repo is not aimed to offend or insult someone (at least not more than each author does it in its article), especially **golang** itself, its authors and the community. It is for educational purpose only. Any contributor can have an absolutely different point of view.

I don't think anyone would deny that **go** has weaknesses: it certainly has. But how do you know, *is it really a language design flaw or is it just you, doing something completely wrong*? This list here to help you quickly answer the question.

## How to use it

You're writing some code. And suddenly you understand you need something that language can't give you. You go here and check if you're the one with that issue or not. If it's a common issue, it'll be here. Then you decide what to do: choose another tool for your task or go find a better solution or a workaround.

# The List

+ https://cowlark.com/2009-11-15-go/ (David Given, 2009)
  - c-style
  - poor design
  - no constructors
  - no user-type iteration
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
+ http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi, 2013)
  - too opinionated
  - unused imports are compile errors
  - unused variables are compile errors
+ http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man, 2013)
  - no OOP
  - no assertions
  - has pointers
  - solves nobody's problem
  - no inheritance
  - no function/operator overloading
  - designers did not consult the literature
  - weird mascot (gopher)
+ http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey, 2013)
  - error handling is mandatory
  - no exceptions
  - inconvenient `range`
  - can't range over user-defined types
  - tabwidth too wide
  - gofmt's style is no one's favourite
  - no generics
+ http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski, 2014)
  - the worst compiler toolchain ever
  - pseudointellectual arrogance of Rob Pike and everything he stands for
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
  - garbage collector requires implicit (hidden) tradeoffs
  - not real-time oriented
  - no unsafe code isolation
  - stuck in '70s
  - no pattern matching
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
  - bad dependency management
+ https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan, 2014)
  - no decent IDE
  - error handling is mandatory
  - unexpected variable shadowing
  - no generics
  - no `map`/`reduce`/`filter`
  - no immutables
  - not-so-obvious slices behaviour
+ http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin, 2014)
  - confusing/stupid syntax
  - immature toolchain
  - immature GC
  - no generics
  - multiple return values have no type checking
  - cumbersome interface
  - goroutine is not original and revolutionary
  - `defer` is abused
  - stuck in Unix thinking
  - types go after identifiers, not before
  - no decent IDE
  - bad dependency management
  - unwieldy to code new collections
  - designers did not consult the literature
  - error handling is mandatory
  - summary: not elegant as Python, not strong as Java
+ http://byrd.im/go-is-poor/ (Ian Byrd, 2015)
  - not-so-obvious slices behaviour
  - nil interfaces are not entirely nil
  - unexpected variable shadowing
  - no first-class support of interfaces
  - questionable compiler rigidity
  - go generate is a quirk
+ https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi, 2015)
  - no generics
  - slow json parsing
  - bad dependency management
  - too verbose
  - GOPATH is a mess
  - no subpackages
+ http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel, 2015)
  - not-so-obvious slices behaviour
  - error handling is mandatory
  - hard to test, hard to mock
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
+ http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg, 2015)
  - un-Googlable name
  - confusing/stupid syntax
  - polymorphism is broken
  - compilation rules are too confining
  - designed for stupid people
  - can't declare/validate implements interface
  - package mechanism is broken
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
+ https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn, 2016)
  - misleading marketing around garbage collector
  - garbage collector is nothing new; concurrent mark/sweep from the '70s
  - garbage collector requires implicit (hidden) tradeoffs
  - garbage collector optimized for pause times at the cost of other desirable gc features
+ http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds, 2016)
  - channels are slow
  - channel API is inconsistent
+ https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus, 2016)
  - not stable
  - no `map`/`reduce`/`filter`
  - bad dependency management
+ https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot, 2016)
  - case-defined scoping is bad
  - no first-class support of interfaces
  - no exceptions
  - difficult to generate code automatically
  - no ternary operator
  - sort.Interface approach is clumsy
  - no versioning model
  - import-based vendoring is terrible
  - no generics
+ https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze, 2017)
  - no sum types
  - no algebraic data types
  - multiple return parameters are overrated
+ https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson, 2017)
  - too opinionated
  - pointers are a mess
+ http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett, 2017)
  - no non-nullable types
  - error handling is mandatory
  - no generics
  - no tagged unions
  - no first-class tuple
  - concurrency and parallelism are mixed
+ https://grimoire.ca/dev/go (Owen Jacobson, 2018)
  - hostile to developer ergonomics
  - too opinionated
  - package mechanism is broken
  - GOPATH is a mess
  - error handling is mandatory


# Reverse complaints index

It's a reverse complaints index, generated by https://github.com/ksimka/go-is-not-good/blob/master/generator.go (thanks to @capoferro)

+ 'rune' is a weird name
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ := is weird
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ Fisher-Price assembly language
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ GOPATH is a mess
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ `defer` is abused
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
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
+ bad dependency management
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ c-style
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ can't change hash function in maps
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ can't declare/validate implements interface
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ can't name nested functions
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ can't range over user-defined types
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://yager.io/programming/go.html (Will Yager 2014)
+ case-defined scoping is bad
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ channel API is inconsistent
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channels are overrated
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ channels are slow
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ closures capture variables by reference, not by value
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ compilation rules are too confining
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ concurrency and parallelism are mixed
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ confusing and undebuggable
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ confusing/stupid syntax
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ constants can only be numbers or strings
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ cumbersome interface
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ designed for stupid people
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ designers did not consult the literature
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ difficult to generate code automatically
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ duck typing of local variables
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ dynamic memory allocation
  - http://yager.io/programming/go.html (Will Yager 2014)
+ error handling is mandatory
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ error handling is not mandatory
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ garbage collector is nothing new; concurrent mark/sweep from the '70s
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector optimized for pause times at the cost of other desirable gc features
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector requires implicit (hidden) tradeoffs
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ go generate is a quirk
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ gofmt's style is no one's favourite
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ goroutine is not original and revolutionary
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
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
+ import-based vendoring is terrible
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ inconvenient `range`
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ internationalization is limited to Unicode support
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ is compiled
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ is garbage-collected
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
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
+ no OOP
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ no `map`/`reduce`/`filter`
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ no `this`
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no algebraic data types
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ no assertions
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ no constructors
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ no decent IDE
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ no exceptions
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no first-class support of interfaces
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no first-class tuple
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no function/operator overloading
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no generics
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no immutables
  - http://yager.io/programming/go.html (Will Yager 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ no incremental or parallel compilation
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ no inheritance
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ no language interoperability (only C)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no macros or templates
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no non-nullable types
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no pattern matching
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no read-eval-print loop
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ no semicolons at line endings
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no subpackages
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ no sum types
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ no tagged unions
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no ternary operator
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no unsafe code isolation
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no user-type iteration
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ no versioning model
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no virtual functions
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no while loops
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
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
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
+ package mechanism is broken
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
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
+ pseudointellectual arrogance of Rob Pike and everything he stands for
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ questionable compiler rigidity
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ return types after function name, not before
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ semicolons are injected by the lexer
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ single-pass compiler
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ slow json parsing
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ solves nobody's problem
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ sort.Interface approach is clumsy
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ string manipulation is a library package
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ stuck in '70s
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ stuck in Plan9 thinking
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ stuck in Unix thinking
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ summary: not elegant as Python, not strong as Java
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ tabwidth too wide
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ technical hubris
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ the worst compiler toolchain ever
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ too opinionated
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ too simple / lack of syntactic sugar
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ too verbose
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ too young
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ type inference is too simple
  - http://yager.io/programming/go.html (Will Yager 2014)
+ types go after identifiers, not before
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ un-Googlable name
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ unexpected variable shadowing
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
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
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ weird mascot (gopher)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)

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
