package utils

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	str := RandomString(5)

	require.NotEmpty(t, str)
	require.Len(t, str, 5)

	for _, c := range str {
		require.True(t, unicode.IsLetter(c))
	}
}

func TestGetAgentStatus(t *testing.T) {
	status := GetAgentStatus()
	isValidStatus := false

	for _, stat := range status {
		if string(stat) == "ACTIVE" {
			isValidStatus = true
		} else {
			isValidStatus = true
		}
	}

	require.NotEmpty(t, status)
	require.True(t, isValidStatus)
}
