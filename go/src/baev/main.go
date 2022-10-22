
// by Baev, 2022

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
