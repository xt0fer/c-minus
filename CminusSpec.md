# Cminus Spec

quick and probably not self-consistent. designed to make the compiler really simple.

##  Comments

comments are // stuff... <eos>

## Variables

### Names for Variables

[a-zA-Z][_a-zA-Z0-9] is the regex for valid variable anad function names.

### var

var implies it might change and be mutable. vars must be within a function (or main) scope.

### simple statement ends (<eos>)

semi-colon will be the end of statement char

### Integers

var j [= int(k)]<eos>

int literals 42 (decimal) and 0x12FF (hexadecimal)

this is a std 32-bit Integer.

### Boolean Class

var canDrive [= boolean([initialer])]<eos>

boolean variable, boolean literals True, False.

### String 

var address [= string("another string")]<eos>

String is immuable.
literals are a run of Runes delimited by double-quotes

- append(string s) string
- runeAt(int idx) rune
- length() int

### Rune (Character) 

Rune is a 32-bit char
var weird = rune('x') // var also works

literals are either UTF-8 or hex literals

### Void

Void is the null type, like nil is the null reference

### Array

var scores = array.int(42)
var flags = array.bool(8)

while 'scores' cannot change, the contents of 'scores' can.

idx is of type int

scores.set(ids, value)

x = score.get(idx)

## Expressions

standard boolean and arithmetic Expressions

&&, ||, ~ - boolean conjunctions(?)
==, !=, >=, >, <, <=  normal comparisons

( ) ^ * / + - %    standard operators

## Assignments

assignment statement

x = x + 1<eos>

## Conditional statement

if condition {
    [statements]
}

if condition {
    [statements]
} else {
    [statements]
}

## Loop statement

while condition {
    [statements]
}

## Loop breakers

break<eos>
continue<eos>
return [ret-value]<eos>

## Function 

function name(type param,) type-return {
    [local Variables]

    [statements]
    return [ret-value?]
}

## Main 

main() is the entry point for cminus programs. 

example:

main() {
    print("Hello World.")
}

## Print & Input

print() prints a string on standard output.

input() requests a string (line) from standard input.

error() prints a string (line) on standard error.

That's All, Folks!

## Maybe someday.

<!-- ### Map type

let symbols = Map[Key-Type]Value-Type

- empty() Void
- isEmpty() Bool
- size() Int

- set(Key-Type k, Key-Type v) Void
- get(Key-Type k) Key-Type
- contains(Key-Type k) Bool
- del(Key-Type k) Bool -->
