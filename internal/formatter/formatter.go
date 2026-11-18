package formatter

import (
	"regexp"
	"strings"
)

const (
	NameWidth = 20
	TypeWidth = 20
	Indent    = "    "
)

var createTableRe = regexp.MustCompile(
	`(?is)(CREATE\s+TABLE\s+IF\s+NOT\s+EXISTS\s+([A-Za-z_][A-Za-z0-9_]*)\s*\()\s*\n(.*?)\n\s*(\))\s*;?`,
)

var columnRe = regexp.MustCompile(
	`^([A-Za-z_][A-Za-z0-9_]*)\s+([A-Za-z_][A-Za-z0-9_]*(?:\([^)]*\))?)\s*(.*)$`,
)

// Format formats a SQL document. It currently applies the custom
// fixed-width layout to CREATE TABLE IF NOT EXISTS statements.
func Format(sql string) string {
	formatted := createTableRe.ReplaceAllStringFunc(sql, formatCreateTable)

	// Remove a final semicolon if it is the last non-whitespace character.
	formatted = strings.TrimRight(formatted, " \t\r\n")
	formatted = strings.TrimSuffix(formatted, ";")

	return formatted + "\n"
}

func formatCreateTable(statement string) string {
	match := createTableRe.FindStringSubmatch(statement)
	if match == nil {
		return statement
	}

	header := match[1]
	body := match[3]
	footer := match[4]

	rows := parseColumns(body)
	if len(rows) == 0 {
		return statement
	}

	var out strings.Builder
	out.WriteString(header)
	out.WriteString("\n")

	for _, row := range rows {
		out.WriteString(Indent)
		out.WriteString(pad(row.Name, NameWidth))
		out.WriteString(" ")
		out.WriteString(pad(row.Type, TypeWidth))
		out.WriteString(" ")
		out.WriteString(row.Constraints)
		out.WriteString("\n")
	}

	out.WriteString(footer)

	return out.String()
}

type Column struct {
	Name        string
	Type        string
	Constraints string
}

func parseColumns(body string) []Column {
	var columns []Column

	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		match := columnRe.FindStringSubmatch(line)
		if match == nil {
			columns = append(columns, Column{Name: line})
			continue
		}

		columns = append(columns, Column{
			Name:        match[1],
			Type:        match[2],
			Constraints: strings.TrimSpace(match[3]),
		})
	}

	return columns
}

func pad(value string, width int) string {
	if len(value) >= width {
		return value
	}

	return value + strings.Repeat(" ", width-len(value))
}