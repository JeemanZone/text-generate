package main

import (
	"flag"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
)

var dataFile, templateFile string
var randomPick bool

func main() {
	getFlags()
	getArgs()

	text, err := os.ReadFile(dataFile)
	check(err)
	data := getSimpleTextData(string(text))
	if randomPick {
		data = pick(data)
	}

	tmpl := template.Must(template.ParseFiles(templateFile))
	check(err)
	err = tmpl.Execute(os.Stdout, struct{ Data [][]string }{data})
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFlags() {
	rp_ptr := flag.Bool("rp", false, "Random pick mode")
	// flag.Parse()被调用后指针所指向的内存地址才会存放从命令行获取的值，在那之前只能取得默认值
	flag.Parse()
	randomPick = *rp_ptr
}

func getArgs() {
	args := flag.Args()
	dataFile = "data.txt"
	templateFile = "template.txt"

	if flag.NArg() > 0 {
		dataFile = args[0]
	}
	if flag.NArg() > 1 {
		templateFile = args[1]
	}
}

func getSimpleTextData(text string) [][]string {
	result := [][]string{}

	text = strings.Replace(text, "\r\n", "\n", -1)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		row := []string{}
		values := strings.Split(line, "\t")
		for _, value := range values {
			value = strings.TrimSpace(value)
			if len(value) > 0 {
				row = append(row, value)
			}
		}
		result = append(result, row)
	}

	return result
}

func pick(data [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(data))
	result := [][]string{}
	result = append(result, data[index])
	return result
}
