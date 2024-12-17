package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/dezween/Calendar/internal/transport"
)

type AsanaUser struct {
	GID   string `json:"gid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AsanaProject struct {
	GID  string `json:"gid"`
	Name string `json:"name"`
}

func FetchUsers() ([]AsanaUser, error) {
	url := fmt.Sprintf("%s/users", "https://app.asana.com/api/1.0")
	resp, err := transport.MakeRequest(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []AsanaUser `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func FetchProjects() ([]AsanaProject, error) {
	url := fmt.Sprintf("%s/projects", "https://app.asana.com/api/1.0")
	resp, err := transport.MakeRequest(url)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data []AsanaProject `json:"data"`
	}
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}
