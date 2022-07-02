package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"

    . "internal/tui"
    . "internal/du"
)

func main() {
    files, err := ListFilesRecursivelyInParallel(".")
    if err != nil {
        log.Fatalln(err)
    }

    if len(files) > 0 {
        log.Println(files[0])
    }

    p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}
