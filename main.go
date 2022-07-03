package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"

    . "internal/du"
    . "internal/tui"
)

func main() {
    directory := "."
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

    initialModel := Model {
        CurrentDirectory: directory,
        ShowHidden: hidden,
        Order: defaultOrdering,
        DirectoryFirst: directoryOrder,
        ShowDiskUsage: diskUsage,
        ShowUniqCol: uniqCol,
        ModifyTime: modifyTime,
        Files: files,
    }

    p := tea.NewProgram(initialModel, tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}
