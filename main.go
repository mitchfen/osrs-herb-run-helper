package main

import (
	"fmt"
	"math"
	"time"

	"github.com/charmbracelet/lipgloss"
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

func main() {
	parsedPriceJson := helpers.GetParsedPriceJson()
	herbs := helpers.BuildHerbsSlice(parsedPriceJson)

	var herbList []string
	for _, herb := range herbs {
		herbName := herbNameStyle.Render(herb.Name)
		profit := profitStyle.Render(fmt.Sprintf("%.0fk", math.Round(herb.ExpectedProfit/1000.0)))
		herbList = append(herbList, fmt.Sprintf("%s: %s", herbName, profit))
	}

	fmt.Println(
		containerStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				titleStyle.Render(fmt.Sprintf("Herb run profitability as of %s", time.Now().Format("01/02 15:04"))),
				lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render("Assumes farming cape, magic secateurs, and 9 herb patches"),
				"",
				lipgloss.JoinVertical(lipgloss.Left, herbList...),
			),
		),
	)
}
