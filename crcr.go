package main

import ( 

	"fmt"
	"flag"
	"bufio"
	"io/ioutil"
	"strings"
	"os"
	
)

var charsLower = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
var charsUpper = []string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
var charsNumber = []string{"0","1","2","3","4","5","6","7","8","9"}
var charsSpecial = []string{"ä","Ä","ö","Ö","ü","Ü","!","\"","§","%","&","/","(",")","=","?","{","[","]","}","\\","`","¸","*","+","-","_",".",":",",",";","<",">"}
var resultMap = map[string]int{}

func Start() string {
	return "Start"
}

func HandleArgs() string {
	var file string
	testtext := "testtext.txt"
	
	flag.StringVar(&file, "file", testtext, "file to analyse")
	flag.Parse()	
	if file != testtext {
		return file
	}
	return testtext
}

func CharInArray(char string, charArray []string) (map[string]int, bool) {
	isPresent := false
	for _, c := range charArray {
		counter := resultMap[c]
		if char == c {
			counter += 1
			resultMap[c] = counter
			isPresent = true
			return resultMap, isPresent
		}
	}
	return resultMap, isPresent
}


func ReadFile(file string) {
 
	filebuffer, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	var isPresent bool
	for data.Scan() {
		resultMap, isPresent = CharInArray(data.Text(), charsLower)
		if isPresent {
			continue
		}
		resultMap, isPresent = CharInArray(data.Text(), charsUpper)
		if isPresent {
			continue
		}
		resultMap, isPresent = CharInArray(data.Text(), charsNumber)
		if isPresent {
			continue
		}
		resultMap, isPresent = CharInArray(data.Text(), charsSpecial)
		if isPresent {
			continue
		}

	}

	fmt.Println(resultMap)
}

func main() {
	fmt.Println(Start())	
	ReadFile(HandleArgs())
}


