package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

type InputResoruce struct {
	Name string `json:"name"`
	Info string `json:"info"`
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
	input, err := exec.Command("./test.sh").Output()
	if err != nil {
		log.Fatal(err)
	}

	inputs := []InputResoruce{}
	if err := json.Unmarshal(input, &inputs); err != nil {
		log.Fatal(err)
	}

	results = []Resource{}
	for idx, input := range inputs {
		results = append(results, Resource{
			Name:   input.Name,
			Info:   input.Info,
			Target: fmt.Sprintf("%d. %s %s", idx+1, input.Name, input.Info),
		})
	}

	return results
}

func (a *App) Search(selected string) []Resource {

	targets := lo.Map(results, func(r Resource, _ int) string {
		return r.Target
	})
	// TODO: 単語数分、ループしてさらに絞り込む
	filteredTargets := fuzzy.FindNormalizedFold(selected, targets)
	filteredResults := lo.FilterMap(results, func(r Resource, _ int) (Resource, bool) {
		if lo.Contains(filteredTargets, r.Target) {
			return r, true
		}
		return r, false
	})

	return filteredResults
}

func (a *App) Exec(selected Resource) {
	shell := "zsh"
	cmdStr := fmt.Sprintf("cd $HOME; cd %s; %s", selected.Info, shell)
	args := []string{"--", shell, "-c", cmdStr}
	cmd := exec.Command("/usr/bin/gnome-terminal", args...)
	cmd.Run()

	runtime.Quit(a.ctx)
}
