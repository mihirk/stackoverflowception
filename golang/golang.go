package stackoverflowceptionCheck

import "github.com/pkg/browser"

func stackoverflowceptionCheck(err error) bool {
	if err != nil {
		url := "http://stackoverflow.com/search?q="
		url += err.Error()
		browser.OpenURL(url)
		return true
	}
	return false
}
