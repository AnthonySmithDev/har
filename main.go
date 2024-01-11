package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const Temp = "tmp/"

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No filename")
		return
	}
	name := args[1]
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	har, err := UnmarshalHar(data)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range har.Log.Entries {
		parse, err := url.Parse(entry.Request.URL)
		if err != nil {
			continue
		}
		dirname := filepath.Dir(parse.Path)
		basename := filepath.Base(parse.Path)
		hasName := pathHasName(parse.Path)
		if entry.Response.Content.MIMEType == FluffyApplicationJSON {
			if !hasName {
				dirname = parse.Path
				basename = "data.json"
			}
		}
		if entry.Response.Content.MIMEType == TextHTML {
			if !hasName {
				dirname = parse.Path
				basename = "index.html"
			}
		}
		dir := Temp + parse.Host + dirname
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			continue
		}
		if entry.Response.Content.Text == nil {
			continue
		}
		filename := path.Join(dir, basename)
		if err := os.WriteFile(filename, []byte(*entry.Response.Content.Text), 0644); err != nil {
			continue
		}
		if parse.RawQuery != "" {
			filename := path.Join(dir, "query.txt")
			if err := os.WriteFile(filename, []byte(parse.RawQuery), 0644); err != nil {
				continue
			}
		}
		fmt.Println("Save: ", dir)
	}
}

func pathHasName(path string) bool {
	return filepath.Ext(path) != ""
}
