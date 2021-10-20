package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

type fileInt struct {
	packagePath string
	filename    string
	iface       string
}

func (fi *fileInt) String() string {
	return fmt.Sprintf("%s %s", fi.packagePath, fi.iface)
}

type fileInts []fileInt

func (fis fileInts) String() string {
	var buffer bytes.Buffer
	for _, fi := range fis {
		buffer.WriteString(fi.String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

var fI fileInts

func scanFile(path string, cb func(line string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cb(line)
	}
}

func parseFile(dir, file string) {
	if strings.HasSuffix(file, "_test.go") {
		return
	}
	if strings.Contains(dir, "mock") {
		return
	}

	scanFile(file, func(line string) {
		if strings.Contains(line, "interface") {
			strs := strings.Split(line, " ")
			if len(strs) == 4 && strs[0] == "type" && strs[2] == "interface" {
				ifaceName := strs[1]
				if unicode.IsUpper([]rune(ifaceName)[0]) {
					fI = append(fI, fileInt{packagePath: dir, filename: file, iface: ifaceName})
				}
			}
		}
	})
}

func parseDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", path, file.Name())
		if file.IsDir() {
			parseDir(filePath)
		} else if strings.HasSuffix(file.Name(), ".go") {
			parseFile(path, filePath)
		}
	}
}

func BuildList() []string {
	fI = make(fileInts, 0)
	parseDir("src/pendo.io")
	pathList := make([]string, len(fI))
	for i, fi := range fI {
		pathList[i] = fi.filename
	}
	return pathList
}
