package main

import (
<<<<<<< Updated upstream
    "log"

    tea "github.com/charmbracelet/bubbletea"
    . "internal/tui"
=======
	"fmt"
	"log"

	"github.com/spf13/cobra"
	//. "internal/tui"
)

/*
	This var sets up the root command and then all other commands. The root command, according to Cobra's structure, is the first thing we hit when we run the program.
	Imagine it as an automatic constructor that's allowing us to run an instance of this program.
*/
var (
	versFlag bool
	rootCmd  = &cobra.Command{
		Use: "put usage example here",
		//TraverseChildren: true,
		Short: "This program shows disk usage",
		Long:  "Put longer version of Short here",
		Run: func(cmd *cobra.Command, args []string) {
			v, _ := cmd.Flags().GetBool("versFlag")
			if v {
				version()
			} else {
				fmt.Println("it brokey")
			}
		},
	}
>>>>>>> Stashed changes
)

func version() {
	fmt.Println("Version goes here")
}

func init() {

	rootCmd.PersistentFlags().BoolVarP(&versFlag, "version", "v", true, "-v or --version: will output the version of godu currently in use")

}

func main() {
<<<<<<< Updated upstream
    p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
=======
	//p := tea.NewProgram(InitialModel(), tea.WithAltScreen())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

>>>>>>> Stashed changes
}
