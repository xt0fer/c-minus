# Cminus Spec

quick and probably not self-consistent. designed to make the compiler really simple.

##  Comments

comments are // stuff... <eol>

## Variables

### Names for Variables

[a-zA-Z][_a-zA-Z0-9] is the regex for valid variable anad function names.

### var

var implies it might change and be mutable

### Integers

var j [= Int(k)]<eol>

int literals 42 (decimal) and 0x12FF (hexadecimal)

this is a std 32-bit Integer.

### Boolean Class

var canDrive [= Bool([initialer])]<eol>

boolean variable, boolean literals True, False.

### String 

var address [= String("another string")]<eol>

String is immuable.
literals are a run of Runes delimited by double-quotes

- append(String s) String
- runeAt(Int idx) Rune
- length() Int

### Rune (Character) 

Rune is a 32-bit char
let weird = Rune('x') // var also works

literals are either UTF-8 or hex literals

### Void

Void is the null type, like nil is the null reference

### Array

let scores = Array.Int(42)
let flags = Array.Bool(8)

while 'scores' cannot change, the contents of 'scores' can.

idx is of type Int

scores.set(ids, value)

x = score.get(idx)

## Expressions

standard boolean and arithmetic Expressions

AND, OR, NOT - boolean conjunctions(?)
==, !=, >=, >, <, <=  normal comparisons

( ) ^ * / + - %    standard operators

## Assignments

assignment statement

x = x + 1<eol>

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

break<eol>
continue<eol>
return [ret-value]<eol>

## Function 

func name(type param,) type-return {
    [local Variables]

    [statements]
    return [ret-value?]
}

<!-- ### Map Class

let symbols = Map[Key-Type]Value-Type

- empty() Void
- isEmpty() Bool
- size() Int

- set(Key-Type k, Key-Type v) Void
- get(Key-Type k) Key-Type
- contains(Key-Type k) Bool
- del(Key-Type k) Bool -->
