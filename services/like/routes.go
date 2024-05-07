package like

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/thangsuperman/bee-happy/db"
	"github.com/thangsuperman/bee-happy/services/auth"
	"github.com/thangsuperman/bee-happy/types"
	"github.com/thangsuperman/bee-happy/utils"
)

type Handler struct {
	store     types.LikeStore
	userStore types.UserStore
}

func NewHandler(store types.LikeStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/post/{id}/likes", h.handleGetTotalLikes).Methods(http.MethodGet)
	router.HandleFunc("/post/{id}/like", auth.WithJWTAuth(h.handleLikePost, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/post/{id}/unlike", auth.WithJWTAuth(h.handleUnlikePost, h.userStore)).Methods(http.MethodPost)
}

// handleGetTotalLikes get total likes
// @Summary Get total likes
// @Description Get total likes
// @Tags Post activites
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id}/likes [get]
func (h *Handler) handleGetTotalLikes(w http.ResponseWriter, r *http.Request) {
	stringPostId := mux.Vars(r)["id"]
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cacheKey := fmt.Sprintf("post:%d:total_likes", postId)
	cacheValue, err := db.RedisClient.Get(context.Background(), cacheKey).Result()

	if cacheValue == "" {
		totalLikes, err := h.store.CountLikes(postId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}

		ttl := 24 * time.Hour
		err = db.RedisClient.Set(context.Background(), cacheKey, totalLikes, ttl).Err()
		if err != nil {
			log.Println("Failed to set key:", err)
			utils.WriteError(w, http.StatusInternalServerError, err)
		}

		utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
			Message:  "Get get total post's likes successfully",
			Metadata: map[string]int{"total_likes": totalLikes},
		})
		return
	}

	totalLikes, err := strconv.Atoi(cacheValue)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Get total post's likes successfully",
		Metadata: map[string]int{"total_likes": totalLikes},
	})
}

// handleLikePost like post
// @Summary Like post
// @Description Like post
// @Tags Post activites
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Post ID"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id}/like [post]
func (h *Handler) handleLikePost(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	stringPostId := mux.Vars(r)["id"]
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.CreateLike(userId, postId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cacheKey := fmt.Sprintf("post:%d:total_likes", postId)
	err = db.RedisClient.Del(context.Background(), cacheKey).Err()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Like post successfully",
		Metadata: nil,
	})
}

// handleUnlikePost unlike post
// @Summary Unlike post
// @Description Unlike post
// @Tags Post activites
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param id path string true "Post ID"
// @Success 200 {object} types.BaseResponse "Success"
// @Router /api/v1/post/{id}/unlike [post]
func (h *Handler) handleUnlikePost(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())
	stringPostId := mux.Vars(r)["id"]
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.DeleteLike(userId, postId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	cacheKey := fmt.Sprintf("post:%d:total_likes", postId)
	err = db.RedisClient.Del(context.Background(), cacheKey).Err()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Unlike post successfully",
		Metadata: nil,
	})

}
