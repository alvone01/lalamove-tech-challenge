# Lalamove Tech Challenge
A coding challenge for the Lalamove Summer Internship Programme, written in Golang, implementing SemVer and Github API.

## Main libraries used
* go-semver
* go-github

## Rundown
* When program runs on the terminal via 'go run', it will request for an input in the form of a text file, this file should be on the working directory. "input.txt" can be used for testing.
* Implemented a parseInput function to parse the text file line by line into an array of strings.
* Based on how many lines there are, processing will run via a loop.
* Function will call repositories page by page via the Github API, until one version preceeds the minVersion.
* Array of repositories is converted into an array of semantic versions, which will then be processed through the LatestVersions function.
* This function will sort the array and filter out those that preceeds the minVersion and is not yet unstable. It will also only return one highest version for each minor version.
* The results will be printed on the console. If there is more than one line, it may take time before the second results appear on the console after the preceeding one.

## Notes
I have to say it was a very interesting journey working on this challenge. Figuring how things fit together in Golang, as it is my first time using this language. I definitely learnt a lot and crave to learn more. I hope there will be more journeys ahead.
