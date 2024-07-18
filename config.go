package main

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/ini.v1"
)

type Config struct {
	Url   string
	Token string
}

func LoadConfig(profile string) (*Config, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	cfg, err := ini.Load(path.Join(userDir, ".gitlab-op/credentials"))
	if err != nil {
		return nil, err
	}

	section, err := cfg.GetSection(profile)
	if err != nil {
		return nil, err
	}

	if section == nil {
		return nil, fmt.Errorf("profile %s not found", profile)
	}

	return &Config{
		Url:   section.Key("url").String(),
		Token: section.Key("token").String(),
	}, nil
}
