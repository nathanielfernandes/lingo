package lingo

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	languages Languages
	bars      *Bars
	langs     LangArray
}

func InitialModel(root string, gitignore bool) Model {
	langs := GetLanguages(root, gitignore)

	bars := Bars{}

	for _, lang := range langs {
		pg := progress.New(progress.WithDefaultGradient())
		bars[lang.Name] = &pg
	}

	return Model{
		languages: langs,
		bars:      &bars,
		langs:     langs.Slice(),
	}
}

type Bars map[string]*progress.Model
type Counted struct{}

func (b Bars) ForEach(fn func(progress.Model)) {
	for _, pg := range b {
		fn(*pg)
	}
}

func (m *Model) Init() tea.Cmd {
	m.languages.CountLines()

	return func() tea.Msg {
		m.languages.Wait()
		return Counted{}
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Counted:
		m.langs = m.languages.Sorted()
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		for _, pg := range *m.bars {
			pg.Width = msg.Width - padding*2 - 4
			if pg.Width > maxWidth {
				pg.Width = maxWidth
			}

			width = pg.Width + 4
		}
		return m, nil

	}

	return m, nil
}

func (m Model) View() string {
	total := m.languages.GetTotal()
	fileCount := m.languages.FileCount()
	doc := strings.Builder{}
	stats := strings.Builder{}

	// title
	{

		desc := lipgloss.JoinVertical(lipgloss.Center,
			lingoStyle.Render("Lingo")+boldSpec.Render(" Summary"),
			infoStyle.Render(bold.Render(fmt.Sprintf(" %d Files", fileCount))+divider+bold.Render(fmt.Sprintf("%d Lines ", total))),
		)

		row := lipgloss.Place(width, 0, lipgloss.Center, lipgloss.Top, desc, lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle))

		doc.WriteString(row + "\n")
	}

	// Stats
	{
		for _, lang := range m.langs {

			info := bold.Render(fmt.Sprintf(" %d Files", lang.FileCount())) + divider + bold.Render(fmt.Sprintf("%d Lines ", lang.Count))

			repeat := width + 64 - (len(info) + len(lang.Name))
			if repeat < 0 {
				repeat = 0
			}
			desc := fmt.Sprintf("%s%s%s", bold.Render(lang.Name), strings.Repeat(" ", repeat), info)
			all := lipgloss.JoinVertical(lipgloss.Left, desc, strings.Repeat(" ", padding)+(*m.bars)[lang.Name].ViewAs(float64(lang.Count)/float64(total)))

			stats.WriteString(strings.Repeat(" ", padding) + all + "\n\n")
		}
	}

	// Stats Box
	{
		ui := lipgloss.JoinVertical(lipgloss.Center, strings.TrimRight(stats.String(), "\n"))
		dialog := lipgloss.Place(0, 0,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		doc.WriteString(dialog + "\n")
	}

	return docStyle.Render(doc.String())

}
