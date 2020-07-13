package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/gmiclotte/gofsplit/of"

	log "github.com/sirupsen/logrus"
	//"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type postType int

const (
	text postType = iota
	image
	video
)

func main() {
	fmt.Println("Welcome to GOFSplit")

	referer := flag.String("referer", "", "")
	target := flag.String("target", "", "")
	token := flag.String("token", "", "")
	cookie := flag.String("cookie", "", "")
	agent := flag.String("agent", "", "")
	flag.Parse()

	// Declare http client
	client := &http.Client{}

	// Declare HTTP Method and Url
	req, err := http.NewRequest("GET", *referer+"api2/v2/users/"+*target+"/posts?limit=10&offset=0&order=desc&app-token="+*token, nil)

	// Set cookie
	req.Header.Set("Cookie", *cookie)
	req.Header.Set("User-Agent", *agent)
	req.Header.Set("Referer", *referer)
	req.Header.Set("Accept", "application/json")
	log.Infof("%+v", req)
	resp, err := client.Do(req)
	// Read response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error = %s ", err)
	}
	var parsed []of.Response
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		log.Error(err)
	}
	// sort files
	directoryName := "output"
	os.Mkdir(directoryName, 0755)
	imgPath := filepath.Join(directoryName, "img")
	vidPath := filepath.Join(directoryName, "vid")
	textPath := filepath.Join(directoryName, "text")
	os.Mkdir(imgPath, 0755)
	os.Mkdir(vidPath, 0755)
	os.Mkdir(textPath, 0755)
	// process response
	for _, entry := range parsed {
		date := strings.Split(entry.PostedAtPrecise, ".")[0]
		if _, err := os.Stat(filepath.Join(textPath, date)); err == nil {
			log.Error(filepath.Join(textPath, date))
			continue
		}
		os.Mkdir(filepath.Join(textPath, date), 0755)
		log.Infof("%s", entry.PostedAt)
		for _, m := range entry.Media {
			switch m.Type {
			case "photo":
				os.Mkdir(filepath.Join(imgPath, date), 0755)
				putFile(m.Full, imgPath, date, client)
			case "video":
				os.Mkdir(filepath.Join(imgPath, date), 0755)
				putFile(m.Preview, imgPath, date, client)
				os.Mkdir(filepath.Join(vidPath, date), 0755)
				putFile(m.Full, vidPath, date, client)
			default:
			}
		}
	}
	return
}

func putFile(src, path, date string, client *http.Client) {
	urlPtr, _ := url.Parse(src)
	split := strings.Split(urlPtr.Path, "/")
	name := split[len(split)-1]
	file, _ := os.Create(filepath.Join(path, date, name))
	log.Info(filepath.Join(path, date, name))
	log.Infof("Downloading %s", src)
	resp, err := client.Get(src)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Infof("Saved at %s, with size %d", file.Name(), size)
}
