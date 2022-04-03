package utils

import (
	"cronicle/ui/constants"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
)

type WriteParams struct {
	Message, Date, Tags string
}

func GetDataFromFile(path string) string {
	d, _ := os.ReadFile(path)
	return string(d)
}

func GetAllFiles(d string) []fs.FileInfo {
	p := GetPath([]string{d})
	f, _ := ioutil.ReadDir(p)
	return f
}

func WriteToFile(m string, fn string) {
	// create or replace
	f, err := os.OpenFile(fn, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
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

func EditFile(path string) {
	v := exec.Command("vim", path)
	v.Stdin = os.Stdin
	v.Stdout = os.Stdout
	if err := v.Run(); err != nil {
		log.Fatal(constants.ERROR_OPEN_FILE, err)
	}
}

func DeleteFile(f string, t string) {
	d := GetPath([]string{t})
	e := os.Remove(filepath.Join(d, f))

	if e != nil {
		log.Fatal(constants.ERROR_DELETE_FILE, e)
	}
}

func ListFiles(t string) {
	path := GetPath([]string{t})

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(constants.ERROR_LIST_FILE, err)
	}

	for i, f := range files {
		fmt.Printf("%v. %s", i+1, f.Name())
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
	var matter Header

	_, err := frontmatter.Parse(strings.NewReader(content), &matter)
	if err != nil {
		log.Println(err)
	}

	return matter
}
