package main

import (
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
)

func main() {
	fmt.Println("Starting Playwright...")

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

	fmt.Println("Launching browser...")

	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	fmt.Println("Creating new page...")

	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	fmt.Println("Navigating to URL...")

	if _, err = page.Goto("https://www.americanexpress.com"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	loadStateOptions := playwright.PageWaitForLoadStateOptions{
		State: (*playwright.LoadState)(playwright.String("networkidle")),
	}
	// Wait for the page to load completely
	if err = page.WaitForLoadState(loadStateOptions); err != nil {
		log.Fatalf("could not wait for load state: %v", err)
	}

	fmt.Println("Capturing cookies...")

	cookies, err := page.Context().Cookies()
	if err != nil {
		log.Fatalf("could not get cookies: %v", err)
	}

	fmt.Println("Printing cookies...")
	fmt.Printf("%v\n", len(cookies))
	fmt.Println("Finished.")
}
