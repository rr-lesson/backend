package responses

import (
	"backend/internal/domains"
	"backend/internal/dto"
)

type CreateVideo struct {
	Video domains.Video `json:"video"`
} //@name CreateVideoRes

type GetAllVideos struct {
	Videos []domains.Video `json:"videos"`
} // @name GetAllVideosRes

type GetAllVideosByLessonId struct {
	Videos []domains.Video `json:"videos"`
} //@name GetAllVideosByLessonIdRes

type GetAllVideosWithDetail struct {
	Videos []dto.VideoDetail `json:"videos"`
} //@name GetAllVideosWithDetailRes

type GetVideoWithDetail struct {
	Video dto.VideoDetail `json:"video"`
} //@name GetVideoWithDetailRes
