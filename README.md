# Prefix

A package for checking which strings match any of a list of prefix strings

## Usage

```
package main

import "github.com/GeorgeNagel/prefix"

func main() {
  // Get the prefixes
  // prefixes := stuff()
  // Get the strings to check
  // stringsToCheck := stuff()
  stringsMachingPrefix := prefix.Match(prefixes, stringsToCheck)
}
```

## Tests

```
go test
```

## Benchmark tests

```
go test -bench .
```
