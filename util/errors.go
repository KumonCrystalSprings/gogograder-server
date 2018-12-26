package util

import "log"

// CheckError is a wrapper for non-fatal error checking
func CheckError(err error, msg string) bool {
	if err != nil {
		if msg != "" {
			log.Println(msg)
		}
		log.Println(err)
		return true
	}
	return false
}

// CheckErrorFatal is a wrapper for fatal error checking
func CheckErrorFatal(err error, msg string) {
	if err != nil {
		if msg != "" {
			log.Println(msg)
		}
		log.Fatalln(err)
	}
}
