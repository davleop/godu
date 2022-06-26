package main

import (
    "os"
    "log"

    "gopkg.in/alecthomas/kingpin.v2"
    tea "github.com/charmbracelet/bubbletea"

    . "internal/tui"
)

var (
    directory = kingpin.Arg("dir", "Directory to view.").Required().String()
)

func main() {
    kingpin.Version("0.0.1a")
    kingpin.Parse()

    dir, err := os.Stat(*directory)

    if err != nil {
        log.Fatalf("failed to open %s", *directory)
    }
    if !dir.IsDir() {
        log.Fatalf("%s is not a directory", dir.Name())
    }

    log.Printf("Using %s\n", dir.Name())

    p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}
