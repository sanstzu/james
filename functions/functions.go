package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sanstzu/james/config"
	tele "gopkg.in/telebot.v3"
)

func Send(recipient tele.Recipient, what interface{}, options ...interface{}) (*tele.Message, error) {
	bot := config.GetBot()
	return bot.Send(recipient, what, options...)
}

func CreateStickerSet(recipient tele.Recipient, stickerConfig tele.StickerSet) error {
	bot := config.GetBot()
	return bot.CreateStickerSet(recipient, stickerConfig)
}

func UploadSticker(to tele.Recipient, png *tele.File) (*tele.File, error) {
	bot := config.GetBot()
	return bot.UploadSticker(to, png)
}

type InputSticker struct {
	Sticker    string   `json:"sticker"`
	Emoji_list []string `json:"emoji_list"`
}

type CreateNewStickerSetParams struct {
	UserId        int64          `json:"user_id"`
	Name          string         `json:"name"`
	Title         string         `json:"title"`
	Stickers      []InputSticker `json:"stickers"`
	StickerFormat string         `json:"sticker_format"`
	StickerType   string         `json:"sticker_type"`
}

type AddStickerToSetParams struct {
	UserId  int64        `json:"user_id"`
	Name    string       `json:"name"`
	Sticker InputSticker `json:"sticker"`
}

func CreateNewStickerSet(stickerConfig *CreateNewStickerSetParams) error {
	b := config.GetBot()

	url := b.URL + "/bot" + b.Token + "/" + "createNewStickerSet"

	JSONStickerConfig, err := json.Marshal(stickerConfig)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(JSONStickerConfig)

	resp, err := http.Post(url, "application/json", reader)

	var resMap map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resMap)

	if err != nil {
		return err
	}
	if resMap["ok"] != true {
		newError := errors.New(resMap["description"].(string))
		return newError
	}
	return nil
}

func AddStickerToSet(sticker *AddStickerToSetParams) error {
	b := config.GetBot()

	url := b.URL + "/bot" + b.Token + "/" + "addStickerToSet"
	JSONStickerConfig, err := json.Marshal(sticker)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(JSONStickerConfig)

	resp, err := http.Post(url, "application/json", reader)

	var resMap map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&resMap)

	if err != nil || resMap["ok"] != true {
		return resMap["description"].(error)
	}
	return nil
}
