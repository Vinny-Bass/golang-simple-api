package comment

import (
	"context"
	"testing"
)

var fakeData = []Comment{
	{
		ID:     "test-id",
		Slug:   "test-slug",
		Body:   "test-body",
		Author: "test-author",
	},
}

type fakeProvider struct{}

func (f fakeProvider) GetCommentById(ctx context.Context, id string) (Comment, error) {
	for i := range fakeData {
		if fakeData[i].ID == id {
			return fakeData[i], nil
		}
	}

	return Comment{}, nil
}

//var service = NewService(fakeProvider{})

func TestGetCommentById(t *testing.T) {
	t.Run("should be able to get a comment by id", func(t *testing.T) {

	})
}
