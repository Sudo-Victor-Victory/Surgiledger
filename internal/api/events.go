package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Sudo-Victor-Victory/surgiledger/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createEvent(c *gin.Context, pool *pgxpool.Pool, episodeID pgtype.UUID, eventType string) {
	queries := db.New(pool)

	var payload map[string]interface{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid JSON payload",
		})
		return
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to encode payload",
		})
		return
	}

	params := db.CreateEventParams{
		EpisodeID: episodeID,
		EventType: eventType,
		Payload:   payloadBytes,
	}

	event, err := queries.CreateEvent(
		context.Background(),
		params,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, event)
}
