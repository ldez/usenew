package main

import (
	"github.com/ldez/usenew"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(usenew.Analyzer)
}
