Relief Family for Feature Selection within Golang
==================

### Table of Content

* [Installation](#Installation)
* [Usage](#Usage)
* [Reference](#Reference)

### Installation

```bash
go get github.com/a2htray/relief
```

### Usage

`Relief`

```go
package main

import (
       	"fmt"
       	"github.com/a2htray/relief"
)

func main() {
    model := relief.NewRelief([][]float64{
        []float64{0, 0},
        []float64{0, 1},
        []float64{1, 0},
        []float64{1, 1},
        []float64{1, 1},
    }, []float64{0, 0, 0, 1, 1}, []int{
        relief.AttributeTypeDiscrete,
        relief.AttributeTypeDiscrete,
    })
    fmt.Println(model.Run(20))
}
```

```bash
[-0.6 0.1]
```

From the output, we confirm that the second feature significantly contributes to the labels in all dataset.

`RelifF`

```go
package main

import (
       	"fmt"
       	"github.com/a2htray/relief"
)

func main() {
    model := relief.NewReliefF([][]float64{
        []float64{0, 0, -1},
        []float64{0, 0, -2},
        []float64{0, 0, -3},
        []float64{0, 0, -4},
        []float64{1, 1, 1},
        []float64{1, 1, 2},
        []float64{1, 1, 3},
        []float64{1, 1, 4},
    }, []float64{0, 0, 0, 0, 1, 1, 1, 1}, []int{
        relief.AttributeTypeDiscrete,
        relief.AttributeTypeDiscrete,
        relief.AttributeTypeDiscrete,
    }, 2)
    fmt.Println(model.Run(100))
}
```

```bash
[-1.0000000000000007 -1.0000000000000007 0]
```

From the output, we should take the third feature into learning model, like SVM, NN, etc.

### Reference

* Robnik-Šikonja, Marko, and Igor Kononenko. "An adaptation of Relief for attribute estimation in regression." Machine Learning: Proceedings of the Fourteenth International Conference (ICML’97). Vol. 5. 1997.