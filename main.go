package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type ManifestData struct {
	DownloadUrl string
	BundleId    string
	Title       string
}

var startPort int
var endPort int
var host string
var hostPath string
var saveDir string

func main() {
	flag.IntVar(&startPort, "startPort", 0, "The inclusive port to start from")
	flag.IntVar(&endPort, "endPort", 10001, "The exclusive port to end at")
	flag.StringVar(&host, "host", "http://127.0.0.1", "The redirection host")
	flag.StringVar(&hostPath, "hostPath", "ota-me", "The redirection host path")
	flag.StringVar(&saveDir, "saveDir", "data", "The generated files save directory")
	flag.Parse()

	t, err := template.New("").ParseFiles("manifest.xml")
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.MkdirAll("data", 0777); err != nil {
		log.Fatalln(err)
	}
	for i := startPort; i < endPort; i++ {
		if err := generate(t, i); err != nil {
			log.Fatalln(err)
		}
	}
}

func generate(t *template.Template, port int) error {
	var result bytes.Buffer
	err := t.ExecuteTemplate(&result, "manifest.xml", ManifestData{
		DownloadUrl: fmt.Sprintf("%s:%d/%s", host, port, hostPath),
		BundleId:    "com.foo.bar",
		Title:       "Local app",
	})
	if err != nil {
		return err
	}
	file := filepath.Join(saveDir, fmt.Sprintf("%d", port))
	if err := ioutil.WriteFile(file, result.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}
