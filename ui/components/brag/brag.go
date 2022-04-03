package brag

import (
	"cronicle/ui/components/section"
	"cronicle/ui/context"
	"cronicle/utils"
	"cronicle/utils/entries"
)

func NewSectionUI(ctx *context.Context) section.Model {
	return section.New(ctx, "brag")
}

func Edit() {
	entries.WriteOrCreateEntry(utils.WriteParams{Message: "What are you proud of?", Tags: ""}, "brag")

	files := utils.GetAllFiles("brag")
	path := utils.GetPath([]string{"brag", files[len(files)-1].Name()})

	utils.EditFile(path)
}
