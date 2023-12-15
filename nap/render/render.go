//go:build !js

package render

import (
	"fmt"
	"github.com/mlctrez/go-nap/nap/enc"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

import (
	"github.com/mlctrez/go-nap/nap"
	"github.com/mlctrez/wasmexec"
)

func Run(router nap.Router) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", page(router))
	mux.HandleFunc("/wasm.js", wasmJs())
	mux.HandleFunc("/runtime.js", static("demo/web/runtime.js"))
	mux.HandleFunc("/logo.svg", static("demo/web/logo.svg"))
	mux.HandleFunc("/bootstrap.min.css", static("demo/web/bootstrap.min.css"))
	mux.HandleFunc("/sign-in.css", static("demo/web/sign-in.css"))
	mux.HandleFunc("/bootstrap.bundle.min.js", static("demo/web/bootstrap.bundle.min.js"))
	mux.HandleFunc("/app.wasm", static("temp/app.wasm"))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func wasmJs() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		wasmexec.WriteLauncher(writer)
	}
}

func page(router nap.Router) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		encoder := enc.New().Indent("  ")
		if err := encoder.EncodePage(router.Page(request.URL)); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if err := encoder.Write(writer); err != nil {
			fmt.Println("writer write", err)
		}
	}
}

func static(path string) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		var staticFile *os.File
		if staticFile, err = os.Open(path); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer closeLog(staticFile)
		writer.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(path)))
		if _, err = io.Copy(writer, staticFile); err != nil {
			fmt.Println(err)
		}
	}
}

func closeLog(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Println(err)
	}
}
