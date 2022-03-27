package utils

import (
	"cronicle/ui/constants"
	"log"
	"os"
	"strings"

	"github.com/adrg/frontmatter"
)

type WriteParams struct {
	Message, Date, Tags string
}

type UpdateParams struct {
	Message, Date, Tags string
	Number              int
}

func WriteToFile(m string, fn string) {
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(constants.ERROR_OPEN_FILE, err)
	}

	if _, err := f.Write([]byte(m)); err != nil {
		f.Close()
		log.Fatal(constants.ERROR_WRITE_FILE, err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(constants.ERROR_CLOSE_FILE, err)
	}
}

func ParseContent(content string) string {
	// Empty struct used to strip frontmatter
	var f struct{}
	var c string

	// Parse frontmatter and remove it
	rest, err := frontmatter.Parse(strings.NewReader(content), &f)
	if err != nil {
		log.Println(err)
		c = content
	} else {
		c = string(rest)
	}

	return c
}
