package main

import (
	"fmt"
	"flag"

)

func main() {
	var username string
	var password string
	flag.StringVar(&username,"username","Gopher","This is a username")
	flag.StringVar(&password,"password","goofyGoober123!","This is a password")
	flag.Parse()

	var res bool = check(username)
	if res {
		fmt.Println("Good username ",username)
	} else {
		fmt.Println("Bad username ", username)
	}

	res = checkPassword(password)
	if res {
		fmt.Println("Good pw ",password)
	} else {
		fmt.Println("Bad pw ", password)
	}
}

func check(uname string) bool {
	if len(uname) > 15 || len(uname) < 5 {
		return false
	}

	for i := 0; i < len(uname); i++ {
		if (uname[i] == '@' && i > 0) {
			return false
		} 
	}

	return true

}

func checkPassword(pw string) bool {
	var length = len(pw)
	if length < 6 || length > 20 {
		return false
	}
	//password needs one uppercase char, one special character, and one number
	var upper, spec, num, space bool
	for i := 0 ; i < length ; i++ {
		if pw[i] >= 65 && pw[i] <= 90 {
			upper = true
		} else if pw[i] >= 48 && pw[i] <= 57 {
			num = true
		} else if pw[i] >= 33 && pw[i] <= 47 {
			spec = true
		} else if pw[i] == ' ' {
			space = true
		}
	}

	if upper && spec && num && !space {
		return true
	} else {
		return false
	}

 }