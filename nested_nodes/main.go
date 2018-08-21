package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
)

var (
	childrenParser = regexp.MustCompile(`\[(.+)\]`)
	nodeParser     = regexp.MustCompile(`^(\w+)(.*)$`)
)

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

// takes a string and returns a slice of string
func splitItems(s string) (items []string) {
	var item string
	var openParens int

	for i, l := 0, len(s); i < l; i++ {
		chr := string(s[i])
		switch chr {
		case "[":
			openParens++
		case "]":
			openParens--
		case ",":
			// if character is a comma, and parentheses open == parentheses closed
			if openParens == 0 {
				items = append(items, item)
				item = ""
				chr = ""
			}
		}

		// reconstruct item
		item = fmt.Sprintf("%s%s", item, chr)
	}

	items = append(items, item)
	return
}

// Accepts a string and returns a slice of
// a pointer to a node
func parse(s string) (desc []*node) {
	// if string is empty return result as
	// there are no more nodes left to process
	if s == "" {
		return
	}

	// find string value within atleast one pair of square brackets
	// convert to a slice of strings and take the internal string at
	// index 1.
	// Pass to splitItems function to analyze and convert to a slice of strings
	children := splitItems(childrenParser.FindStringSubmatch(s)[1])

	for _, child := range children {
		// match any word and use word at index 1 as new node name,
		// pass the rest of values (at index 2) as Children of this
		// new node, and again process.
		details := nodeParser.FindStringSubmatch(child)
		desc = append(desc, &node{Name: details[1], Children: parse(details[2])})
	}

	return
}

var examples = []string{
	"[a,b,c]",
	"[t,h[is],h[a,r],d]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func main() {
	for i, example := range examples {
		// Send each string in examples slice
		// to parse function for processing
		r := parse(example)

		// convert r to a slice of byte with attached indenting
		jj, _ := json.MarshalIndent(r, " ", " ")

		// print result, converting jj bytes to string.
		log.Printf("Example %d: %s - \n%s", i, example, string(jj))
	}
}
