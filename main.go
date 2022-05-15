package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	//loop
	for {
		fetchAction()
		checkOutAction()
		mergeAction()
		gitAdd()
		gitCommit()
		gitPush()

		time.Sleep(time.Duration(24) * time.Hour)
	}
}

func fetchAction() {
	defer func() {
		if r := recover(); r != nil {
			println("Recovered for", r)
			fetchAction()
		}
	}()
	app := "git"
	arg0 := "fetch"
	arg1 := "upstream"
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
		//return
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
