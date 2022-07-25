package comment

import (
	"context"
	"errors"
	"fmt"
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

var (
	ErrFetchingCommentById = errors.New("failed to fetch comment by id")
	ErrNotImplemented      = errors.New("not implemented")
)

type Provider interface {
	GetCommentById(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
}

type Service struct {
	Provider Provider
}

func NewService(provider Provider) *Service {
	return &Service{
		Provider: provider,
	}
}

func (s *Service) GetCommentById(ctx context.Context, id string) (Comment, error) {
	comment, err := s.Provider.GetCommentById(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingCommentById
	}

	return comment, nil
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Provider.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}

	return insertedCmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Provider.UpdateComment(ctx, id, cmt)
	if err != nil {
		fmt.Println("Error updating the comment")
		return Comment{}, err
	}
	return updatedCmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Provider.DeleteComment(ctx, id)
	if err != nil {
		fmt.Println("Error deleting the comment")
		return err
	}
	return nil
}
