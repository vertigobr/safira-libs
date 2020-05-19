package repository

import "testing"

func TestGetRepositoryURL(t *testing.T)  {
	listTemplates()
	getRepositoryURL("https://google.com")
	getRepositoryURL("google.com")
	getRepositoryURL("google")
	getRepositoryURL("/foo/bar")
}
