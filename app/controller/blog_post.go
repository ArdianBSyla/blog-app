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

type BlogPostController struct {
	blogPostService service.BlogPostService
}

func NewBlogPostController(blogPostService service.BlogPostService) app.Controller {
	return &BlogPostController{blogPostService}
}

func (c *BlogPostController) Register(r app.Router) {

	r.
		Route("/api/v1/blog-post", func(r app.Router) {
			r.With(middleware.JSONHeader()).
				Post("/", c.CreateBlogPost)

			r.With(middleware.JSONHeader()).
				Get("/list", c.GetBlogPostList)

		})
}

func (c *BlogPostController) CreateBlogPost(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), time.Duration(5000000000))
	defer cancel()

	var newPost dto.BlogPostToCreate

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&newPost); err != nil {
		http.Error(writer, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	createdPost, err := c.blogPostService.CreateBlogPost(ctx, &newPost)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	responseJSON, _ := json.Marshal(createdPost)

	writer.Write(responseJSON)
}

func (c *BlogPostController) GetBlogPostList(writer http.ResponseWriter, request *http.Request) {
	ctx, cancel := context.WithTimeout(request.Context(), time.Duration(5000000000))
	defer cancel()

	limit, err := queryParamInt(request, "limit", 1000)
	if err != nil {
		http.Error(writer, "cannot interpret 'limit' as integer", http.StatusBadRequest)
		return
	}

	if limit < 0 {
		http.Error(writer, "invalid 'limit' value", http.StatusBadRequest)
		return
	}

	offset, err := queryParamInt(request, "offset", 0)
	if err != nil {
		http.Error(writer, "cannot interpret 'offset' as integer", http.StatusBadRequest)
		return
	}

	if offset < 0 {
		http.Error(writer, "invalid 'offset' value", http.StatusBadRequest)
		return
	}

	posts, err := c.blogPostService.GetListOfBlogPosts(ctx, limit, offset)
	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	responseJSON, _ := json.Marshal(posts)

	writer.Write(responseJSON)
}
