package handlers

import (
	"bytes"
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

	emojiList := strings.Split(args["emojis"], "")
	log.Printf("emojiList: %v", emojiList)
	if !utils.IsAllEmoji(emojiList) || len(emojiList) > 20 {
		log.Printf("IsAllEmoji: %v", utils.IsAllEmoji(emojiList))
		return c.Reply("Invalid emoji(s). The number of emojis must be between 1 and 20, and must be all emojis.")
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

	var photo *tele.Photo
	if c.Message().Photo != nil {
		photo = c.Message().Photo
	} else if c.Message().ReplyTo != nil && c.Message().ReplyTo.Photo != nil {
		photo = c.Message().ReplyTo.Photo
	} else {
		return c.Reply("No photo found in message/replied message.")
	}

	photoId := photo.FileID

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

	_, resizedImg, err := utils.ResizeImage(rawFile)
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
