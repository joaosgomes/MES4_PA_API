package controllers

import (
	"api/models"
	"api/repository"
	util "api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ImageController handles operations related to images.
type ImageController struct {
	imageRepo *repository.ImageRepository
}

// NewImageController creates a new instance of ImageController.
func NewImageController(imageRepo *repository.ImageRepository) *ImageController {
	return &ImageController{
		imageRepo: imageRepo,
	}
}

// GetImage returns an image by ID.
// @Summary Get an image by ID
// @Description Get an image from the database by its ID
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "Image ID" Format(uuid) // <-- Define the format as UUID
// @Success 200 {object} models.Image
// @Failure 404 {object} models.ErrorResponseModel
// @Router /image/{id} [get]
func (h *ImageController) GetImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the image ID from the request parameters
		id := c.Params("id")
		imageID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidImageId, nil, fiber.StatusBadRequest))
		}

		// Retrieve the image from the repository using the UUID
		image, err := h.imageRepo.GetImageByID(imageID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Error(err.Error(), nil, fiber.StatusBadRequest))
		}

		return c.JSON(image)
	}
}

// GetImages returns all images.
// @Summary Get all images
// @Description Get all images from the database
// @Tags images
// @Accept json
// @Produce json
// @Success 200 {array} models.Image
// @Router /image [get]
func (h *ImageController) GetImages() fiber.Handler {
	return func(c *fiber.Ctx) error {
		images := h.imageRepo.GetAllImages()
		return c.JSON(images)
	}
}

// PostImage creates a new image.
// @Summary Create a new image
// @Description Create a new image and add it to the database
// @Tags images
// @Accept json
// @Produce json
// @Param image body models.Image true "Image object"
// @Success 201 {object} models.Image
// @Failure 400 {object} models.ErrorResponseModel
// @Router /image [post]
func (h *ImageController) PostImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var image models.Image
		if err := c.BodyParser(&image); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidRequestPayload, nil, fiber.StatusBadRequest))
		}

		h.imageRepo.AddImage(&image)
		return c.JSON(image)
	}
}

// DeleteImage deletes an image by ID.
// @Summary Delete an image by ID
// @Description Delete an image from the database by its ID
// @Tags images
// @Accept json
// @Produce json
// @Param id path string true "Image ID" Format(uuid) // <-- Define the format as UUID
// @Success 204
// @Failure 400 {object} models.ErrorResponseModel
// @Failure 404 {object} models.ErrorResponseModel
// @Router /image/{id} [delete]
func (h *ImageController) DeleteImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse the image ID from the request parameters
		id := c.Params("id")
		imageID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.Error(util.ErrInvalidImageId, nil, fiber.StatusBadRequest))
		}

		// Delete the image from the repository using the UUID
		if err := h.imageRepo.DeleteImage(imageID); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(models.Error(err.Error(), nil, fiber.StatusBadRequest))
		}

		// Return a success response
		return c.SendStatus(fiber.StatusNoContent)
	}
}
