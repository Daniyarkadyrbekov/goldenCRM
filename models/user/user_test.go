package user

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserConstructor(t *testing.T) {

	sep := "_"
	generatedIDs := make(map[int64]struct{})

	for i := 0; i < 1000; i++ {
		iterStr := strconv.Itoa(i)
		u := New(iterStr, sep+iterStr)

		_, ok := generatedIDs[u.id]
		require.False(t, ok)
		generatedIDs[u.id] = struct{}{}

		require.Equal(t, iterStr+sep+iterStr, u.GetFullName())
	}
}
