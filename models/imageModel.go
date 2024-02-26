package models

import (
	"github.com/google/uuid"
)

// Image represents the structure of an image
// @Description Image information
type Image struct {
	// ID of the image
	// @json:id
	Id uuid.UUID `json:"id"`

	// Alt text for the image
	// @json:alt
	Alt string `json:"alt"`

	// Description of the image
	// @json:description
	Description string `json:"description"`

	// Source URL of the image
	// @json:src_image
	Src_Image string `json:"src_image"`

	// Size of the image in bytes
	// @json:size
	Size int `json:"size"`
}
