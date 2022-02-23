package lingo

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Lingo(cmd *cobra.Command, args []string) {

	root := "."
	if len(args) > 0 {
		_, err := os.Stat(args[0])
		if err != nil {
			fmt.Printf("'%s' is not a valid path!", args[0])

			os.Exit(0)
		}

		root = args[0]
	}

	ng, err := cmd.Flags().GetBool("nogitignore")

	if err != nil {
		ng = false
	}

	m := InitialModel(root, !ng)
	if err := tea.NewProgram(&m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
