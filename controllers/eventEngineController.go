package controllers

import (
	"api/repository"
	"log"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v2"
)

type EventTemplate struct {
	eventRepo *repository.EventRepository
}

func EventTemplateRepository(eventRepo *repository.EventRepository) *EventTemplate {
	return &EventTemplate{
		eventRepo: eventRepo,
	}
}

// GetEventsTemplate renders the template for displaying all events.
// @Summary Render template for displaying all events
// @Description Retrieve all events from the repository and render the template with the events data
// @Tags events
// @Accept html
// @Produce html
// @Success 200 {string} html "HTML content with events data"
// @Router /events [get]
func (h *EventTemplate) GetEventsTemplate(c *fiber.Ctx) error {

	events := h.eventRepo.GetAllEvents()

	return c.Render("index", fiber.Map{
		"Title":  "Events",
		"Events": events,
	})

}

// GetEventTemplate renders the template for displaying a specific event.
// @Summary Render template for displaying a specific event
// @Description Retrieve the event specified by its ID from the repository and render the template with the event data
// @Tags events
// @Accept html
// @Produce html
// @Param id path string true "Event ID" Format(uuid) // <-- Define the format as UUID
// @Success 200 {string} html "HTML content with event data"
// @Failure 400 {string} html "Error message"
// @Router /events/{id} [get]
func (h *EventTemplate) GetEventTemplate(c *fiber.Ctx) error {

	id := c.Params("id")
	eventID, err := uuid.Parse(id)
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Error converting event ID to integer: ", err)
		// Pass error flag and message to template
		return c.Render("index", fiber.Map{
			"Title": "Error",
			"Error": "Invalid event ID",
		})
	}

	event, err := h.eventRepo.GetEventByID(eventID)
	if err != nil {
		log.Println("Error: ", err)
		log.Println("Error fetching event from repository: ", err)
		// Pass error flag and message to template
		return c.Render("index", fiber.Map{
			"Title": "Error",
			"Error": "Failed to retrieve event",
		})
	}

	return c.Render("index", fiber.Map{
		"Title": "Events",
		"Event": event,
	})

}
