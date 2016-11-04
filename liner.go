package main

import "github.com/mitchellh/go-homedir"
import "github.com/peterh/liner"
import "path/filepath"
import "fmt"
import "log"
import "io"
import "os"

const (
	promptDefault = "kiku> "
)

type wsLiner struct {
	*liner.State
	buffer string
}

func newWsLiner() *wsLiner {
	l := liner.NewLiner()
	l.SetCtrlCAborts(true)
	return &wsLiner{State: l}
}

func (wl *wsLiner) Prompt(s string) (string, error) {
	if s == "" {
		s = promptDefault
	}
	l, e := wl.State.Prompt(s)

	if e == io.EOF {
		if wl.buffer != "" {
			// cancel line continuation
			wl.Accepted()
			fmt.Println()
			e = nil
		}
	} else if e == liner.ErrPromptAborted {
		e = nil
		if wl.buffer != "" {
			wl.Accepted()
		} else {
			fmt.Println("(^D to quit)")
		}
	} else if e == nil {
		wl.buffer = l
	}

	return wl.buffer, e
}

func (wl *wsLiner) Accepted() {
	wl.State.AppendHistory(wl.buffer)
	wl.buffer = ""
}

func (wl *wsLiner) readInput() string {
	i, e := wl.Prompt("")
	if e != nil {
		if e == io.EOF {
			// break
		}
		log.Fatal(e)
		os.Exit(1)
	}

	return i
}

func (wl *wsLiner) initHistoryFile() {
	name, e := homedir.Dir()
	if e != nil {
		return
	}

	name = filepath.Join(name, ".ws_history")

	file, e := os.Create(name)
	if e != nil {
		return
	}

	_, e = wl.WriteHistory(file)
	if e != nil {
		fmt.Errorf("init history file error: %s", e)
	}
}
