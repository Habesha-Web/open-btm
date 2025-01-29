package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm/clause"
	"open-btm.com/auth"
	"open-btm.com/database"
	"open-btm.com/graph"
	"open-btm.com/models"
	"open-btm.com/observe"
)

// ##################################################
// Helper function for test
// ##################################################
func otelechospanstarter(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		routeName := ctx.Path() + "_" + strings.ToLower(ctx.Request().Method)
		tracer, span := observe.EchoAppSpanner(ctx, fmt.Sprintf("%v-root", routeName))
		ctx.Set("tracer", &observe.RouteTracer{Tracer: tracer, Span: span})

		// Process request
		err := next(ctx)
		if err != nil {
			return err
		}

		span.End()
		return nil
	}
}

func nextAuthValidator(key string, ctx echo.Context) (bool, error) {
	// Print the current route's path (URL pattern)
	// fmt.Println("Route path:", ctx.Path())

	if key != "login" {

		// fmt.Println("Role required: ", models.EndpointJSON[ctx.Path()])
		// Parse JWT token
		usr_claim, _ := models.ParseJWTToken(key)

		// Print user roles (assuming usr_claim.Roles is a field in your JWT claim structure)
		fmt.Println("User Roles:", usr_claim.Roles)
	}

	return true, nil
}

func dbsessioninjection(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		db, err := database.ReturnSession()
		if err != nil {
			return err
		}
		ctx.Set("db", db)

		nerr := next(ctx)
		if nerr != nil {
			return nerr
		}

		return nil
	}
}

func setupRoutes(app *echo.Echo) {

	// the Otel spanner middleware
	app.Use(otelechospanstarter)

	// Authentication middleware
	app.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		KeyLookup: "header:x-app-token",
		Validator: nextAuthValidator,
	}))

	app.Use(middleware.BodyDump(func(ctx echo.Context, reqBody, resBody []byte) {
		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)
		tracer.Span.SetAttributes(attribute.String("request", string(reqBody)))
		tracer.Span.SetAttributes(attribute.String("response", string(resBody)))
	}))

	// db session injection
	app.Use(dbsessioninjection)

	// hello world add
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}).Name = "hello_world"

	// creating api group
	gapp := app.Group("/api/v1")

	// login route mounting
	gapp.POST("/login", auth.LoginUser).Name = "login_btm"

	// project directive cfg
	cfg := graph.Config{
		Directives: graph.DirectiveRoot{
			HasRole: graph.HasRoleDirective,
		},
	}

	//  This is GraphQL project graphql route
	gapp.POST("/project/:project_id", func(ctx echo.Context) error {

		// Parsing project ID
		project_id, err := strconv.Atoi(ctx.Param("project_id"))
		if err != nil {
			return err
		}

		//  Getting Project Database Session
		project_db, err := database.ReturnSession()
		if err != nil {
			fmt.Println(err)
			return nil
		}

		// Getting Project database Name
		var project models.Project
		if res := project_db.Model(&models.Project{}).Preload(clause.Associations).Where("id = ?", uint(project_id)).First(&project); res.Error != nil {
			fmt.Println(res.Error)
			return res.Error
		}

		//  Connecting to Databse
		db, err := database.ReturnSessionDatabase(project.DatabaseName)
		if err != nil {
			log.Errorf("Error Connecting to Database: %v\n", err)
			return err
		}

		//  Geting tracer
		tracer := ctx.Get("tracer").(*observe.RouteTracer)

		// Providing reslover trancer and database connection
		cfg.Resolvers = &graph.Resolver{DB: db, Tracer: tracer}

		//  Schema handler
		graphqlHandler := handler.NewDefaultServer(
			graph.NewExecutableSchema(cfg),
		)
		graphqlHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	}).Name = "btm_project"

}

// ##################################################
// Helper function for test
// ##################################################

var (
	TestApp   *echo.Echo
	groupPath = "/api/v1"
)

func setupTestApp() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}
	TestApp = echo.New()
	setupRoutes(TestApp)
}

type GraphQLRequest struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

