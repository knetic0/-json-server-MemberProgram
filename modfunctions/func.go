package modfunctions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Member struct {
	MemberId      int    `json:"id"`
	MemberName    string `json:"name"`
	MemberSurname string `json:"surname"`
}

func CheckUp() { // file Control - db.json -
	if _, err := os.Stat("db.json"); err == nil { // file exists
		fmt.Println("No Problems. Continue...")
	} else {
		if _, err := os.Stat("../db.json"); err == nil { // file exists
			fmt.Println("No Problems. Continue...")
		} else {
			os.Create("db.json") // file create
		}
	}

	fmt.Println("All is done.")

	fmt.Println(SetUser())
}

func GetUser() (string, error) { // get Member Information
	response, err := http.Get("http://localhost:3000/member")

	if err != nil {
		return "Some error is here -> ", err
	}
	defer response.Body.Close()

	read_bytes, _ := ioutil.ReadAll(response.Body)

	var members []Member
	json.Unmarshal(read_bytes, &members)
	fmt.Println(members)

	return "Finish!", nil
}

func SetUser() (string, error) { // add New Member
	var id, timer int
	var name, surname string

	fmt.Print("How many Member do you want to add json file : ")
	fmt.Scan(&timer)

	for i := 0; i < timer; i++ {
		fmt.Println("Please enter id, Name, Surname -> ")
		fmt.Scan(&id, &name, &surname)

		fmt.Println(id, name, surname)

		nextMember := Member{
			MemberId:      id,
			MemberName:    name,
			MemberSurname: surname,
		}

		jsonMember, _ := json.Marshal(nextMember)

		response, err := http.Post("http://localhost:3000/member", "application/json;charset=utf-8",
			bytes.NewBuffer(jsonMember))

		if err != nil {
			return "Some error is here -> ", err
		}
		defer response.Body.Close()

		processing, _ := GetUser()
		fmt.Println(processing)
	}

	if timer < 0 {
		return "Please enter number of bigger than zero ->", nil
	} else if timer == 1 {
		return "Saved!", nil
	} else {
		return "All is Saved", nil
	}

}
