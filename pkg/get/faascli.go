package get

import (
	"fmt"
	c "github.com/vertigobr/safira-libs/pkg/config"
	"io"
	"net/http"
	"net/url"
	o "os"
)

const faasVersion = "0.12.2"

func getFaasCliUrl() string {
	var suffix string
	arch := getARCH()
	os := getOS()

	if os == "darwin" {
		suffix = "-darwin"
	} else if os == "mingw" || os == "windows" {
		suffix = ".exe"
	} else {
		if arch == "aarch64" {
			suffix = "-arm64"
		} else if arch == "armv6l" || arch == "armv7l" {
			suffix = "-armhf"
		}
	}

	return fmt.Sprintf("https://github.com/openfaas/faas-cli/releases/download/%s/faas-cli%s", faasVersion, suffix)
}

func DownloadFaasCli() error {
	faasCliUrl := getFaasCliUrl()
	parsedURL, _ := url.Parse(faasCliUrl)

	res, err := http.DefaultClient.Get(parsedURL.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	dest, err := c.CreateInBinDir()
	if err != nil {
		return err
	}

	// Criar arquivo
	out, err := o.Create(fmt.Sprintf("%s/faas-cli", dest))
	if err != nil {
		return err
	}
	defer out.Close()

	// Escreve o corpo da resposta no arquivo
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	if err := o.Chmod(fmt.Sprintf("%s/faas-cli", dest), 0700); err != nil {
		return err
	}

	return nil
}