func TestGraphQLAPI(t *testing.T) {
	// creating database for test
	models.InitDatabase()
	defer models.CleanDatabase()
	setupTestApp()

	// Define test cases for sprints
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
					"status":      "Open",
					"start_date":  "2023-01-01",
					"duration":    14,
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Sprints",
			description: "Get Sprints",
			query: `query($page: Int!, $size: Int!) {
				sprints(page: $page, size: $size) {
					total
					sprints {
						id
						name
						description
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for requirements
	testsRequirement := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Requirement",
			description: "Create Requirement 1",
			query: `mutation($input: CreateRequirementInput!) {
				createrequirement(input: $input) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":            "Requirement 1",
					"description":     "First requirement",
					"purpose":         "Purpose 1",
					"status":          "Open",
					"bussiness_value": 10,
					"assigned_to":     "User 1",
					"sprint_id":       1,
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Requirements",
			description: "Get Requirements",
			query: `query($page: Int!, $size: Int!) {
				requirements(page: $page, size: $size) {
					total
					requirements {
						id
						name
						description
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for tests
	testsTest := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Test",
			description: "Create Test 1",
			query: `mutation($input: CreateTestInput!) {
				createtest(input: $input) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":            "Test 1",
					"description":     "First test",
					"steps":           "Step 1",
					"expected_result": "Result 1",
					"requirement_id":  1,
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Tests",
			description: "Get Tests",
			query: `query($page: Int!, $size: Int!) {
				tests(page: $page, size: $size) {
					total
					tests {
						id
						name
						description
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for testsets
	testsTestset := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Testset",
			description: "Create Testset 1",
			query: `mutation($input: CreateTestsetInput!) {
				createtestset(input: $input) {
					id
					name
					description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":        "Testset 1",
					"description": "First testset",
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Testsets",
			description: "Get Testsets",
			query: `query($page: Int!, $size: Int!) {
				testsets(page: $page, size: $size) {
					total
					testsets {
						id
						name
						description
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for testinstances
	testsTestInstance := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create TestInstance",
			description: "Create TestInstance 1",
			query: `mutation($input: CreateTestInstanceInput!) {
				createtestinstance(input: $input) {
					id
					test_id
					testset_id
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"test_id":    1,
					"testset_id": 1,
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve TestInstances",
			description: "Get TestInstances",
			query: `query($page: Int!, $size: Int!) {
				testinstances(page: $page, size: $size) {
					total
					test_instances {
						id
						test_id
						testset_id
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for testruns
	testsTestRun := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create TestRun",
			description: "Create TestRun 1",
			query: `mutation($input: CreateTestRunInput!) {
				createtestrun(input: $input) {
					id
					run_status
					result
					sevierity
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"run_status": "Passed",
					"result":     "Result 1",
					"sevierity":  "High",
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve TestRuns",
			description: "Get TestRuns",
			query: `query($page: Int!, $size: Int!) {
				testruns(page: $page, size: $size) {
					total
					test_runs {
						id
						run_status
						result
						sevierity
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for issues
	testsIssue := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Issue",
			description: "Create Issue 1",
			query: `mutation($input: CreateIssueInput!) {
				createissue(input: $input) {
					id
					issue_name
					issue_status
					issue_description
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"issue_name":        "Issue 1",
					"issue_status":      "Open",
					"issue_description": "First issue",
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Issues",
			description: "Get Issues",
			query: `query($page: Int!, $size: Int!) {
				issues(page: $page, size: $size) {
					total
					issues {
						id
						issue_name
						issue_status
						issue_description
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	// Define test cases for documents
	testsDocument := []struct {
		name        string
		query       string
		description string
		variables   interface{}
		expected    int
	}{
		{
			name:        "Create Document",
			description: "Create Document 1",
			query: `mutation($input: CreateDocumentInput!) {
				createdocument(input: $input) {
					id
					name
					file_url
				}
			}`,
			variables: map[string]interface{}{
				"input": map[string]interface{}{
					"name":     "Document 1",
					"file_url": "http://example.com/doc1",
				},
			},
			expected: 200,
		},
		{
			name:        "Retrieve Documents",
			description: "Get Documents",
			query: `query($page: Int!, $size: Int!) {
				documents(page: $page, size: $size) {
					total
					documents {
						id
						name
						file_url
					}
				}
			}`,
			variables: map[string]interface{}{
				"page": 1,
				"size": 10,
			},
			expected: 200,
		},
	}

	test_app := TestApp
	// Run test cases for sprints
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
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for requirements
	for _, tt := range testsRequirement {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for tests
	for _, tt := range testsTest {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for testsets
	for _, tt := range testsTestset {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for testinstances
	for _, tt := range testsTestInstance {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for testruns
	for _, tt := range testsTestRun {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for issues
	for _, tt := range testsIssue {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}

	// Run test cases for documents
	for _, tt := range testsDocument {
		t.Run(tt.name, func(t *testing.T) {
			body := GraphQLRequest{
				Query:     tt.query,
				Variables: tt.variables,
			}
			jsonBody, _ := json.Marshal(body)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/project/1", bytes.NewBuffer(jsonBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			test_app.ServeHTTP(rec, req)

			assert.Equalf(t, tt.expected, rec.Result().StatusCode, tt.description)
		})
	}
}
