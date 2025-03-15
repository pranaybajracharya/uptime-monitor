package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"uptime_monitor/db"
)

type Uptime struct {
	url         string
	status      bool
	status_code int
}

var urls = []string{
	"https://www.pranaybajracharya.com.np",
	"https://www.pranaybajracharya.com.np/about",
	"https://www.pranaybajracharya.com.np/projects",
	"https://www.pranaybajracharya.com.np/blog",
	"https://www.pranaybajracharya.com.np/blog/debounce-control",
}

var checkInterval = 60 * time.Second

func main() {
	db := db.InitDB()
	defer db.Close()

	fmt.Println("Starting uptime monitor...")

	for {
		var wg sync.WaitGroup
		uptimeGroup := make(chan *Uptime, len(urls))

		for _, url := range urls {
			wg.Add(1)
			go func() {
				defer wg.Done()
				uptimeGroup <- checkUptime(url)
			}()
		}
		fmt.Println("Waiting uptime...")
		wg.Wait()
		fmt.Println("Waiting done...")
		for range len(urls) {
			up := <-uptimeGroup
			if up != nil {
				insertUptime(db, *up)
			}
		}
		time.Sleep(checkInterval)
	}

}

func checkUptime(url string) *Uptime {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error checking %s: %v\n", url, err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		fmt.Printf("%s is up (Status: %d)\n", url, res.StatusCode)
		uptime := Uptime{url: url, status: true, status_code: res.StatusCode}
		return &uptime

	} else {
		fmt.Printf("%s is down (Status: %d)\n", url, res.StatusCode)
		uptime := Uptime{url: url, status: false, status_code: res.StatusCode}
		return &uptime
	}
}

func insertUptime(db *sql.DB, uptime Uptime) {
	query := `INSERT INTO uptime (url, status, status_code) VALUES (?, ?, ?)`
	_, err := db.Exec(query, uptime.url, uptime.status, uptime.status_code)
	if err != nil {
		log.Fatal(err)
	}
}
