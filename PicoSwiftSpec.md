Picoswift Spec.md
# Picoswift Spec

quick and probably not self-consistent. designed to make the compiler really simple.

##  Comments

comments are // stuff... <eol>

## Variables

### Names for Variables

[a-zA-Z][_a-zA-Z0-9] is the regex for valid variable names.

### Names for Classes

[A-Z][_a-zA-Z0-9] is the regex for valid variable names.

### let versus var

let implies that this will never change once initialized (immutable)
var implies it might change and be mutable

### Integers Class

let i [= Int(k)]<eol>
var j [= Int(k)]<eol>

int literals 42 (decimal) and 0x12FF (hexadecimal)

this is a std 32-bit Integer.

### Boolean Class

let isAlive [= Bool() or Bool(True or False)]<eol>
var canDrive [= Bool([initialer])]<eol>

boolean variable, boolean literals True, False.

### String Class

let name [= String("a string")]<eol>
var address [= String("another string")]<eol>

String is immuable.
literals are a run of Runes delimited by double-quotes

- append(String s) String
- runeAt(Int idx) Rune
- length() Int

### Rune (Character) Class

let weird = Rune('x') // var also works

literals are either UTF-8 or hex literals

### Void Class

Void is the null type, like nil is the null reference

### Array Class

let scores = Array.Int(42)

while 'scores' cannot change, the contents of 'scores' can.

idx is of type Int

scores.set(ids, value)
var x = score.get(idx)
// maybe someday you could do
// scores[idx] = IntLiteral or IntVar //set
// var tmp = scores[idx] //get

### Map Class

let symbols = Map[Key-Type]Value-Type

- empty() Void
- isEmpty() Bool
- size() Int

- set(Key-Type k, Key-Type v) Void
- get(Key-Type k) Key-Type
- contains(Key-Type k) Bool
- del(Key-Type k) Bool

## Expressions

standard boolean and arithmetic Expressions

AND, OR, NOT - boolean conjunctions(?)
==, !=, >=, >, <, <=  normal comparisons

( ) ^ * / + - %    standard operators

## Assignments

assignment statement

x = x + 1<eol>

OR

x += 1<eol>

## Conditional statement

if condition {

}

if condition {

} else {

}

## Loop statement

while condition {

}

## Loop breakers

break<eol>
continue<eol>
return<eol>

## Function 

func name(type param,) type-return {
    local Variables

    statements
    return
}

## Classes

class Name {
    instance Variables

    Name() {
        // null constructor
    }
    Name(type param ,) {

    }

    method name(type param,) type-return {

    }
}