package cmdshared

import (
	"fmt"

	"github.com/spf13/cobra"
)

var CommandGroups = []commandGroup{MetaRepoGroup, RepositoryGroup}

// A group to which a top-level command may optionally belong.
type commandGroup string

const (
	MetaRepoGroup   commandGroup = "meta-repo-group"
	RepositoryGroup commandGroup = "repository-group"
)

func (group commandGroup) Id() string {
	return string(group)
}

func (group commandGroup) title() string {
	switch group {
	case MetaRepoGroup:
		return "Meta Repo Commands"
	case RepositoryGroup:
		return "Repository Commands"
	default:
		return fmt.Sprintf("Unknown group <%s>", group)
	}
}

func (group commandGroup) ToCobraGroup() *cobra.Group {
	return &cobra.Group{ID: group.Id(), Title: group.title()}
}
