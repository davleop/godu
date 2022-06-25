package main

import (
    "os"
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/davleop/godu/internal/tui"
)

func main() {
    p := tea.NewProgram(initialModel())
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}
