package cukesupport

import messages "github.com/cucumber/messages/go/v21"

func findTag(name string, tags []*messages.PickleTag) *messages.PickleTag {
	for _, tag := range tags {
		if tag.Name == name {
			return tag
		}
	}

	return nil
}
