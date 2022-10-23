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
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"crypto/md5"
	"encoding/hex"
	"path/filepath"
	"encoding/json"
	"strconv"
)

func uplFile(w http.ResponseWriter, r *http.Request) {
	
	if rqkeyCheck(r.FormValue("rqkey")) != true {
		uplFile_err(w, "0")
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		uplFile_err(w, "0")
		return
	}
	
	file, handler, err := r.FormFile("uplFile")
	if err != nil {
		uplFile_err(w, "2")
		return
	}
	defer file.Close()
	
	hash := md5.New()
	tm := time.Now().String()
	flsz := strconv.FormatInt(handler.Size, 10)
	
	flnm_pr := tm + r.FormValue("login") + handler.Filename + flsz
	hash.Write([]byte(flnm_pr))
	flhash := hex.EncodeToString(hash.Sum(nil))
	flext := filepath.Ext(handler.Filename)
	
	dst, err := os.Create("/var/www/fldb/"+flhash+flext)
	defer dst.Close()
	if err != nil {
		uplFile_err(w, "2")
		return
	}
	
	if _, err := io.Copy(dst, file); err != nil {
		uplFile_err(w, "2")
		return
	}
	
	resp_src := map[string]string{
		"status": "1",
		"md5": flhash,
		"name": handler.Filename,
		"ext": flext,
		"size": flsz}
	resp_json, err := json.Marshal(resp_src)
	fmt.Fprintf(w, string(resp_json))
	
}

func uplFile_err(w http.ResponseWriter, status string) {
	resp_src := map[string]string{}
	if status == "0" {
		resp_src = map[string]string{
			"status": status,
			"err": "The 403 Forbidden Error!"}
	}
	if status == "2" {
		resp_src = map[string]string{
			"status": status,
			"err": "Internal Server Error!"}
	}
	resp_json, _ := json.Marshal(resp_src)
	fmt.Fprintf(w, string(resp_json))
}
