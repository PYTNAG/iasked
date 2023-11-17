package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/PYTNAG/iasked/util"
	"github.com/stretchr/testify/require"
)

func createRandomRFC(t *testing.T, user *User) *Rfc {
	authorId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	if user != nil {
		authorId.Int32 = user.ID
		authorId.Valid = true
	}

	params := CreateRFCParams{
		AuthorID: authorId,
		Message:  util.RandomMessage(),
	}

	rfcId, err := testQueries.CreateRFC(context.Background(), params)
	require.NoError(t, err)

	return &Rfc{
		ID:        rfcId,
		AuthorID:  authorId,
		Message:   params.Message,
		CreatedAt: time.Now(),
		Archived:  false,
	}
}

func compareRfcs(t *testing.T, expected, actual *Rfc) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.AuthorID, actual.AuthorID)
	require.Equal(t, expected.Message, actual.Message)
	require.Equal(t, expected.Archived, actual.Archived)

	require.True(t, actual.CreatedAt.Add(time.Minute).After(expected.CreatedAt))
}

func TestCreateRFC(t *testing.T) {
	user := createRandomUser(t, sql.NullString{
		String: util.RandomUsername(),
		Valid:  true,
	})

	rfc := createRandomRFC(t, user)

	params := GetLastRFCsParams{
		Limit:  1,
		Offset: 0,
	}

	lastRfcs, err := testQueries.GetLastRFCs(context.Background(), params)

	require.NoError(t, err)

	compareRfcs(t, rfc, &lastRfcs[0])

	deleteRandomUser(t, user.ID)
}

func TestGetLastRFCs(t *testing.T) {
	user := createRandomUser(t, sql.NullString{
		String: util.RandomUsername(),
		Valid:  true,
	})

	rfcCount := 5
	offset := 1
	limit := 3

	rfcs := make([]Rfc, rfcCount)

	for i := 0; i < rfcCount; i++ {
		rfcs[i] = *createRandomRFC(t, user)
	}

	params := GetLastRFCsParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	}

	actualRfcs, err := testQueries.GetLastRFCs(context.Background(), params)

	require.NoError(t, err)

	for i := offset; i < offset+limit; i++ {
		compareRfcs(t, &rfcs[rfcCount-i-1], &actualRfcs[i-offset])
	}

	deleteRandomUser(t, user.ID)
}

func TestDeleteRFC(t *testing.T) {
	user := createRandomUser(t, sql.NullString{
		String: util.RandomUsername(),
		Valid:  true,
	})

	rfc := createRandomRFC(t, user)

	err := testQueries.DeleteRFC(context.Background(), rfc.ID)

	require.NoError(t, err)

	params := GetLastRFCsParams{
		Limit:  1,
		Offset: 0,
	}

	lastRFCs, err := testQueries.GetLastRFCs(context.Background(), params)

	if len(lastRFCs) == 0 {
		require.EqualError(t, err, sql.ErrNoRows.Error())
	} else {
		require.NotEqual(t, rfc.ID, lastRFCs[0].ID)
	}

	deleteRandomUser(t, user.ID)
}
