package models

import (
	"time"

	"github.com/google/uuid"
)

// Event represents the structure of an event
// @Description Event information
type Event struct {
	// The unique identifier of the event.
	// @json:id
	Id uuid.UUID `json:"id"`

	// The name of the event.
	// @json:name
	Name string `json:"name"`

	// The description of the event.
	// @json:description
	Description string `json:"description"`

	// The temperature of the event.
	// @json:temperature
	Temperature float64 `json:"temperature"`

	// The location of the event.
	// @json:location
	Location string `json:"location"`

	// The URL of the event image.
	// @json:src_image
	Src_Image string `json:"src_image"`

	// The date and time of the event.
	// @json:date
	Date time.Time `json:"date"`

	// The creation date of the event.
	// @json:createdate
	CreateDate time.Time `json:"createdate"`

	// The version date of the event.
	// @json:versiondate
	VersionDate time.Time `json:"versiondate"`

	// Indicates whether the event is active.
	// @json:isactive
	IsActive bool `json:"isactive"`

	// The capacity of the event.
	// @json:capacity
	Capacity int `json:"capacity"`
}
