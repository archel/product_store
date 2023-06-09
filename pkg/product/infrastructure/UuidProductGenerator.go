package product

import "github.com/google/uuid"

func UuidProductGenerator() (string, error) {
	return uuid.New().String(), nil
}
