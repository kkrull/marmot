package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	metaRepoGrp   = "meta-repo-group"
	repositoryGrp = "repository-group"
)

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

/* Child commands */

func addCommandGroups(cobraCmd *cobra.Command) {
	cobraCmd.AddGroup(
		metaRepoGroup.toCobraGroup(),
		repositoryGroup.toCobraGroup(),
	)
}

func AddMetaRepoCommand(parent *cobra.Command, child cobra.Command) {
	child.GroupID = metaRepoGrp
	parent.AddCommand(&child)
}

func AddRepositoryCommand(parent *cobra.Command, child cobra.Command) {
	child.GroupID = repositoryGrp
	parent.AddCommand(&child)
}
