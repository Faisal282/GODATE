package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
import "html/template"
import "exercises/app/models"

type Time struct {
	TimeNow string `json:"time_now"`
	TimeFuture string `json:"time_future"`
}

func GetDate(w http.ResponseWriter, r *http.Request) {

	var date = models.Datetimes{
		Current: time.Now().Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
		After : time.Now().Add(12*time.Hour).Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
	}

	var data = map[string]string{
		"Times":    date.Current,
		"TimesAfter" : date.After,
	}

	var t, err = template.ParseFiles("app/views/view.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}
func GetDateAfter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		//fmt.Print(decoder)
		payload := struct {
			Time   string `json:"time"`
		}{}
		if err := decoder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := Time{
			TimeNow: time.Now().Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
			TimeFuture: time.Now().Add(12*time.Hour).Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
		}
		json.NewEncoder(w).Encode(data)
		//message := fmt.Sprintf(
		//	` [{"date" : "%s"}, {"dateAfter"" : "%s"}]`,
		//	payload.Time,
		//	time.Now().Add(12*time.Hour).Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
		//	//payload.Age,
		//	//payload.Gender,
		//)
		//w.Write([]byte(message))
		//return

		//var date_result  = models.Datetimes{
		//	//Current : payload.Time,
		//	After : time.Now().Add(12*time.Hour).Format("2006-01-02 15:04:05 PM -07:00 Jan Mon MST"),
		//}
		//
		//var dataFinal = map[string]string{
		//	//"Times":        date_result.Current,
		//	"TimesAfter":   date_result.After,
		//}
		////fmt.Print(date_result.After)
		////fmt.Print(date_result.Current)
		//var t, err = template.ParseFiles("app/views/view.html")
		//if err != nil {
		//	fmt.Println(err.Error())
		//	return
		//}
		//t.Execute(w, dataFinal)
	}

	//http.Error(w, "Only accept POST request", http.StatusBadRequest)
}
