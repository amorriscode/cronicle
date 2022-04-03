package daily

import (
	"cronicle/ui/components/section"
	"cronicle/ui/context"
	"cronicle/utils"
)

func NewSectionUI(ctx *context.Context) section.Model {
	return section.New(ctx, "daily")
}

func Edit() {
	utils.WriteOrCreateDaily(utils.WriteDailyParams{Message: "What did you work on?", Tags: ""})

	files := utils.GetAllFiles("daily")
	path := utils.GetPath([]string{"daily", files[len(files)-1].Name()})

	utils.EditFile(path)
}
