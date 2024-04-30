package post

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/thangsuperman/bee-happy/services/auth"
	"github.com/thangsuperman/bee-happy/types"
	"github.com/thangsuperman/bee-happy/utils"
)

type Handler struct {
	store     types.PostStore
	userStore types.UserStore
}

func NewHandler(store types.PostStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/posts", h.handleGetPosts).Methods(http.MethodGet)
	router.HandleFunc("/post/{id}", h.handleGetPost).Methods(http.MethodGet)
	router.HandleFunc("/post", auth.WithJWTAuth(h.handleCreateProduct, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/post/{id}", auth.WithJWTAuth(h.handleUpdatePost, h.userStore)).Methods(http.MethodPatch)
	router.HandleFunc("/post/{id}", auth.WithJWTAuth(h.handleDeletePost, h.userStore)).Methods(http.MethodDelete)
}

// handleGetPost get a post
// @Summary Get a post
// @Description Get a post by id
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id} [get]
func (h *Handler) handleGetPost(w http.ResponseWriter, r *http.Request) {
	stringPostId := mux.Vars(r)["id"]
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	p, err := h.store.GetPostById(postId)
	if p == nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Get post successfully",
		Metadata: p,
	})
}

// handleGetProducts get all posts
// @Summary Get all posts
// @Description Get all post
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/posts [get]
func (h *Handler) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.store.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Get all posts successfully",
		Metadata: posts,
	})
}

// handleUpdatePost update a post
// @Summary Update a post
// @Description Update a post
// @Tags Post
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Post ID"
// @Param payload body types.UpdatePostPayload  true "Post payload"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id} [patch]
func (h *Handler) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	authorId := auth.GetUserIdFromContext(r.Context())
	stringPostId := mux.Vars(r)["id"]
	var payload types.UpdatePostPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %w", errors))
		return
	}

	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.store.UpdatePost(payload, postId, authorId); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Update post successfully",
		Metadata: payload,
	})
}

// handleDeletePost delete a post
// @Summary Delete a post
// @Description Delete a post
// @Tags Post
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Post ID"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id} [delete]
func (h *Handler) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	authorId := auth.GetUserIdFromContext(r.Context())
	stringPostId := mux.Vars(r)["id"]
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.DeletePostById(postId, authorId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Delete post successfully",
		Metadata: nil,
	})
}

// handleCreatePost creates a new post.
// @Summary Create a new post
// @Description Create a new post
// @Tags Post
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param payload body types.CreatePostPayload true "Post payload"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post [post]
func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	authorId := auth.GetUserIdFromContext(r.Context())
	var payload types.CreatePostPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %w", errors))
		return
	}

	err := h.store.CreatePost(payload, authorId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Create a post successfully",
		Metadata: payload,
	})
}
