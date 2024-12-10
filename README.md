# Advent Of Code 2024

Giving Go another try, haven't done anything with since AoC 2023.

Code is up for grabs, public domain.

Every day will be a directory, with 1 `main.go`.

Puzzles on https://adventofcode.com/2024/


## Day 6, part 2
Takes about 6s to run... Hugely inefficient

## Day 7
Didn't want to go down the `eval` rabbit hole, but apparently it is fairly easy to evaluate a mathematical expression in Go:

```Go
fs := token.NewFileSet()
tv, _ := types.Eval(fs, nil, token.NoPos, "(1+4) * 5")
fmt.Print(tv.Value.String())
```

## Day 9
I most definitely got bitten by the "modify array up the wazoo vs. use some well chosen counters because I followed the puzzle question too literally".

Once I saw the map { id: {pos, size}, ...} I got the strong feeling that we could build up the necessary data much easier...


<!-- vim: set spell: -->
