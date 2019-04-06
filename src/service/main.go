package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mask = [...]byte{128, 64, 32, 16, 8, 4, 2, 1}

func EncodeMessage(img []byte, message string) ([]byte, error) {
	m := []byte(message)

	if len(message)*8 > len(img)-54 {
		return nil, errors.New("error: image is not large enough to hold this message")
	}

	for i := 0; i < len(m); i++ {
		index := 55 + i*8
		b := m[i]

		for j := 0; j < 8; j++ {
			if b&mask[j] == 0 {
				img[index+j] = setLSB(0, img[index+j])
			} else {
				img[index+j] = setLSB(1, img[index+j])
			}
		}

		if i == len(m)-1 {
			for j := 8; j < 16; j++ {
				img[index+j] = setLSB(0, img[index+j])
			}
		}
	}

	return img, nil
}

func DecodeMessage(img []byte) string {
	message := ""

	for i := 55; i < len(img)-9; i += 8 {
		var letter byte
		for j := 0; j < 8; j++ {
			b := img[i+j]
			if b%2 == 0 {
				letter &^= 1
			} else {
				letter |= 1
			}

			if j != 7 {
				letter = letter << 1
			}
		}

		if letter == 0 {
			break
		}

		message += string(letter)
	}
	return message
}

func HandleError(err error, c *gin.Context) {
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		panic(err)
	}
}

func setLSB(b byte, val byte) byte {
	if b != 0 {
		val |= 1
	} else {
		val &^= 1
	}
	return val
}