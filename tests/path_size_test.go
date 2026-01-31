package tests

import (
	"code/pkg/du"
	"testing"

	"github.com/c2fo/testify/require"
)

func TestGetPathSize_File(t *testing.T) {
	size, err := du.GetSize("../testdata/8.bin")
	require.NoError(t, err)
	require.Equal(t, int64(8), size)
}

func TestGetPathSize_Directory(t *testing.T) {
	size, err := du.GetSize("../testdata")
	require.NoError(t, err)
	require.Equal(t, int64(120), size)
}
