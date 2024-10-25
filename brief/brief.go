package brief

import (
	"fmt"
	"regexp"
	"strings"
)

func isUnchangedAttributeLine(line string) bool {
	rightSides := []string{
		`".*"`,  // string value
		"true",  // boolean value
		"false", // boolean value
		`\d+`,   // integer value
		"{}",    // empty map
		`\[\]`,  // empty list
		`\(sensitive value\)`,
	}

	for _, rightSide := range rightSides {
		regex := fmt.Sprintf(`^\s*\w+\s*=\s*%s\s*$`, rightSide)
		matched, _ := regexp.MatchString(regex, line)
		if matched {
			return true
		}
		regex = fmt.Sprintf(`^\s*"\w+"\s*=\s*%s\s*$`, rightSide)
		matched, _ = regexp.MatchString(regex, line)
		if matched {
			return true
		}
	}

	return false
}

func Plan(lines []string) []string {
	var result []string

	insideMultiline := false
	prevLine := ""
	insideCurlyBraces := false
	insideSquareBraces := false
	sawFirstLine := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			result = append(result, line)
			continue
		}
		if trimmedLine == "Terraform will perform the following actions:" {
			sawFirstLine = true
			result = append(result, line)
			continue
		}
		if !sawFirstLine {
			continue
		}
		if strings.HasSuffix(trimmedLine, "<<-EOT") {
			if insideCurlyBraces {
				insideCurlyBraces = false
				result = append(result, prevLine)
			}
			result = append(result, line)
			insideMultiline = true
			prevLine = line
			continue
		} else if strings.HasPrefix(trimmedLine, "EOT") {
			insideMultiline = false
		}

		if insideMultiline {
			if strings.HasPrefix(trimmedLine, "+ ") || strings.HasPrefix(trimmedLine, "- ") {
				result = append(result, line)
			}
			continue
		}
		if insideCurlyBraces {
			insideCurlyBraces = false
			if trimmedLine == "}" {
				continue
			} else {
				result = append(result, prevLine)
			}
		}
		if insideSquareBraces {
			insideSquareBraces = false
			if trimmedLine == "]" {
				continue
			} else {
				result = append(result, prevLine)
			}
		}

		if matched, _ := regexp.MatchString(`^\s*\w+\s*{\s*$`, line); matched {
			insideCurlyBraces = true
		} else if matched, _ := regexp.MatchString(`^\s*\w+\s*=\s*{\s*$`, line); matched {
			insideCurlyBraces = true
		} else if matched, _ := regexp.MatchString(`^\s*\w+\s*\[\s*$`, line); matched {
			insideSquareBraces = true
		} else if matched, _ := regexp.MatchString(`^\s*\w+\s*=\s*\[\s*$`, line); matched {
			insideSquareBraces = true
		} else if matched, _ := regexp.MatchString(`^\s*".+",\s*$`, line); matched {
			continue
		} else if !isUnchangedAttributeLine(line) {
			result = append(result, line)
		}
		prevLine = line
	}

	return result
}
