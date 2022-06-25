package main

import (
    "log"

    tea "github.com/charmbracelet/bubbletea"
    . "internal/tui"
)

func main() {
    p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
    if err := p.Start(); err != nil {
        log.Fatal(err)
    }
}
