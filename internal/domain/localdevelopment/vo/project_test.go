package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProject_GetServiceLabels(t *testing.T) {
	project := Project{
		Services: []Service{
			{Label: "service1"},
			{Label: "service2"},
			{Label: "service3"},
		},
	}

	expectedLabels := []string{"service1", "service2", "service3"}
	labels := project.GetServiceLabels()

	assert.Equal(t, expectedLabels, labels, "they should be equal")
}

func TestProject_FindServiceByLabel_Found(t *testing.T) {
	project := Project{
		Services: []Service{
			{Label: "service1"},
			{Label: "service2"},
		},
	}

	service, err := project.FindServiceByLabel("service2")

	assert.NoError(t, err)
	assert.Equal(t, "service2", service.Label)
}

func TestProject_FindServiceByLabel_NotFound(t *testing.T) {
	project := Project{
		Services: []Service{
			{Label: "service1"},
			{Label: "service3"},
		},
	}

	_, err := project.FindServiceByLabel("service2")

	assert.Error(t, err)
	assert.EqualError(t, err, "service with label service2 not found")
}
