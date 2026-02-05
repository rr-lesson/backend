package handlers

import (
	"backend/internal/domains"
	"backend/internal/dto/requests"
	"backend/internal/dto/responses"
	"backend/internal/repositories"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type VideoHandler struct {
	minio     *minio.Client
	videoRepo *repositories.VideoRepository
}

func NewVideoHandler(
	minio *minio.Client,
	videoRepo *repositories.VideoRepository,
) *VideoHandler {
	return &VideoHandler{
		minio:     minio,
		videoRepo: videoRepo,
	}
}

func (h *VideoHandler) RegisterRoutes(router fiber.Router) {
	router.Get("/hls/*", h.handleHLS)

	g1 := router.Group("/classes/subjects/lessons/videos")
	g1.Get("/", h.getAllVideosWithDetail)

	g2 := router.Group("/classes/:classId/subjects/:subjectId/lessons/:lessonId/videos")
	g2.Get("/", h.getAllVideosByLessonId)

	g3 := router.Group("/videos")
	g3.Post("/", h.createVideo)
	g3.Get("/", h.getAllVideos)
	g3.Get("/:videoId", h.GetVideoWithDetail)
}

func (h *VideoHandler) handleHLS(c *fiber.Ctx) error {
	filepath := c.Params("*")
	bucket := os.Getenv("MINIO_BUCKET")

	obj, err := h.minio.GetObject(
		c.Context(),
		bucket,
		filepath,
		minio.GetObjectOptions{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Error{
			Message: err.Error(),
		})
	}
	defer obj.Close()

	stat, err := obj.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	switch {
	case strings.HasSuffix(filepath, ".m3u8"):
		c.Set("Content-Type", "application/vnd.apple.mpegurl")
	case strings.HasSuffix(filepath, ".ts"):
		c.Set("Content-Type", "video/mp2t")
	default:
		c.Set("Content-Type", stat.ContentType)
	}

	c.Set("Accept-Ranges", "bytes")
	c.Set("Cache-Control", "public, max-age=31536000, immutable")

	_, err = io.Copy(c.Context().Response.BodyWriter(), obj)
	return err
}

// @id 					CreateVideo
// @tags 				video
// @accept 			json
// @produce 		json
// @param 			body body requests.CreateVideo true "body"
// @success 		200 {object} responses.CreateVideo
// @router 			/api/v1/videos [post]
func (h *VideoHandler) createVideo(c *fiber.Ctx) error {
	lessonId, err := c.ParamsInt("lessonId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	var req requests.CreateVideo
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.videoRepo.Create(domains.Video{
		LessonId:    uint(lessonId),
		FilePath:    req.FilePath,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(responses.CreateVideo{
		Video: *res,
	})
}

// @id 					GetAllVideos
// @tags 				video
// @accept 			json
// @produce 		json
// @param 			lessonId query int false "lessonId"
// @success 		200 {object} responses.GetAllVideos
// @router 			/api/v1/videos [get]
func (h *VideoHandler) getAllVideos(c *fiber.Ctx) error {
	lessonId := c.QueryInt("lessonId", 0)

	res, err := h.videoRepo.GetAll(repositories.VideoFilter{
		LessonId: uint(lessonId),
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllVideos{
		Videos: *res,
	})
}

// @deprecated
// @id 					GetAllVideosByLessonId
// @tags 				video
// @accept 			json
// @produce 		json
// @param 			subjectId path int true "subjectId"
// @param 			lessonId path int true "lessonId"
// @success 		200 {object} responses.GetAllVideosByLessonId
// @router 			/api/v1/subjects/{subjectId}/lessons/{lessonId}/videos [get]
func (h *VideoHandler) getAllVideosByLessonId(c *fiber.Ctx) error {
	lessonId, err := strconv.Atoi(c.Params("lessonId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.videoRepo.GetAllByLessonId(uint(lessonId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllVideosByLessonId{
		Videos: *res,
	})
}

// @id 					GetAllVideosWithDetail
// @tags 				video
// @accept 			json
// @produce 		json
// @success 		200 {object} responses.GetAllVideosWithDetail
// @router 			/api/v1/classes/subjects/lessons/videos [get]
func (h *VideoHandler) getAllVideosWithDetail(c *fiber.Ctx) error {
	res, err := h.videoRepo.GetAllWithDetail()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetAllVideosWithDetail{
		Videos: *res,
	})
}

// @id 					GetVideoWithDetail
// @tags 				video
// @accept 			json
// @produce 		json
// @param 			videoId path int true "videoId"
// @success 		200 {object} responses.GetVideoWithDetail
// @router 			/api/v1/videos/{videoId} [get]
func (h *VideoHandler) GetVideoWithDetail(c *fiber.Ctx) error {
	videoId, err := c.ParamsInt("videoId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	res, err := h.videoRepo.GetWithDetail(uint(videoId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.GetVideoWithDetail{
		Video: *res,
	})
}
