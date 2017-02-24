package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	//"time"
	//"strconv"
	"strconv"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "returns a specific article")
	fmt.Println("Endpoint Hit: returnArticle")
}


type GoPoint struct {
	Amount float32 `json:"amount"`//question
	Voucher Voucher `json:"voucher"`
}

type Voucher struct {
	Code int `json:"code"`
	Transaction_amount int `json:"transaction_amount"`
	Transaction_time int64 `json:"transaction_time"`
	Transaction_currency string `json:"transaction_currency"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Document_type string `json:"document_type"`
	Merchant_id string `json:"merchant_id"`
	Terminal_id string `json:"terminal_id"`
	Stan int `json:"stan"`
}

type GoPointNew struct{
	Name string `json:"name"`
	Age int `json:"age"`
}

func returnAllArticles(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "{\"success\": true,\"code\": \"00\",\"message\": \"O Hai, you redeemed your voucher, kthnxbai!\",\"redemption\" : {\"id\": \"GH657576576678\",\"redeemed_time\": \"20-12-2016 08:00:00\",\"voucher\": {\"code\": 12345678,\"transaction_amount\": 10000,\"transaction_time\": \"20-12-2016 08:00:00\",\"transaction_currency\": \"IDR\",\"metadata\": {\"document_type\": \"financial transaction\",\"merchant_id\": \"abcde1234567890\",\"terminal_id\": \"abcd1234\",\"stan\": 123456}}}}{\"success\": true,\"code\": \"00\",\"message\": \"O Hai, you redeemed your voucher, kthnxbai!\",\"redemption\" : {\"id\": \"GH657576576678\",\"redeemed_time\": \"20-12-2016 08:00:00\",\"voucher\": {\"code\": 12345678,\"transaction_amount\": 10000,\"transaction_time\": \"20-12-2016 08:00:00\",\"transaction_currency\": \"IDR\",\"metadata\": {\"document_type\": \"financial transaction\",\"merchant_id\": \"abcde1234567890\",\"terminal_id\": \"abcd1234\",\"stan\": 123456}}}}")

	//fmt.Println("Endpoint Hit: returnAllArticles")
	decoder := json.NewDecoder(req.Body)
	var t GoPoint
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	fmt.Println("Code: " + strconv.Itoa(t.Voucher.Code) +
		"\nDoc_Type : " + t.Voucher.Metadata.Document_type)
		//"\nTime : " + time.Unix(t.Voucher.Transaction_time,0).String())
	fmt.Printf("Unix Time: ")
	fmt.Println(time.Unix(0, t.Voucher.Transaction_time*int64(time.Millisecond)))
	//fmt.Println(r.Body.Read())

}

func addArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Adds an article to list of articles")
	fmt.Println("Endpoint Hit: addArticle")
}

func delArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "deletes a specific article")
	fmt.Println("Endpoint Hit: delArticle")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/gopoint/redeem", returnAllArticles)
	http.HandleFunc("/single", returnArticle)
	http.HandleFunc("/delete", delArticle)
	http.HandleFunc("/add", addArticle)
	log.Fatal(http.ListenAndServe(":1818", nil))
}

func main() {
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Print("Sec : ")
	fmt.Println(secs)
	fmt.Print("Mil : ")
	fmt.Println(millis)
	fmt.Print("Nsn : ")
	fmt.Println(nanos)

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
	handleRequests()
}