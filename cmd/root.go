/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"local/plsfile/internal/plsfile"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pls",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.Flags().GetString("file")
		if err != nil {
			panic(err)
		}

		plsF, err := plsfile.ReadFile(file)
		if err != nil {
			panic(err)
		}

		if len(args) == 0 {
			fmt.Println("Please provide a job name. Available are:")
			for name, _ := range plsF.Jobs {
				fmt.Printf("%s, ", name)
			}
			fmt.Println("")
			return
		}

		job, found := plsF.Jobs[args[0]]
		if !found {
			fmt.Println("Job not found.")
			return
		}

		err = runCommands(job.Commands)
		if err != nil {
			panic(err)
		}

	},
}

func runCommands(commands []string) error {
	script := "set -e\n" + strings.Join(commands, "\n")
	cmd := exec.Command("sh", "-c", script)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plsfile.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "plsfile", "which file/path should be used. Default is ./plsfile")
}
