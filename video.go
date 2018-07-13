package igdb

import (
	"fmt"

	"github.com/nineteen72/igdb/models"
)

func parseVideos(videos *[]models.Video) {
	var vList []models.Video

	for _, i := range *videos {
		v := models.Video{
			Name: i.Name,
			URL:  fmt.Sprintf("https://www.youtube.com/watch?v=%s", i.VideoID),
		}
		vList = append(vList, v)
	}

	videos = &vList
}
