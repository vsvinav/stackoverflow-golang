// package main

// import (
// 	"net/http"
// 	"testing"
// )

// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }
// func TestEmptyTable(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/posts", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)
// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }
