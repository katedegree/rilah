// infrastructure/graphql/scalars.go

package graphql

import (
	"fmt"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// MarshalUintID はGoのuint32 → GraphQLのstringに変換します
func MarshalUintID(u uint32) graphql.Marshaler {
	return graphql.MarshalString(strconv.FormatUint(uint64(u), 10))
}

// UnmarshalUintID はGraphQLのstring → Goのuint32に変換します
func UnmarshalUintID(v interface{}) (uint32, error) {
	s, err := graphql.UnmarshalString(v)
	if err != nil {
		return 0, err
	}
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid UintID: %v", err)
	}
	return uint32(n), nil
}
