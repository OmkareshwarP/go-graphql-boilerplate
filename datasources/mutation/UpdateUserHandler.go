package mutation

import (
	"context"
	"go-graphql-boilerplate/graph/model"
	"go-graphql-boilerplate/pkg/utils"
)

func UpdateUserHandler(ctx context.Context, userID string, username *string, firstName *string, lastName *string) (*model.GenericResponse, error) {
	utils.Mu.Lock()
	defer utils.Mu.Unlock()

	isUsernameExists := username != nil && len(*username) > 0
	isFirstNameExists := firstName != nil && len(*firstName) > 0
	isLastNameExists := lastName != nil && len(*lastName) > 0

	if !isUsernameExists && !isFirstNameExists && !isLastNameExists {
		return &model.GenericResponse{
			Error:              true,
			Message:            "Some information missing while validating your request",
			ErrorCodeForClient: "inputParamsValidationFailed",
			StatusCode:         403,
			Data:               "",
		}, nil
	}

	if isUsernameExists && !utils.IsUsernameUnique(*username) {
		return &model.GenericResponse{
			Error:              true,
			Message:            "username already exists",
			ErrorCodeForClient: "usernameAlreadyExists",
			StatusCode:         400,
			Data:               "",
		}, nil
	}

	for i, user := range utils.UsersData {
		if user.UserID == userID {
			_user := utils.UsersData[i]
			if isUsernameExists {
				_user.Username = *username
			}
			if isFirstNameExists {
				_user.FirstName = *firstName
			}
			if isLastNameExists {
				_user.LastName = *lastName
			}
			return &model.GenericResponse{
				Error:              false,
				Message:            "User updated successfully.",
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
