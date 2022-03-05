package lingo

import "github.com/charmbracelet/lipgloss"

const (
	padding  = 2
	maxWidth = 200
)

var (
	width       = 96
	columnWidth = 30
)

var (
	// General.
	subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	text      = lipgloss.AdaptiveColor{Light: "#212121", Dark: "#FFF7DB"}

	divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(subtle).
		String()

	infoStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderTop(true).
			BorderForeground(subtle)

	// Dialog.
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(highlight).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	left  = lipgloss.NewStyle().Align(lipgloss.Left)
	right = lipgloss.NewStyle().Align(lipgloss.Right)

	// Page.
	lingoStyle = lipgloss.NewStyle().
			Bold(true).
			Italic(true).
			Foreground(special).
			SetString("Lip Gloss").
			BorderBottom(true)

	bold = lipgloss.NewStyle().Bold(true).Foreground(text)

	boldSpec = lipgloss.NewStyle().Bold(true).Foreground(special)

	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

	languageBar = lipgloss.NewStyle()
)
