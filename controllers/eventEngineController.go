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

func (h *EventTemplate) GetEventsTemplate(c *fiber.Ctx) error {

	events := h.eventRepo.GetAllEvents()

	return c.Render("index", fiber.Map{
		"Title":  "Events",
		"Events": events,
	})

}

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
