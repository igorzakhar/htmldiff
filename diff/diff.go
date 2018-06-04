package diff

import (
	"bytes"
	"fmt"
	"github.com/aryann/difflib"
	"html"
	"html/template"
	"io/ioutil"
	"strings"
)

type DiffLine struct {
	Opcode    string
	LeftNum   int
	LeftLine  string
	RightNum  int
	RightLine string
}

var templateString = `
  <table class="table table-bordered table-sm">
    {{.Diff}}
  </table>
`

func DiffHTML(text1, text2 string) (string, error) {
	unescapedText1 := html.UnescapeString(text1)
	unescapedText2 := html.UnescapeString(text2)
	sequence1 := strings.Split(unescapedText1, "\n")
	sequence2 := strings.Split(unescapedText2, "\n")

	diffs := GetUnifiedDiffLines(sequence1, sequence2)
	diffsTable := DiffHTMLTable(diffs)
	tmpl, _ := template.New("diffTemplate").Parse(templateString)

	var output bytes.Buffer
	err := tmpl.Execute(&output, map[string]interface{}{
		"Diff": template.HTML(diffsTable),
	})
	if err != nil {
		return "", err
	}
	result := output.String()
	return result, err
}

func fileToLines(filename string) []string {
	content, _ := ioutil.ReadFile(filename)
	return strings.Split(string(content), "\n")
}

func GetUnifiedDiffLines(seq1, seq2 []string) []DiffLine {
	diffLines := []DiffLine{}
	i, j := 0, 0
	for _, d := range difflib.Diff(seq1, seq2) {
		opcode := "equal"
		leftNum, rightNum := 0, 0
		leftLine := d.Payload
		rightLine := d.Payload
		if d.Delta == difflib.Common || d.Delta == difflib.LeftOnly {
			i++
			leftNum = i
			if d.Delta == difflib.LeftOnly {
				opcode = "deleted"
				leftLine = d.Payload
				rightLine = ""
			}
		}
		if d.Delta == difflib.Common || d.Delta == difflib.RightOnly {
			j++
			rightNum = j
			if d.Delta == difflib.RightOnly {
				opcode = "added"
				leftLine = ""
				rightLine = d.Payload
			}
		}
		diffLines = append(diffLines, DiffLine{opcode, leftNum, leftLine, rightNum, rightLine})
	}
	return diffLines
}

func DiffHTMLTable(diffs []DiffLine) string {
	buf := bytes.NewBufferString("")
	clearColumn := "</td><td class=\"clear\">"
	for _, diff := range diffs {
		buf.WriteString(`<tr><td class="line-num">`)
		if diff.LeftNum != 0 {
			escapeLine := html.EscapeString(diff.LeftLine)
			fmt.Fprintf(buf, "%d</td><td class=\"%s\">%s", diff.LeftNum, diff.Opcode, escapeLine)
		} else {
			fmt.Fprint(buf, clearColumn)
		}
		buf.WriteString(`</td><td class="line-num">`)
		if diff.RightNum != 0 {
			escapeLine := html.EscapeString(diff.RightLine)
			fmt.Fprintf(buf, "%d</td><td class=\"%s\">%s", diff.RightNum, diff.Opcode, escapeLine)
		} else {
			fmt.Fprint(buf, clearColumn)
		}

		buf.WriteString("</td></tr>\n")
	}

	return buf.String()
}
