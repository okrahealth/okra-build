package main

import "fmt"
import "github.com/okrahealth/okra-build/pypidata"

func main() {
	fmt.Println("Starting downloads...")
	siteMapIndex := pypidata.FetchSitemaps("https://pypi.org/sitemap.xml")
	urlSets := pypidata.FetchUrlSets(siteMapIndex)
	projectMap := pypidata.GetProjectNameAndUrls(urlSets)
	pypidata.SaveProjectJSON(projectMap)
}
