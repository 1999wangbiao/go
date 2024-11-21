package service

import "gvb_server/service/image_service"

type ServiceGropu struct {
	ImageService image_service.ImageService
}

var ServiceApp = new(ServiceGropu)
