<h1 align="center">go-tart (go-TA RealTime)</h1>

[![Build Status](https://travis-ci.com/iamjinlei/go-tart.svg?branch=master)](https://travis-ci.com/iamjinlei/go-tart)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamjinlei/go-tart)](https://goreportcard.com/report/github.com/iamjinlei/go-tart)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Reference](https://pkg.go.dev/badge/github.com/iamjinlei/go-tart.svg)](https://pkg.go.dev/github.com/iamjinlei/go-tart)

## Intro

[go-talib](https://github.com/markcheno/go-talib) project provides a nice port of the C-implementation [TA-Lib](http://ta-lib.org/).
However, its code is very verbose and hard to read due to the straight translation from the "C" code of [TA-Lib](http://ta-lib.org/).
More often, streaming incremental updates to indicators is desirable.
We don't want to recalculate the full result if the sliding window only moves forward a single value while history values are not changed at all.
Recalculation of the history for every new value is expensive.
With those in mind, [go-tart](https://github.com/iamjinlei/go-tart) intends to support streaming updates of new values with a concise implementation.
Most of the indicators from [go-talib](https://github.com/markcheno/go-talib) are realized except a few trivial ones and those involves *Hilbert Transform*.
Results are verified using [TA-Lib](http://ta-lib.org/)'s output.
MACD's "stability" issue from [go-talib](https://github.com/markcheno/go-talib) is fixed.

## Performance

A benchmark is written to compare the performance for the set of indicators from both [go-talib](https://github.com/markcheno/go-talib) and [go-talib](https://github.com/markcheno/go-talib).
This is not a scientific evaluation but gives a sense of how they perform in case of a full recalculation.
As shown below, [go-tart](https://github.com/iamjinlei/go-tart) is about 3X slower in full recalculation, which is due to function call to *Update* in each iteration.
[go-talib](https://github.com/markcheno/go-talib) holds advantage because its code is inlined.
```bash
gotest -run ^$ -bench=.
BenchmarkTart-4             2019            496703 ns/op
BenchmarkTalib-4            7944            150311 ns/op
```
After the initial full calculation, [go-tart](https://github.com/iamjinlei/go-tart) should perform better as it only needs a single iteration cost to compute a new update.

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
