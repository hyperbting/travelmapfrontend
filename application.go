package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	//"os"
	"net/http"
	//"encoding/json"
	"html/template"

	b64 "encoding/base64"

	"github.com/gorilla/mux"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/info", Info)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	router.HandleFunc("/tmp/{stringToEncode}", ShowQRCode)

	log.Fatal(http.ListenAndServe(":5000", router))
}

func Info(w http.ResponseWriter, r *http.Request) {
	endpoint := os.Getenv("API_ENDPOINT")
	fmt.Fprintln(w, endpoint)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}

func ShowQRCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vars := mux.Vars(r)

	var png []byte
	png, _ = qrcode.Encode("https://Hooloop.com/"+vars["stringToEncode"], qrcode.Medium, 256)

	sEnc := b64.StdEncoding.EncodeToString(png)

	template, err := parseTemplate("public/templates/ShowQRCode.html", map[string]interface{}{
		"Base64png": sEnc,
		"Name":      "Phil Hughes",
		"Data": []*User{
			&User{UserId: 2, Username: "username1"},
			&User{UserId: 3, Username: "username2"},
			&User{UserId: 4, Username: "username3"}}})

	//cookie := http.Cookie{Name: "test", Value: "1", Expires: time.Now().Add(2 * time.Hour)}
	//http.SetCookie(w, &cookie)

	if err == nil {
		fmt.Fprintf(w, string(template))
	} else {
		fmt.Println(err)
	}
}

func parseTemplate(fileName string, data interface{}) (output []byte, err error) {
	var buf bytes.Buffer
	template, err := template.ParseFiles(fileName)

	if err != nil {
		return nil, err
	}

	err = template.Execute(&buf, data)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type User struct {
	UserId   int     `json:"id"`
	Username string  `json:"username"`
	Friends  []*User `json:"friends"`
}
