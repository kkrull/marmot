package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commandGroups = []commandGroup{metaRepoGroup, repositoryGroup}

// A group to which a top-level command may optionally belong.
type commandGroup string

const (
	metaRepoGroup   commandGroup = "meta-repo-group"
	repositoryGroup commandGroup = "repository-group"
)

func (group commandGroup) id() string {
	return string(group)
}

func (group commandGroup) title() string {
	switch group {
	case metaRepoGroup:
		return "Meta Repo Commands"
	case repositoryGroup:
		return "Repository Commands"
	default:
		return fmt.Sprintf("Unknown group <%s>", group)
	}
}

func (group commandGroup) toCobraGroup() *cobra.Group {
	return &cobra.Group{ID: group.id(), Title: group.title()}
}
