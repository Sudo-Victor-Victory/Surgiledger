package api

import (
	"context"
	"net/http"

	"github.com/Sudo-Victor-Victory/surgiledger/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createEpisode(c *gin.Context, pool *pgxpool.Pool) {
	queries := db.New(pool)

	episode, err := queries.CreateEpisode(
		context.Background(),
		"created",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, episode)
}

func deleteEpisode(c *gin.Context, pool *pgxpool.Pool, id pgtype.UUID) {
	queries := db.New(pool)

	err := queries.DeleteEpisode(
		context.Background(),
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, 204)
}

func getEpisode(c *gin.Context, pool *pgxpool.Pool, id pgtype.UUID) {
	queries := db.New(pool)
	var dbEpisode db.Episode
	dbEpisode, err := queries.GetEpisode(
		context.Background(),
		id,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, dbEpisode)
}

func updateEpisodeStatus(c *gin.Context, pool *pgxpool.Pool, id pgtype.UUID, status string) {
	queries := db.New(pool)
	var idk db.UpdateEpisodeStatusParams
	idk.ID = id
	idk.Status = status
	dbEpisode, err := queries.UpdateEpisodeStatus(
		context.Background(),
		idk,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, dbEpisode)
}

func listEpisodes(c *gin.Context, pool *pgxpool.Pool) {
	queries := db.New(pool)
	dbEpisode, err := queries.ListEpisodes(
		context.Background(),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, dbEpisode)
}
