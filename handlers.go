package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/windows/registry"
)

// CmdHandler handles the Wails command to disable access to cmd
func CmdHandler() error {
	keyPath := `Software\Policies\Microsoft\Windows\System`

	key, err := registry.OpenKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
	if err != nil {
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
	keyPath := `SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU`

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		key, _, err = registry.CreateKey(registry.CURRENT_USER, keyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer key.Close()
		fmt.Println("Parent registry key created.")
	}
	defer key.Close()

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
		if scanner.Text() == "127.0.0.1 "+website {
			fmt.Printf("%s is already blocked.\n", website)
			return nil
		}
	}

	newLine := "\n127.0.0.1 " + website
	_, err = fmt.Fprintln(file, newLine)
	if err != nil {
		return err
	}

	cmd := exec.Command("ipconfig", "/flushdns")
	err = cmd.Run()
	if err != nil {
		return err
	}

	fmt.Printf("%s blocked successfully.\n", website)
	return nil
}
func ScreenTimeoutHandler(timeout uint) error {
	keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System`

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		key, _, err = registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.ALL_ACCESS)
		if err != nil {
			fmt.Printf("Error creating or opening registry key: %v\n", err)
			return err
		}
		defer key.Close()
		fmt.Println("Parent registry key created.")
	}
	defer key.Close()

	err = key.SetDWordValue("InactivityTimeoutSecs", uint32(180))
	if err != nil {
		fmt.Println("Error setting registry value:", err)
		return err
	}

	fmt.Println("Lock Screen Timeout changed. Please restart your system for changes to take effect.")
	return nil
}
