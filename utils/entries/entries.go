package entries

import (
	"cronicle/ui/constants"
	"cronicle/utils"
	"cronicle/utils/prompts"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

func getDate() string {
	time := time.Now()
	return time.Format("2006-01-02")
}

func WriteOrCreateEntry(w utils.WriteParams, t string) {
	path := utils.GetPath([]string{t, getDate() + ".md"})
	if _, err := os.Stat(path); err == nil {
		// current Entry file exists
		appendEntry(w, t)
	} else if errors.Is(err, os.ErrNotExist) {
		// no Entry file created yet
		createdDateTime := time.Now().Format("2006-01-02 15:04")
		m := fmt.Sprintf("- %s\n", w.Message)
		newEntry := composeEntry(utils.WriteParams{Message: m, Tags: w.Tags, Date: createdDateTime}, t)
		utils.WriteToFile(newEntry, utils.GetPath([]string{t, getDate() + ".md"}))
	} else {
		log.Fatal(constants.ERROR_WRITE_FILE, err)
	}

}

func appendEntry(w utils.WriteParams, t string) {
	// merge tags
	path := utils.GetPath([]string{t, getDate() + ".md"})
	entry := utils.GetDataFromFile(path)
	tags := utils.ParseHeader(entry).Tags
	content := utils.ParseContent(entry)

	// add tags only if they don't already exist
	if w.Tags != "" {
		for _, t := range strings.Split(w.Tags, ",") {
			if !utils.Contains(tags, t) {
				tags = append(tags, t)
			}
		}
	}

	// join previous messages and new entry together
	newContent := strings.Builder{}
	newContent.WriteString(content)
	m := fmt.Sprintf("\n- %s\n", w.Message)
	newContent.WriteString(m)

	// compose new entry with new content, joined tags and previous created at date
	newEntry := composeEntry(utils.WriteParams{Message: newContent.String(), Tags: strings.Join(tags, ","), Date: utils.ParseHeader(entry).Date}, t)

	utils.WriteToFile(newEntry, path)
}

func composeEntry(w utils.WriteParams, t string) string {
	output := strings.Builder{}
	// header
	output.WriteString("---\n")

	// date created
	date := fmt.Sprintf("date: %s\n", w.Date)
	output.WriteString(date)

	//type
	entryType := fmt.Sprintf("type: %s\n", t)
	output.WriteString(entryType)

	// tags
	tags := fmt.Sprintf("tags: [%s]\n", w.Tags)
	output.WriteString(tags)

	// footer
	output.WriteString("---\n")

	output.WriteString(strings.TrimSpace(w.Message))

	return output.String()
}

func EditEntry(args []string, t string) {
	files := utils.GetAllFiles(t)
	var id int

	if len(args) > 0 {
		// user has given an arguement
		argId := utils.GetIdFromArg(args, files)
		if argId == -1 {
			fmt.Printf("Invalid argument")
			return
		}
		id = argId
	} else {
		// user selects from options
		id = GetIdFromOptions(files, t, "update")
	}

	path := utils.GetPath([]string{t, files[id].Name()})

	utils.EditFile(path)
}

func DeleteEntry(args []string, t string) {
	files := utils.GetAllFiles(t)
	var id int

	if len(args) > 0 {
		// user has given an arguement
		argId := utils.GetIdFromArg(args, files)
		if argId == -1 {
			fmt.Printf("Invalid argument")
			return
		}
		id = argId
	} else {
		// user selects from options
		id = GetIdFromOptions(files, t, "delete")
	}

	utils.DeleteFile(files[id].Name(), t)
}

func GetIdFromOptions(files []fs.FileInfo, t string, action string) int {
	var id int
	if t == "todo" {
		id = GetIdFromTodoOptions(files, t, action)
	} else {
		id = GetIdFromEntryOptions(files, t, action)
	}
	return id
}

func GetIdFromEntryOptions(files []fs.FileInfo, t string, action string) int {
	options := prompts.GetEntryDisplayOptions(t, files)
	i, err := prompts.SelectEntry(options, action)

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	return i
}

func GetIdFromTodoOptions(files []fs.FileInfo, t string, action string) int {
	options := prompts.GetTodoDisplayOptions(files)
	i, err := prompts.SelectTodo(options, action)

	if err != nil {
		log.Fatal(constants.ERROR_PROMPT, err)
	}

	return i
}
