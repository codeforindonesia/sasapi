package controller

import (
	"log"
	"net/http"
	"app/model"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"app/shared/session"
)

// MonitorReadGET displays the monitors in the monitor
func Index(w http.ResponseWriter, r *http.Request) {
	// Get session
	result := map[string]string{"name": "SAS API", "version": "1.0.2"}
	b, err := json.MarshalIndent(result, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

	log.Println(err)

}

// MonitorReadGET displays the monitors in the monitor
func RealisasiJSON(w http.ResponseWriter, r *http.Request) {
	// Get session
	var err error
	var params httprouter.Params
	params = context.Get(r, "params").(httprouter.Params)

	//startDate := r.FormValue("startDate")
	//endDate := r.FormValue("endDate")

	startDate := params.ByName("startDate")
	endDate := params.ByName("endDate")
	log.Printf(startDate)
	log.Println(endDate)

	realisasi,err := model.RealisasiByDate(startDate, endDate)

	if err != nil {
		log.Println(err)
	}
	b, err := json.MarshalIndent(realisasi, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

	log.Println(err)

}

func RealisasiPOSTJSON(w http.ResponseWriter, r *http.Request) {
	// Get session
	var err error
	//var params httprouter.Params
	//params = context.Get(r, "params").(httprouter.Params)
	action := r.FormValue("action")
	startDate := r.FormValue("startDate")
	endDate := r.FormValue("endDate")
	token := r.FormValue("token")

	tokenSetting := session.Auth

	if token == tokenSetting {

		var realisasi []model.Realisasi
		var pagu []model.Pagu
		var b []byte
		// method

		switch action {
		case "realisasi":
			realisasi,err = model.RealisasiByDate(startDate, endDate)
			b, err = json.MarshalIndent(realisasi, "", "    ")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(b)


		case "pagu":
			pagu,err = model.PaguByDate()
			b, err = json.MarshalIndent(pagu, "", "    ")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(b)


		}
	} else {
		InvalidToken(w,r)
	}
	if err != nil {
		log.Println(err)
	}

}