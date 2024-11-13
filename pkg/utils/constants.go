package utils

import (
	"go-graphql-boilerplate/pkg/types"
	"regexp"
	"sync"
)

var (
	UsersData = []*types.User{} // in-memory store
	Mu        sync.Mutex        // protects usersData
)

// Helper function to check if a username is unique
func IsUsernameUnique(username string) bool {
	for _, user := range UsersData {
		if user.Username == username {
			return false
		}
	}
	return true
}

func RemoveSpacesAndSpecialChars(input string) string {
	// Create a regular expression to match non-alphanumeric characters
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	// Replace matched characters with an empty string
	return re.ReplaceAllString(input, "")
}
