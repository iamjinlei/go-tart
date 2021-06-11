<h1 align="center">go-tart (go-ta realtime)</h1>

[go-talib](https://github.com/markcheno/go-talib) project provides a nice port of the python [TA-Lib](http://ta-lib.org/).
But it lacks the capability to stream new updates into indicators.
In some circumstance, re-computation of the full history could be costly.
Thus the birth of go-tart (go-TA Realtime).
The codebase is heavily influenced by [go-talib](https://github.com/markcheno/go-talib).
Still a work-in-project!

## Install

Install the package with:

```bash
go get github.com/iamjinlei/go-tart
```

Import it with:

```go
import "github.com/iamjinlei/go-tart"
```

and use `tart` as the package name

## Exmaple

```go
package main

import (
	"fmt"

	"github.com/iamjinlei/go-tart"
)

func main() {
	sma := tart.NewSma(5)
	for i := 0; i < 20; i++ {
		fmt.Printf("sma[%v] = %.4f\n", i, sma.Update(float64(i%7)))
	}
}
```
