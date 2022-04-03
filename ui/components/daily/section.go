package daily

import (
	"cronicle/ui/components/section"
	"cronicle/ui/context"
	"cronicle/utils"
	"cronicle/utils/entries"
)

func NewSectionUI(ctx *context.Context) section.Model {
	return section.New(ctx, "daily")
}

func Edit() {
	entries.WriteOrCreateEntry(utils.WriteParams{Message: "What did you work on?", Tags: ""}, "daily")

	files := utils.GetAllFiles("daily")
	path := utils.GetPath([]string{"daily", files[len(files)-1].Name()})

	utils.EditFile(path)
}
