package daily

import (
	"cronicle/ui/constants"
	"cronicle/utils/entries"
	"cronicle/utils/prompts"
	"log"
)

func EditDaily() {
	dailyOptions := entries.GetEntryDisplayOptions("daily")
	id, err := prompts.SelectEntry(dailyOptions, "update")

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	entries.EditEntry("daily", id)
}

func DeleteDaily() {
	dailyOptions := entries.GetEntryDisplayOptions("daily")
	id, err := prompts.SelectEntry(dailyOptions, "delete")

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	entries.DeleteEntry("daily", id)
}
