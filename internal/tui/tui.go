package tui

import (
	. "internal/du"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("FFFDF5")).
			Background(lipgloss.Color("25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{"Light": "#04B575", "Dark": "#04B575"}).
				Render
)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
}

func newListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
	}
}

type Model struct {
	// This section is for maintaining the `du` content
	CurrentDirectory string
	Cursor           int
	Subcursor        int
	ShowHidden       bool
	Order            string
	DirectoryFirst   bool
	ShowDiskUsage    bool
	ShowUniqCol      bool
	ModifyTime       bool
	Files            []File
	currentFiles     []File

	// the rest is for actually maintaining the TUI display
	list         list.Model
	keys         *listKeyMap
	delegateKeys *delegateKeyMap
}

func NewModel(m Model) model {
	//
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

/*func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.Cursor > 0 {
				m.Cursor--
				m.Subcursor--
			}

		case "down", "j":
			// cycle items
			if m.Cursor < len(m.currentFiles)-1 {
				m.Cursor++
				m.Subcursor++
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
			m.Viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.Viewport.SetContent(m.Content)
			m.Ready = true

			m.Viewport.YPosition = headerHeight + 1
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.Viewport))
		}
	}

	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) fileView() string {
	s := ""
	for _, file := range m.Files {
		if file.HighDir == m.CurrentDirectory {
			m.currentFiles = append(m.currentFiles, file)
		}
	}

	m.Content = ""
	for i, file := range m.currentFiles {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %d\t%s\n", cursor, file.Size, file.Name)
	}
	s += fmt.Sprintf("\nCursor: %d\n", m.Cursor)
	for _, file := range m.currentFiles {
		s += fmt.Sprintln(file)
	}
	return s
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

	m.Viewport.SetContent(m.fileView())
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
}*/
