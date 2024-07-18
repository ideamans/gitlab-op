package main

import (
	"fmt"
	"path"

	gitlab "github.com/xanzy/go-gitlab"
)

func CreateProject(slug, name string) error {
	app, err := NewApp()
	if err != nil {
		return err
	}

	group, err := app.FindGroup(path.Dir(slug))
	if err != nil {
		return err
	}
	if group == nil {
		return fmt.Errorf("group %s not found", path.Dir(slug))
	}

	p, _, err := app.client.Projects.CreateProject(&gitlab.CreateProjectOptions{
		Name:                             &name,
		Path:                             gitlab.Ptr(path.Base(slug)),
		NamespaceID:                      gitlab.Ptr(group.ID),
		Description:                      &name,
		IssuesAccessLevel:                gitlab.Ptr(gitlab.DisabledAccessControl),
		MergeRequestsAccessLevel:         gitlab.Ptr(gitlab.DisabledAccessControl),
		ForkingAccessLevel:               gitlab.Ptr(gitlab.DisabledAccessControl),
		AnalyticsAccessLevel:             gitlab.Ptr(gitlab.DisabledAccessControl),
		SecurityAndComplianceAccessLevel: gitlab.Ptr(gitlab.DisabledAccessControl),
		WikiAccessLevel:                  gitlab.Ptr(gitlab.PrivateAccessControl),
		BuildsAccessLevel:                gitlab.Ptr(gitlab.DisabledAccessControl),
		SnippetsAccessLevel:              gitlab.Ptr(gitlab.DisabledAccessControl),
		ModelRegistryAccessLevel:         gitlab.Ptr(gitlab.DisabledAccessControl),
		ModelExperimentsAccessLevel:      gitlab.Ptr(gitlab.DisabledAccessControl),
		PagesAccessLevel:                 gitlab.Ptr(gitlab.DisabledAccessControl),
		MonitorAccessLevel:               gitlab.Ptr(gitlab.DisabledAccessControl),
		FeatureFlagsAccessLevel:          gitlab.Ptr(gitlab.DisabledAccessControl),
		InfrastructureAccessLevel:        gitlab.Ptr(gitlab.DisabledAccessControl),
		ReleasesAccessLevel:              gitlab.Ptr(gitlab.DisabledAccessControl),
		EmailsEnabled:                    gitlab.Ptr(false),
		ContainerRegistryAccessLevel:     gitlab.Ptr(gitlab.DisabledAccessControl),
		SharedRunnersEnabled:             gitlab.Ptr(false),
		LFSEnabled:                       gitlab.Ptr(false),
		RequestAccessEnabled:             gitlab.Ptr(false),
		PrintingMergeRequestLinkEnabled:  gitlab.Ptr(false),
		AutoDevopsEnabled:                gitlab.Ptr(false),
		PackagesEnabled:                  gitlab.Ptr(false),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Project %s (%s) created successfully\n", name, slug)
	fmt.Printf("Open: %s\n", p.WebURL)
	fmt.Printf("git branch -M main\n")
	fmt.Printf("git remote add origin %s\n", p.SSHURLToRepo)
	fmt.Printf("git push -u origin main\n")

	return nil
}
