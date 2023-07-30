package main

import (
	"context"
	"fmt"

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

var results = []string{
	"item1",
	"item2",
	"item3",
	"item4",
	"item5",
	"item6",
	"item7",
	"item8",
	"item9",
}

func (a *App) GetSources() []string {
	return results
}

func (a *App) Search(selected string) []string {
	return fuzzy.FindNormalizedFold(selected, results)
}
