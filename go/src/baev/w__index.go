
// by Baev, 2022

package main

import (
	"net/http"
	"strings"
	"time"
	"strconv"
)

func start_index(w http.ResponseWriter, r *http.Request) {
	
	ruri := r.URL.Path
	ruriGo := strings.Split(ruri, "/")
	ruriGo_len := len(ruriGo)
	if (ruriGo[1]=="") { ruriGo[1]="/" }

	title := ""
	if ruriGo_len>3 && ruriGo[3]!="" {
		title = pg_title[ruriGo[1]+"/"+ruriGo[2]+"/"+ruriGo[3]]
	} else if ruriGo_len>2 && ruriGo[2]!="" {
		title = pg_title[ruriGo[1]+"/"+ruriGo[2]]
		ruriGo = append(ruriGo, "")
	} else {
		title = pg_title[ruriGo[1]]
		ruriGo = append(ruriGo, "")
		ruriGo = append(ruriGo, "")
	}

	if (title=="") {
		hdl_redirect_home(w, r)
	} else {
		tpl_data := map[string]string{
			"title":		title,
			"myurl":		myurl,
			"author":		author,
			"copyright":	copyright,
			"ruriGo_1":		ruriGo[1],
			"ruriGo_2":		ruriGo[2],
			"ruriGo_3":		ruriGo[3],
			"time":			strconv.FormatInt(time.Now().Sub(time.Unix(0,0)).Nanoseconds(),10)}
		
		display(w, "main", tpl_data)
		//display(w, "403", nil)
	}
}
