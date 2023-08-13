package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/personal/blog-app/app"
	"github.com/personal/blog-app/app/dto"
	"github.com/personal/blog-app/app/middleware"
	"github.com/personal/blog-app/app/service"
)

type PostCommentController struct {
	postCommentService service.PostCommentService
	blogPostService    service.BlogPostService
}

func NewPostCommentController(postCommentService service.PostCommentService, blogPostService service.BlogPostService) app.Controller {
	return &PostCommentController{postCommentService, blogPostService}
}

func (c *PostCommentController) Register(r app.Router) {

	r.
		Route("/api/v1/post-comment", func(r app.Router) {
			r.With(middleware.JSONHeader()).
				Post("/", c.CreatePostComment)

			r.With(middleware.JSONHeader()).
				Delete("/{commentID:[1-9][0-9]*}", c.DeletePostComment)

		})
}

func (c *PostCommentController) CreatePostComment(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), time.Duration(5000000000))
	defer cancel()

	var newComment dto.PostCommentToCreate

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&newComment); err != nil {
		http.Error(writer, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	_, err := c.blogPostService.GetBlogPostByID(ctx, newComment.PostID)
	if err != nil {
		http.Error(writer, "post with this id not found", http.StatusNotFound)
		return
	}

	createdPost, err := c.postCommentService.CreatePostComment(ctx, &newComment)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	responseJSON, _ := json.Marshal(createdPost)

	writer.Write(responseJSON)
}

func (c *PostCommentController) DeletePostComment(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), time.Duration(5000000000))
	defer cancel()

	commentID, err := URLParamInt(request, "commentID")
	if err != nil {
		http.Error(writer, "cannot convert commentID to integer", http.StatusBadRequest)
		return
	}

	err = c.postCommentService.DeletePostComment(ctx, commentID)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}
