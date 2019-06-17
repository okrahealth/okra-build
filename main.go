package main

import "fmt"
import "github.com/okrahealth/okra-build/xmlparser"

func main() {
	fmt.Println("Starting downloads...")
	siteMapIndex := xmlparser.FetchSitemaps("https://pypi.org/sitemap.xml")
	urlSets := xmlparser.FetchUrlSets(siteMapIndex)
	projectMap := xmlparser.GetProjectNameAndUrls(urlSets)
	xmlparser.SaveProjectJSON(projectMap)
}
