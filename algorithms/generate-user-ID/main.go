package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	"errors"
)

func main() {	
	
	// test with different values
	fmt.Println(GetNewMemID("aab19"))
	fmt.Println(GetNewMemID("aab18"))
	fmt.Println(GetNewMemID("zzz19"))
	fmt.Println(GetNewMemID("zz19"))
	fmt.Println(GetNewMemID("zz19a"))


}

// GetNewMemID - gets the next member id
func GetNewMemID(recentMemID string) (string, error) {
	var newMemID string
	alpha := "abcdefghijklmnopqrstuvwxyz"
	recentMemID = strings.ToLower(recentMemID)
	id1 := make([]string, 3)	
	
	// if the length of the member is not up to 5
	if len(recentMemID) < 5 {
		return "", errors.New("invalid member ID")
	}

	// if the final two digits in the id are not integers 
	if _, err := strconv.Atoi(recentMemID[3:]); err != nil {
		return "", errors.New("invalid member ID")
	}

	// if we are in a new year, set the recent member Id to aaa...
	currentYear := strconv.Itoa(time.Now().Year())
	if recentMemID[3:] != currentYear[2:] {
		recentMemID = "AAA" + currentYear[2:]

		// check that it doesn't exist in the db
		return recentMemID, nil
	}

	for i:= range alpha {
		id1[0] = string(alpha[i])

		for j := range alpha {
			id1[1] = string(alpha[j])
			
			for k := range alpha {
				id1[2] = string(alpha[k])
				if recentMemID[:3] < strings.Join(id1, "") {
					newMemID = strings.Join(id1, "") + currentYear[2:]

					// check that it doesn't exist in the db
					return strings.ToUpper(newMemID), nil
				}
			}
		}
	}

	return "", errors.New("Out of range")
}