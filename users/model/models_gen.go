// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateProjectInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Mutation struct {
}

type Project struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UUID        string `json:"uuid"`
}

type Query struct {
}

type UpdateProjectInput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type User struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type UserGet struct {
	ID    uint   `json:"ID"`
	Email string `json:"Email"`
	UUID  string `json:"UUID"`
}

type UserInput struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
