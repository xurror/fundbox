package models

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

//type UUID uuid.UUID

func MarshalID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

func UnmarshalID(v interface{}) (uuid.UUID, error) {
	idAsString, ok := v.(string)
	if !ok {
		return uuid.Nil, errors.New("id should be a valid UUID")
	}
	return uuid.Parse(idAsString)
}

func MarshalInputID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

func UnmarshalInputID(v interface{}) (uuid.UUID, error) {
	idAsString, ok := v.(string)
	if !ok {
		return uuid.Nil, errors.New("id should be a valid UUID")
	}
	return uuid.Parse(idAsString)
}
