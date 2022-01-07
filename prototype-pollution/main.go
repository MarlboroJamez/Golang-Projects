package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

var (
	concurrency int
	urls        bool
	outputFile  string
)

const (
	//InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor = "\033[1;36m%s\033[0m"
	//WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor = "\033[1;31m%s\033[0m"
	//DebugColor   = "\033[0;36m%s\033[0m"
)

func banner() {
	fmt.Println("|_@Marlboro_Jamez_|")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		// create context
		url := sc.Text()
		//fmt.Println(url)
		ctx, cancel := chromedp.NewContext(context.Background())

		// run task list
		var res string
		if urls == true {
			err := chromedp.Run(ctx,
				chromedp.Navigate(url+"&__proto__[protoscan]=protoscan"),
				chromedp.Evaluate(`window.protoscan`, &res),
			)
			cancel()
			if err != nil {
				log.Printf(ErrorColor, url+" [Not Vulnerable]")
				continue
			}
		} else {
			err := chromedp.Run(ctx,
				chromedp.Navigate(url+"/"+"?__proto__[protoscan]=protoscan"),
				chromedp.Evaluate(`window.protoscan`, &res),
			)
			cancel()
			if err != nil {
				log.Printf(ErrorColor, url+" [Not Vulnerable]")
				continue
			}
		}
		if outputFile != "" {
			f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Println(err)
			}
			if _, err := f.WriteString(url + "\n"); err != nil {
				log.Fatal(err)
			}
			f.Close()
		}
		log.Printf(NoticeColor, url+" [Vulnerable]")
	}
}

