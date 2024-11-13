package query

import (
	"context"
	"go-graphql-boilerplate/graph/model"
	"go-graphql-boilerplate/pkg/utils"
)

func GetUserByUsernameHandler(ctx context.Context, username string) (*model.GetUserResponse, error) {
	usersData := utils.UsersData

	for _, user := range usersData {
		if user.Username == username {
			return &model.GetUserResponse{
				Error:              false,
				Message:            "User fetched successfully",
				ErrorCodeForClient: "",
				StatusCode:         200,
				Data: &model.User{
					UserID:    user.UserID,
					Username:  user.Username,
					FirstName: user.FirstName,
					LastName:  user.LastName,
				},
			}, nil
		}
	}

	return &model.GetUserResponse{
		Error:              true,
		Message:            "User does not exists.",
		ErrorCodeForClient: "userNotFound",
		StatusCode:         404,
		Data:               nil,
	}, nil
}
