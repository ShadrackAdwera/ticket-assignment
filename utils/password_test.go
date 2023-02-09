package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := RandomString(10)
	hashedPw, err := HashPassword(password)
	hashedPw1, err1 := HashPassword(password)

	require.NoError(t, err)
	require.NoError(t, err1)
	require.NotEmpty(t, hashedPw)
	require.NotEqual(t, password, hashedPw)
	require.NotEqual(t, hashedPw, hashedPw1)

	err = ComparePasswords(hashedPw, password)
	require.NoError(t, err)
	require.Nil(t, err)

	wrongPassword := RandomString(9)

	err = ComparePasswords(hashedPw, wrongPassword)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
