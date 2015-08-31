# Dungeon Game

Simple implementation of the [Dungeon Game](https://leetcode.com/problems/dungeon-game/) in GO.

Contains 2 strategies: Recursive and Dijkstra.
Dijkstra is hard coded in main.go at the moment.

### Running

It takes a dungeon map YAML as parameter:

```YAML
# Example dungeon map:
---
rooms:
- [1, 2, -5, -4]
- [-5, -6, -7, 9]
- [9, 10, 11, 12]
```

To run:
```bash
./dungeongame -dungeon ./dungeonmap.yml
# => MIN HP: 4
# => BEST ROUTE: RIGHT -> DOWN -> DOWN -> RIGHT -> RIGHT
```

### Tests

```bash
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
make test
```
