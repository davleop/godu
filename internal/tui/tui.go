package tui

import (
    "fmt"

    tea "github.com/charmbracelet/bubbletea"

    . "internal/du"
)

type Model struct {
    currentDirectory string
    cursor  int
    showHidden bool
    order string
    directoryFirst bool
    showDiskUsage bool
    showUniqCol bool
    modifyTime bool
    directories []File
}

func InitialModel(cd string, c int, sh bool, o string, df bool, sdu bool, suc bool, mt bool, f []File) Model {
    return Model {
        currentDirectory: cd,
        cursor: c,
        showHidden: sh,
        order: o,
        directoryFirst: df,
        showDiskUsage: sdu,
        showUniqCol: suc,
        modifyTime: mt,
        directories: f,
    }
}

func (m Model) Init() tea.Cmd {
    return tea.Batch(tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
            if m.cursor < len(m.directories)-1 {
                m.cursor++
            }

        case "enter", " ", "l", "right":
            // open selected directory

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

func (m Model) View() string {
    /*s := "What should we buy at the market?\n\n"
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }*/

    s := "Header somehow here"
    s += fmt.Sprintf("Current directory: %s\n")
    s += "Footer somehow here"

    return s
}


