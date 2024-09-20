package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"open-btm.com/models"
)

type GraphQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

func TestGraphQLAPI(t *testing.T) {
	// creating database for test
	models.InitDatabase()
	defer models.CleanDatabase()
	setupTestApp()

	// Define test cases
	testsSprint := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Sprint",
			description: "Create Sprint 1",
			query: `mutation($input: CreateSprintInput!) {
				createsprint(input: $input) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":        "Sprint 1",
					"description": "First sprint",
				},
			},
			expected: 200,
		},
		{
			name:        "Create Sprint 2",
			description: "Create Sprint 2",
			query: `mutation($input: CreateSprintInput!) {
				createsprint(input: $input) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":        "Sprint 2",
					"description": "Second sprint",
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Sprints",
			description: "Get Sprint 2",
			query: `query($page: Int!, $size: Int!) {
				sprints(page: $page, size: $size) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200, // Adjust this based on your expected data
		},
		// Add more tests for other queries and mutations...
	}
	// Define test cases
	testsProject := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Project",
			description: "Create Project 1",
			query: `mutation createproject ($input: CreateProjectInput!) {
				    createproject (input: $input) {
				        id
				        name
				        description
				        uuid
				    }
				}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":        "Project One",
					"description": "First Project",
				},
			},
			expected: 200,
		},
	}

	// first creating the projects
	for _, tt := range testsProject {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/admin", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			TestApp.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// testing creating sprint
	for _, tt := range testsSprint {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			TestApp.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

}
