package cmdshared

import "github.com/spf13/cobra"

func NewCommandGroup(id string, title string) CommandGroup {
	return CommandGroup{Id: id, Title: title}
}

type CommandGroup struct {
	Id    string
	Title string
}

func (group CommandGroup) ToCobraGroup() *cobra.Group {
	return &cobra.Group{
		ID:    group.Id,
		Title: group.Title,
	}
}
