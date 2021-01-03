package handler

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/google/uuid"
	"pancake.maker/gen/api"
)

type ImageUploadHandler struct {
	sync.Mutex
	files map[string][]byte
}

func NewImageUploadHandler() *ImageUploadHandler {
	return &ImageUploadHandler{
		files: make(map[string][]byte),
	}
}

// Upload image
func (h *ImageUploadHandler) Upload(stream api.ImageUploadService_UploadServer) error {
	// receive first request(metadata)
	req, err := stream.Recv()
	if err != nil {
		return err
	}
	meta := req.GetFileMeta()
	filename := meta.Filename

	// generate UUID
	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	id := u.String()

	// buffer for image data
	buf := &bytes.Buffer{}

	// loop for receive chunked binary data
	for {
		// receive after second requests
		r, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		chunk := r.GetData()
		_, err = buf.Write(chunk)
		if err != nil {
			return err
		}
	}

	data := buf.Bytes()
	mimeType := http.DetectContentType(data)

	h.files[filename] = data

	// Return a response and close the stream
	err = stream.SendAndClose(&api.ImageUploadResponse{
		Uuid:        id,
		Size:        int32(len(data)),
		ContentType: mimeType,
		Filename:    filename,
	})

	file, err := os.Create("uploaded/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return err
}
