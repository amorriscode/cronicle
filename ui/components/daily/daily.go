package daily

import (
	"cronicle/ui/components/section"
	"cronicle/ui/context"
)

func NewSectionUI(ctx *context.Context) section.Model {
	return section.New(ctx, "daily")
}
