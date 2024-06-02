package controllers

import (
	"encoding/json"
	response "goravel/app/helpers/responses"
	"goravel/app/models"
	"io"
	net "net/http"

	"github.com/goravel/framework/contracts/http"
)

type AnimeController struct {
}

func NewAnimeController() *AnimeController {
	return &AnimeController{}
}

func (r *AnimeController) ShowTopAnime(ctx http.Context) http.Response {
	//fetch data from external api
	req, err := net.NewRequest("GET", "https://api.jikan.moe/v4/top/anime", nil)

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	client := net.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response.ErrInternalServerError(err, ctx)
	}

	var Anime models.Animes

	json.Unmarshal(body, &Anime)

	return response.SuccessOK(ctx, Anime)
}
