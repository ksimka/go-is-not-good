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
+ http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby, 2015)
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

+ <a id='gopath-is-a-mess'>GOPATH is a mess</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ <a id='new-and-make-instead-of-one'>`new` and `make` instead of one</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ <a id='panic-instead-of-exceptions'>`panic` instead of exceptions</a>
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
+ <a id='bad-dependency-management'>bad dependency management</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ <a id='bad-unicode-support'>bad unicode support</a>
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ <a id='c-style'>c-style</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
+ <a id='cant-change-hash-function-in-maps'>can't change hash function in maps</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='cant-declare-validate-implements-interface'>can't declare/validate implements interface</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='cant-import-packages-relatively'>can't import packages relatively</a>
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ <a id='case-defined-exports-is-bad'>case defined exports is bad</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='channel-api-is-inconsistent'>channel API is inconsistent</a>
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ <a id='channels-are-slow'>channels are slow</a>
  - http://www.jtolds.com/writing/2016/03/go-channels-are-bad-and-you-should-feel-bad/ (JT Olds 2016)
+ <a id='compilation-rules-are-too-confining'>compilation rules are too confining</a>
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ <a id='concurrency-and-parallelism-are-mixed'>concurrency and parallelism are mixed</a>
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ <a id='confusing-and-undebuggable'>confusing and undebuggable</a>
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ <a id='confusing-stupid-syntax'>confusing/stupid syntax</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ <a id='cumbersome-interface'>cumbersome interface</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='defer-is-abused'>defer is abused</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='designed-for-stupid-people'>designed for stupid people</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ <a id='different-build-approaches'>different build approaches</a>
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ <a id='difficult-to-generate-code-automatically'>difficult to generate code automatically</a>
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='error-handling'>error handling</a>
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ <a id='garbage-collector-is-nothing-new%3B-concurrent-mark-sweep-from-the-70s'>garbage collector is nothing new; concurrent mark/sweep from the '70s</a>
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ <a id='garbage-collector-optimized-for-pause-times-at-the-cost-of-other-desirable-gc-features'>garbage collector optimized for pause times at the cost of other desirable gc features</a>
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ <a id='garbage-collector-requires-implicit-hidden-tradeoffs'>garbage collector requires implicit (hidden) tradeoffs</a>
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ <a id='go-generate-is-a-quirk'>go generate is a quirk</a>
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ <a id='goroutine-is-not-original-and-revolutionary'>goroutine is not original and revolutionary</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='hard-to-extend'>hard to extend</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ <a id='hard-to-test%2C-hard-to-mock'>hard to test, hard to mock</a>
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
+ <a id='has-pointers'>has pointers</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ <a id='hidden-types'>hidden types</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
+ <a id='hostile-to-developer-ergonomics'>hostile to developer ergonomics</a>
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ <a id='immature-gc'>immature GC</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='immature-toolchain'>immature toolchain</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='import-based-vendoring-is-terrible'>import-based vendoring is terrible</a>
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='inconvenient-range'>inconvenient `range`</a>
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
+ <a id='is-compiled'>is compiled</a>
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ <a id='lack-of-basic-data-structures'>lack of basic data structures</a>
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ <a id='misleading-marketing-around-garbage-collector'>misleading marketing around garbage collector</a>
  - https://blog.plan99.net/modern-garbage-collection-911ef4f8bd8e#.62yek82xg (Mike Hearn 2016)
+ <a id='multiple-return-parameters-are-overrated'>multiple return parameters are overrated</a>
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ <a id='multiple-return-values-have-no-type-checking'>multiple return values have no type checking</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='nil-as-a-failure-marker'>nil as a failure marker</a>
  - http://yager.io/programming/go.html (Will Yager 2014)
+ <a id='nil-interfaces-are-not-entirely-nil'>nil interfaces are not entirely nil</a>
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ <a id='no-oop'>no OOP</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ <a id='no-map-reduce-filter'>no `map`/`reduce`/`filter`</a>
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ <a id='no-this'>no `this`</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ <a id='no-algebraic-data-types'>no algebraic data types</a>
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ <a id='no-asserts'>no asserts</a>
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://magicmakerman.blogspot.ru/2013/07/why-googles-go-programming-language.html (Magic Maker Man 2013)
+ <a id='no-constructors'>no constructors</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ <a id='no-decent-ide'>no decent IDE</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
+ <a id='no-exceptions'>no exceptions</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='no-first-class-support-of-interfaces'>no first-class support of interfaces</a>
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='no-first-class-tuple'>no first-class tuple</a>
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ <a id='no-function-overloading'>no function overloading</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='no-function-operator-overloading'>no function/operator overloading</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - http://yager.io/programming/go.html (Will Yager 2014)
+ <a id='no-generics'>no generics</a>
  - http://how-bazaar.blogspot.ru/2013/04/the-go-language-my-thoughts.html (Tim Penhey 2013)
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
  - http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ <a id='no-immutables'>no immutables</a>
  - http://yager.io/programming/go.html (Will Yager 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ <a id='no-language-interoperability-only-c'>no language interoperability (only C)</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ <a id='no-macros-or-templates'>no macros or templates</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='no-non-nullable-types'>no non-nullable types</a>
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ <a id='no-pattern-matching'>no pattern matching</a>
  - http://yager.io/programming/go.html (Will Yager 2014)
