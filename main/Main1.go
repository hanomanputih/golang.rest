package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "returns a specific article")
	fmt.Println("Endpoint Hit: returnArticle")
}


type test_struct struct {
	Test string
}

func returnAllArticles(w http.ResponseWriter, req *http.Request){

	fmt.Fprintf(w, "{\"success\": true,\"code\": \"00\",\"message\": \"O Hai, you redeemed your voucher, kthnxbai!\",\"redemption\" : {\"id\": \"GH657576576678\",\"redeemed_time\": \"20-12-2016 08:00:00\",\"voucher\": {\"code\": 12345678,\"transaction_amount\": 10000,\"transaction_time\": \"20-12-2016 08:00:00\",\"transaction_currency\": \"IDR\",\"metadata\": {\"document_type\": \"financial transaction\",\"merchant_id\": \"abcde1234567890\",\"terminal_id\": \"abcd1234\",\"stan\": 123456}}}}{\"success\": true,\"code\": \"00\",\"message\": \"O Hai, you redeemed your voucher, kthnxbai!\",\"redemption\" : {\"id\": \"GH657576576678\",\"redeemed_time\": \"20-12-2016 08:00:00\",\"voucher\": {\"code\": 12345678,\"transaction_amount\": 10000,\"transaction_time\": \"20-12-2016 08:00:00\",\"transaction_currency\": \"IDR\",\"metadata\": {\"document_type\": \"financial transaction\",\"merchant_id\": \"abcde1234567890\",\"terminal_id\": \"abcd1234\",\"stan\": 123456}}}}")
	fmt.Println("Endpoint Hit: returnAllArticles")
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	log.Println(t.Test)
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
	handleRequests()
}