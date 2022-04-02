package utils

import (
	"cronicle/ui/constants"
	"io/fs"
	"io/ioutil"
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

func GetDataFromFile(path string) string {
	dat, _ := os.ReadFile(path)
	return string(dat)
}

func GetAllFiles(dataType string) []fs.FileInfo {
	dirPath := GetPath([]string{dataType})
	files, _ := ioutil.ReadDir(dirPath)
	return files
}

func WriteToFile(m string, fn string) {
	// create or replace
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
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

type Header struct {
	Date string   `yaml:"date"`
	Due  string   `yaml:"due"`
	Type string   `yaml:"type"`
	Tags []string `yaml:"tags"`
}

func ParseHeader(content string) Header {
	var matter struct {
		Date string   `yaml:"date"`
		Due  string   `yaml:"due"`
		Type string   `yaml:"type"`
		Tags []string `yaml:"tags"`
	}

	_, err := frontmatter.Parse(strings.NewReader(content), &matter)
	if err != nil {
		log.Println(err)
	}

	return matter
}
