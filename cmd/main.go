package main

import (
	"fmt"

	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
)

var _ v1.ConstantSampler

func main() {
	fmt.Println("hello world")
}
