package main

import (
	"github.com/mlctrez/go-nap/demo"
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/render"
)

func main() {
	render.Run(nap.NewRouter(demo.Routes))
}
