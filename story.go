package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	storyDir := "story"
	datDir := "../memu/dats"
	iniDir := "../memu/folders"

	fileInfoList, err := ioutil.ReadDir(storyDir)
	CheckErr(err)

	var storyBuffer bytes.Buffer
	var storyIniBuffer bytes.Buffer
	storyIniBuffer.WriteString("[story]\n")
	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() == true || !strings.HasSuffix(fileInfo.Name(), ".txt") {
			continue
		}
		storyBytes, err := ioutil.ReadFile(storyDir + "/" + fileInfo.Name())
		CheckErr(err)
		storyBuffer.WriteString(string(storyBytes) + "\n")

		// storyIniBuffer.WriteString(fileInfo.Name() + "\n")
		storyIniBuffer.WriteString(strings.TrimSuffix(fileInfo.Name(), ".txt") + "\n")
	}

	ioutil.WriteFile(datDir+"/story.dat", storyBuffer.Bytes(), os.ModePerm)
	ioutil.WriteFile(iniDir+"/Story.ini", storyIniBuffer.Bytes(), os.ModePerm)
}

func CheckErr(err error) {
	if nil != err {
		log.Println(err)
		//log.Fatal(err)
	}
}
