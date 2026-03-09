package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// PrettyPrint prints a struct in pretty JSON format (for debugging)
func PrettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error printing:", err)
		return
	}
	fmt.Println(string(b))
}

// Contains checks if a string slice contains a specific string
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
