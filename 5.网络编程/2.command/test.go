package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//cmdStr := string("ipconfig /all")
	cmd := exec.Command("ipconfig", "all")
	ret, err := cmd.Output()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("net", ret)
}
