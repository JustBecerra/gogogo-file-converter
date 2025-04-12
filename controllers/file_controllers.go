package controllers

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatFile(c *gin.Context) {
	format := c.PostForm("format")// format needed to change

 	file, _, err := c.Request.FormFile("image")// get file from client
 	if err != nil {
 		c.String(http.StatusBadRequest, "Failed to read file: %v", err)
 		return
 	}
 	defer file.Close()

 	// Decode image of any supported format
 	img, _, err := image.Decode(file)
 	if err != nil {
 		c.String(http.StatusBadRequest, "Failed to decode image: %v", err)
 		return
 	}

 	// Encode image to new format
 	var buf bytes.Buffer
 	var contentType string

	switch format {
	case "png":
		err = png.Encode(&buf, img)
		contentType = "image/png"
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90})
		contentType = "image/jpeg"
	// case "webp":
	// 	err = webp.Encode(&buf, img, &webp.Options{Lossless: true})
	// 	contentType = "image/webp"
	default:
		c.String(http.StatusBadRequest, "Unsupported format: %s", format)
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to encode image: %v", err)
		return
	}

	c.Data(http.StatusOK, contentType, buf.Bytes())
}