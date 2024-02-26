// Package repository provides a repository implementation for managing events.
package repository

import (
	"api/db"
	"api/models"
	util "api/utils"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

// EventRepository is a repository for managing events.
type EventRepository struct {
	mu     sync.RWMutex    // Mutex for synchronization
	events []*models.Event // Slice to store events
}

// NewEventRepository creates a new instance of EventRepository.
func NewEventRepository() *EventRepository {
	// Initialize EventRepository with events from the database
	return &EventRepository{
		events: db.DbEvents(),
	}
}

// GetEventByID retrieves an event by its ID.
func (r *EventRepository) GetEventByID(Id uuid.UUID) (*models.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Iterate through events to find the one with the specified ID
	for _, event := range r.events {
		if event.Id == Id {
			return event, nil // Return the event if found
		}
	}
	return nil, errors.New(util.ErrEventNotFound) // Return error if event is not found
}

// GetAllEvents retrieves all events from the repository.
func (r *EventRepository) GetAllEvents() []*models.Event {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.events // Return all events
}

// AddEvent adds a new event to the repository.
func (r *EventRepository) AddEvent(event *models.Event) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generate a UUID for the new event
	event.Id = uuid.New()

	r.events = append(r.events, event) // Add the event to the slice
}

// UpdateEvent updates an existing event in the repository.
func (r *EventRepository) UpdateEvent(event *models.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Iterate through events to find the one with the specified ID
	for i, e := range r.events {
		if e.Id == event.Id {
			// Update the event's properties
			r.events[i].Name = event.Name
			r.events[i].Description = event.Description
			r.events[i].Location = event.Location
			r.events[i].Capacity = event.Capacity
			r.events[i].IsActive = event.IsActive
			r.events[i].VersionDate = time.Now()
			r.events[i].Date = event.Date

			return nil // Return nil if update is successful
		}
	}

	return errors.New(util.ErrEventNotFound) // Return error if event is not found
}

// DeleteEvent deletes an event from the repository by its ID.
func (r *EventRepository) DeleteEvent(Id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Iterate through events to find the one with the specified ID
	for i, event := range r.events {
		if event.Id == Id {
			// Remove the event from the slice
			r.events = append(r.events[:i], r.events[i+1:]...)
			return nil // Return nil if deletion is successful
		}
	}

	return errors.New(util.ErrEventNotFound) // Return error if event is not found
}
