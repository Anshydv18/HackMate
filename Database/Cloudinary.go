package database

import (
	constants "Hackmate/Constants"
	env "Hackmate/Env"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

func CloudinaryConnect(ctx *context.Context) (*cloudinary.Cloudinary, error) {
	cloudinaryName := env.Get(constants.CLOUDINARY_NAME)
	cloudinaryApiKey := env.Get(constants.CLOUDINARY_API_KEY)
	cloudinarySecretKey := env.Get(constants.CLOUDINARY_API_SECRET_KEY)
	cld, err := cloudinary.NewFromParams(cloudinaryName, cloudinaryApiKey, cloudinarySecretKey)
	if err != nil {
		return nil, err
	}
	return cld, nil
}
