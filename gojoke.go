package gojoke

import (
	"net/http"
	"strings"
	"io"
)

func clen(n []byte) int {
    for i := 0; i < len(n); i++ {
        if n[i] == 0 {
            return i
        }
    }
    return len(n)
}
func joke(endpoint string) (string, error) {
	res, err := http.Get("https://v2.jokeapi.dev/joke/"+endpoint)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	joke := body[:clen(body)]
	return string(joke), nil
}

func Any() (string, error) {
	return joke("Any?format=txt")
}

func Safe() (string, error) {
	return joke("Any?format=txt&safe-mode")
}

func Language(language string) (string, error) {
	return joke("Any?format=txt&lang="+language)
}

func Category(category string) (string, error) {
	return joke(category+"?format=txt")
}

func Joke(language string, category string, blacklist []string, single bool, twopart bool, safe bool) (string, error) {
	endpoint := category + "?format=txt"
	if !(single && twopart) {
		typ := ""
		if single {
			typ = "single"
		} else {
			typ = "twopart"
		}
		endpoint = endpoint+"&type="+typ
	}
	if len(blacklist) > 0 {
		endpoint = endpoint+"&blacklistFlags="+strings.Join(blacklist, ",")
	}
	if language != "en" {
		endpoint = endpoint+"&lang="+language
	}
	if safe {
		endpoint = endpoint+"&safe-mode"
	}
	return joke(endpoint)
}


