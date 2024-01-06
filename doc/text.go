/*
 * Source text
 * Copyright 2024 John Douglas Pritchard, Syntelos
 */
package main

import (
	"fmt"
	"log"
	"os"
	notes "github.com/syntelos/go-notes"
)

func Enumerate(this notes.Page) {
	var bar int
	for tix, txt := range this {

		if txt.IsText() {

			bar = tix
			break
		}
	}

	var cc int = 0

	switch operand() {

	case "head":
		fmt.Print("var encodehead Page = Page{")
		for tix, txt := range this {

			if tix < bar {
				if 0 != cc {
					fmt.Print(",")
				}
				fmt.Print(" Text{")
				for cx, ch := range txt {
					if 0 != cx {
						fmt.Print(",")
					}
					fmt.Printf(" 0x%02X",ch)
				}
				fmt.Print(" }")
				cc += 1
			}
		}
		fmt.Println("}")

		os.Exit(0)

	case "tail":
		fmt.Print("var encodetail Page = Page{")
		for tix, txt := range this {

			if bar < tix {
				if 0 != cc {
					fmt.Print(",")
				}
				fmt.Print(" Text{")
				for cx, ch := range txt {
					if 0 != cx {
						fmt.Print(",")
					}
					fmt.Printf(" 0x%02X",ch)
				}
				fmt.Print(" }")
				cc += 1
			}
		}
		fmt.Println("}")

		os.Exit(0)
	default:
		usage()
	}

}

func List(this notes.Page) {
	var bar int
	for tix, txt := range this {

		if txt.IsText() {

			bar = tix
			break
		}
	}

	switch operand() {

	case "head":
		for tix, txt := range this {

			if tix < bar {

				fmt.Println(string(txt))
			}
		}

		os.Exit(0)

	case "tail":
		for tix, txt := range this {

			if bar < tix {

				fmt.Println(string(txt))
			}
		}

		os.Exit(0)
	default:
		usage()
	}

}

func open() (fi *os.File) {
	var er error
	fi, er = os.Open("text.svg")
	if nil != er {
		fi, er = os.Open("doc/text.svg")
		if nil != er {
			return nil
		}
	}
	return fi
}

func operator() (opr string) {
	if 1 < len(os.Args) {
		return os.Args[1]
	} else {
		return opr
	}
}

func operand() (opd string) {
	if 2 < len(os.Args) {
		return os.Args[2]
	} else {
		return opd
	}
}

func usage() {
	fmt.Println(`
Synopsis

  text list                   -- List (head|tail).

  text enumerate              -- Enumerate (head|tail).

`)
	os.Exit(1)
}

func main(){

	var file *os.File = open()
	if nil != file {
		defer file.Close()

		var page notes.Page
		var er error

		page, er = page.Read(file)
		if nil != er {
			log.Fatal(er)
		} else {
			switch operator() {

			case "enumerate":
				Enumerate(page)

			case "list":
				List(page)

			default:
				usage()
			}
		}
	} else {
		log.Fatal("Missing source text.")
	}
}