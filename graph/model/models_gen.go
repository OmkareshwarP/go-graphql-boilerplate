// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateUserResponse struct {
	Error              bool    `json:"error"`
	Message            string  `json:"message"`
	StatusCode         int     `json:"statusCode"`
	ErrorCodeForClient string  `json:"errorCodeForClient"`
	Data               *UserID `json:"data,omitempty"`
}

type GenericResponse struct {
	Error              bool   `json:"error"`
	Message            string `json:"message"`
	StatusCode         int    `json:"statusCode"`
	ErrorCodeForClient string `json:"errorCodeForClient"`
	Data               string `json:"data"`
}

type GetAllUsersResponse struct {
	Error              bool    `json:"error"`
	Message            string  `json:"message"`
	StatusCode         int     `json:"statusCode"`
	ErrorCodeForClient string  `json:"errorCodeForClient"`
	Data               []*User `json:"data,omitempty"`
}

type GetUserResponse struct {
	Error              bool   `json:"error"`
	Message            string `json:"message"`
	StatusCode         int    `json:"statusCode"`
	ErrorCodeForClient string `json:"errorCodeForClient"`
	Data               *User  `json:"data,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type User struct {
	UserID    string `json:"userId"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type UserID struct {
	UserID string `json:"userId"`
}
