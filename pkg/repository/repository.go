package repository

import (
	"fmt"
	"net/url"
	"strings"
)

var templates = map[string]string{
	"all": "https://github.com/vertigobr/openfaas-templates",
}

func getRepositoryURL(repository string) (string, error) {
	if strings.Contains(repository, "/") || strings.Contains(repository, ".") {
		_, err := url.ParseRequestURI(repository)
		if err != nil {
			repository = strings.ToLower(strings.TrimRight(repository, "://"))

			if !strings.HasPrefix(repository, "http") && !strings.HasPrefix(repository, "https") {
				repository = fmt.Sprintf("https://%s", repository)
				_, _ = getRepositoryURL(repository)
				return "", nil
			}

			return "", err
		}

		return repository, nil
	} else {
		u := templates[repository]
		if u == "" {
			return "", fmt.Errorf("repositório inválido")
		}

		return u, nil
	}
}

func listTemplates() {
	for template, url := range templates {
		fmt.Printf("%s    %s\n", template, url)
	}
}
