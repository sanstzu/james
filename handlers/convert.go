package handlers

import (
	"bytes"
	"image"
	"log"
	"strings"

	"github.com/sanstzu/james/consts"
	fnc "github.com/sanstzu/james/functions"
	"github.com/sanstzu/james/models"
	"github.com/sanstzu/james/utils"
	tele "gopkg.in/telebot.v3"
)

func Convert(c tele.Context) error {
	rawArgs := strings.Split(c.Message().Text, " ")
	if len(rawArgs) < 2 {
		return c.Reply("Usage: /convert <emojis>\nemojis: Emoji(s) that represents the sticker (no spaces between)")
	}
	args := make(map[string]string)

	args["emojis"] = rawArgs[1]

	emojiList := []string{args["emojis"]}
	log.Printf("emojiList: %v", emojiList)
	if !utils.IsAllEmoji(emojiList) || len(emojiList) > 20 {
		return c.Reply("Invalid emoji")
	}

	chatDetails, err := models.GetChat(c.Chat().ID)
	if err != nil {
		return err
	}

	if chatDetails == nil {
		return c.Reply("This chat have not initialize any sticker set yet. Please use /initialize first.")
	}

	stickerSetDetails, err := models.GetSticker(chatDetails["sticker_id"].(string))
	if err != nil {
		return err
	}

	var photoId string = ""
	var photoSource string = ""
	if photoId, photoSource = utils.ExtractFileId(c.Message()); photoId == "" && c.Message().ReplyTo != nil {
		photoId, photoSource = utils.ExtractFileId(c.Message().ReplyTo)
	}
	if photoId == "" {
		return c.Reply("Please reply or send to a photo/sticker to convert it.")
	}

	var resp map[string]interface{}
	err = utils.Get("https://api.telegram.org/bot"+consts.ENV("TELEGRAM_APITOKEN")+"/getFile?file_id="+photoId, &resp)
	if err != nil || resp["ok"] == false {
		return err
	}

	filePath := resp["result"].(map[string]interface{})["file_path"].(string)

	var rawFile []byte

	err = utils.GetRaw("https://api.telegram.org/file/bot"+consts.ENV("TELEGRAM_APITOKEN")+"/"+filePath, &rawFile)
	if err != nil {
		return err
	}

	var resizeImage func(raw []byte) ([]byte, image.Image, error)

	if photoSource == "sticker" {
		resizeImage = utils.ResizeImage
	} else if photoSource == "photo" {
		resizeImage = utils.ResizeImageJpeg
	}
	_, resizedImg, err := resizeImage(rawFile)
	if err != nil {
		return err
	}

	resizedRawWebp, err := utils.ConvertToWebp(resizedImg)
	if err != nil {
		return err
	}

	file := tele.FromReader(bytes.NewReader(resizedRawWebp))
	uploadedFile, err := fnc.UploadSticker(c.Bot().Me, &file)

	sticker := &fnc.InputSticker{
		Sticker:    uploadedFile.FileID,
		Emoji_list: emojiList,
	}

	if err != nil {
		return err
	}

	request := &fnc.AddStickerToSetParams{
		UserId:  stickerSetDetails["owner_id"].(int64),
		Name:    stickerSetDetails["sticker_id"].(string),
		Sticker: *sticker,
	}

	err = fnc.AddStickerToSet(request)
	if err != nil {
		return err
	}

	return c.Reply("Sticker added to set successfully!")
}
