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

type EventResponse struct {
	ID        pgtype.UUID        `json:"id"`
	EpisodeID pgtype.UUID        `json:"episode_id"`
	EventType string             `json:"event_type"`
	Payload   interface{}        `json:"payload"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

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

func getEpisodeEvents(c *gin.Context, pool *pgxpool.Pool, id pgtype.UUID) {
	queries := db.New(pool)

	dbEvents, err := queries.GetEpisodeEvents(
		context.Background(),
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := make([]EventResponse, 0, len(dbEvents))

	for _, event := range dbEvents {
		var payload interface{}

		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to decode event payload",
			})
			return
		}

		responses = append(responses, EventResponse{
			ID:        event.ID,
			EpisodeID: event.EpisodeID,
			EventType: event.EventType,
			Payload:   payload,
			CreatedAt: event.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func getEventEpisode(c *gin.Context, pool *pgxpool.Pool, episode_id pgtype.UUID, event_id pgtype.UUID) {
	queries := db.New(pool)
	var dbEvents db.Event
	eventParams := db.GetEventByIdParams{EpisodeID: episode_id, ID: event_id}
	dbEvents, err := queries.GetEventById(
		context.Background(),
		eventParams,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var payload interface{}

	if err := json.Unmarshal(dbEvents.Payload, &payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to decode event payload",
		})
		return
	}

	response := EventResponse{
		ID:        dbEvents.ID,
		EpisodeID: dbEvents.EpisodeID,
		EventType: dbEvents.EventType,
		Payload:   payload,
		CreatedAt: dbEvents.CreatedAt,
	}

	c.JSON(http.StatusOK, response)
}
