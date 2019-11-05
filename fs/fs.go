package fs

import (
	"bytes"
	"io/ioutil"
	"strings"

	"../sh"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// NewPort change the port in the file
func NewPort(port string, path string) {

	dat, err := ioutil.ReadFile(path)
	check(err)

	//dat = contenu du fichier
	contenu := string(dat)
	listenIndex := strings.Index(contenu, "listen") + 8 //isolation après listen
	tempString := contenu[listenIndex:]

	semiconIndex := strings.Index(tempString, ";")

	firstPart := contenu[:listenIndex]

	secondPart := contenu[listenIndex+semiconIndex:]

	//concat the 2 parts and change the port
	buf := bytes.Buffer{}
	buf.WriteString(firstPart)
	buf.WriteString(port)
	buf.WriteString(secondPart)
	result := buf.String()

	//fmt.Println(result)

	//write in file
	d1 := []byte(result)
	err = ioutil.WriteFile(path, d1, 0644)
	check(err)

	sh.RunShFile("./test/script.sh") //Users/masterplow/Desktop/GoWriteFile/test/script.sh
}
