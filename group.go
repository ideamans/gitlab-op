package main

import (
	"fmt"
	"path"

	gitlab "github.com/xanzy/go-gitlab"
)

func CreateGroup(slug, name string) error {
	app, err := NewApp()
	if err != nil {
		return err
	}

	opts := &gitlab.CreateGroupOptions{
		Name:                 &name,
		Path:                 gitlab.Ptr(path.Base(slug)),
		Description:          gitlab.Ptr(fmt.Sprintf("%s 様専用グループ", name)),
		MembershipLock:       gitlab.Ptr(false),
		AutoDevopsEnabled:    gitlab.Ptr(false),
		EmailsEnabled:        gitlab.Ptr(false),
		MentionsDisabled:     gitlab.Ptr(true),
		LFSEnabled:           gitlab.Ptr(false),
		RequestAccessEnabled: gitlab.Ptr(false),
	}

	dirSlug := path.Dir(slug)
	if dirSlug != "" && dirSlug != "." {
		group, err := app.FindGroup(dirSlug)
		if err != nil {
			return err
		}
		if group == nil {
			return fmt.Errorf("parent group %s not found", dirSlug)
		}
		opts.ParentID = gitlab.Ptr(group.ID)
	}

	group, _, err := app.client.Groups.CreateGroup(opts)
	if err != nil {
		return err
	}

	fmt.Printf("Group %s (%s) created successfully\n", name, slug)
	fmt.Printf("Open: %s\n", group.WebURL)

	return nil
}
