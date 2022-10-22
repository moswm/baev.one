
// by Baev, 2022

package main

import (
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles(
	//system
	tpl_path+"main.html",
	tpl_path+"403.html",
	tpl_path+"upload.html",
	tpl_path+"crypto.html",
	//pg
	tpl_path+"pg/page_main.html",
	tpl_path+"pg/page_dev.html",
	tpl_path+"pg/page_pfl.html",
	tpl_path+"pg/page_privacy_policy.html",
	tpl_path+"pg/page_vidpro.html",
	tpl_path+"pg/page_codeexample.html",
	//pg sub folders
	tpl_path+"pg/dev/calculator/form.html",
	tpl_path+"pg/sendforms/callform.html"))

func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}
