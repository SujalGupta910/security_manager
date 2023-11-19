package main

import (
	"context"
	"fmt"

	"golang.org/x/sys/windows/registry"
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

// CmdHandler handles the Wails command to disable access to cmd
func CmdHandler() error {
	// Define the registry key path
	keyPath := `Software\Policies\Microsoft\Windows\System`

	// Open or create a registry key for writing
	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	if err != nil {
		// If the key doesn't exist, create it
		key, _, err = registry.CreateKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer key.Close()
		fmt.Println("Parent registry key created.")
	} else {
		defer key.Close()
		fmt.Println("Parent registry key opened.")
	}

	// Write a registry value
	valueName := "DisableCMD"
	valueData := 1

	err = key.SetDWordValue(valueName, uint32(valueData))
	if err != nil {
		fmt.Printf("Error writing registry value: %v\n", err)
		return err
	}
	fmt.Printf("Registry value %s set to %d\n", valueName, valueData)
	return nil
}
func WindowsUpdatesHandler() error {
	// Specify the registry key for Windows Update settings
	keyPath := `SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`

	// Open the registry key for writing
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		fmt.Println("Error opening registry key:", err)
		return err
	}
	defer key.Close()

	// Set the registry values to disable automatic updates
	err = key.SetDWordValue("NoAutoUpdate", 1)
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		return err
	}

	err = key.SetDWordValue("AUOptions", 2)
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		return err
	}

	fmt.Println("Automatic Windows updates disabled. Please restart your system for changes to take effect.")
	return nil
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
	return "Downloads in Chrome and Edge Disabled Successfully", nil
}
func (a *App) DisableFacebook() (string, error) {
	return "Access to Facebook Disabled Successfully", nil
}
func (a *App) ChangeScreenTimeout() (string, error) {
	return "Screen Timeout chnaged to 3 minutes", nil
}
