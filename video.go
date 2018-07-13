package igdb

import (
	"fmt"

	"github.com/nineteen72/igdb/models"
)

func parseVideos(videos []models.Video) []models.Video {
	var vList []models.Video

	for _, i := range videos {
		v := models.Video{
			Name:    i.Name,
			VideoID: i.VideoID,
			URL:     fmt.Sprintf("https://www.youtube.com/watch?v=%s", i.VideoID),
		}
		vList = append(vList, v)
	}

	return vList
}
