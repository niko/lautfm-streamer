package main

import (
	"fmt"
	"log"
	"io"
	"os"
	"time"
	"regexp"
	"strings"
	"io/ioutil"
	"github.com/howeyc/fsnotify"
	"github.com/systemfreund/go-libshout"
)

func createLautShout(pass string, mount string) (s shout.Shout) {
	s = shout.Shout{
		Host:     "live.laut.fm",
		Password: pass,
		Port:     uint(8080),
		User:     "source",
		Mount:    fmt.Sprintf("/%s.live", mount),
		Format:   shout.FORMAT_MP3,
		Protocol: shout.PROTOCOL_HTTP,
	}

	return
}

func readMetadataFile(filename string, metadataChange chan string) {
	md, err := ioutil.ReadFile(filename)
	if err != nil { log.Fatal(err) }
	metadataChange <- strings.TrimSpace(string(md))
}

func watchMetadataFile(filename string, metadataChange chan string) {
	fmt.Printf("watching '%s'", filename)
	watcher, err := fsnotify.NewWatcher()
	if err != nil { log.Fatal(err) }

	exp, _ := regexp.Compile( ".*(MODIFY|MOVE).*$" )
	deleted := false

	for {
		if _, err := os.Stat(filename); err == nil {
			if deleted {
				readMetadataFile(filename, metadataChange)
			}

			err = watcher.Watch(filename)
			if err != nil { log.Fatal(err) }
			ev := <-watcher.Event

			if len( exp.FindStringSubmatch(ev.String()) ) > 0 {
				readMetadataFile(filename, metadataChange)
			}
			deleted = false
		} else {
			deleted = true
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println("Arguments are station name and password. Metadata file as 3rd argument is optional.")
		os.Exit(1)
	}

	metadataFile := os.Args[3]

	s := createLautShout(os.Args[2], os.Args[1])

	input := os.Stdin

	stream, err := s.Open()
	if err != nil {
		fmt.Printf("Couldn't connect. Check station name (%s), password (%s).", os.Args[1], os.Args[2])
		os.Exit(2)
	}
	defer s.Close()

	go func() {
		buffer := make([]byte, shout.BUFFER_SIZE)
		for {
			n, err := input.Read(buffer)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}
			stream <- buffer
		}
	}()

	metadataChange := make(chan string)
	go watchMetadataFile(metadataFile, metadataChange)

	for {
		md := <- metadataChange
		fmt.Printf(md)
		s.UpdateMetadata("title", md)
	}
}
