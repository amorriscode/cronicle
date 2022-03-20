package daily

import (
	"cronicle/ui/components/section"
	"cronicle/ui/context"
)

func New(ctx *context.Context) section.Model {
	return section.New(ctx, "daily")
}
