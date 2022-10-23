/*
 * baev.one / Personal website
 * Copyright (C) 2022 Baev
 *
 * MIT License
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 * 
 * GNU General Public License, version 2
 * This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later version.
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.
 * You should have received a copy of the GNU General Public License along with this program; if not, write to the Free Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301, USA.
 * 
*/

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
