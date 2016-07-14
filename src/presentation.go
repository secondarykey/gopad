package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func init() {
}

func render(m *Memo) string {

	var rtn = bytes.NewBuffer(make([]byte, 0, 100))

	rtn.WriteString("class: center,middle\n")
	rtn.WriteString("# " + m.Title + "\n\n")

	marks := rendering(m.Content, 1)
	rtn.WriteString(build(marks, m.Title))

	return rtn.String()
}

func build(ms []*mark, t string) string {

	var rtn = bytes.NewBuffer(make([]byte, 0, 100))

	for _, m := range ms {

		if m == nil {
			continue
		}

		rtn.WriteString("---\n\n")

		rtn.WriteString("class: top,left\n")
		rtn.WriteString("## " + m.title + "\n\n")
		rtn.WriteString(m.content + "\n\n")
		rtn.WriteString(".footnote[" + t + "]\n\n")
		if m.children != nil {
			rtn.WriteString(build(m.children, t+"/"+m.title))
		}
	}
	return rtn.String()
}

func rendering(s string, prefix int) []*mark {

	m := make([]*mark, 0)
	r := strings.NewReader(s)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Printf("%v\n", serr)
		panic(serr)
	}

	header := strings.Repeat("#", prefix) + " "

	data := ""
	flg := false
	var mk *mark

	for _, line := range lines {

		if strings.Trim(line, " ") == "" {
			continue
		}

		idx := strings.Index(line, header)
		if idx == 0 {

			if data != "" {
				mk.children = rendering(data, prefix+1)
				data = ""
			}

			if mk != nil {
				m = append(m, mk)
			}

			mk = &mark{
				title:    "",
				content:  "",
				children: nil,
			}

			mk.title = line[len(header):]
			flg = false

		} else {

			if idx > -1 {
				flg = true
			}

			if flg {
				data = data + line + "\n"
			} else {
				mk.content = mk.content + line + "\n"
			}
		}
	}

	if mk != nil {

		if data != "" {
			mk.children = rendering(data, prefix+1)
		}
		m = append(m, mk)
	}

	return m
}

type mark struct {
	title    string
	content  string
	children []*mark
}
