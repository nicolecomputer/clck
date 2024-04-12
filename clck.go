package main

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	timer       timer.Model
	currentTime time.Time
}

func initializeModel() model {
	return model{
		currentTime: time.Now(),
		timer:       timer.NewWithInterval(time.Hour*24, 1),
	}
}

func (m model) Init() tea.Cmd {
	return m.timer.Init()
}

func (m model) View() string {
	return m.currentTime.Format("3:04:05 PM")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.currentTime = time.Now()
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func main() {
	p := tea.NewProgram(initializeModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
