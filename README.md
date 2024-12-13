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

## Day 11
Part 2: aha moment of caching results (think there was a similar thing in previous years)

TODO: Does Go support memoization (cf. Python's @functools.cache decorator)?

## Day 13
Had some initial pointers, like Dijkstra (max 2 neighbours; both buttons always move the gripper up-right, there wasn't a minus sign in the input and no buttons that would only move on 1 axis).  Might work for <=100 presses, but would probably not suffice for part2.

It probably could be brute forced, but that would _definitely_ not be adequate for part 2.

Maybe we could do something with vectors, but knowledge--

I also saw equations with 2 unknowns (ie. "Stelsel van vergelijkingen met 2 onbekenden" in Dutch), but couldn't quite work it out all the way.

It occurred to me that it wouldn't matter in which order you press the buttons; in the end, you'd to have moved a certain amount on the X and another amount on the Y axis.
So, "winning the price" could be expressed as `X*a + Y*b = p` where `X` is the number of presses on `a` and `Y` the number of presses on `b`.  Because pressing a button moves a certain amount on the x and y axis, this can be rewritten as

```math
X*ax + Y*bx = px
X*ay + Y*by = px
```

... Giving the system of equations with 2 unknowns.
Solving this was a frustrating time; I started out with the substitution method, but that yielded horrible formulas making it extremely error prone to write down and then type out correctly...
In the end I looked up a video about solving this and came across the "combination" method, which was much less error prone and resulted in cleaner formulas.

Luckily, all this made part2 a breeze.

<!-- vim: set spell: -->
