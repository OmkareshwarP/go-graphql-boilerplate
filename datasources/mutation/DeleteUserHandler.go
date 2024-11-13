package mutation

import (
	"context"
	"go-graphql-boilerplate/graph/model"
	"go-graphql-boilerplate/pkg/utils"
)

func DeleteUserHandler(ctx context.Context, userID string) (*model.GenericResponse, error) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	for i, user := range utils.UsersData {
		if user.UserID == userID {
			utils.UsersData = append(utils.UsersData[:i], utils.UsersData[i+1:]...)
			return &model.GenericResponse{
				Error:              false,
				Message:            "User deleted successfully.",
				ErrorCodeForClient: "",
				StatusCode:         200,
				Data:               "done",
			}, nil
		}
	}

	return &model.GenericResponse{
		Error:              true,
		Message:            "User does not exists.",
		ErrorCodeForClient: "userNotFound",
		StatusCode:         404,
		Data:               "",
	}, nil
}
