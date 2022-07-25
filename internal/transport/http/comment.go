package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"simple-api/internal/comment"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CommentService interface {
	CreateComment(context.Context, comment.Comment) (comment.Comment, error)
	GetCommentById(context.Context, string) (comment.Comment, error)
	UpdateComment(context.Context, string, comment.Comment) (comment.Comment, error)
	DeleteComment(context.Context, string) error
}

type Response struct {
	Message string
}

type CreateCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func FormatCreateCommentRequest(c CreateCommentRequest) comment.Comment {
	return comment.Comment{
		Slug:   c.Slug,
		Author: c.Author,
		Body:   c.Body,
	}
}

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var cmt CreateCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmt)
	if err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	formatedComment := FormatCreateCommentRequest(cmt)
	createdComment, err := h.Service.CreateComment(r.Context(), formatedComment)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(createdComment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetCommentById(w http.ResponseWriter, r *http.Request) {
	queryString := mux.Vars(r)
	id := queryString["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}

	cmt, err := h.Service.GetCommentById(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	queryString := mux.Vars(r)
	id := queryString["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	queryString := mux.Vars(r)
	id := queryString["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is required"))
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "successful deleted"}); err != nil {
		panic(err)
	}
}
