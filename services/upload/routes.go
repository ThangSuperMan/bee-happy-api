package upload

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gorilla/mux"
	AppSetting "github.com/thangsuperman/bee-happy/config"
	"github.com/thangsuperman/bee-happy/services/auth"
	"github.com/thangsuperman/bee-happy/types"
	"github.com/thangsuperman/bee-happy/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/upload", auth.WithJWTAuth(h.handleUploadImage, h.store))
}

// handleUploadImage  Login
// @Summary		        Upload image
// @Description       Upload an image for the post/profile
// @Tags			        Upload
// @Accept			      json
// @Produce		        json
// @Param Authorization header string true "JWT Token"
// @Param             file formData file true "Image file to upload"
// @Success		        200 {object} types.BaseResponse
// @Router			      /api/v1/upload [post]
func (h *Handler) handleUploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(AppSetting.Envs.AWSRegion),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				AppSetting.Envs.AWSBucketAccessKey,
				AppSetting.Envs.AWSBucketSecretKey, ""),
		),
	)

	if err != nil {
		log.Fatalf("failed to load the SDK %v", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	client := s3.NewFromConfig(cfg)

	filename := generateFilename(header.Filename)
	input := s3.PutObjectInput{
		Bucket: aws.String(AppSetting.Envs.AWSBucketName),
		Key:    aws.String(filename),
		Body:   file,
	}

	_, err = client.PutObject(context.TODO(), &input)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	imageUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", AppSetting.Envs.AWSBucketName, AppSetting.Envs.AWSRegion, filename)
	utils.WriteJSON(w, http.StatusCreated, types.BaseResponse{
		Message:  "Upload image successfully",
		Metadata: map[string]string{"image_url": imageUrl},
	})
}

func generateFilename(originalFilename string) string {
	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	ext := filepath.Ext(originalFilename)
	return fmt.Sprintf("%s_%s%s", timestamp, originalFilename[:len(originalFilename)-len(ext)], ext)
}
