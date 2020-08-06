package main

import (
	"encoding/json"
	"flag"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/0xAX/notificator"
	"github.com/fsnotify/fsnotify"
	hel "github.com/thejini3/go-helper"
)

// (required) will set in flags
var cfgFilepath string

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var scheduledPaths []string
var contents []theContent

var notify = notificator.New(notificator.Options{
	DefaultIcon: "default.png",
	AppName:     "git-sync",
})

func main() {
	// flags
	flags()

	hel.Pl("Starting git-sync")

	err := json.Unmarshal(hel.GetFileBytes(cfgFilepath), &contents)

	if err != nil {
		panic("Error filedata - " + cfgFilepath)
	}

	watch()
}

func flags() {
	flag.StringVar(&cfgFilepath, "f", "", "config file path, ex: -f ~/.config.git-sync.json")

	flag.Parse()

	if cfgFilepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func watch() {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		hel.Pl(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:

				if !ok {
					return
				}

				dir := filepath.Dir(event.Name)
				base := filepath.Base(dir)

				if content, found := getContentFromDirPath(dir); found {
					if hel.StrContains(content.IgnoreFiles, base) {
						return
					}
					execute(content)
				}

			case err, ok := <-watcher.Errors:

				if !ok {
					return
				}

				hel.Pl("watcher.Errors", "error:", err)
			}
		}
	}()

	for i := range contents {
		err = watcher.Add(contents[i].DirPath)
		if err != nil {
			hel.Pl("error in adding watcher in", contents[i], "err:", err)
		} else {
			hel.Pl("adding watcher", i+1)
			hel.PrettyPrint(&contents[i])
		}
	}

	<-done
}

func getContentFromDirPath(path string) (theContent, bool) {
	var tc theContent
	for i := range contents {
		if contents[i].DirPath == path {
			tc = contents[i]
			break
		}
	}
	return tc, len(tc.DirPath) > 0
}

func execute(c theContent) {

	if !hel.StrContains(scheduledPaths, c.DirPath) {
		hel.Pl("scheduling", c)
		scheduledPaths = append(scheduledPaths, c.DirPath)
	} else {
		return
	}

	time.AfterFunc(c.Delay*time.Second, func() {
		var outStr = ""
		for _, command := range c.Commands {

			for i := range command.Args {
				command.Args[i] = strings.ReplaceAll(command.Args[i], "$CURRENT_TIME$", time.Now().String())
				command.Args[i] = strings.ReplaceAll(command.Args[i], "$RANDOM$", getRandomStr(10))
			}

			var cmd *exec.Cmd

			if len(command.Args) == 0 {
				cmd = exec.Command(command.Command)
			} else {
				cmd = exec.Command(command.Command, command.Args...)
			}

			cmd.Dir = c.DirPath

			out, err := cmd.Output()

			hel.Pl(
				"`Command`", command.Command+" "+hel.ArrToStr(command.Args, " "),
				"\n`Err`", err,
				"\n`Output`", strings.TrimSpace(string(out)),
			)

			outStr += strings.TrimSpace(string(out)) + "\n"

		}

		no(c.DirPath, outStr)
		scheduledPaths = removeFromArray(scheduledPaths, c.DirPath)

	})

}

func removeFromArray(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getRandomStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func no(title string, desc string) {
	notify.Push(title, desc, "default.png", notificator.UR_NORMAL)
}
