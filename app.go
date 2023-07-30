package main

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/samber/lo"
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

type Resource struct {
	Name string `json:"name"`
	Info string `json:"info"`
}
type Plugin struct {
	List []Resource
}

var results []Resource

func (a *App) GetInitialList() []Resource {
	results = []Resource{}

	urls := []Resource{
		{
			Name: "/home/takeshi-miyajima/workspace_highway/product1/backend-api",
		},
		{
			Name: "/home/takeshi-miyajima/workspace_highway/product1/frontend-web",
		},
		{
			Name: "/home/takeshi-miyajima/workspace_highway/memo",
		},
		{
			Name: "/home/takeshi-miyajima/private/memo",
		},
	}
	results = append(results, urls...)
	return results
}

func (a *App) Search(selected string) []Resource {

	sources := lo.Map(results, func(r Resource, _ int) string {
		return r.Name
	})
	filteredSources := fuzzy.FindNormalizedFold(selected, sources)
	filteredResults := lo.FilterMap(results, func(r Resource, _ int) (Resource, bool) {
		if lo.Contains(filteredSources, r.Name) {
			return r, true
		}
		return r, false
	})

	return filteredResults
}

func (a *App) Exec(selected string) {
	shell := "zsh"
	cmdStr := fmt.Sprintf("cd $HOME; cd %s; %s", selected, shell)
	args := []string{"--", shell, "-c", cmdStr}
	cmd := exec.Command("/usr/bin/gnome-terminal", args...)
	cmd.Run()

	runtime.Quit(a.ctx)
}
