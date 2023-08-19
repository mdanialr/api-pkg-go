package helper

import "strings"

// StrYNToBool return true only if given s is 'yes' otherwise will return
// false.
func StrYNToBool(s string) bool {
	return s == "yes"
}

// BoolYNToStr return 'yes' if true otherwise return 'no' instead.
func BoolYNToStr(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// DropBlankSpace return -1 if given rune is blank space ' ', this intended to
// be used in strings.Map to remove any blank space from string.
func DropBlankSpace(r rune) rune {
	if r == ' ' {
		return -1 // drop blank space
	}
	return r
}

// ToSnake converts a string to snake_case
func ToSnake(s string) string {
	return toDelimited(s, '_')
}

// toDelimited converts a string to delimited.snake.case
// (in this case `delimiter = '.'`)
//
// Shamelessly copied from https://github.com/iancoleman/strcase/blob/master/snake.go.
func toDelimited(s string, delimiter uint8) string {
	return toScreamingDelimited(s, delimiter, "", false)
}

// toScreamingDelimited converts a string to SCREAMING.DELIMITED.SNAKE.CASE
// (in this case `delimiter = '.'; screaming = true`)
// or delimited.snake.case
// (in this case `delimiter = '.'; screaming = false`)
//
// Shamelessly copied from https://github.com/iancoleman/strcase/blob/master/snake.go.
func toScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
	s = strings.TrimSpace(s)
	n := strings.Builder{}
	n.Grow(len(s) + 2) // nominal 2 bytes of extra space for inserted delimiters
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(s) {
			next := s[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			// add underscore if next letter case type is changed
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(s[i-1]), ignore)
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
							n.WriteByte(delimiter)
						}
					}
					n.WriteByte(v)
					if vIsLow || vIsNum || nextIsNum {
						n.WriteByte(delimiter)
					}
					continue
				}
			}
		}

		if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
			// replace space/underscore/hyphen/dot with delimiter
			n.WriteByte(delimiter)
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}
