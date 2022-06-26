package tui

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    choices []string
    cursor  int
    selected map[int]struct{}
}

func InitialModel() model {
    return model {
        choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

        selected: make(map[int]struct{}),
    }
}

func (m model) Init() tea.Cmd {
    return tea.Batch(tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "esc":
            // exit
            return m, tea.Quit

        case "up", "k":
            // cycle items
            if m.cursor > 0 {
                m.cursor--
            }

        case "down", "j":
            // cycle items
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "enter", " ", "l", "right":
            // open selected directory
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }

        case "left", "h":
            // go to parent directory

        case "n":
            // order by filename

        case "s":
            // order by size

        case "C":
            // order by number of items

        case "a":
            // toggle disk usage/apparent size

        case "M":
            // order by latest modify time

        case "d":
            // delete selected file or directory

        case "t":
            // toggles directories before files when sorting

        case "g":
            // toggles percentage, graph, both, or none

        case "u":
            // toggle display of shared/unique size column for directories

        case "c":
            // toggle display of child counts

        case "m":
            // toggle display of latest modify time

        case "e":
            // Show/hide hidden files and directories

        case "i":
            // show information about selected item

        case "r":
            // refresh calculations in the current directory

        case "b":
            // spawn shell in current directory
        }
    }

    return m, nil
}

func (m model) View() string {
    s := "What should we buy at the market?\n\n"

    for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nPress q to quit.\n"

    return s
}


