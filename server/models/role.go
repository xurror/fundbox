package models

import (
	"fmt"
	"github.com/lib/pq"
	"io"
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

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (r *Role) UnmarshalGQL(v interface{}) error {
	role, ok := v.(string)
	if !ok {
		return fmt.Errorf("Role must be a string")
	}
	*r = Role(role)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (r Role) MarshalGQL(w io.Writer) {
	w.Write([]byte(`"` + r + `"`))
}
