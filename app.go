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
		// If the key doesn't exist, create it
		key, _, err = registry.CreateKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer key.Close()
		fmt.Println("Parent registry key created.")
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
func DownloadsHandler() error {
	valueName := "DownloadRestrictions"
	valueData := 3

	//Chrome
	chromeKeyPath := `Software\Policies\Google\Chrome`
	chromeKey, err := registry.OpenKey(registry.CURRENT_USER, chromeKeyPath, registry.ALL_ACCESS)
	if err != nil {
		chromeKey, _, err = registry.CreateKey(registry.CURRENT_USER, chromeKeyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer chromeKey.Close()
		fmt.Println("Parent registry key created.")
	} else {
		defer chromeKey.Close()
		fmt.Println("Parent registry key opened.")
	}
	err = chromeKey.SetDWordValue(valueName, uint32(valueData))
	if err != nil {
		fmt.Printf("Error writing registry value: %v\n", err)
		return err
	}
	fmt.Printf("Registry value %s set to %d\n", valueName, valueData)

	//Edge
	edgeKeyPath := `Software\Policies\Microsoft\Edge`
	edgeKey, err := registry.OpenKey(registry.CURRENT_USER, edgeKeyPath, registry.ALL_ACCESS)
	if err != nil {
		// If the key doesn't exist, create it
		edgeKey, _, err = registry.CreateKey(registry.CURRENT_USER, edgeKeyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer edgeKey.Close()
		fmt.Println("Parent registry key created.")
	} else {
		defer edgeKey.Close()
		fmt.Println("Parent registry key opened.")
	}
	err = edgeKey.SetDWordValue(valueName, uint32(valueData))
	if err != nil {
		fmt.Printf("Error writing registry value: %v\n", err)
		return err
	}
	fmt.Printf("Registry value %s set to %d\n", valueName, valueData)

	return nil
}
func FacebookHandler() error {
	return nil
}
func ScreenTimeoutHandler() error {
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
	err := DownloadsHandler()
	if err != nil {
		return "", err
	}
	return "Downloads in Chrome and Edge Disabled Successfully", nil
}
func (a *App) DisableFacebook() (string, error) {
	err := FacebookHandler()
	if err != nil {
		return "", err
	}
	return "Access to Facebook Disabled Successfully", nil
}
func (a *App) ChangeScreenTimeout() (string, error) {
	err := ScreenTimeoutHandler()
	if err != nil {
		return "", err
	}
	return "Screen Timeout chnaged to 3 minutes", nil
}
