package post

import (
	"fmt"
	"net/http"

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
	router.HandleFunc("/posts", auth.WithJWTAuth(h.handleGetProducts, h.userStore)).Methods(http.MethodGet)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIdFromContext(r.Context())

	// TODO: should remove this one
	fmt.Println("userId :", userId)

	products, err := h.store.GetPosts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types.BaseResponse{
		Message:  "Get all posts successfully",
		Metadata: products,
	})
}
