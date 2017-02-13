package main

import "os/exec"

func main() {
	fetchAction()
	checkOutAction()
	mergeAction()
}

func fetchAction() {
	app := "git"
	arg0 := "fetch"
	arg1 := "upstream"
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func checkOutAction() {
	app := "git"
	arg0 := "checkout"
	arg1 := "gh-pages"
	cmd := exec.Command(app, arg0, arg1)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}

func mergeAction() {
	app := "git"
	arg0 := "merge"
	arg1 := "--allow-unrelated-histories"
	arg2 := "upstream/gh-pages"
	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}

	print(string(out))
}
