package controllers

import (
	"api/models"
	"api/repository"
	util "api/utils"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// EventController handles operations related to events.
type EventController struct {
	eventRepo *repository.EventRepository
}

// NewEventController creates a new instance of EventController.
func NewEventController(eventRepo *repository.EventRepository) *EventController {
	return &EventController{
		eventRepo: eventRepo,
	}
}

// GetEvent returns an event by ID.
// @Summary Get an event by ID
// @Description Get an event from the database by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID" Format(uuid) // <-- Define the format as UUID
// @Success 200 {object} models.Event
// @Failure 404 {object} models.ErrorResponseModel
// @Router /event/{id} [get]
func (h *EventController) GetEvent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the event ID from the request parameters
		id := c.Params("id")
		eventID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidEventId, nil, fiber.StatusBadRequest))
		}

		// Retrieve the event from the repository using the UUID
		event, err := h.eventRepo.GetEventByID(eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Error(err.Error(), nil, fiber.StatusBadRequest))
		}

		return c.JSON(event)
	}
}

// GetEvents returns all events.
// @Summary Get all events
// @Description Get all events from the database
// @Tags events
// @Accept json
// @Produce json
// @Success 200 {array} models.Event
// @Router /event [get]
func (h *EventController) GetEvents() fiber.Handler {
	return func(c *fiber.Ctx) error {
		events := h.eventRepo.GetAllEvents()
		return c.JSON(events)
	}
}

// PostEvent creates a new event.
// @Summary Create a new event
// @Description Create a new event and add it to the database
// @Tags events
// @Accept json
// @Produce json
// @Param event body models.Event true "Event object"
// @Success 201 {object} models.Event
// @Failure 400 {object} models.ErrorResponseModel
// @Router /event [post]
func (h *EventController) PostEvent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var event models.Event
		if err := c.BodyParser(&event); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidRequestPayload, nil, fiber.StatusBadRequest))
		}

		h.eventRepo.AddEvent(&event)
		return c.JSON(event)
	}
}

// PutEvent updates an existing event.
// @Summary Update an existing event
// @Description Update an existing event in the database
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID" Format(uuid) // <-- Define the format as UUID
// @Param event body models.Event true "Event object"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 404 {object} models.ErrorResponseModel
// @Router /event/{id} [put]
func (h *EventController) PutEvent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get event ID from URL params
		id := c.Params("id")
		eventID, err := uuid.Parse(id)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidEventId, nil, fiber.StatusBadRequest))
		}

		// Set the event ID to the ID extracted from the URL path
		var event models.Event
		event.Id = eventID

		// Parse request body into Event object
		if err := c.BodyParser(&event); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidRequestPayload, nil, fiber.StatusBadRequest))
		}

		// Update the event in the repository
		if err := h.eventRepo.UpdateEvent(&event); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Error(err.Error(), nil, fiber.StatusBadRequest))
		}

		return c.JSON(event)
	}
}

// DeleteEvent deletes an event by ID.
// @Summary Delete an event by ID
// @Description Delete an event from the database by its ID
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID" Format(uuid) // <-- Define the format as UUID
// @Success 204
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 404 {object} models.ErrorResponseModel
// @Router /event/{id} [delete]
func (h *EventController) DeleteEvent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the event ID from the request parameters
		id := c.Params("id")
		eventID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidEventId, nil, fiber.StatusBadRequest))
		}

		// Delete the event from the repository using the UUID
		if err := h.eventRepo.DeleteEvent(eventID); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Error(err.Error(), nil, fiber.StatusBadRequest))
		}

		// Return a success response
		return c.SendStatus(fiber.StatusNoContent)
	}
}

// calculateTimeUntilEvent calculates the time until the event occurs
func calculateTimeUntilEvent(event *models.Event) time.Duration {
	timeUntilEvent := event.Date.Sub(time.Now())
	if timeUntilEvent < 0 {
		timeUntilEvent = 0
	}
	return timeUntilEvent
}

// WSHandler handles WebSocket connections for getting remaining time of an event.
// @Summary Get remaining time of an event via WebSocket
// @Description Establishes a WebSocket connection to get the remaining time until an event occurs
// @Tags events
// @Accept json
// @Produce json
// @Param id path string true "Event ID" Format(uuid) // <-- Define the format as UUID
// @Router /ws/event/{id} [get]
func (h *EventController) WSHandler(c *websocket.Conn) {
	id := c.Params("id")
	eventID, err := uuid.Parse(id)

	if err != nil {
		return
	}
	// Retrieve event details based on ID

	event, err := h.eventRepo.GetEventByID(eventID)
	if err != nil {
		return
	}
	if err != nil {
		log.Println("Error retrieving event details:", err)
		return
	}

	// Send updates
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// Calculate
		remainingTime := calculateTimeUntilEvent(event)

		// Send
		if err := c.WriteJSON(remainingTime.String()); err != nil {
			log.Println(util.ErrSendingWsMessage, err)
			return
		}
	}
}
