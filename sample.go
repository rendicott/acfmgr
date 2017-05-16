package main

import (
	"strings"

	"./acfmgr"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// build the inputs
	credFileName := "./creds"
	entryName := "[testdat2]"
	entryContentsString := `output = json
		region = us-east-1
		aws_access_key_id = ASIAJAVEFOINVCBXOLA
		aws_secret_access_key = zgylMqe64havoaoinweofnviUHqQKYHMGzFMA8CI
		aws_session_token = FQoDYXdzEGYaDNYfEnCsHW/8rG3zpiKwAfS8TctytN2YJXv7a80q3PCz/Rak/muP8OZKnUGcYNM51++TR7UvJc8i8PZ9szaJFzTbWeXxkJKMVyWrasdfawefafawefawefawfeCp+6War/2MvoAavoiansdovinawpeoirheowiv7N6pISm/geT+LShtRq5vgq8yC095taHhm5OKY614qKL/ZyMgF

		`
	entryContentsSlice := strings.Split(entryContentsString, "\n")

	// build the credfile session with the inputs
	c, err := acfmgr.NewCredFileSession(credFileName)
	check(err)
	c.NewEntry(entryName, entryContentsSlice)
	c.NewEntry("[dev-account-1]", []string{"output = json", "region = us-east-1", "...", ""})
	c.NewEntry("[dev-account-2]", []string{"output = json", "region = us-west-1", "...", ""})
	err = c.AssertEntries()
	err = c.AssertEntries()
	// err = c.DeleteEntries()
	check(err)
}
