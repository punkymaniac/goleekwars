package apileek

import (
)

type forumService struct {
    apiService
}

// Search post on forum
// TODO Incomplete, missing documentation
// Always return a error
func (s *forumService) Get(
    query string, // Query to search
    farmer string, // TODO Missing documentation
    category uint, // TODO Missing documentation
    page uint, // TODO Missing documentation
    order string, // TODO Missing documentation
    admin bool, // TODO Missing documentation
    moderator bool, // TODO Missing documentation
) error {
    return newError("Api not implemented: Missing documentation")
}

