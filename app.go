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
	Name   string   `json:"name"`
	Info   string   `json:"info"`
	Target string   `json:"target"` // TODO: これを検索対象とする。ついでに必ずユニークになるように内部的に番号を振る
	Tag    []string `json:"tag"`
}
type Plugin struct {
	List []Resource
}

var results []Resource

func (a *App) GetInitialList() []Resource {
	results = []Resource{}

	// targetを自動生成すること
	urls := []Resource{
		{
			Name:   "backend-api",
			Info:   "/home/takeshi-miyajima/workspace_highway/product1/backend-api",
			Target: "1. backend-api /home/takeshi-miyajima/workspace_highway/product1/backend-api",
			Tag:    []string{"cd"},
		},
		{
			Name:   "frontend-web",
			Info:   "/home/takeshi-miyajima/workspace_highway/product1/frontend-web",
			Target: "2. frontend-web /home/takeshi-miyajima/workspace_highway/product1/frontend-web",
			Tag:    []string{"cd"},
		},
		{
			Name:   "workspace_highway memo",
			Info:   "/home/takeshi-miyajima/workspace_highway/memo",
			Target: "3. workspace_hiway memo /home/takeshi-miyajima/workspace_highway/memo",
			Tag:    []string{"cd"},
		},
		{
			Name:   "private memo",
			Info:   "/home/takeshi-miyajima/private/memo",
			Target: "4. private memo /home/takeshi-miyajima/private/memo",
			Tag:    []string{"cd"},
		},
	}
	results = append(results, urls...)
	return results
}

func (a *App) Search(selected string) []Resource {

	sources := lo.Map(results, func(r Resource, _ int) string {
		return r.Target
	})
	// 単語数分、ループしてさらに絞り込む
	filteredSources := fuzzy.FindNormalizedFold(selected, sources)
	filteredResults := lo.FilterMap(results, func(r Resource, _ int) (Resource, bool) {
		if lo.Contains(filteredSources, r.Target) {
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
