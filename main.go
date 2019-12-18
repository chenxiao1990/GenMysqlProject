package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
func main() {

	GinInit(8008)
	fmt.Println("浏览器访问 http://localhost:8008/vue")

	openbrowser("http://localhost:8008/vue")
	select {}
}
