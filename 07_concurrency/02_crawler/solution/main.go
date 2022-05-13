package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/exp/slices"
)

func RequestLinks(URL string) ([]string, error) {
	var links []string
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		if href, exists := s.Attr("href"); exists && strings.HasPrefix(href, "http") {
			links = append(links, href)
		}
	})
	return links, nil
}

type Crawler struct {
	seenLinks []string
	mutex     sync.Mutex
}

func (c *Crawler) crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 1 {
		return
	}

	fmt.Printf("Making request to: %s\n", url)
	links, err := RequestLinks(url)
	if err != nil {
		fmt.Println(err)
	}

	for _, link := range links {
		if !slices.Contains(c.seenLinks, link) {
			c.mutex.Lock()
			c.seenLinks = append(c.seenLinks, link)
			c.mutex.Unlock()
			wg.Add(1)
			go c.crawl(link, depth-1, wg)
		}
	}
}

func (c *Crawler) Crawl(url string, depth int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go c.crawl(url, depth, &wg)
	wg.Wait()
}

/*
	Tasks:
		Create a crawler, which requests multiple websites at the same time with goroutines.
		Use WaitGroups, Channels, Mutexes to receive the links.
		Allow a specific "depth" to crawl.
		This means, that the crawler will follow "depth" websites until it returns.
		For example with a depth of 2:
			https://bykovski.de -> https://www.instagram.com/bykof/ -> stop
		Start with the website: https://bykovski.de

	Output with a depth of 2:
		Crawled 3 links
		- http://bykovski.de
		- https://www.instagram.com/bykof/
		- https://github.com/bykof

	Output with a depth of 3:
		Crawled 31 links
		- http://bykovski.de
		- https://www.instagram.com/bykof/
		- https://github.com/bykof
		- https://github.com/
		- https://lab.github.com/
		- https://opensource.guide
		- https://github.community
		- https://education.github.com
		- https://stars.github.com
		- https://github.com/enterprise/contact
		- https://avatars.githubusercontent.com/u/5264141?v=4
		- https://github.com/bykof?tab=followers
		- https://github.com/bykof?tab=following
		- https://github.com/aoe
		- https://archiveprogram.github.com/
		- https://docs.github.com/en/articles/blocking-a-user-from-your-personal-account
		- https://docs.github.com/en/articles/reporting-abuse-or-spam
		- https://docs.github.com/articles/why-are-my-contributions-not-showing-up-on-my-profile
		- https://skyline.github.com/bykof
		- https://docs.github.com/categories/setting-up-and-managing-your-github-profile
		- https://github.com
		- https://docs.github.com/en/github/site-policy/github-terms-of-service
		- https://docs.github.com/en/github/site-policy/github-privacy-statement
		- https://github.com/security
		- https://www.githubstatus.com/
		- https://docs.github.com
		- https://support.github.com?tags=dotcom-footer
		- https://github.com/pricing
		- https://services.github.com
		- https://github.blog
		- https://github.com/about

*/
func main() {
	crawler := Crawler{}
	crawler.Crawl("https://bykovski.de", 3)
	fmt.Println("Done")
	fmt.Printf("Crawled %d links\n", len(crawler.seenLinks))
	for _, seenLink := range crawler.seenLinks {
		fmt.Printf("- %s\n", seenLink)
	}
}
