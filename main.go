// main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "hexcli",
		Short: "A CLI tool for initializing hexagonal architecture projects",
		Long:  "A CLI tool to bootstrap new Go projects with hexagonal architecture.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	var initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project with hexagonal architecture",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Initializing project with hexagonal architecture...")
			if err := initProject(); err != nil {
				fmt.Println("Error initializing project:", err)
				return
			}
			fmt.Println("Project initialized successfully.")
		},
	}

	var createCmd = &cobra.Command{
		Use:   "create [component] [name]",
		Short: "Create a new file for a specific component (port, adapter, etc.)",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			component := args[0]
			name := args[1]
			if err := createFile(component, name); err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			fmt.Printf("%s created successfully.\n", name)
		},
	}

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initProject() error {
	dirs := []string{
		"cmd",
		"pkg",
		"internal/core/domain",
		"internal/core/ports",
		"internal/core/services/service1",
		"internal/handlers",
		"internal/repositories",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	
	files := map[string]string{
		"internal/core/domain/domain1.go":             "",
		"internal/core/ports/repositories.go":         "",
		"internal/core/ports/services.go":             "",
		"internal/core/services/service1/service1.go": "",
	}

	for file, content := range files {
		if err := createFileWithContent(file, content); err != nil {
			return err
		}
	}

	return nil
}

func createFileWithContent(filepath, content string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		return err
	}

	return nil
}

func createFile(component, name string) error {
	var dir string

	switch component {
	case "port":
		dir = "internal/core/ports"
	case "adapter":
		dir = "internal/adapters"
	case "handler":
		dir = "internal/handlers"
	case "repository":
		dir = "internal/repositories"
	default:
		return fmt.Errorf("unknown component: %s", component)
	}

	filename := filepath.Join(dir, name+".go")
	if _, err := os.Create(filename); err != nil {
		return err
	}

	return nil
}
