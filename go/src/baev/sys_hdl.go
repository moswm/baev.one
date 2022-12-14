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

import (
	"net/http"
)

func hdl_CORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func hdl_redirect_home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+myhost+":443", 301)
}

func hdl_80(w http.ResponseWriter, r *http.Request) {
	hdl_redirect_home(w, r)
}

func hdl_443(w http.ResponseWriter, r *http.Request) {
	start_index(w, r)
}

func hdl_5001(w http.ResponseWriter, r *http.Request) {
	start_crypto(w, r)
}

func hdl_7001(w http.ResponseWriter, r *http.Request) {
	DispHtml(w, "403", nil)
}
func hdl_7001_upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		DispHtml(w, "403", nil)
		//DispHtml(w, "upload", nil)
	case "POST":
		hdl_CORS(w)
		uplFile(w, r)
	default:
		DispHtml(w, "403", nil)
	}
}
func hdl_7001_download(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		hdl_CORS(w)
		dwlFile(w, r)
	case "POST":
		DispHtml(w, "403", nil)
	default:
		DispHtml(w, "403", nil)
	}
}
