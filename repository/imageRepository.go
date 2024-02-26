// Package repository provides a repository implementation for managing images.
package repository

import (
	"api/db"
	"api/models"
	util "api/utils"
	"errors"
	"sync"

	"github.com/google/uuid"
)

// ImageRepository is a repository for managing images.
type ImageRepository struct {
	mu     sync.RWMutex    // Mutex for synchronization
	images []*models.Image // Slice to store images
}

// NewImageRepository creates a new instance of ImageRepository.
func NewImageRepository() *ImageRepository {
	// Initialize ImageRepository with images from the database
	return &ImageRepository{
		images: db.DbImages(),
	}
}

// GetImageByID retrieves an Image by its ID.
func (r *ImageRepository) GetImageByID(Id uuid.UUID) (*models.Image, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Iterate through images to find the one with the specified ID
	for _, image := range r.images {
		if image.Id == Id {
			return image, nil // Return the Image if found
		}
	}
	return nil, errors.New(util.ErrImageNotFound) // Return error if Image is not found
}

// GetAllimages retrieves all images from the repository.
func (r *ImageRepository) GetAllImages() []*models.Image {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.images // Return all images
}

// AddImage adds a new Image to the repository.
func (r *ImageRepository) AddImage(Image *models.Image) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Generate a UUID for the new Image
	Image.Id = uuid.New()

	r.images = append(r.images, Image) // Add the Image to the slice
}

// DeleteImage deletes an Image from the repository by its ID.
func (r *ImageRepository) DeleteImage(Id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Iterate through images to find the one with the specified ID
	for i, Image := range r.images {
		if Image.Id == Id {
			// Remove the Image from the slice
			r.images = append(r.images[:i], r.images[i+1:]...)
			return nil // Return nil if deletion is successful
		}
	}

	return errors.New(util.ErrImageNotFound) // Return error if Image is not found
}
