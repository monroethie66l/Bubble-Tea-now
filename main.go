package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Identifiable interface to support selection persistence
type Identifiable interface {
	ID() string
}

// RefreshListMsg is the message sent when data is updated
type RefreshListMsg struct {
	Items []list.Item
}

func updateList(m *Model, newItems []list.Item) {
	// 1. Store current selection
	var selectedID string
	if item, ok := m.list.SelectedItem().(Identifiable); ok {
		selectedID = item.ID()
	}

	// 2. Update items
	m.list.SetItems(newItems)

	// 3. Restore selection
	if selectedID != "" {
		for i, item := range newItems {
			if identifiable, ok := item.(Identifiable); ok && identifiable.ID() == selectedID {
				m.list.Select(i)
				return
			}
		}
	}

	// 4. Fallback: ensure index is within bounds if selection was lost
	if m.list.Index() >= len(newItems) && len(newItems) > 0 {
		m.list.Select(len(newItems) - 1)
	}
}

type Model struct {
	list list.Model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case RefreshListMsg:
		updateList(&m, msg.Items)
	}
	return m, nil
}

func main() {}