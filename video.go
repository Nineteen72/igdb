package igdb

import (
	"net/http"

	"github.com/Nineteen72/go-get-youtube/youtube"
	"github.com/Nineteen72/igdb/models"
)

func parseVideos(videos []models.Video, client *http.Client) ([]models.Video, error) {
	var vList []models.Video

	for _, i := range videos {
		vi, err := youtube.Get(i.VideoID, client)
		if err != nil {
			return nil, err
		}

		for i, val := range vi.Formats {
			if ext := vi.GetExtension(i); ext == "mp4" {
				v := models.Video{
					Name:      vi.Title,
					URL:       val.Url,
					Extension: ext,
				}
				vList = append(vList, v)
				break
			}
		}
	}

	return vList, nil
}
