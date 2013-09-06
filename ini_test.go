package ini

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	src := `
  # Comments are ignored

  herp = derp

  [foo]
  hello=world
  whitespace should   =   not matter   
  ; sneaky semicolon-style comment
  multiple = equals = signs

  [bar]
  this = that
  `

	file, err := Load(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	check := func(section, key, expect string) {
		if value, _ := file.Get(section, key); value != expect {
			t.Errorf("Get(%q, %q): expected %q, got %q", section, key, expect, value)
		}
	}

	check("", "herp", "derp")
	check("foo", "hello", "world")
	check("foo", "whitespace should", "not matter")
	check("foo", "multiple", "equals = signs")
	check("bar", "this", "that")
}
