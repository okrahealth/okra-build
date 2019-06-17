package xmlparser

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

type SitemapIndex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"sitemap"`
	Loc     string   `xml:"loc"`
}

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func FetchSitemaps(url string) SitemapIndex {
	var sitemapIndex SitemapIndex
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(body, &sitemapIndex)
	return sitemapIndex
}

func FetchUrlSets(si SitemapIndex) []Urlset {
	urls := []string{}
	for i := 0; i < len(si.Sitemaps); i++ {
		urls = append(urls, si.Sitemaps[i].Loc)
	}
	urlSets := make([]Urlset, len(urls))
	for j := 0; j < len(urls); j++ {
		resp, _ := http.Get(urls[j])
		body, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(body, &urlSets[j])
	}
	return urlSets
}

func SaveProjectJSON(projects map[string]string) {
	fmt.Println("Starting to save files...")
	for name, url := range projects {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		currentTime := time.Now()
		file, err := os.Create(name + "_" + currentTime.Format("01-02-2006") + ".json")
		if err != nil {
			return
		}
		defer file.Close()

		file.WriteString(string(body))
	}
}

func GetProjectNameAndUrls(urlSets []Urlset) map[string]string {
	projects := map[string]string{}
	for i := 0; i < len(urlSets); i++ {
		for j := 0; j < len(urlSets[i].Urls); j++ {
			if IsProjectUrl(urlSets[i].Urls[j].Loc) {
				temp := parseProjectNameAndUrl(urlSets[i].Urls[j].Loc)
				projectName := temp["projectName"]
				// name => url
				projects[projectName] = temp[projectName]
			}
		}
	}
	return projects
}

func parseProjectNameAndUrl(url string) map[string]string {
	expression := regexp.MustCompile(`https://pypi.org/project/(?P<projectName>[a-zA-Z0-9-]+)/`)
	matches := expression.FindStringSubmatch(url)
	results := map[string]string{}
	projectNameUrlMap := map[string]string{}
	for i, name := range matches {
		results[expression.SubexpNames()[i]] = name
		projectNameUrlMap[expression.SubexpNames()[i]] = name
		projectNameUrlMap[name] = "https://pypi.org/pypi/" + name + "/json"
	}
	return projectNameUrlMap
}

func IsProjectUrl(url string) bool {
	r, err := regexp.Compile("pypi.org/project")
	if err != nil {
		fmt.Println(err)
	}
	matched := r.MatchString(url)
	return matched
}
