package validate

import "github.com/google/uuid"

func GenerateID() string {
	return uuid.NewString()
}

func CheckID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidID
	}
	return nil
}
