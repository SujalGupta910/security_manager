package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

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
func FacebookHandler(website string) error {
	hostsFilePath := "C:\\Windows\\System32\\drivers\\etc\\hosts"
	// Open hosts file with append permissions
	file, err := os.OpenFile(hostsFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer file.Close()

	// Check if the entry already exists in the hosts file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "127.0.0.1\t"+website {
			fmt.Printf("%s is already blocked.\n", website)
			return nil
		}
	}

	// If the entry doesn't exist, add it to the hosts file
	_, err = file.WriteString("127.0.0.1\t" + website + "\n")
	if err != nil {
		return err
	}

	fmt.Printf("%s blocked successfully.\n", website)
	return nil
}
func ScreenTimeoutHandler(timeout uint) error {
	const (
		SPI_SETSCREENSAVETIMEOUT = 15
	)

	var (
		user32               = syscall.NewLazyDLL("user32.dll")
		systemParametersInfo = user32.NewProc("SystemParametersInfoW")
	)

	success, _, err := systemParametersInfo.Call(
		uintptr(SPI_SETSCREENSAVETIMEOUT),
		uintptr(timeout),
		0,
		0,
	)
	if success == 0 {
		return err
	}
	return nil
}
