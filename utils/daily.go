package utils

import (
	"cronicle/ui/constants"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type WriteDailyParams struct {
	Message, Tags, Date string
}

func GetDate() string {
	time := time.Now()
	return time.Format("2006-01-02")
}

func AddToCurrentDaily(w WriteDailyParams) {
	path := GetPath([]string{"daily", GetDate() + ".md"})
	if _, err := os.Stat(path); err == nil {
		// current daily file exists
		AddEntry(w)
	} else if errors.Is(err, os.ErrNotExist) {
		// no daily file created yet
		createdDateTime := time.Now().Format("2006-01-02 15:04")
		m := fmt.Sprintf("- %s\n", w.Message)
		newDaily := ComposeDaily(WriteDailyParams{Message: m, Tags: w.Tags, Date: createdDateTime})
		WriteToFile(newDaily, GetPath([]string{"daily", GetDate() + ".md"}))
	} else {
		log.Fatal(constants.ERROR_WRITE_FILE, err)
	}

}

func AddEntry(w WriteDailyParams) {
	// merge tags
	path := GetPath([]string{"daily", GetDate() + ".md"})
	daily := GetDataFromFile(path)
	tags := ParseHeader(daily).Tags
	content := ParseContent(daily)

	// add tags only if they don't already exist
	if w.Tags != "" {
		for _, t := range strings.Split(w.Tags, ",") {
			if !Contains(tags, t) {
				tags = append(tags, t)
			}
		}
	}

	// join previous daily messages and new entry together
	newContent := strings.Builder{}
	newContent.WriteString(content)
	m := fmt.Sprintf("\n- %s\n", w.Message)
	newContent.WriteString(m)

	// compose new daily with new content, joined tags and previous created at date
	newDaily := ComposeDaily(WriteDailyParams{Message: newContent.String(), Tags: strings.Join(tags, ","), Date: ParseHeader(daily).Date})

	WriteToFile(newDaily, path)
}

func ComposeDaily(w WriteDailyParams) string {
	output := strings.Builder{}
	// header
	output.WriteString("---\n")

	// date created
	date := fmt.Sprintf("date: %s\n", w.Date)
	output.WriteString(date)

	//type
	output.WriteString("type: daily\n")

	// tags
	tags := fmt.Sprintf("tags: [%s]\n", w.Tags)
	output.WriteString(tags)

	// footer
	output.WriteString("---\n")

	output.WriteString(strings.TrimSpace(w.Message))

	return output.String()
}
