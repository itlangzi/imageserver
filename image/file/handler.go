package file

import (
	"github.com/itlangzi/imageserver"
)

type Handler struct {
	// Processor Processor
}

// Handle implements imageserver.Handler.
func (hdr *Handler) Handle(im *imageserver.Image, params imageserver.Params) (*imageserver.Image, error) {
	return im, nil
}

type FallbackHandler struct {
	*Handler
	Fallback imageserver.Handler
}

// Handle implements imageserver.Handler.
func (hdr *FallbackHandler) Handle(im *imageserver.Image, params imageserver.Params) (*imageserver.Image, error) {
	h, err := hdr.getHandler(im, params)
	if err != nil {
		return nil, err
	}
	return h.Handle(im, params)
}

func (hdr *FallbackHandler) getHandler(im *imageserver.Image, params imageserver.Params) (imageserver.Handler, error) {
	if im.Format == "" || im.Format == "file" {
		return hdr.Handler, nil
	}
	return hdr.Fallback, nil
}
