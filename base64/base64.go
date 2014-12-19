package base64

import (
	"fmt"
	"os"
	"bytes"
	"net/http"
	"encoding/base64"
)

func ReadImage(filename string) (string, error) {
	var b bytes.Buffer
	fileExists, _ := exists(filename)
	if !fileExists {
		return "", fmt.Errorf("The file '%s' does not exist\n", filename)
	}
	
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Error opening '%s'\n", filename)
	}
	_, err = b.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("Error reading the '%s' file to buffer\n", filename)
	}

	enc := encode(b.Bytes())
	mime := http.DetectContentType(b.Bytes())
	return format(enc, mime), nil
	
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}



func encode(bin []byte) []byte {
	e64 := base64.StdEncoding
	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)
	e64.Encode(encBuf, bin)
	return encBuf
}

func format(enc []byte, mime string) string {
	switch mime {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;base64,%s", mime, enc)
	default:
	}
	return fmt.Sprintf("data:image/png;base64,%s", enc)
}
