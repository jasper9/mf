package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Username   string `json:"username"`
	DeviceID   string `json:"deviceid"`
	DeviceName string `json:"devicename"`
	AccessKey  string `json:"accesskey"`
}

var users Users

// Main function
func main() {

	// https://tutorialedge.net/golang/parsing-json-with-golang/
	jsonFile, err := os.Open("config.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	//var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	/*for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Key: " + users.Users[i].Key)
		fmt.Println("User Name: " + users.Users[i].Name)
	} */

	// Init the mux router
	router := mux.NewRouter()

	// Route handles & endpoints
	//router.HandleFunc("/hello/", GetHello).Methods("GET")
	router.HandleFunc("/checkin/", PostCheckIn).Methods("POST")

	// serve the app
	fmt.Println("Serving on port 80")
	log.Fatal(http.ListenAndServe(":80", router))

}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//func GetHello(w http.ResponseWriter, r *http.Request) {

//	printMessage("Getting movies...")

//}

func PostCheckIn(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n\n")
	u_deviceid := r.FormValue("deviceid")
	u_accesskey := r.FormValue("accesskey")

	// this should be re-usable
	for i := range users.Users {
		if users.Users[i].AccessKey == u_accesskey {
			if users.Users[i].DeviceID == u_deviceid {
				fmt.Println("Keys and DeviceID match!")

				var sec int64
				now := time.Now() // current local time
				sec = now.Unix()
				fmt.Printf("Seconds: %d\n", sec)

				//err := os.WriteFile("./checks/"+checkid, sec, 0644)
				//checkErr(err)
				f, err := os.Create("./checks/" + u_deviceid + ".txt")
				if err != nil {
					log.Fatal(err)
				}

				defer f.Close()

				bs := []byte(strconv.Itoa(int(sec)))
				//fmt.Println(bs)

				_, err2 := f.Write(bs)

				if err2 != nil {
					log.Fatal(err2)
				}

			}

			//fmt.Println("Keys match")
			//fmt.Printf("config key = %v\n", users.Users[i].Key)
			//fmt.Printf("u_deviceid = %v\n", u_deviceid)
			//fmt.Printf("u_accesskey = %v\n", u_accesskey)

		} //else {
		//	fmt.Println("Keys dont match")
		//fmt.Printf("config key = %v\n", users.Users[i].Key)
		//fmt.Printf("u_deviceid = %v\n", u_deviceid)
		//fmt.Printf("u_accesskey = %v\n", u_accesskey)
		//}
	}

}
