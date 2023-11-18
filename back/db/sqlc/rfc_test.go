package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/PYTNAG/iasked/util"
	"github.com/stretchr/testify/require"
)

func createRandomRFC(t *testing.T, userId sql.NullInt32) *Rfc {
	params := CreateRFCParams{
		AuthorID: userId,
		Message:  util.RandomMessage(),
	}

	rfcId, err := testQueries.CreateRFC(context.Background(), params)
	require.NoError(t, err)

	return &Rfc{
		ID:        rfcId,
		AuthorID:  userId,
		Message:   params.Message,
		CreatedAt: time.Now(),
		Archived:  false,
	}
}

func deleteRandomRFC(t *testing.T, rfcId int32) {
	err := testQueries.DeleteRFC(context.Background(), rfcId)
	require.NoError(t, err)
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

	rfc := createRandomRFC(t, sql.NullInt32{
		Int32: user.ID,
		Valid: true,
	})

	deleteRandomRFC(t, rfc.ID)
	deleteRandomUser(t, user.ID)
}

func TestGetLastRFCs(t *testing.T) {
	user := createRandomUser(t, sql.NullString{
		String: util.RandomUsername(),
		Valid:  true,
	})

	defer deleteRandomUser(t, user.ID)

	rfcCount := 5
	offset := 1
	limit := 3

	rfcs := make([]*Rfc, rfcCount)

	for i := 0; i < rfcCount; i++ {
		rfcs[i] = createRandomRFC(t, sql.NullInt32{
			Int32: user.ID,
			Valid: true,
		})
		defer deleteRandomRFC(t, rfcs[i].ID)
	}

	params := GetLastRFCsParams{
		Offset: int32(offset),
		Count:  int32(limit),
	}

	actualRfcs, err := testQueries.GetLastRFCs(context.Background(), params)
	require.NoError(t, err)

	for i := offset; i < offset+limit; i++ {
		compareRfcs(t, rfcs[rfcCount-i-1], &actualRfcs[i-offset])
	}
}

func TestDeleteRFC(t *testing.T) {
	user := createRandomUser(t, sql.NullString{
		String: util.RandomUsername(),
		Valid:  true,
	})

	defer deleteRandomUser(t, user.ID)

	olderRfc := createRandomRFC(t, sql.NullInt32{
		Int32: user.ID,
		Valid: true,
	})

	defer deleteRandomRFC(t, olderRfc.ID)

	rfc := createRandomRFC(t, sql.NullInt32{
		Int32: user.ID,
		Valid: true,
	})

	deleteRandomRFC(t, rfc.ID)

	params := GetLastRFCsParams{
		Offset: 0,
		Count:  1,
	}

	lastRFCs, err := testQueries.GetLastRFCs(context.Background(), params)

	require.NoError(t, err)
	require.NotEqual(t, rfc.ID, lastRFCs[0].ID)
}
