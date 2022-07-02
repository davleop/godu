package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"

    . "internal/du"
    . "internal/tui"
)

func main() {
    directory := "."
    cPos := 0
    hidden := false
    defaultOrdering := "name"
    directoryOrder := true
    diskUsage := true
    // percentage := true, graph, both, none
    uniqCol := false
    modifyTime := false

    files, err := ListFilesRecursivelyInParallel(directory)
    if err != nil {
        log.Fatalln(err)
    }

    if len(files) > 0 {
        log.Println(files[0])
    }

    initialModel := InitialModel(directory, cPos, hidden, defaultOrdering, directoryOrder, diskUsage, uniqCol, modifyTime, files)

    p := tea.NewProgram(initialModel, tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}
