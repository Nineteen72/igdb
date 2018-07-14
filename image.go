package igdb

import (
	"strings"

	"github.com/Nineteen72/igdb/models"
)

func parseImage(images ...models.Image) []models.Image {
	var im []models.Image

	for _, i := range images {
		img := strings.Replace(i.URL, "t_thumb", "t_1080p", 1)
		img = strings.Replace(img, "//", "https://", 1)
		ima := models.Image{
			URL:    img,
			Width:  i.Width,
			Height: i.Height,
		}
		im = append(im, ima)
	}

	return im
}
