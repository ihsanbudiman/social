package thread_handler_rest

import (
	"fmt"
	"social/constant"
	"social/domain"
	"social/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PhotoSaver(c *fiber.Ctx, photos *[]domain.ThreadPhoto, images ...string) error {

	var err error = nil

	for i, image := range images {
		image, err := c.FormFile(image)
		if err != nil {

			err = fmt.Errorf("error when get image %d: %w", i, err)
			break
		} else {

			extImage, err := helper.GetFileExtensionFromUrl(image.Filename)
			if err != nil {

				err = fmt.Errorf("error when get image %d extension: %w", i, err)
				break
			}
			uuid := uuid.New()
			imageName := fmt.Sprintf("%s.%s", uuid.String(), extImage)

			err = c.SaveFile(image, fmt.Sprintf(".%s/%s", constant.THREAD_IMAGE, imageName))

			if err != nil {

				err = fmt.Errorf("error when save image %d: %w", i, err)
				break
			}

			*photos = append(*photos, domain.ThreadPhoto{
				FileUrl: imageName,
			})
		}
	}

	return err
}
