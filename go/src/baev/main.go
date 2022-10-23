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
	"log"
)

func main() {
	
	mux80:=http.NewServeMux()
	mux80.HandleFunc("/",hdl_80)
	
	mux443:=http.NewServeMux()
	mux443.HandleFunc("/",hdl_443)
	fs:=http.FileServer(http.Dir(static_path))
	mux443.Handle("/css/",http.StripPrefix("/",fs))
	mux443.Handle("/js/",http.StripPrefix("/",fs))
	mux443.Handle("/img/",http.StripPrefix("/",fs))
	
	mux5001:=http.NewServeMux()
	mux5001.HandleFunc("/",hdl_5001)
	
	mux7001:=http.NewServeMux()
	mux7001.HandleFunc("/",hdl_7001)
	mux7001.HandleFunc("/upload",hdl_7001_upload)
	mux7001.HandleFunc("/download",hdl_7001_download)
	
	go func() { log.Fatal(http.ListenAndServe(":80", mux80)) }()
	go func() { log.Fatal(http.ListenAndServeTLS(":443",ssl_cert,ssl_key,mux443)) }()
	//go func() { log.Fatal(http.ListenAndServeTLS(":5001",ssl_cert,ssl_key,mux5001)) }()
	//go func() { log.Fatal(http.ListenAndServeTLS(":7001",ssl_cert,ssl_key,mux7001)) }()
	select {}
	
}
