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
  - stuck in 70's
  - no OOP
  - too simple / lack of syntactic sugar
+ https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis, 2010)
  - no language interoperability (only C)
  - no versioning model
  - no OOP
  - has pointers
  - no semicolons at line endings
  - no `this`
  - no function/operator overloading
  - no exceptions
+ https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield, 2012)
  - error handling
  - stuck in 70's
  - `panic` instead of exceptions
+ http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King, 2012)
  - un-googlable name
  - stuck in 70's
  - no exceptions
  - is compiled
  - too young
+ http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish, 2012)
  - c-style
  - no unused imports
  - bad unicode support
  - no asserts
+ http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man, 2013)
  - no OOP
  - no asserts
  - has pointers
  - weird mascot (gopher)
+ http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey, 2013)
  - error handling
  - inconvenient `range`
  - no generics
+ http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi, 2013)
  - too opinionated
+ http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta, 2014)
  - project layout is bad
  - can't import packages relatively
  - different build approaches
  - no sugar for slices
  - lack of basic data structures
+ https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman, 2014)
  - no decent IDE
  - GOPATH is a mess
  - pointers are a mess
  - upper-case/lower-case scoping is bad
  - hidden types
  - immature GC
  - hard to extend
  - no exceptions
  - no generics
  - bad dependency management
+ http://yager.io/programming/go.html (Will Yager, 2014)
  - no generics
  - no function/operator overloading
  - nil as a failure marker
  - type inference is too simple
  - no immutables
  - no pattern matching
+ http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin, 2014)
  - confusing/stupid syntax
  - immature toolchain
  - immature GC
  - no generics
  - multiple return values have no type checking
  - cumbersome interface
  - goroutine is not original and revolutionary
  - defer is abused
  - stuck in Unix thinking
  - summary: not elegant as Python, not strong as Java
+ https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan, 2014)
  - no decent IDE
  - error handling
  - unexpected variable shadowing
  - no generics
  - not-so-obvious slices behaviour
+ http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski, 2014)
  - the worst compiler toolchain ever
  - psuedointellectual arrogance of Rob Pike and everything he stands for
  - confusing and undebuggable
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
  - no subpackages
+ http://www.evanmiller.org/four-days-of-go.html (Evan Miller, 2015)
  - no unused imports
  - too opinionated
  - poor std math lib
  - weird mascot (gopher)
+ http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel, 2015)
  - not-so-obvious slices behaviour
  - error handling
  - hard to test, hard to mock
+ http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov, 2015)
  - designed for stupid people
  - too verbose
  - no ternary operator
  - no macros or templates
  - can't change hash function in maps
  - unwieldy to code new collections
  - no function overloading
  - can't declare/validate implements interface
  - no virtual functions
  - case defined exports is bad
+ http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg, 2015)
  - un-googlable name
  - confusing/stupid syntax
  - polymorphism is broken
  - compilation rules are too confining
  - designed for stupid people
  - package mechanism is broken
+ http://nomad.uk.net/articles/why-gos-design-is-a-disservice-to-intelligent-programmers.html (Gary Willoughby, 2015)
  - too simple / lack of syntactic sugar
  - no generics
  - bad dependency management
  - stuck in 70's
+ http://blog.goodstuff.im/golang (David Pollak, 2015)
  - no immutables
  - too simple / lack of syntactic sugar
  - confusing/stupid syntax
  - too opinionated
  - error handling
  - stuck in 70's
  - no generics
  - upper-case/lower-case scoping is bad
  - no `map`/`reduce`/`filter`
+ https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn, 2016)
  - misleading marketing around garbage collector
  - garbage collector is nothing new; concurrent mark/sweep from the '70s
  - garbage collector requires implicit (hidden) tradeoffs
  - garbage collector optimized for pause times at the cost of other desirable gc features
+ https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot, 2016)
  - upper-case/lower-case scoping is bad
  - case defined exports is bad
  - no first-class support of interfaces
  - no exceptions
  - difficult to generate code automatically
  - no ternary operator
  - sort.Interface approach is clumsy
  - no versioning model
  - import-based vendoring is terrible
  - no generics
