package igdb

import (
	"strings"

	"github.com/nineteen72/igdb/models"
)

func parseImage(images []models.Image) []models.Image {
	var im []models.Image

	for _, i := range images {
		ima := models.Image{
			URL:    strings.Replace(i.URL, "t_thumb", "t_1080p", 1),
			Width:  i.Width,
			Height: i.Height,
		}
		im = append(im, ima)
	}

	return im
}
