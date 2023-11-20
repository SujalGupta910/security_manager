package main

import (
	"context"
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
}

func (a *App) DisableWindowsUpdates() (string, error) {
	err := WindowsUpdatesHandler()
	if err != nil {
		return "", err
	}
	return "Disabled Auto Udates Successfully", nil
}
func (a *App) DisableCmd() (string, error) {
	err := CmdHandler()
	if err != nil {
		return "", err
	}
	return "Cmd access disabled Successfully", nil
}
func (a *App) DisableDownloads() (string, error) {
	err := DownloadsHandler()
	if err != nil {
		return "", err
	}
	return "Downloads in Chrome and Edge Disabled Successfully", nil
}
func (a *App) DisableFacebook() (string, error) {
	err := FacebookHandler("facebook.com")
	if err != nil {
		return "", err
	}
	return "Access to Facebook Disabled Successfully", nil
}
func (a *App) ChangeScreenTimeout() (string, error) {
	err := ScreenTimeoutHandler(uint(180))
	if err != nil {
		return "", err
	}
	return "Screen Timeout chnaged to 3 minutes", nil
}
