# Yewolf's Advent Of Code 2023

![GitHub last commit (main)](https://img.shields.io/github/last-commit/yyewolf/goaoc-2023/main)

[Advent Of Code](https://adventofcode.com/) is an advent calendar filled with challenges that are gradually harder.

Personal rules (2023) :
 - Use Golang for the challenges.
 - Use the least import possible (not counting `runtime`/`debug`).
 - Do everything possible to run in under 100 000 ns. (yes, this will lead to some spaghetti code)

## 2023 Results

All benchmarks are run with the same command: `go test -bench=. -benchtime=10s -benchmem`

| Benchmark                       | Results                                   | Satisfied ?                              | 
|---------------------------------|-------------------------------------------|------------------------------------------|
| [Day 01 - P1](./day01/)          | `14046 ns/op 4144 B/op 2 allocs/op`      | ✅ |
| [Day 01 - P2](./day01/)          | `49475 ns/op 4144 B/op 2 allocs/op`      | ✅ |
||| |
| [Day 02 - P1](./day02/)          | `12011 ns/op 4144 B/op 2 allocs/op`      | ✅ |
| [Day 02 - P2](./day02/)          | `18509 ns/op 4144 B/op 2 allocs/op`      | ✅ |
||| |
| [Day 03 - P1](./day03/)          | `32054 ns/op 4144 B/op 2 allocs/op`      | ✅ |
| [Day 03 - P2](./day03/)          | `42310 ns/op 37135 B/op 2 allocs/op`     | ✅ |
||| |
| [Day 04 - Parsing](./day04/)     | `60.41 ns/op 0 B/op 0 allocs/op`         | ✅ |
| [Day 04 - P1](./day04/)          | `14435 ns/op 4144 B/op 2 allocs/op`      | ✅ |
| [Day 04 - P2](./day04/)          | `12767 ns/op 4144 B/op 2 allocs/op`      | ✅ |
||| |
| [Day 05 - P1](./day05/)          | `7220 ns/op 3520 B/op 46 allocs/op`      | ✅ |
| [Day 05 - P2](./day05/)          | `19216 ns/op 23928 B/op 65 allocs/op`    | ✅ |
||| |
| [Day 06 - P1](./day06/)          | `173.0 ns/op 112 B/op 3 allocs/op`       | ✅ |
| [Day 06 - P2](./day06/)          |  `90.19 ns/op 0 B/op 0 allocs/op`        | ✅ |
| [Day 06 - P1 (Formula)](./day06/)| `155.0 ns/op 112 B/op 3 allocs/op`       | ✅ |
| [Day 06 - P2 (Formula)](./day06/)|  `62.56 ns/op 0 B/op 0 allocs/op`        | ✅ |
||| |
| [Day 07 - P1](./day07/)          |`83805 ns/op 40 B/op 3 allocs/op`         | ✅ |
| [Day 07 - P2](./day07/)          |`99430 ns/op 40 B/op 3 allocs/op`         | ✅ |
||| |
| [Day 08 - Parsing](./day08/)     |`2597 ns/op 0 B/op 0 allocs/op`           | ✅ |
| [Day 08 - P1](./day08/)          |`33814 ns/op 0 B/op 0 allocs/op`          | ✅ |
| [Day 08 - P2](./day08/)          |`179994 ns/op 1896 B/op 102 allocs/op`    | ❌ |
||| |
| [Day 09 - P1](./day09/)          |`39380 ns/op 4144 B/op 2 allocs/op`       | ✅ |
| [Day 09 - P2](./day09/)          |`43172 ns/op 4144 B/op 2 allocs/op`       | ✅ |
||| |
| [Day 10 - P1](./day10/)          |`20604 ns/op 24304 B/op 142 allocs/op`| ✅ |
| [Day 10 - P2](./day10/)          |`92380 ns/op 24304 B/op 142 allocs/op`| ✅ |
||| |
| [Day 11 - P1](./day11/)          |`97037 ns/op 0 B/op 0 allocs/op`| ✅ |
| [Day 11 - P2](./day11/)          |`91704 ns/op 0 B/op 0 allocs/op`| ✅ |
||| |
| [Day 12 - P1](./day12/)          |                                          | |
| [Day 12 - P2](./day12/)          |                                          | |
||| |
| [Day 13 - P1](./day13/)          |                                          | |
| [Day 13 - P2](./day13/)          |                                          | |
||| |
| [Day 14 - P1](./day14/)          |                                          | |
| [Day 14 - P2](./day14/)          |                                          | |
||| |
| [Day 15 - P1](./day15/)          |                                          | |
| [Day 15 - P2](./day15/)          |                                          | |
||| |
| [Day 16 - P1](./day16/)          |                                          | |
| [Day 16 - P2](./day16/)          |                                          | |
||| |
| [Day 17 - P1](./day17/)          |                                          | |
| [Day 17 - P2](./day17/)          |                                          | |
||| |
| [Day 18 - P1](./day18/)          |                                          | |
| [Day 18 - P2](./day18/)          |                                          | |
||| |
| [Day 19 - P1](./day19/)          |                                          | |
| [Day 19 - P2](./day19/)          |                                          | |
||| |
| [Day 20 - P1](./day20/)          |                                          | |
| [Day 20 - P2](./day20/)          |                                          | |
||| |
| [Day 21 - P1](./day21/)          |                                          | |
| [Day 21 - P2](./day21/)          |                                          | |
||| |
| [Day 22 - P1](./day22/)          |                                          | |
| [Day 22 - P2](./day22/)          |                                          | |
||| |
| [Day 23 - P1](./day23/)          |                                          | |
| [Day 23 - P2](./day23/)          |                                          | |
||| |
| [Day 24 - P1](./day24/)          |                                          | |
| [Day 24 - P2](./day24/)          |                                          | |
