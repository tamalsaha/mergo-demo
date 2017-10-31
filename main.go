package main

import (
	"github.com/appscode/mergo"
	"log"
	"encoding/json"
	"fmt"
)

type Out struct {
	I Inner
	Z string
}

type Inner struct {
	X string
	Y string
}

func main() {
	a := &Out{
		I: Inner{
			X: "x2",
		},
		Z: "z2",
	}
	b := Out{
		I: Inner{
			X: "x",
			Y: "y",
		},
		Z: "z",
	}

	err := mergo.Merge(a, b)
	if err != nil {
		log.Fatal(err)
	}
	v, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(v))
}
