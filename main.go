package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
)

// struct that defines the request data format
type PostReqFormat struct {
	Id *int
	FirstName *string
	LastName *string
	EmailAddress *string
	ValidUser *bool
  }

// helper func that checks for error
func errorChecking(err error) {
	if err != nil {
		panic(err)
	}
}

// variadic func that checks for nil pointers, returns true when nil pointer is detected
func nilFieldCheck(fields ...any) bool  {
	for _, value := range fields {
		if reflect.ValueOf(value).IsNil() {
			return true
		}
	}
	return false
}

// helper function that writes the request data in a file named data.json
func writeFile(content []byte) {
	file, err := os.OpenFile("./data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	errorChecking(err)
	defer file.Close()

	_, err = file.Write(content)
	errorChecking(err)
	_, err = file.WriteString("\n")
	errorChecking(err)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// respond with 400 if a different http method is used
	if r.Method == "POST" {
		var req PostReqFormat
		// decode req body into PostReqFormat format
		err := json.NewDecoder(r.Body).Decode(&req)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
        	return
    	}
		// respond with 400 if any required field is undefined in the request
		if nilFieldCheck(req.Id, req.FirstName, req.LastName, req.EmailAddress, req.ValidUser) == true {
			http.Error(w, "REQUIRED FIELDS UNDEFINED", http.StatusBadRequest)
			return
		}
		// convert request body into []byte for writing file and response
		reqJson, err := json.Marshal(req)
		errorChecking(err)

		// write request body content into a file
		writeFile(reqJson)

		// respond with 201 and content if all previous operations are successful 
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(reqJson)
	} else {
		http.Error(w, "SERVER HANDLES POST REQUESTS ONLY", http.StatusBadRequest)
        return
	}
}

func main() {
	// define route handler for POST method
	http.HandleFunc("/postEndpoint", postHandler)
	// starting the server
	fmt.Println("Server starts listening on port 3000!")
	err := http.ListenAndServe(":3000", nil)
	errorChecking(err)
}