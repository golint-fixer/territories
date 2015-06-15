package application

import (
	"fmt"
	"html/template"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/zenazn/goji"

	"github.com/Quorumsco/contact/components/database"
	"github.com/Quorumsco/contact/views"
)

type Application struct {
	Templates map[string]*template.Template
	DB        *sqlx.DB
	Urls      map[string]string
}

func (application *Application) Init() error {
	application.Templates = make(map[string]*template.Template)
	application.Urls = make(map[string]string)

	funcMap := template.FuncMap{
		"path": func(name string, params ...interface{}) string {
			return fmt.Sprintf(application.Urls[name], params...)
		},
	}

	var err error
	if application.DB, err = database.Init(); err != nil {
		return err
	}
	application.Templates = views.Templates(&funcMap)

	return nil
}

func (application *Application) Name(url_format string, name string) {
	application.Urls[name] = url_format
}

func (application *Application) Get(pattern interface{}, controller interface{}, name string) {
	switch v := pattern.(type) {
	case string:
		application.Name(v, name)
	case regexp.Regexp:
		r := regexp.MustCompile("(\\(.+\\))")
		format := r.ReplaceAllLiteralString(v.String(), "%v")
		application.Name(format, name)
		fmt.Println(format)
	}

	goji.Get(pattern, controller)
}

func (application *Application) Path(name string, params ...interface{}) string {
	return fmt.Sprintf(application.Urls[name], params...)
}
