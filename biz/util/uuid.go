package util

import (
	googleUUid "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() string {
	return uuid.NewV4().String()
}

func GenerateUID() int64 {
	return int64(googleUUid.New().ID())
}
