package views

import (
	"html/template"
)

func Templates(funcMap *template.FuncMap) map[string]*template.Template {
	var T = make(map[string]*template.Template)

	//T["sample/sample_template"] = template.Must(template.ParseFiles("views/sample/base.tmpl", "views/sample/sample_template.tmpl"))

	return T
}
