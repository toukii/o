package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/everfore/exc"
	"github.com/toukii/bytes"
	"github.com/toukii/goutils"
)

func init() {
	dic = make(map[string]*Note)
	err := toml.Unmarshal(goutils.ReadFile(notesFile), &dic)
	if err != nil {
		log.Fatal(err)
	}
}

type Note struct {
	Val   string
	Exced bool
}

var (
	dic map[string]*Note

	notesFile = "notes.toml"
)

func refresh() {
	wr := bytes.NewWriter(make([]byte, 0, 1024))
	err := toml.NewEncoder(wr).Encode(dic)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(goutils.ToString(wr.Bytes()))

	goutils.WriteFile(notesFile, wr.Bytes())
}

func main() {
	args := os.Args
	size := len(args)
	switch size {
	case 0:
	case 1:
		fmt.Println("nil args")
	case 2:
		GetNote(args[1])
	default:
		SetNote(args[1:])
	}
}

func SetNote(args []string) {
	exced := false
	if args[1] == "-e" {
		exced = true
	}
	val := ""
	if exced {
		val = strings.Join(args[2:], " ")
	} else {
		val = strings.Join(args[1:], " ")
	}
	dic[args[0]] = &Note{
		Val:   val,
		Exced: exced,
	}
	refresh()
}

func GetNote(key string) {
	note, ex := dic[key]
	if !ex {
		fmt.Println("note nil")
		return
	}
	if !note.Exced {
		fmt.Printf("%s\n", note.Val)
		return
	}
	exc.Bash(note.Val).Debug(true).Execute()
}
