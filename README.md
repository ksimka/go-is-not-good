# What's this

This repository is a list of articles that complain about **golang**'s imperfection.

## Motivation

Seems like complaining about **go**'s flaws is becoming a trend. Any newbie must have a chance to read all the **go**-is-bad arguments before they go too far. So here it is.

## What it's for

This repo is not aimed to offend or insult someone (at least not more than each author does it in its article), especially **golang** itself, its authors and the community. It is for educational purpose only. Any contributor can have an absolutely different point of view.

I don't think anyone would deny that **go** has weaknesses: it certainly has. But how do you know, *is it really a language design flaw or is it just you, doing something completely wrong*? This list here to help you quickly answer the question.

## How to use it

You're writing some code. And suddenly you understand you need something that language can't give you. You go here and check if you're the one with that issue or not. If it's a common issue, it'll be here. Then you decide what to do: choose another tool for your task or go find a better solution or a workaround.

# The list

+ https://dzone.com/articles/i-don%E2%80%99t-much-get-go (Jon Davis, 2010)
  - no language interoperability (only C)
  - no versioning model
  - no OOP
  - has pointers
  - no semicolons at line endings
  - no `this`
  - no overloading
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
+ http://jozefg.bitbucket.org/posts/2013-08-23-leaving-go.html (Danny Gratzer, 2013)
  - no generics
  - is not extensible
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
+ http://yager.io/programming/go.html (Will Yager, 2014)
  - no generics
  - no operator overloading
  - nil as a failure marker
  - type inference is too simple
  - no immutables
  - no pattern matching
+ https://gist.github.com/kachayev/21e7fe149bc5ae0bd878 (Alexey Kachayev, 2014)
  - no generics
  - not everything is an expression
+ http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta, 2014)
  - project layout is bad
  - can't import packages relatively
  - different build approaches
  - no sugar for slices
  - lack of basic data structures
+ http://dtrace.org/blogs/wesolows/2014/12/29/golang-is-trash/ (Keith Wesolowski, 2014)
  - the worst compiler toolchain ever
  - psuedointellectual arrogance of Rob Pike and everything he stands for
  - confusing and undebuggable
+ http://www.evanmiller.org/four-days-of-go.html (Evan Miller, 2015)
  - no unused imports
  - too opinionated
  - poor std math lib
  - weird mascot (gopher)
+ http://blog.goodstuff.im/golang (David Pollak, 2015)
  - no immutables
  - too simple
  - stupid syntax
  - too opinionated
  - error handling
  - stuck in 70's
  - no generics
  - method access modifiers (upper/lower case for public/private)
  - no `map` and `filter`
+ http://nomad.so/2015/03/why-gos-design-is-a-disservice-to-intelligent-programmers/ (Gary Willoughby, 2015)
  - too simple
  - no generics
  - bad dependency management
  - stuck in 70's
+ http://spaces-vs-tabs.com/4-weeks-of-golang-the-good-the-bad-and-the-ugly/ (Freddy Rangel, 2015)
  - not-so-obvious slices behaviour
  - error handling
  - hard to test, hard to mock
+ http://byrd.im/go-is-poor (Ian Byrd, 2015)
  - slice manipulations are broken
  - nil interfaces are not entirely nil
  - funny variable shadowing
  - no first-class support of interfaces
  - questionable compiler rigidity
  - go generate is a quirk
+ https://kaushalsubedi.com/blog/2015/11/10/golang-sucks-heres-why/ (Kaushal Subedi, 2015)
  - no generics
  - slow json parsing
  - bad dependency management
  - no subpackages

# Get involved

Feel free to add a PR with a new or old article you found on the internet. The structure is simple, just look at existing entries. Sort items chronologically. 

+ link (Author, year of publication)
 - common complaint 1
 - common complaint 2

## todo

- make reverse index: each common complaint to list of links to articles
