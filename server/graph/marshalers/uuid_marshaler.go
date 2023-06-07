package marshallers

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
	"io"
)

func MarshalUUIDScalar(u uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(`"` + u.String() + `"`))
	})
}

func UnmarshalUUIDScalar(v interface{}) (uuid.UUID, error) {
	switch v := v.(type) {
	case string:
		return uuid.Parse(v)
	case uuid.UUID:
		return v, nil
	default:
		return uuid.New(), fmt.Errorf("%T is not a github.com/google/uuid.UUID", v)
	}
}
