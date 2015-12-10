# What's this

This repository is a list of articles that complain about **golang**'s imperfection.

## Motivation

Seems like complaining about **go**'s flaws is becoming a trend. Any newbie must have a chance to read all the **go**-is-bad arguments before he goes too far. So here it is.

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
  - mascot sucks (gopher)
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
+ http://blog.mattbasta.com/things_that_make_me_sad_in_go.html (Matt Basta, 2014)
  - project layout is bad
  - can't import packages relatively
  - different build approaches
  - no sugar for slices
  - lack of basic data structures
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
