package model

import (
	"fmt"
	"github.com/lib/pq"
	"io"
	"log"
)

type Role string

const (
	Initiator   Role = "INITIATOR"
	Contributor      = "CONTRIBUTOR"
)

func (r *Role) String() pq.StringArray {
	array := pq.StringArray{string(*r)}
	return array
}

func FromString(role string) Role {
	switch role {
	case "INITIATOR":
		return Initiator
	case "CONTRIBUTOR":
		return Contributor
	default:
		log.Panic(fmt.Sprintf("Unknown role: %s", role))
		return ""
	}
}

func FromStringArray(roles []string) []Role {
	var convertedRoles []Role
	for _, role := range roles {
		convertedRoles = append(convertedRoles, FromString(role))
	}
	return convertedRoles
}

func ConvertToPQStringArray(roles []Role) pq.StringArray {
	var convertedRoles []string
	for _, role := range roles {
		convertedRoles = append(convertedRoles, string(role))
	}
	return convertedRoles
}

func ConvertToRoleArray(stringArray pq.StringArray) []Role {
	roles := []string(stringArray)
	convertedRoles := make([]Role, len(roles))
	for i, role := range roles {
		convertedRoles[i] = Role(role)
	}
	return convertedRoles
}

func (r *Role) UnmarshalGQL(v interface{}) error {
	role, ok := v.(string)
	if !ok {
		return fmt.Errorf("Role must be a string")
	}
	*r = Role(role)
	return nil
}

func (r Role) MarshalGQL(w io.Writer) {
	w.Write([]byte(`"` + r + `"`))
}
