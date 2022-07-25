//go:build integration

package db

import (
	"context"
	"simple-api/internal/comment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create and getById comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Slug",
			Author: "Author",
			Body:   "Body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetCommentById(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Slug",
			Author: "Author",
			Body:   "Body",
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		cmt, err = db.GetCommentById(context.Background(), cmt.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.CreateComment(context.Background(), comment.Comment{
			Slug:   "Slug",
			Author: "Author",
			Body:   "Body",
		})
		assert.NoError(t, err)

		testSlug := "diferent slug"

		_, err = db.UpdateComment(context.Background(), cmt.ID, comment.Comment{
			Slug: testSlug,
		})
		assert.NoError(t, err)

		cmt, err = db.GetCommentById(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, testSlug, cmt.Slug)
	})
}
