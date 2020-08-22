# Contributing

When contributing to this repository, please first discuss the change you wish to make via issue,
email, or any other method with the owners of this repository before making a change.

Please note we have a [code of conduct](/CODE_OF_CONDUCT.md), please follow it in all your interactions with the project.

## Pull Request Process

1. Write proper Documentation in the Source Code so the `godoc` can pick them up and use in the API Documentation
2. Update the README.md with details of changes to the interface, this includes new environment variables, exposed ports, useful file locations and container parameters.
3. Increase the version numbers in any examples files and the README.md to the new version that this Pull Request would represent. The versioning scheme we use is [SemVer](http://semver.org/).
4. Your Pull Request may be auto-rejected if it breaks the build-process (github workflow). can take a look at it [here](/.github/workflows/go.yml)
5. To just add a new Spinner update the [spinnerList.go](/spinnerList.go) file with the spinner symbols and an unique ID (below 1000 if all symbols are ASCII). Then add a gif of the Spinner in the README.md file.
