package tests

import (
	"fmt"
	"testing"

	"github.com/fatih/color"
	"github.com/imthaghost/goclone/parser"
)

func TestURLExtension(t *testing.T) {
	tables := []struct {
		url      string
		expected string
	}{
		{"https://tesla.com/main.css", ".css"},
		{"https://tesla.com/main.css?Asf341", ".css"},
		{"https://dribbble.com/css/home", ""},
	}
	for _, table := range tables {
		result := parser.URLExtension(table.url)
		expectedresult := table.expected
		if result != expectedresult {
			t.Error()
			red := color.New(color.FgRed).SprintFunc()
			fmt.Printf("%s URLFilename Failed: %s , expected %s got %s \n", red("[-]"), table.url, expectedresult, result)

		} else {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("%s Passing: %s \n", green("[+]"), table.url)
		}
	}
}
