package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCnNum(t *testing.T) {
	require.Equal(t, "一", CnNum(1))
}
