# Cminus Code generation patterns

_this stuff needs a v1.3 of ZipRISC1_

## Register Usage

x0 zero
x1  - RA  - return address Caller saved
x2, x3, x4 - function call arguments
x2 - function return agument

x7, x8, x9, xA - temp regs (used inside of functions)
rd, rs, rt, ru - temp regs alt names 

xF - PC, xE - IR, xD - SP, xC - FP

/*
 * Just for reference
 * INSTRUCTIONs
- ADD rd, rs, rt | 1dst | rd <- rs + rt
- SUB rd, rs, rt | 2dst | rd <- rs - rt
- SUBI rd, rs, k | 3dsk | rd ← rs - k
- LSH rd, rs, k | 4dsk | rd <- rs / k ??
- RSH rd, rs, k | 5dsk | rd <- rs * k ??
- BRZ rd, aa | 6daa | branch to aa on rd == 0
- BGT rd, aa | 7daa | branch to aa on rd > 0
- LD rd, aa | 8daa | load rd with value of memory loc aa
- ST rs, aa | 9saa | store rd value to memory loc aa
- HLT | 0000 | halt.
- HCF | 0FFF | halt and catch fire.
- IN rd | Ad00 | read in a number to rd
- OUT rd | Bd00 | output a number from rd
- ADDI rd, rs, k │ Cdsk │ rd ← rs + k
- DUMP | F000 | print out registers, machine state and memory
- LDR rd, rs | 11dr0 | load rd with contents of address held by rs
- STR rd, rs | 12dr0 | store rd into the address held by rs

* PSEUDOs
- MOV rd, rs │ ADD rd, rs, x0 │ rd ← rs
- CLR rd │ ADD rd, x0, x0 │ rd ← 0
- DEC rd │ SUBI rd, rd, 1 │ rd ← rd - 1
- INCR rd |ADD rd, rd, 1  | rd <- rd + 1
- BRA aa │ BRZ x0, aa │ next instruction to read is at aa

* the calling convention for subroutines/functions.
- CALL aa | ADDI x1 xPC 1; BRA aa | ra <- PC + 1, jump to aa
- RET | ADD xPC x1 x0 | pc <- ra (ra is "return address")
- PUSH rd | DECR SP; STR rd SP | sp <- sp - 1, store rd to contents of SP
- POP rd | LDR rd, SP; INCR SP | load rd with contents SP, sp <- sp + 1

*/

### Function call setup
- setup regs to hold function args
- call function

- inside function
  - make room on stack for locals
  - use arguments from regs
  - use temp regs as needed
  - load return reg with result
  - drop locals from stack
  - return

```

### Integer variable declaration

```
var name = int(0);
```

```
name:
.WD 0

// fs: "\n%s:\n.WD %d\n", name, constantInt
```

### Boolean variable declaration

```
var name = boolean(true);
var name2 = boolean(false);
```

```
name:
.WD 1
name2:
.WD 0
// fs: "\n%s:\n.WD %d\n", name, constantInt
```

### String variable declaration

```
var name = string("hello");
```

```
name:
.DW &(name+2)
.DW len (5)
.DW 'h'
.DW 'e'
.DW 'l'
.DW 'l'
.DW 'o'
// fs: "\n%s:\n.WD %d\n", name, constantInt
```

### Rune variable declaration

```
var name = rune('x');
```

```
name:
.WD utf-8('x')
// fs: "\n%s:\n.WD %d\n", name, utf-8('x')
```

### Arrays - not yet

### Expressions

### Assignment Statement

```
x = y * z ;
```

```
LD t0 y
LD t1 z
ADD t0 t0 t1
ST t0 x
```



