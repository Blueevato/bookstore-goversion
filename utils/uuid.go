package utils

import "github.com/google/uuid"

func CreatUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "can't create UUID", err
	}
	return u.String(), nil
}
