# Advent of Code 2017 to practice Go

## Flags:

```sh
-Day int
    Number of the day(1 to 25) (default 1)
-FilePath string
    Input file Path (default "input.txt")
-Part int
    Assignment part (1 or 2) (default 1)
```

## Example usage

Without debug:
`./aoc -FilePath=$(pwd)/input.txt -Part 1 -Day 1`

With debug:
`DEBUG=1 ./aoc -FilePath=$(pwd)/input.txt -Part 1 -Day 1`
