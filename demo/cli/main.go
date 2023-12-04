package main

import (
	"github.com/mlctrez/go-nap/demo/web"
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/go-nap/nap/render"
)

func main() {
	render.Run(nap.NewRouter(web.Routes))
}
