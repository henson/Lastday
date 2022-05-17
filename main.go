package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var readme, timeStamp string
	//loop
	for {
		timeStamp = time.Now().Format("2006/01/02")
		scrape(timeStamp)

		readme = "# Lastday\n\nhttps://henson.github.io/Lastday/" + timeStamp + ".json for visit"
		writeMarkDown("README", readme)

		fetchAction()
		checkOutAction()
		mergeAction()

		gitAdd()
		gitCommit()
		gitPush()

		//waiting for nextday
		NowStamp, _ := time.ParseInLocation("2006/01/02", timeStamp, time.Local)
		time.Sleep(time.Until(NowStamp.AddDate(0, 0, 1)))
	}
}

func scrape(times string) {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if r := recover(); r != nil {
			println("Recovered for", interface2string(r))
			//Waiting for about 5 Minutes
			time.Sleep(time.Duration(60*rand.Intn(5000)) * time.Millisecond)
			scrape(times)
		}
	}()

	// Request the HTML page.
	res, err := http.Get("https://github.com/sanddudu/LastDay/blob/gh-pages/" + times + ".json")
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err.Error())
	}

	if doc.Find("head > title").Text() == "Page not found · GitHub · GitHub" {
		panic("Page not found")
	}
}

func interface2string(inter interface{}) string {
	var tempStr string
	switch inter.(type) {
	case string:
		tempStr = inter.(string)
	case float64:
		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
	case int64:
		tempStr = strconv.FormatInt(inter.(int64), 10)
	case int:
		tempStr = strconv.Itoa(inter.(int))
	}
	return tempStr
}

func fetchAction() {
	app := "git"
	arg0 := "fetch"
	arg1 := "upstream"
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func checkOutAction() {
	app := "git"
	arg0 := "checkout"
	arg1 := "gh-pages"
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func mergeAction() {
	app := "git"
	arg0 := "merge"
	arg1 := "--allow-unrelated-histories"
	arg2 := "upstream/gh-pages"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func gitAdd() {
	app := "git"
	arg0 := "add"
	arg1 := "."
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func gitCommit() {
	app := "git"
	arg0 := "commit"
	arg1 := "-am"
	arg2 := time.Now().Format("2006-01-02 15:04:05")
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func gitPush() {
	app := "git"
	arg0 := "push"
	arg1 := "origin"
	arg2 := "gh-pages"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(out))
}

func writeMarkDown(fileName, content string) {
	// open output file
	fo, err := os.Create(fileName + ".md")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)
	if _, err := w.WriteString(content); err != nil {
		println(err.Error())
	}
	w.Flush()
}
