package main

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mitchfen/osrs-herb-run-helper/internal/helpers"
)

var (
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#25A065")).
		Padding(0, 1)

	herbNameStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Width(12)

	profitStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B50FF"))

	containerStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#25A065")).
		Padding(1, 2)
)

type model struct {
	spinner  spinner.Model
	loading  bool
	herbs    []helpers.Herb
	err      error
}

type dataLoadedMsg struct {
	herbs []helpers.Herb
	err   error
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchData())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case dataLoadedMsg:
		m.loading = false
		m.herbs = msg.herbs
		m.err = msg.err
		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v", m.err)
	}
	if m.loading {
		return fmt.Sprintf("%s Loading herb data...", m.spinner.View())
	}

	var herbList []string
	for _, herb := range m.herbs {
		herbName := herbNameStyle.Render(herb.Name)
		profit := profitStyle.Render(fmt.Sprintf("%.0fk", math.Round(herb.ExpectedProfit/1000.0)))
		herbList = append(herbList, fmt.Sprintf("%s: %s", herbName, profit))
	}

	return containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			titleStyle.Render(fmt.Sprintf("Herb run profitability as of %s", time.Now().Format("01/02 15:04"))),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render("Assumes farming cape, magic secateurs, and 9 herb patches"),
			"",
			lipgloss.JoinVertical(lipgloss.Left, herbList...),
		),
	) + lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render("\nCtrl+C to exit")
}

func fetchData() tea.Cmd {
	return func() tea.Msg {
		parsedPriceJson := helpers.GetParsedPriceJson()
		herbs := helpers.BuildHerbsSlice(parsedPriceJson)
		return dataLoadedMsg{herbs: herbs}
	}
}

func main() {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := model{
		spinner: s,
		loading: true,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}