package controllers

import (
	"fmt"
	"testing"

	_ "github.com/dezween/Calendar/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestFetchUsers(t *testing.T) {
	users, err := FetchUsers()

	assert.NoError(t, err, "Error fetching users")
	assert.NotEmpty(t, users, "No users fetched")

	fmt.Println("Полученные пользователи:")
	for _, user := range users {
		fmt.Printf("ID: %v, Name: %v\n", user.GID, user.Name)
	}
}

func TestFetchProjects(t *testing.T) {
	projects, err := FetchProjects()

	assert.NoError(t, err, "Error fetching projects")
	assert.NotEmpty(t, projects, "No projects fetched")

	fmt.Println("Полученные проекты:")
	for _, project := range projects {
		fmt.Printf("ID: %v, Name: %v\n", project.GID, project.Name)
	}
}

