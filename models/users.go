package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/http/httptrace"

	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"open-btm.com/configs"
	"open-btm.com/users/model"
)

var AdminAccessToken string

// createHTTPClient sets up an HTTP client with OpenTelemetry instrumentation
func createHTTPClient() *http.Client {
	return &http.Client{
		Transport: otelhttp.NewTransport(
			http.DefaultTransport,
			otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
				return otelhttptrace.NewClientTrace(ctx)
			}),
		),
	}
}

func CreateUser(ctx context.Context, user model.User) (*model.UserGet, error) {
	//  define var for response
	var response_user model.UserGet

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	postDataBytes, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%v/user", url), bytes.NewReader(postDataBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return nil, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parsing response body
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return nil, err
	}

	user_map := responseMap["data"].(map[string]interface{})

	// Marshal map to JSON
	jsonData, err := json.Marshal(user_map)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &response_user)
	if err != nil {
		return nil, err
	}

	return &response_user, nil
}

func GetUsers(ctx context.Context, page, size uint) ([]model.UserGet, error) {
	//  define var for response
	var response struct {
		Data []model.UserGet `json:"data"`
	}

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	app_uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/appusers?page=%v&size=%v&app_uuid=%v", url, page, size, app_uuid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return nil, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parsing response body
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return nil, err
	}

	// Marshal map to JSON
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func GetUser(ctx context.Context, user_id uint) (*model.UserGet, error) {
	//  define var for response
	var response struct {
		Data model.UserGet `json:"data"`
	}

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	app_uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/appuser/%v?app_uuid=%v", url, user_id, app_uuid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return nil, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parsing response body
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return nil, err
	}

	// Marshal map to JSON
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}

func UpdateUser(ctx context.Context, user model.UserUdateInput, user_id int) (bool, error) {

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	putDataBytes, err := json.Marshal(user)
	if err != nil {
		return false, err
	}

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "PATCH", fmt.Sprintf("%v/user/%v", url, user_id), bytes.NewReader(putDataBytes))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func ResetPasswordUser(ctx context.Context, password string, email string) (bool, error) {

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()
	user := passData{
		Email:    email,
		Password: password,
	}

	putDataBytes, err := json.Marshal(user)
	if err != nil {
		return false, err
	}

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("%v/user", url), bytes.NewReader(putDataBytes))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}
	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func DeleteUser(ctx context.Context, user_id uint) (bool, error) {
	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	app_uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%v/appuser/%v?app_uuid=%v", url, user_id, app_uuid), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func CheckUser(ctx context.Context, uuid string) (uint, error) {
	//  define var for response
	var response struct {
		Data model.UserGet `json:"data"`
	}

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/useruuid?uuid=%v", url, uuid), nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return 0, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// parsing response body
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return 0, err
	}

	// Marshal map to JSON
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		return 0, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return 0, err
	}

	//return user if parsed with no error
	return response.Data.ID, nil
}

func ActivateDeactivateUser(ctx context.Context, user_id uint, status bool) (bool, error) {
	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("%v/user/%v?status=%v", url, user_id, status), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func AddRoleToUser(ctx context.Context, role_id, user_id int) (bool, error) {
	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	app_uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%v/approleuser/%v/%v?app_uuid=%v", url, role_id, user_id, app_uuid), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func RemoveRoleFromUser(ctx context.Context, role_id, user_id int) (bool, error) {
	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	app_uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%v/approleuser/%v/%v?app_uuid=%v", url, role_id, user_id, app_uuid), nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return false, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("Error at IAM: %v\n", resp.Status)
}

func GetAppRoles(ctx context.Context) ([]model.Role, error) {
	//  define var for response
	var response struct {
		Data []model.Role `json:"data"`
	}

	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	uuid := configs.AppConfig.Get("BLUE_ADMIN_UUID")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Build and send the request
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/appruid/%v", url, uuid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if AdminAccessToken == "" {
		if _, ok := LoginBlueAdmin(); ok != true {
			return nil, fmt.Errorf("Error logging in to IAM please correct your config credntial files")
		}
	}

	//  Set Token Header
	req.Header.Set("X-APP-TOKEN", AdminAccessToken)

	//  Make request to IAM
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// parsing response body
	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return nil, err
	}

	// Marshal map to JSON
	jsonData, err := json.Marshal(responseMap)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON to struct
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

type PostData struct {
	GrantType string `json:"grant_type"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}
type passData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginBlueAdmin() (string, bool) {
	// Retrieve configuration values
	url := configs.AppConfig.Get("BLUE_ADMIN_URI")
	email := configs.AppConfig.Get("BLUE_ADMIN_USER")
	password := configs.AppConfig.Get("BLUE_ADMIN_PASSWORD")

	// Create an HTTP client with OpenTelemetry middleware
	client := createHTTPClient()

	// Prepare POST request data
	postData := PostData{
		GrantType: "authorization_code",
		Email:     email,
		Password:  password,
		Token:     "token",
	}

	// change post data to json strings
	postDataBytes, err := json.Marshal(postData)
	if err != nil {
		return "Error marshalling request data", false
	}

	// creating blank context for Loging in
	ctx := context.Background()

	// Build and send the request
	//
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%v/login", url), bytes.NewReader(postDataBytes))
	if err != nil {
		return "Error creating request", false
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("#########: %v\n", err)
		return "Error executing request", false
	}
	defer resp.Body.Close()

	// Read and unmarshal response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body", false
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		return "Error unmarshalling response", false
	}

	accessToken, ok := responseMap["data"].(map[string]interface{})["access_token"].(string)
	if !ok {
		return "Access token not found in response", false
	}

	AdminAccessToken = accessToken
	return accessToken, true
}
