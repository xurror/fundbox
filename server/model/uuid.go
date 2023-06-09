package model

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func MarshalUUID(id uuid.UUID) graphql.Marshaler {
	return graphql.MarshalString(id.String())
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	idAsString, ok := v.(string)
	if !ok {
		return uuid.Nil, errors.New("id should be a valid UUID")
	}
	return uuid.Parse(idAsString)
}
