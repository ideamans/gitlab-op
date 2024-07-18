package main

import (
	"os"
	"slices"
	"strings"

	gitlab "github.com/xanzy/go-gitlab"
)

type App struct {
	config *Config
	client *gitlab.Client
}

func NewApp() (*App, error) {
	profile := os.Getenv("GITLAB_OP_PROFILE")
	if profile == "" {
		profile = "default"
	}

	config, err := LoadConfig(profile)
	if err != nil {
		return nil, err
	}

	client, err := gitlab.NewClient(config.Token, gitlab.WithBaseURL(config.Url))
	if err != nil {
		return nil, err
	}

	return &App{config, client}, nil
}

func (a *App) FindGroup(slugPath string) (*gitlab.Group, error) {
	for page := 1; ; page++ {
		groups, _, err := a.client.Groups.ListGroups(&gitlab.ListGroupsOptions{
			Search: &slugPath,
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(groups) == 0 {
			break
		}

		for _, group := range groups {
			if group.FullPath == slugPath {
				return group, nil
			}
		}
	}

	return nil, nil
}

func (a *App) FindUser(email string) (*gitlab.User, error) {
	for page := 1; ; page++ {
		users, _, err := a.client.Users.ListUsers(&gitlab.ListUsersOptions{
			Search: &email,
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			break
		}

		for _, user := range users {
			if user.Email == email {
				return user, nil
			}
		}
	}

	return nil, nil
}

func (a *App) FindUsers(emails []string) (map[string]*gitlab.User, error) {
	var result = make(map[string]*gitlab.User)
	for page := 1; ; page++ {
		users, _, err := a.client.Users.ListUsers(&gitlab.ListUsersOptions{
			ListOptions: gitlab.ListOptions{
				Page:    page,
				PerPage: 100,
			},
		})
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			break
		}

		for _, user := range users {
			if slices.Contains(emails, user.Email) {
				result[user.Email] = user
			}
		}
	}

	return result, nil
}

func (a *App) CreateUser(email string) (*gitlab.User, error) {
	username := strings.Replace(email, "@", "_", -1)
	user, _, err := a.client.Users.CreateUser(&gitlab.CreateUserOptions{
		Email:            &email,
		Username:         &username,
		Name:             &email,
		ResetPassword:    gitlab.Ptr(true),
		CanCreateGroup:   gitlab.Ptr(false),
		External:         gitlab.Ptr(true),
		ProjectsLimit:    gitlab.Ptr(0),
		Admin:            gitlab.Ptr(false),
		SkipConfirmation: gitlab.Ptr(true),
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
