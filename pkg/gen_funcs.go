package pkg

import (
	"fmt"
	"strings"
	"text/template"
)

var (
	c1 int
)

func leftPad(number int, padding int) string {
	/*
		i, err := strconv.Atoi(number)
		if err != nil {
			return number
		}

	*/
	format := fmt.Sprintf("%%0%vd", padding)
	return fmt.Sprintf(format, number)
}

/// c1gin counter 1 get and increase
func c1gi() int {
	//s := fmt.Sprintf("%v", c1)
	c1_copy := c1
	c1++
	return c1_copy
}

/// c1g counter 1 get
func c1g() int {
	return c1
	// return fmt.Sprintf("%v", c1)
}

func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

func snakeToInitCaps(s string) string {
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
}

var FuncMap template.FuncMap = template.FuncMap{
	"title":           strings.Title,
	"upper":           strings.ToUpper,
	"split":           strings.Split,
	"trimSpace":       strings.TrimSpace,
	"SnakeToInitCaps": snakeToInitCaps,
	"c1gi":            c1gi,
	"c1g":             c1g,
	"leftPad":         leftPad,
	"replaceAll":      strings.ReplaceAll,
}

/*
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
*/
