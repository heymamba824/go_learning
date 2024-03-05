package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"mygithub/pkg/config"
	"mygithub/pkg/models"
	"net/http"
	"path/filepath"
)

/*
var tc = make(map[string]*template.Template)

	func RenderTmpt(w http.ResponseWriter, t string) {
		var tmpl *template.Template
		var err error
		//check to see if we already have the template in our cache
		_, inMap := tc[t]
		if !inMap {
			fmt.Println("creating template and adding to cache")
			//need to create the template
			err = createTemplateCache(t)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			//we have the template in the cache
			fmt.Println("using cached template")
		}
		tmpl = tc[t]
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(err)
		}

}

	func createTemplateCache(t string) error {
		templates := []string{
			fmt.Sprintf("./templates/%s", t),
			"./templates/base.layout.tmpl",
		}
		//parsethe template
		tmpl, err := template.ParseFiles(templates...)
		if err != nil {
			return err
		}
		//add template to cache
		tc[t] = tmpl
		return nil
	}

// renderTemplate renders templates using html

	func RenderTmptTest(w http.ResponseWriter, tmpl string) {
		parsedTmpt, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		err := parsedTmpt.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template", err)
		}
	}
*/
var app *config.AppConfig

// set the conig for the template
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTmpt(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// creat a template cache
	/*	tc, err := CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tc)
		// get requested template from cache
		t, ok := tc[tmpl]
		if !ok {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		err = t.Execute(buf, nil)
		if err != nil {
			log.Println(err)
		}*/
	var tc map[string]*template.Template
	if app.UserCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	// get the template cache from the app config
	//tc := app.TemplateCache
	//render the template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing teplate to browser", err)
	}
	/*
		parsedTmpt, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
		err := parsedTmpt.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template", err)
		}*/
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	//获取文件所有相关的pages
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	fmt.Println(pages)
	if err != nil {
		return myCache, err
	}
	//for循环遍历找到的每个页面模板文件
	for _, page := range pages {
		name := filepath.Base(page)
		//创建一个命名为name的模版
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			fmt.Println("#######需要这个吗")
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
