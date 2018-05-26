package utils

import (
    "golang.org/x/net/html"
    "log"
    "net/url"
    "path"
    "strings"
)

//Given an html.Token, returns an string the value of the href attribute of the token
//If the Token doesn't have any href parameter, it returns false in the second value.
func GetHref(t html.Token) (string, bool) {
    for _, a := range t.Attr {
        if a.Key == "href" {
            return a.Val, true
        }
    }
    return "", false
}

//Given an original path, a base path and the new path, generates a new path
func AppendPath(o string, b string, p string) string {
    var s string
    if path.IsAbs(p) {
        s = o
    } else {
        s = b
    }

    u, err := url.Parse(s)
    if err != nil {
        log.Fatal(err)
    }
    u.Path = path.Join(u.Path, p)
    s = u.String()
    return s
}

//Given an URL it returns a boolean indicating if we can navigate to that URL
func IsValidURL(u string) bool {
    return !strings.Contains(u, ":") && !strings.Contains(u, "//") && !strings.HasPrefix(u, "#")
}
