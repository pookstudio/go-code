package main

import (
		"fmt"
		"os"
		"os/exec"
		"io/ioutil"
		"runtime"
		"strings"
		"log"
)

var (
	Println = fmt.Println
	Print = fmt.Print
)

var (
	OBJ = "obj"
	OUT = "bin"
	SRC = "src"
	LIB = "lib"
)

var clear map[string]func()
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		Println("Your platform is unsupported! :(")
	}
}

func menu() {
	clearScreen()
	Println("\n         (\\/)         tom-code")
	Println("         (..)    Compile <Code/> C++")
	Println("        c(\")(\")      pookstudio\n")
	Println("   _______________ Menu _______________ ")
	Println("  |                                    |")
	Println("  |   New Template     >>    new       |")
	Println("  |   Compile          >>    compile   |")
	Println("  |   Build            >>    build     |")
	Println("  |   Compile&Build    >>    b         |")
	Println("  |   Build&Run        >>    a         |")
	Println("  |   Clean            >>    clean     |")
	Println("  |   Run              >>    run       |")
	Println("  |   Contact          >>    contact   |")
	Println("  |____________________________________|\n")
	Println("    *** Changes OS clean Before Compile")     
	Println("        Example >> clean\n")
}

func mkdir() {
	os.MkdirAll(OUT, os.ModePerm)
	os.MkdirAll(OBJ, os.ModePerm)
	os.MkdirAll(SRC, os.ModePerm)
	os.MkdirAll(LIB + "/lib", os.ModePerm)
	os.MkdirAll(LIB + "/include", os.ModePerm)
}

func clean() {
	files, err := ioutil.ReadDir("bin/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		var err_del = os.Remove("bin/" + f.Name())
		if isError(err_del) {
			Println("Error")
		} else {
			Println("done deleting file")
		}
	}
}

func newTemplate() {

}

func isError(err error) bool {
	return (err != nil)
}

func main() {
	for {
		menu()
		/*
		files, err := ioutil.ReadDir("bin/")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			Println(f.Name())
		}
		/*  Input Select  */
		fmt.Print("  Enter Select >> ")
		var inselect string
		fmt.Scanf("%s", &inselect)
		if strings.Compare("exit", inselect) == 0 {
			Println()
			break
		} else if inselect == "new" {
			mkdir()
		}
		//fmt.Printf("%q",inselect)
	}
}