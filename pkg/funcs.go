package pkg

import (
	"strings"
	"text/template"
)

func isLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + 32
	}
	return ch
}

func isUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

var FuncMap template.FuncMap = template.FuncMap{
	"title": strings.Title,
	"upper": strings.ToUpper,
	"SnakeToInitCaps": func(s string) string {
		// Init the array with a lot of room to growth
		r := make([]rune, 0, len(s)+3)
		var prev rune = '_'
		for _, c := range s {
			if c == '_' {
				prev = c
				continue
			}
			if prev == '_' {
				r = append(r, toUpper(c))
			} else {
				r = append(r, c)
			}
			prev = c
		}
		return string(r)
	},
}
