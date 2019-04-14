package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	elementPrefix = "│─ "
	depthPrefix   = "│  "
	emptyPrefix   = "  "
	lastPrefix    = "└─ "
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tree <directory>")
		return
	}
	err := List(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func List(dirpath string) error {
	fmt.Println(dirpath)
	return list(dirpath, "")
}

func list(dirpath string, prefix string) error {
	infos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return err
	}

	for idx, info := range infos {
		fmt.Print(prefix)
		if idx == len(infos)-1 {
			fmt.Print(lastPrefix)
		} else {
			fmt.Print(elementPrefix)
		}
		if info.IsDir() {
			fmt.Println(color.BlueString("%v", info.Name()))
		} else {
			fmt.Println(info.Name())
		}
		if info.IsDir() {
			if idx == len(infos)-1 {
				err = list(filepath.Join(dirpath, info.Name()), prefix+emptyPrefix)
				if err != nil {
					return err
				}
			} else {
				err = list(filepath.Join(dirpath, info.Name()), prefix+depthPrefix)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
