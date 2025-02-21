package image

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
)

func EncodeImageToString(imagePath string) string {
	fp, err := os.Open(imagePath)
	if err != nil {
		return ""
	}
	defer fp.Close()

	// build a base64 encoded string from the uploaded file
	buffer := new(bytes.Buffer)
	io.Copy(buffer, fp)
	return base64.StdEncoding.EncodeToString(buffer.Bytes())
}
