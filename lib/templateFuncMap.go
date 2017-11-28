package funcMap

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

// template内変数の独自変換
func FunctionForTemplate() template.FuncMap {
	funcMap := template.FuncMap{
		"htmlRelease": func(s string) template.HTML {
			// package bluemonday でscriptなどを排除
			p := bluemonday.UGCPolicy()
			r := p.Sanitize(string(s))
			r = strings.Replace(r, "\n", "<br>", -1)
			return template.HTML(r)
		},
		"dateRelease": func(t time.Time) string {
			year, month, day := t.Date()
			h := t.Hour()
			m := t.Minute()
			var zerominutes string
			if m < 10 {
				zerominutes = "0" + strconv.Itoa(m)
			} else {
				zerominutes = strconv.Itoa(m)
			}
			wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}
			w := wdays[t.Weekday()]
			return fmt.Sprintf("%d/%d/%d(%s) %d:%s", year, month, day, w, h, zerominutes)
		},
		"runeRelease": func(s string) string {
			var numRunes = 0
			for i := range s {
				numRunes++
				if numRunes > 60 {
					return fmt.Sprintf("%s…", s[:i])
				}
			}
			return s
		},
		"runeReleaseTitle": func(s string) string {
			var numRunes = 0
			for i := range s {
				numRunes++
				if numRunes > 8 {
					return fmt.Sprintf("%s…", s[:i])
				}
			}
			return s
		},
		"runeRelease120": func(s string) string {
			var numRunes = 0
			for i := range s {
				numRunes++
				if numRunes > 120 {
					return fmt.Sprintf("%s…", s[:i])
				}
			}
			return s
		},
	}
	return funcMap
}