+ http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds, 2016)
  - channels are slow
  - channel API is inconsistent
+ https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus, 2016)
  - not stable
  - no `map`/`reduce`/`filter`
  - bad dependency management
+ http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett, 2017)
  - no non-nullable types
  - error handling
  - no generics
  - no tagged unions
  - no first-class tuple
  - concurrency and parallelism are mixed
+ https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson, 2017)
  - too opinionated
  - pointers are a mess
+ https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze, 2017)
  - no sum types
  - no algebraic data types
  - multiple return parameters are overrated
+ https://grimoire.ca/dev/go (Owen Jacobson, 2018)
  - hostile to developer ergonomics
  - too opinionated
  - package mechanism is broken
  - GOPATH is a mess
  - error handling


# Reverse complaints index

It's a reverse complaints index, generated by https://github.com/ksimka/go-is-not-good/blob/master/generator.go (thanks to @capoferro)

+ GOPATH is a mess
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ `new` and `make` instead of one
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ `panic` instead of exceptions
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
+ bad dependency management
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://nomad.uk.net/articles/why-gos-design-is-a-disservice-to-intelligent-programmers.html (Gary Willoughby 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ bad unicode support
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ c-style
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ can't change hash function in maps
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ can't declare/validate implements interface
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ can't import packages relatively
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ case defined exports is bad
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ channel API is inconsistent
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ channels are slow
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
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
+ cumbersome interface
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ defer is abused
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ designed for stupid people
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ different build approaches
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ difficult to generate code automatically
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ error handling
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ garbage collector is nothing new; concurrent mark/sweep from the '70s
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector optimized for pause times at the cost of other desirable gc features
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ garbage collector requires implicit (hidden) tradeoffs
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ go generate is a quirk
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ goroutine is not original and revolutionary
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ hard to extend
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ hard to test, hard to mock
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
+ has pointers
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ hidden types
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ hostile to developer ergonomics
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ immature GC
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ immature toolchain
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ import-based vendoring is terrible
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ inconvenient `range`
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ is compiled
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ lack of basic data structures
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
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
+ no OOP
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ no `map`/`reduce`/`filter`
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ no `this`
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no algebraic data types
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ no asserts
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ no constructors
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ no decent IDE
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
+ no exceptions
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no first-class support of interfaces
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no first-class tuple
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no function overloading
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no function/operator overloading
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no generics
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://nomad.uk.net/articles/why-gos-design-is-a-disservice-to-intelligent-programmers.html (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no immutables
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ no language interoperability (only C)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no macros or templates
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ no non-nullable types
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no pattern matching
  - http://yager.io/programming/go.html (Will Yager 2014)
+ no semicolons at line endings
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ no subpackages
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ no sugar for slices
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ no sum types
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ no tagged unions
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ no ternary operator
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no unused imports
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ no user-type iteration
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ no versioning model
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ no virtual functions
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
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
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
+ polymorphism is broken
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ poor design
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ poor std math lib
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ project layout is bad
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ psuedointellectual arrogance of Rob Pike and everything he stands for
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ questionable compiler rigidity
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ slow json parsing
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ sort.Interface approach is clumsy
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ stuck in 70's
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://nomad.uk.net/articles/why-gos-design-is-a-disservice-to-intelligent-programmers.html (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ stuck in Unix thinking
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ summary: not elegant as Python, not strong as Java
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ the worst compiler toolchain ever
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ too opinionated
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ too simple / lack of syntactic sugar
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://nomad.uk.net/articles/why-gos-design-is-a-disservice-to-intelligent-programmers.html (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ too verbose
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ too young
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ type inference is too simple
  - http://yager.io/programming/go.html (Will Yager 2014)
+ un-googlable name
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ unexpected variable shadowing
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ unwieldy to code new collections
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ upper-case/lower-case scoping is bad
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
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
