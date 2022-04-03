package brag

import (
	"cronicle/ui/constants"
	"cronicle/utils/entries"
	"cronicle/utils/prompts"
	"log"
)

func EditBrag() {
	bragOptions := entries.GetEntryDisplayOptions("brag")
	id, err := prompts.SelectEntry(bragOptions, "update")

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	entries.EditEntry("brag", id)
}

func DeleteBrag() {
	bragOptions := entries.GetEntryDisplayOptions("brag")
	id, err := prompts.SelectEntry(bragOptions, "delete")

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	entries.DeleteEntry("brag", id)
}
