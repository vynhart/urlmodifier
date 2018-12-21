package urlmodifier

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

func SubDomain(urls string) (string, error) {
	var subdomain string

	urlObj, err := url.Parse(urls)
	if err == nil {
		subdomain = strings.Split(urlObj.Hostname(), ".")[0]
	}

	return subdomain, err
}

func ReplaceSubDomain(urls string, subdomain string) (string, error) {
	urlObj, err := url.Parse(urls)
	if err != nil {
		return urls, err
	}

	host := urlObj.Hostname()
	port := urlObj.Port()

	domains := strings.Split(host, ".")
	if len(domains) < 1 {
		return urls, errors.New("subdomain not found")
	}

	domains[0] = subdomain
	host = strings.Join(domains, ".")
	if port != "" {
		urlObj.Host = fmt.Sprintf("%v:%v", host, port)
	} else {
		urlObj.Host = host
	}
	return urlObj.String(), nil
}

func RandomizeSubDomain(urls string, subdomains []string) (string, error) {
	subdomain := randomize(subdomains)

	return ReplaceSubDomain(urls, subdomain)
}

func ReplacePathComponent(urls string, component string, idx int) (string, error) {
	if idx < 1 {
		return urls, errors.New("index must greater than 0")
	}

	urlObj, err := url.Parse(urls)
	if err != nil {
		return urls, err
	}

	path := urlObj.Path
	paths := strings.Split(path, "/")
	if len(paths) <= idx {
		return urls, errors.New("index exceeds length of path")
	}

	paths[idx] = component
	urlObj.Path = strings.Join(paths, "/")

	return urlObj.String(), nil
}

func randomize(s []string) string {
	if len(s) < 2 {
		return s[0]
	} else {
		idx := time.Now().Nanosecond() % len(s)
		return s[idx]
	}
}
