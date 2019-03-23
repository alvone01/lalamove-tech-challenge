package main

import (
	"context"
	"fmt"
  "bufio"
  "os"
	"strings"
	"path/filepath"
	"github.com/coreos/go-semver/semver"
	"github.com/google/go-github/github"
)

// LatestVersions returns a sorted slice with the highest version as its first element and the highest version of the smaller minor versions in a descending order
func LatestVersions(releases []*semver.Version, minVersion *semver.Version) []*semver.Version {
	var versionSlice []*semver.Version

	//check if input is nil
  if releases == nil {
    fmt.Printf("Error : no releases found.")
  }

  if minVersion == nil {
    fmt.Printf("Error : please input minimum version")
  }

	//sort version array in descending order
	for i := 0; i<len(releases); i++ {
		for j := 0; j<len(releases)-1; j++ {
			if releases[j].Compare(*releases[j+1]) == -1 {
				temp := releases[j+1]
				releases[j+1] = releases[j]
				releases[j] = temp
			}
		}
	}

	//vars to keep track of succeeding filter function
	var currentMajor int64 = -1
	var currentMinor int64 = -1

	//append all highest stable minor versions to versionSlice
	for _, release := range releases {
		if release.Compare(*minVersion) == -1 || release.PreRelease != "" {
			continue
		}

		if release.Major != currentMajor || release.Minor != currentMinor {
			currentMajor = release.Major
			currentMinor = release.Minor

			versionSlice = append(versionSlice, release)

		}
	}

	return versionSlice

}

// function to parse lines of input from file
func parseInput(path string) [][]string {

	absPath, _ := filepath.Abs(path)

	file, err := os.Open(absPath)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtArr [][]string

	for scanner.Scan() {
		arr := strings.FieldsFunc(scanner.Text(), Split)
		txtArr = append(txtArr, arr)
	}

	return txtArr
}

//function to support string-splitting with multiple delimiters
func Split(r rune) bool {
	return r == '/' || r == ','
}

//function to check and print errors
func check(err error) {
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
}


//function to process each line of input from the file
func findLatest(input []string) {

	client := github.NewClient(nil)
	ctx := context.Background()
	opt := &github.ListOptions{PerPage: 10}

	minVersion := semver.New(input[2])

	var releases []*github.RepositoryRelease

	//call repositories only until page where ~minVersion is found
  for {
    repos, resp, err := client.Repositories.ListReleases(ctx, input[0], input[1], opt)
  	check(err)

	releases = append(releases, repos...)

	tagName := *repos[0].TagName
	if tagName[0] == 'v' {
		tagName = tagName[1:]
	}

	if semver.New(tagName).Compare(*minVersion) == -1 {
		break
	}

	opt.Page = resp.NextPage
}

allReleases := make([]*semver.Version, len(releases))

	for i, release := range releases {
		versionString := *release.TagName
		if versionString[0] == 'v' {
			versionString = versionString[1:]
		}
		allReleases[i] = semver.New(versionString)
	}
	versionSlice := LatestVersions(allReleases, minVersion)

	fmt.Printf("latest version of %s/%s: %s\n", input[0], input[1], versionSlice)

}


func main() {

	reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter path: ")
  text, _ := reader.ReadString('\n')
	text = "./" + text
	a := text[:len(text)-1]


	arrOfInputs := parseInput(a)

	for _, input := range arrOfInputs {
		findLatest(input)
	}

}
