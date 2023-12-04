package main

import (
	"github.com/magefile/mage/mage"
	"os"
)

// main allows GoLand ide to run mage targets.
func main() {
	os.Exit(mage.Main())
}
