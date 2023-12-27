package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/mlctrez/go-nap/magefiles/generator"
	"os"
	"os/exec"
	"os/signal"
)

// to suppress unused function warnings in IDE
var _ = []any{Build, ChromeDp, ChromeDpReload, DevServer, Default}

var Default = ChromeDp

func Build(ctx context.Context) error {
	return run(ctx, makeTemp, compileHtml, buildWasm, buildApp)
}

func DevServer(ctx context.Context) (err error) {
	if err = Build(ctx); err != nil {
		return err
	}

	notifyCtx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	server := exec.CommandContext(notifyCtx, "web/app.bin")
	server.Stdout = os.Stdout

	err = server.Start()
	if err != nil {
		panic(err)
	}
	_, err = server.Process.Wait()
	return err
}

func ChromeDpReload(ctx context.Context) (err error) {
	for {
		if err = ChromeDp(ctx); err != nil {
			if err.Error() == "window closed" {
				return nil
			}
			return err
		}
	}
}

func ChromeDp(ctx context.Context) (err error) {
	if err = Build(ctx); err != nil {
		return err
	}

	notifyCtx, notifyCancel := signal.NotifyContext(ctx, os.Interrupt)
	defer notifyCancel()

	server := exec.CommandContext(notifyCtx, "temp/app.bin")
	server.Stdout = os.Stdout

	err = server.Start()
	if err != nil {
		panic(err)
	}

	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts, chromedp.Flag("profile-directory", "Profile 1"))
	opts = append(opts, chromedp.UserDataDir("temp/chrome"))
	opts = append(opts, chromedp.WindowSize(1920, 1080))
	opts = append(opts, chromedp.Flag("headless", false))
	opts = append(opts, chromedp.Flag("auto-open-devtools-for-tabs", true))
	allocCtx, cancelAllocCtx := chromedp.NewExecAllocator(notifyCtx, opts...)

	defer cancelAllocCtx()
	chromeDpContext, cancelCtx := chromedp.NewContext(allocCtx)
	defer cancelCtx()

	tasks := []chromedp.Action{
		chromedp.Navigate("http://127.0.0.1:8080"),
		chromedp.WaitEnabled(`#chromeDpDone`),
		chromedp.ActionFunc(chromedp.Cancel),
	}
	err = chromedp.Run(chromeDpContext, tasks...)
	if !errors.Is(err, context.Canceled) {
		return err
	}

	err = chromedp.Cancel(ctx)
	if err != nil {
		// The chromedp browser window was closed so cancel the process
		err = server.Cancel()
		if err != nil {
			return err
		}
		return errors.New("window closed")
	}

	_, err = server.Process.Wait()
	return err
}

func makeTemp(ctx context.Context) (err error) {
	return os.MkdirAll("temp", 0755)
}

func compileHtml(ctx context.Context) (err error) {
	return generator.GenerateDirs("demo", "nap")
}

func buildWasm(ctx context.Context) error {
	return goCmd(true, "build", "-o", "temp/app.wasm", "demo/cli/main.go")
}

func buildApp(ctx context.Context) error {
	return goCmd(false, "build", "-o", "temp/app.bin", "demo/cli/main.go")
}

func goCmd(wasm bool, args ...string) error {
	cmd := exec.Command("go", args...)
	if wasm {
		cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	}
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Println(string(output))
	}
	return err
}

func run(ctx context.Context, commands ...func(ctx context.Context) error) error {
	for _, command := range commands {
		if err := command(ctx); err != nil {
			return err
		}
	}
	return nil
}
