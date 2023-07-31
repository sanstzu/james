package utils

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"image/png"

	"io/ioutil"
	"net/http"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func Get(url string, resp *map[string]interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, resp)
	return err
}

func GetFile(url string, resp *map[string]interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(body, resp)
	return err
}

func GetRaw(url string, resp *[]byte) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	*resp = body
	return err
}

func ConvertToWebp(img image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := webp.Encode(buf, img, &webp.Options{Lossless: true}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ResizeImage(raw []byte) ([]byte, image.Image, error) {
	img, err := jpeg.Decode(bytes.NewReader(raw))
	if err != nil {
		return nil, nil, err
	}

	b := img.Bounds()

	var isWidthAuto bool
	width := b.Max.X
	height := b.Max.Y

	if width > height {
		isWidthAuto = false
	} else {
		isWidthAuto = true
	}
	var res image.Image
	if isWidthAuto {
		res = resize.Resize(0, 512, img, resize.MitchellNetravali)
	} else {
		res = resize.Resize(512, 0, img, resize.MitchellNetravali)
	}
	rawRes := new(bytes.Buffer)
	err = png.Encode(rawRes, res)
	if err != nil {
		return nil, nil, err
	}
	return rawRes.Bytes(), res, nil
}

func IsAllEmoji(s []string) bool {
	/*
		for _, v := range s {
			if !gomoji.ContainsEmoji(v) {
				return false
			}
		}
		return true
	*/

	// Return true until a solution is found to detect emojis in a string
	return true
}
