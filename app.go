package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"

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

type Plugin struct {
	Command        string          `json:"command"`
	CommandArgs    []string        `json:"command_args"`
	InputResources []InputResoruce `json:"input_resources"`
}

var plugin Plugin

// 	Command:        "/usr/bin/gnome-terminal",
// 	CommandArgs:    []string{"--", "zsh", "-c", "cd $HOME; cd ${info}; zsh"},
// 	InputResources: []InputResoruce{},
// }

type InnerResource struct {
	Name   string `json:"name"`
	Info   string `json:"info"`
	Target string `json:"target"` // TODO: これを検索対象とする。ついでに必ずユニークになるように内部的に番号を振る
	Tag    string `json:"tag"`
}

var innerResources []InnerResource

func (a *App) GetInitialList() []InnerResource {
	input, err := exec.Command("./change_directory.sh").Output()
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(input, &plugin); err != nil {
		log.Fatal(err)
	}

	innerResources = []InnerResource{}
	for idx, inputResource := range plugin.InputResources {
		innerResources = append(innerResources, InnerResource{
			Name:   inputResource.Name,
			Info:   inputResource.Info,
			Target: fmt.Sprintf("%d. %s %s", idx+1, inputResource.Name, inputResource.Info),
		})
	}

	return innerResources
}

func (a *App) Search(selected string) []InnerResource {

	targets := lo.Map(innerResources, func(r InnerResource, _ int) string {
		return r.Target
	})
	// TODO: 単語数分、ループしてさらに絞り込む
	filteredTargets := fuzzy.FindNormalizedFold(selected, targets)
	filteredResults := lo.FilterMap(innerResources, func(r InnerResource, _ int) (InnerResource, bool) {
		if lo.Contains(filteredTargets, r.Target) {
			return r, true
		}
		return r, false
	})

	return filteredResults
}

func (a *App) Exec(selected InnerResource) {
	newArgs := []string{}
	for _, arg := range plugin.CommandArgs {
		newArg := strings.Replace(arg, "${name}", selected.Name, -1)
		newArg = strings.Replace(newArg, "${info}", selected.Info, -1)
		newArgs = append(newArgs, newArg)
	}
	cmd := exec.Command(plugin.Command, newArgs...)
	cmd.Run()

	runtime.Quit(a.ctx)
}
