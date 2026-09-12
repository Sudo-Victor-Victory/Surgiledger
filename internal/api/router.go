package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	router.POST("/episodes", func(c *gin.Context) {
		createEpisode(c, pool)
	})

	router.DELETE("/episodes/:id", func(c *gin.Context) {
		var pgUUID pgtype.UUID
		err := pgUUID.Scan(c.Param("id"))
		if err != nil {
			fmt.Println("Sorry bad user data")
			return
		}
		deleteEpisode(c, pool, pgUUID)
	})

	router.GET("/episodes/:id", func(c *gin.Context) {
		var pgUUID pgtype.UUID
		err := pgUUID.Scan(c.Param("id"))
		if err != nil {
			fmt.Println("Sorry bad user data")
			return
		}
		getEpisode(c, pool, pgUUID)
	})

	router.PATCH("/episodes/:id/:status", func(c *gin.Context) {
		var pgUUID pgtype.UUID
		var status string
		status = c.Param("status")
		err := pgUUID.Scan(c.Param("id"))

		if err != nil {
			fmt.Println("Sorry bad user data")
			return
		}
		updateEpisodeStatus(c, pool, pgUUID, status)
	})

	router.GET("/episodes", func(c *gin.Context) {
		listEpisodes(c, pool)
	})

	router.POST("/episodes/:episode_id/events/:event_type", func(c *gin.Context) {
		var episodeID pgtype.UUID

		if err := episodeID.Scan(c.Param("episode_id")); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid episode ID",
			})
			return
		}

		eventType := c.Param("event_type")

		createEvent(c, pool, episodeID, eventType)
	})
	return router
}
