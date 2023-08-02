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

func Initialize(c tele.Context) error {
	rawArgs := strings.Split(c.Message().Text, " ")
	if len(rawArgs) < 4 {
		return c.Reply("Usage: /initialize <sticker_name> <emojis> <title>\nsticker_name: ID of the sticker (alphanumeric and underscore only)\nemojis: Emoji(s) that represents the sticker (no spaces between)\ntitle: Title of the sticker pack")
	}

	args := make(map[string]string)

	args["name"] = rawArgs[1]
	stickerName := args["name"] + "_by_" + c.Bot().Me.Username
	args["emojis"] = rawArgs[2]
	args["title"] = strings.Join(rawArgs[3:], " ")

	emojiList := []string{args["emojis"]}
	log.Printf("emojiList: %v", emojiList)
	if !utils.IsAllEmoji(emojiList) || len(emojiList) > 20 {
		return c.Reply("Invalid emoji")
	}

	chatDetails, err := models.GetChat(c.Chat().ID)
	if err != nil {
		return err
	}
	if chatDetails != nil {
		return c.Reply("This chat already has a sticker set. Please use /convert instead, or /getpack to get the sticker set.")
	}

	stickerSetDetails, err := models.GetSticker(stickerName)
	if err != nil {
		return err
	}
	if stickerSetDetails != nil {
		return c.Reply("This sticker name already exists. Please use another sticker name.")
	}

	var photoId string = ""
	if photoId = utils.ExtractFileId(c.Message()); photoId == "" && c.Message().ReplyTo != nil {
		photoId = utils.ExtractFileId(c.Message().ReplyTo)
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
	if err != nil {
		return err
	}

	log.Printf(uploadedFile.FileID)

	startingSticker := &fnc.InputSticker{
		Sticker:    uploadedFile.FileID,
		Emoji_list: emojiList,
	}

	stickerSet := &fnc.CreateNewStickerSetParams{
		UserId:        c.Message().Sender.ID,
		Name:          stickerName,
		Title:         args["title"] + "[made with" + c.Bot().Me.Username + "]",
		Stickers:      []fnc.InputSticker{*startingSticker},
		StickerFormat: "static",
		StickerType:   "regular",
	}

	addChat := &models.AddChatParams{
		ChatId:    c.Message().Chat.ID,
		StickerId: stickerName,
	}
	addSticker := &models.AddStickerParams{
		StickerId: stickerName,
		OwnerId:   c.Message().Sender.ID,
	}

	err = fnc.CreateNewStickerSet(stickerSet)
	if err != nil {
		return err
	}

	err = models.AddChat(*addChat)
	if err != nil {
		return err
	}

	err = models.AddSticker(*addSticker)
	if err != nil {
		return err
	}

	return c.Send("Sticker pack created!\nLink: https://t.me/addstickers/" + stickerName)

}
