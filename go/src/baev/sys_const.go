
// by Baev, 2022

package main

const (
	//ssl
	ssl_cert		= "/var/www/ssl/cert.pem"
	ssl_key			= "/var/www/ssl/key.pem"
	//system
	tpl_path		= "/var/www/go/src/baev/tpl/"
	static_path		= "/var/www/static/"
	//main
	copyright		= "1999 - 2022 Baev"
	author			= "Baev"
	prodname		= "AyDaka Content Management System v10.0"
	myhost			= "baev.one"
	myhostidn		= myhost
	myurl			= "https://"+myhost+"/"
	myurlidn		= "https://"+myhostidn+"/"
	myemail			= "...@..."
	smalltitle		= "Golang / PHP & JavaScript Full stack developer"
)

//pages
var pg_title = map[string]string{
	"/":				"Golang / PHP & JavaScript Full stack developer",
	"privacy_policy":	"Privacy Policy",
	"pfl":				"Some completed projects",
	"dev":				"Web service examples and useful tools",
	"dev/api-crypto":	"Prices for some crypto assets by API",
	"dev/calculator":	"Interactive application-calculator",
	"dev/geo-cities":	"Selection of cities / regions by API"}
