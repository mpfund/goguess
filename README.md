# goguess
brute forcer in golang

Easy to use brute forcer:
``` go
// goGuessTester.go
package main

import (
	"fmt"
	"github.com/BlackEspresso/goguess"
)

func main() {
	// Setup, use a,b,c and start from length 1
	bfs := &goguess.BruteForceSetup{
		[]rune{'a', 'b', 'c'}, // our characters or symbols
		1, // min string length  => "a"
		2, // max string length  => "cc"
	}

	// here we store our current position
	st := goguess.NewState(bfs)

	for {
		// convert state as string
		fmt.Println(goguess.StateToString(st, bfs))

		if goguess.HasNext(st, bfs) {
			st = goguess.MoveNext(st, bfs)
		} else {
			break
		}
	}
	fmt.Println("done")
}
```

prints:
```
a
b
c
aa
ba
ca
ab
bb
cb
ac
bc
cc
done
```
