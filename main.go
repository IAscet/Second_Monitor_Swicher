package main

import (
	//"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for {

		k := <-keyboardChan
		if k.Message == types.WM_KEYDOWN {
			char := rune(k.VKCode)
			//fmt.Println(string(char))
			// to take inputs from keyboard in console if needed

			// q is f2
			if string(char) == "q" {
				pc := exec.Command("cmd", "/C", "%windir%/System32/DisplaySwitch.exe /internal")
				pc.Run()

			}

			// `` on input key on keyboard is À

			if string(char) == "À" {
				tv := exec.Command("cmd", "/C", "%windir%/System32/DisplaySwitch.exe /external")

				tv.Run()
			}
		}

	}
}