+ <a id='no-semicolons-at-line-endings'>no semicolons at line endings</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
+ <a id='no-subpackages'>no subpackages</a>
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ <a id='no-sugar-for-slices'>no sugar for slices</a>
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ <a id='no-sum-types'>no sum types</a>
  - https://awalterschulze.github.io/blog/post/sum-types-over-multiple-returns/ (Walter Schulze 2017)
+ <a id='no-tagged-unions'>no tagged unions</a>
  - http://sitr.us/2017/02/21/changes-i-would-make-to-go.html (Jesse Hallett 2017)
+ <a id='no-ternary-operator'>no ternary operator</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='no-unused-imports'>no unused imports</a>
  - http://ridiculousfish.com/blog/posts/go_bloviations.html (ridiculousfish 2012)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ <a id='no-user-type-iteration'>no user-type iteration</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ <a id='no-versioning-model'>no versioning model</a>
  - https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis 2010)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='no-virtual-functions'>no virtual functions</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='not-stable'>not stable</a>
  - https://medium.com/@rgausnet/3-reasons-why-go-isnt-the-perfect-language-yet-25e0da5ec04c (Ryan Gaus 2016)
+ <a id='not-so-obvious-slices-behaviour'>not-so-obvious slices behaviour</a>
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
  - http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel 2015)
+ <a id='package-mechanism-is-broken'>package mechanism is broken</a>
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ <a id='pointers-are-a-mess'>pointers are a mess</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
+ <a id='polymorphism-is-broken'>polymorphism is broken</a>
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ <a id='poor-design'>poor design</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
+ <a id='poor-std-math-lib'>poor std math lib</a>
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
+ <a id='project-layout-is-bad'>project layout is bad</a>
  - http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta 2014)
+ <a id='psuedointellectual-arrogance-of-rob-pike-and-everything-he-stands-for'>psuedointellectual arrogance of Rob Pike and everything he stands for</a>
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ <a id='questionable-compiler-rigidity'>questionable compiler rigidity</a>
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ <a id='slow-json-parsing'>slow json parsing</a>
  - https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi 2015)
+ <a id='sort.interface-approach-is-clumsy'>sort.Interface approach is clumsy</a>
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='stuck-in-70s'>stuck in 70's</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - https://uberpython.wordpress.com/2012/09/23/why-im-not-leaving-python-for-go/ (Yuval Greenfield 2012)
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ <a id='stuck-in-unix-thinking'>stuck in Unix thinking</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='summary%3A-not-elegant-as-python%2C-not-strong-as-java'>summary: not elegant as Python, not strong as Java</a>
  - http://www.yinwang.org/blog-cn/2014/04/18/golang (Wang Yin 2014)
+ <a id='the-worst-compiler-toolchain-ever'>the worst compiler toolchain ever</a>
  - http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski 2014)
+ <a id='too-opinionated'>too opinionated</a>
  - http://corte.si/posts/code/go/go-rant.html (Aldo Cortesi 2013)
  - http://www.evanmiller.org/four-days-of-go.html (Evan Miller 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://maryvilledevcenter.com/golang-thinks-you-are-a-bad-programmer/ (Dane Johnson 2017)
  - https://grimoire.ca/dev/go (Owen Jacobson 2018)
+ <a id='too-simple---lack-of-syntactic-sugar'>too simple / lack of syntactic sugar</a>
  - https://cowlark.com/2009-11-15-go/ (David Given 2009)
  - http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby 2015)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
+ <a id='too-verbose'>too verbose</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='too-young'>too young</a>
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
+ <a id='type-inference-is-too-simple'>type inference is too simple</a>
  - http://yager.io/programming/go.html (Will Yager 2014)
+ <a id='un-googlable-name'>un-googlable name</a>
  - http://www.darkcoding.net/software/go-lang-after-four-months/ (Graham King 2012)
  - http://valuedrivenit.blogspot.ru/2015/12/to-go-language-is-mess.html (Cliff Berg 2015)
+ <a id='unexpected-variable-shadowing'>unexpected variable shadowing</a>
  - https://www.upguard.com/blog/our-experience-with-golang (Mark Sheahan 2014)
  - http://byrd.im/go-is-poor/ (Ian Byrd 2015)
+ <a id='unwieldy-to-code-new-collections'>unwieldy to code new collections</a>
  - http://tmikov.blogspot.com/2015/02/you-dont-like-googles-go-because-you.html (Tzvetan Mikov 2015)
+ <a id='upper-case-lower-case-scoping-is-bad'>upper-case/lower-case scoping is bad</a>
  - https://rule1.quora.com/Golang-Not-yet (Jordan Zimmerman 2014)
  - http://blog.goodstuff.im/golang (David Pollak 2015)
  - https://www.teamten.com/lawrence/writings/why-i-dont-like-go.html (Lawrence Kesteloot 2016)
+ <a id='weird-mascot-gopher'>weird mascot (gopher)</a>
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
