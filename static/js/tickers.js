var request;

function CreateRequest() {
	var request=null;
	try {
		request = new XMLHttpRequest();
	}
	catch (e) {
		try {
			request=new ActiveXObject("Msxml2.XMLHTTP");
		}
		catch (e) {
			request=new ActiveXObject("Microsoft.XMLHTTP");
		}
	}
	return request;
}

function sendtickers(url,comm) {
	request=CreateRequest();
	if(request==null) { return; }
	request.onreadystatechange = function(){loadtickers(comm)};
	request.open("GET", url, true);
	request.send(null);
}

function loadtickers(comm) {
	if (request.readyState == 4){
		if (request.status == 200){
			var gethtml = request.responseText;
			if (gethtml == "no") {
			} else {
				var btcusdtvalue_old=0;
				if (document.getElementById("usdtbtc")) {
					btcusdtvalue_old=parseFloat(document.getElementById("usdtbtc").innerHTML);
				}
				var pfltickersblock = gethtml.split("#");
				var pfltickerslist = pfltickersblock[0].split("|");
				var pfltickervalues = pfltickerslist[0].split("!");
				if (pfltickervalues[1]) {
					var floatvalue = parseFloat(pfltickervalues[1]);
					var btcusdtvalue = floatvalue.toFixed(2);
					var triangulum="";
					if (parseFloat(pfltickervalues[4])>parseFloat(pfltickervalues[5])) {
						triangulum="triangulumas_b";
					} else {
						triangulum="triangulumdes_b";
					}
					var btcusdthl="";
					btcusdthl+="h: -"+parseFloat(pfltickervalues[5]).toFixed(2)+"% / "+parseFloat(pfltickervalues[2]).toFixed(2)+"<br />";
					btcusdthl+="l: +"+parseFloat(pfltickervalues[4]).toFixed(2)+"% / "+parseFloat(pfltickervalues[3]).toFixed(2);
				}
				if (comm == "onlybtc") {
					var tbl="<table><tr><td>BTC</td>";
					tbl+="<td id=\"usdtbtc\">"+btcusdtvalue+"</td><td><div class=\""+triangulum+"\"></div><span>24h</span></td></tr>";
					tbl+="<tr><td colspan=\"3\" id=\"usdtbtc_hl\" class=\"hlbtc\">"+btcusdthl+"</td>";
					tbl+="</tr></table>";
					document.getElementById("btctickers").innerHTML=tbl;
					if (btcusdtvalue_old!=0) {
						var clrch="noch";
						if (btcusdtvalue_old<parseFloat(btcusdtvalue)) {
							clrch="#00CC00";
						}
						if (btcusdtvalue_old>parseFloat(btcusdtvalue)) {
							clrch="#EE0000";
						}
						if (clrch!="noch") {
							document.getElementById("usdtbtc").style.backgroundColor=clrch;
							let timerclrId = setTimeout(function checkclrtimer() {
								document.getElementById("usdtbtc").style.backgroundColor="#fff";
								clearTimeout(timerclrId);
							}, 1000);
						}
					}
				}
				if (comm == "prices") {
					document.getElementById("USDT_BTC").innerHTML=btcusdtvalue+"<span> USDT</span>";
					document.getElementById("USDT_BTC_hl").innerHTML=btcusdthl;
					document.getElementById("USDT_BTC_tgl").innerHTML="<div class=\""+triangulum+"\"></div><span>24h</span>";

/*document.getElementById("test_tickers").innerHTML=pfltickersblock[1]+pfltickersblock[2]+"<br/>";*/

					for (ilst = 1; ilst < 3; ++ilst) { 
						var pfltickerslist = pfltickersblock[ilst].split("|");
						for (index = 0; index < pfltickerslist.length; ++index) {
							var pfltickervalues = pfltickerslist[index].split("!");
							if (pfltickervalues[1]) {
								var floatvalue = parseFloat(pfltickervalues[1]);
								var hp_value = parseFloat(pfltickervalues[2]);
								var lp_value = parseFloat(pfltickervalues[3]);
								var strvalue = floatvalue.toFixed(8);
								var str_hp_value = hp_value.toFixed(8);
								var str_lp_value = lp_value.toFixed(8);
								if (pfltickervalues[0].substr(0,3)=="BTC") {
									if (floatvalue<0.00000010) {
										var strvalue = floatvalue.toFixed(10);
										var str_hp_value = hp_value.toFixed(10);
										var str_lp_value = lp_value.toFixed(10);
									}
									strvalue += "<span> BTC</span>";
								} else {
									if (floatvalue>=0.1) {
										var strvalue = floatvalue.toFixed(2);
										var str_hp_value = hp_value.toFixed(2);
										var str_lp_value = lp_value.toFixed(2);
									}
									strvalue += "<span> USDT</span>";
								}
								document.getElementById(pfltickervalues[0]).innerHTML=strvalue;
								var triangulum="";
								if (parseFloat(pfltickervalues[4])>parseFloat(pfltickervalues[5])) {
									triangulum="triangulumas_b";
								} else {
									triangulum="triangulumdes_b";
								}
								var resulthtml="";
								resulthtml+="h: -"+parseFloat(pfltickervalues[5]).toFixed(2)+"% / "+str_hp_value+"<br />";
								resulthtml+="l: +"+parseFloat(pfltickervalues[4]).toFixed(2)+"% / "+str_lp_value;
								document.getElementById(pfltickervalues[0]+"_hl").innerHTML=resulthtml;
								var resulthtml_tr="<div class=\""+triangulum+"\"></div><span>24h</span>";
								document.getElementById(pfltickervalues[0]+"_tgl").innerHTML=resulthtml_tr;
							}
						}
					}

				}
			}
		}
	}
}
