package main

import (
	"github.com/mikerybka/util"
)

type Root string

func main() {
	panic(util.NewAPI([]string{}).Start())
}
