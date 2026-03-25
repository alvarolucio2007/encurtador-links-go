// Package verifiers verifica os links
package verifiers

import (
	"net/url"
	"strings"
)

func ValidateURL(link string) bool {
	u, err := url.Parse(link)
	if err != nil {
		return false
	}
	if u.Scheme == "" || u.Host == "" {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	return scheme == "http" || scheme == "https"
}
