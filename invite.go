package main

import (
	"fmt"

	gitlab "github.com/xanzy/go-gitlab"
)

func InviteEmails(groupSlug string, emails []string) error {
	app, err := NewApp()
	if err != nil {
		return err
	}

	group, err := app.FindGroup(groupSlug)
	if err != nil {
		return err
	}
	if group == nil {
		return fmt.Errorf("group %s not found", groupSlug)
	}

	users, err := app.FindUsers(emails)
	if err != nil {
		return err
	}

	for _, e := range emails {
		user, ok := users[e]
		if !ok {
			user, err = app.CreateUser(e)
			if err != nil {
				return err
			}
			fmt.Printf("User %s created successfully\n", e)
		}

		_, _, err = app.client.GroupMembers.AddGroupMember(group.ID, &gitlab.AddGroupMemberOptions{
			UserID:      gitlab.Ptr(user.ID),
			AccessLevel: gitlab.Ptr(gitlab.ReporterPermissions),
		})
		if err != nil {
			fmt.Printf("Failed to invite %s: %v\n", e, err)
			continue
		}

		fmt.Printf("Invited %s to %s\n", e, groupSlug)
	}

	return nil
}
