package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/PYTNAG/iasked/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, username sql.NullString) *User {
	hash, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)

	params := CreateUserParams{
		Email:    util.RandomEmail(),
		Username: username,
		Hash:     hash,
	}

	userInfo, err := testQueries.CreateUser(context.Background(), params)
	require.NoError(t, err)

	return &User{
		ID:       userInfo.ID,
		Email:    userInfo.Email,
		Username: userInfo.Username,
		Hash:     hash,
	}
}

func deleteRandomUser(t *testing.T, userId int32) {
	err := testQueries.DeleteUser(context.Background(), userId)

	require.NoError(t, err)
}

func TestCreateUserWithUsername(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	deleteRandomUser(t, createdUser.ID)
}

func TestDeleteUser(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	deleteRandomUser(t, createdUser.ID)

	userInfo, err := testQueries.GetUserInfo(context.Background(), createdUser.ID)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, userInfo)
}

func TestGetUserHash(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	defer deleteRandomUser(t, createdUser.ID)

	userHash, err := testQueries.GetUserHash(context.Background(), createdUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, userHash)

	require.Equal(t, createdUser.Hash, userHash)
}

func TestGetUserInfo(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	defer deleteRandomUser(t, createdUser.ID)

	userInfo, err := testQueries.GetUserInfo(context.Background(), createdUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, userInfo)

	require.Equal(t, createdUser.Email, userInfo.Email)
	require.Equal(t, createdUser.Username, userInfo.Username)
}

func TestUpdateEmail(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	defer deleteRandomUser(t, createdUser.ID)

	params := UpdateEmailParams{
		NewEmail: util.RandomEmail(),
		UserID:   createdUser.ID,
	}

	err := testQueries.UpdateEmail(context.Background(), params)

	require.NoError(t, err)

	actualUserInfo, err := testQueries.GetUserInfo(context.Background(), createdUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualUserInfo)

	require.Equal(t, params.NewEmail, actualUserInfo.Email)
}

func TestUpdateHash(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	defer deleteRandomUser(t, createdUser.ID)

	expectedHash, err := util.HashPassword(util.RandomPassword())
	require.NoError(t, err)

	params := UpdateHashParams{
		NewHash: expectedHash,
		UserID:  createdUser.ID,
	}

	err = testQueries.UpdateHash(context.Background(), params)

	require.NoError(t, err)

	actualHash, err := testQueries.GetUserHash(context.Background(), createdUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualHash)

	require.Equal(t, expectedHash, actualHash)
}

func TestUpdateUsername(t *testing.T) {
	username := util.RandomUsername()
	createdUser := createRandomUser(t, sql.NullString{
		String: username,
		Valid:  true,
	})

	defer deleteRandomUser(t, createdUser.ID)

	params := UpdateUsernameParams{
		NewUsername: util.RandomUsername(),
		UserID:      createdUser.ID,
	}

	err := testQueries.UpdateUsername(context.Background(), params)

	require.NoError(t, err)

	actualUserInfo, err := testQueries.GetUserInfo(context.Background(), createdUser.ID)

	require.NoError(t, err)
	require.NotEmpty(t, actualUserInfo)

	require.Equal(t, params.NewUsername, actualUserInfo.Username)
}
