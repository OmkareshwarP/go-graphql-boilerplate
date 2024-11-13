package query

import (
	"context"
	"go-graphql-boilerplate/graph/model"
	"go-graphql-boilerplate/pkg/utils"
)

func GetAllUsersHandler(ctx context.Context) (*model.GetAllUsersResponse, error) {
	usersData := utils.UsersData

	var users []*model.User

	for _, user := range usersData {
		users = append(users, &model.User{
			UserID:    user.UserID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}

	return &model.GetAllUsersResponse{
		Error:              false,
		Message:            "Users fetched successfully",
		ErrorCodeForClient: "",
		StatusCode:         200,
		Data:               users,
	}, nil
}
