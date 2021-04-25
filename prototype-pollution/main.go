package main

import (
	"context"
	"log"
	"bufio"
	"sync"
	"flag"
	"os"
	"github.com/chromedp/chromedp"
)

func main() {
	var userJS string
	flag.StringVar(&userJS, "j", "", "the JS to run on each page")
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	
	var wg sync.WaitGroup
	urls := make(chan string)
	for i := 0; i < 5; i++ {
	wg.Add(1)
		go func(){
			for u := range urls{
				ctx, cancel := chromedp.NewContext(context.Background())


				var res string
				err := chromedp.Run(ctx,
					chromedp.Navigate(u),
					chromedp.Evaluate(userJS, &res),
				)
				cancel()
				
				if err != nil {
					log.Printf("error on %s: %s", u,err)
					continue
				}

				log.Printf("%s: %v", u, res)
				}
				wg.Done()	
		}()
	}
	
	for sc.Scan() {
		u := sc.Text()	
		urls <- u
	}
	wg.Wait()
	
}
