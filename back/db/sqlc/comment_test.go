package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T, authorId sql.NullInt32, rfcId int32) *Comment {
	params := CreateCommentParams{
		AuthorID: authorId,
		RfcID:    rfcId,
	}

	commentId, err := testQueries.CreateComment(context.Background(), params)
	require.NoError(t, err)

	return &Comment{
		ID:        commentId,
		AuthorID:  authorId,
		RfcID:     rfcId,
		CreatedAt: time.Now(),
	}
}

func compareComments(t *testing.T, expected, actual *Comment) {
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.AuthorID, actual.AuthorID)
	require.Equal(t, expected.RfcID, actual.RfcID)

	// expected.CreatedAt initialized after actual.CreatedAt
	require.True(t, actual.CreatedAt.Add(time.Minute).After(expected.CreatedAt))
}

func TestCreateComment(t *testing.T) {
	nullableUserId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	rfc := createRandomRFC(t, nullableUserId)
	defer deleteRandomRFC(t, rfc.ID)

	_ = createRandomComment(t, nullableUserId, rfc.ID)

	// there is no deleteRandomComment because all comments will be deleted with deleting rfc
}

func TestGetRFCComments(t *testing.T) {
	nullableUserId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	rfc := createRandomRFC(t, nullableUserId)
	defer deleteRandomRFC(t, rfc.ID)

	commentsCount := 5
	offset := 1
	countToGet := 3

	require.GreaterOrEqual(t, commentsCount, offset+countToGet)

	comments := make([]*Comment, commentsCount)

	for i := 0; i < commentsCount; i++ {
		comments[i] = createRandomComment(t, nullableUserId, rfc.ID)
	}

	params := GetRFCCommentsParams{
		RfcID:  rfc.ID,
		Offset: int32(offset),
		Count:  int32(countToGet),
	}

	actualComments, err := testQueries.GetRFCComments(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, actualComments)

	for i := offset; i < offset+countToGet; i++ {
		compareComments(t, comments[i], &actualComments[i-offset])
	}
}

func TestCountRFCComments(t *testing.T) {
	nullableUserId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	rfc := createRandomRFC(t, nullableUserId)
	defer deleteRandomRFC(t, rfc.ID)

	commentsCount := 5

	for i := 0; i < commentsCount; i++ {
		createRandomComment(t, nullableUserId, rfc.ID)
	}

	actualCommentsCount, err := testQueries.CountRFCComments(context.Background(), rfc.ID)

	require.NoError(t, err)
	require.EqualValues(t, commentsCount, actualCommentsCount)
}

func TestDeleteComment(t *testing.T) {
	nullableUserId := sql.NullInt32{
		Int32: 0,
		Valid: false,
	}

	rfc := createRandomRFC(t, nullableUserId)
	defer deleteRandomRFC(t, rfc.ID)

	commentsCount := 5

	for i := 0; i < commentsCount; i++ {
		createRandomComment(t, nullableUserId, rfc.ID)
	}

	comment := createRandomComment(t, nullableUserId, rfc.ID)

	err := testQueries.DeleteComment(context.Background(), comment.ID)
	require.NoError(t, err)

	actualCommentsCount, err := testQueries.CountRFCComments(context.Background(), rfc.ID)

	require.NoError(t, err)
	require.EqualValues(t, commentsCount, actualCommentsCount)
}
