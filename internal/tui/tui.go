package tui

import (
    "fmt"
    "strings"

    "github.com/charmbracelet/bubbles/viewport"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"

    . "internal/du"
)

type Model struct {
    CurrentDirectory string
    Cursor  int
    subcursor int
    ShowHidden bool
    Order string
    DirectoryFirst bool
    ShowDiskUsage bool
    ShowUniqCol bool
    ModifyTime bool
    Files []File
    Viewport viewport.Model
    currentFiles []int
    Ready bool
    Content string
}

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q", "esc":
            // exit
            return m, tea.Quit

        case "up", "k":
            // cycle items
            if m.subcursor > 0 {
                m.subcursor--
            }

        case "down", "j":
            // cycle items
            if m.subcursor < len(m.currentFiles)-1 {
                m.subcursor++
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

    case tea.WindowSizeMsg:
        headerHeight := lipgloss.Height(m.headerView())
        footerHeight := lipgloss.Height(m.footerView())
        verticalMarginHeight := headerHeight + footerHeight

        if !m.Ready {
            m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
            m.Viewport.YPosition = headerHeight
            //m.Viewport.HighPerformanceRendering = useHighPerformanceRenderer
            m.Viewport.SetContent(m.Content)
            m.Ready = true

            m.Viewport.YPosition = headerHeight + 1
        } else {
            m.Viewport.Width = msg.Width
            m.Viewport.Height = msg.Height - verticalMarginHeight
        }
    }

    m.Viewport, cmd = m.Viewport.Update(msg)
    cmds = append(cmds,cmd)

    return m, tea.Batch(cmds...)
}

func (m Model) headerView() string {
    title := titleStyle.Render(fmt.Sprintf("godu: %d ~ %d", m.Cursor, m.subcursor))
    line := strings.Repeat("-", max(0, m.Viewport.Width-lipgloss.Width(title)))
    return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
    info := infoStyle.Render(fmt.Sprintf("Total disk usage: %d/%v(%d)", m.Cursor, m.currentFiles, len(m.currentFiles)))
    line := strings.Repeat("-", max(0, m.Viewport.Width-lipgloss.Width(info)))
    return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (m Model) View() string {
    if !m.Ready {
        return "\n Initializing..."
    }

    // here is the base screen of listing the current directory's contents
    m.Content = ""
    for i, file := range m.Files {
        if file.HighDir == m.CurrentDirectory {
            cursor := " "
            if i == m.currentFiles[m.subcursor] {
                cursor = ">"
            }

            m.currentFiles = append(m.currentFiles, i)
            m.Content += fmt.Sprintf("%s %d\t%s\n", cursor, file.Size, file.Name)
        }
    }
    m.Viewport.SetContent(m.Content)

    return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
}


