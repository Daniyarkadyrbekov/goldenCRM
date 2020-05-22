package aws

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestS3PutAndGetObject(t *testing.T) {
	storage := New()

	content := "hello world2!!"
	key := "firstfile"
	bucket := "testgoldencrm"

	require.NoError(t, storage.PutObject(bucket, key, content))

	gottenContend, err := storage.GetObject(bucket, key)
	require.NoError(t, err)
	require.Equal(t, content, string(gottenContend))
}
