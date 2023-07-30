package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.WindowCenter(ctx)
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Quit the application
func (a *App) Quit() {
	runtime.Quit(a.ctx)
}

var results []string

func (a *App) GetSources() []string {
	results = []string{}

	urls := []string{
		"/home/takeshi-miyajima/workspace_highway/product1/backend-api",
		"/home/takeshi-miyajima/workspace_highway/product1/frontend-web",
		"/home/takeshi-miyajima/workspace_highway/memo",
		"/home/takeshi-miyajima/private/memo",
	}
	results = append(results, urls...)
	return results
}

func (a *App) Search(selected string) []string {
	return fuzzy.FindNormalizedFold(selected, results)
}

func (a *App) Exec(selected string) {
	shell := "zsh"
	cmdStr := fmt.Sprintf("cd $HOME; cd %s; %s", selected, shell)
	args := []string{"--", shell, "-c", cmdStr}
	cmd := exec.Command("/usr/bin/gnome-terminal", args...)
	cmd.Run()

	runtime.Quit(a.ctx)
}
