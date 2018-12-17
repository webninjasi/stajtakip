package templates

import (
  "html/template"
)

var tpl = template.Must(template.ParseFiles("templates/main.html"))

type Main struct {
  Title string
}

func Load(filenames ...string) *template.Template {
  return template.Must(template.Must(tpl.Clone()).ParseFiles(filenames...))
}
