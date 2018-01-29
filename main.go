package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

type Files []string

func (f *Files) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *Files) Set(value string) error {
	*f = append(*f, value)
	return nil
}

var files Files

func main() {
	flag.Var(&files, "f", "List of files")
	flag.Parse()

	var wg sync.WaitGroup

	for _, path := range files {
		wg.Add(1)
		go convertGray(path, &wg)
	}

	wg.Wait()
}

func convertGray(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	cmdName := "/usr/local/bin/convert"
	cmdArgs := []string{path, "-colorspace", "Gray", "-separate", "-average", path}
	cmd := exec.Command(cmdName, cmdArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Exit(1)
	}

	optimize(path)
}

func optimize(path string) {
	cmdName := "/Applications/ImageOptim.app/Contents/MacOS/ImageOptim"
	cmdArgs := []string{path}
	cmd := exec.Command(cmdName, cmdArgs...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Exit(1)
	}
}
