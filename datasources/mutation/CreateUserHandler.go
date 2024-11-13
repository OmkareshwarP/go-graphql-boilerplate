package mutation

import (
	"context"
	"go-graphql-boilerplate/graph/model"
	"go-graphql-boilerplate/pkg/types"
	"go-graphql-boilerplate/pkg/utils"
)

func CreateUserHandler(ctx context.Context, firstName string, lastName string) (*model.CreateUserResponse, error) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	userID := utils.GenerateNanoIdWithLength(15)
	fullName := firstName + " " + lastName
	username := utils.RemoveSpacesAndSpecialChars(fullName) + "_" + userID

	if !utils.IsUsernameUnique(username) {
		return &model.CreateUserResponse{
			Error:              true,
			Message:            "username already exists",
			ErrorCodeForClient: "usernameAlreadyExists",
			StatusCode:         400,
			Data:               nil,
		}, nil
	}

	user := &types.User{
		UserID:    userID,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}
	utils.UsersData = append(utils.UsersData, user)

	return &model.CreateUserResponse{
		Error:              false,
		Message:            "User created successfully",
		ErrorCodeForClient: "",
		StatusCode:         200,
		Data: &model.UserID{
			UserID: userID,
		},
	}, nil
}
