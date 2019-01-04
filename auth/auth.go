package auth

import (
	"math/rand"
	"time"

	"../db"
)

const sessionExpirationTime = int64(60 * 60 * 24)
const sessionIDLength = 32

// const sessionCleanupInterval = 1 * time.Hour

// Session is a temporary user session given after a successful login
type Session struct {
	Name        string
	Expires     int64
	RecordSheet string
}

func (s Session) expired() bool {
	return s.Expires < time.Now().Unix()
}

var sessions map[string]Session

func init() {
	sessions = make(map[string]Session)

	rand.Seed(time.Now().UnixNano())

	// sessionCleanupTicker := time.NewTicker(sessionCleanupInterval)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-sessionCleanupTicker.C:
	// 			removeExpiredSessions()
	// 		}
	// 	}
	// }()
}

// `letters` and `randString()` are used to generate the token for a new session
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// newSession takes in a name and creates a new session
func newSession(name string, recordSheet string) string {
	removeOlderSessions(name)

	var id string

	unique := false
	for !unique {
		id = randString(sessionIDLength)
		if _, ok := sessions[id]; !ok {
			unique = true
		}
	}

	s := Session{
		Name:        name,
		Expires:     time.Now().Unix() + sessionExpirationTime,
		RecordSheet: recordSheet,
	}
	// fmt.Println(s)
	sessions[id] = s

	return id
}

func removeOlderSessions(name string) bool {
	for id, s := range sessions {
		if s.Name == name {
			delete(sessions, id)
			return true
		}
	}
	return false
}

// // removeExpiredSessions is run periodically to remove expired sessions
// func removeExpiredSessions() {
// 	for id, s := range sessions {
// 		if s.expired() {
// 			delete(sessions, id)
// 		}
// 	}
// }

// Login verifies if the name and id match records, and returns a new session ID if it does
func Login(name string, password string) (string, bool, error) {
	recordSheet, err := db.FetchStudent(name, password)
	if err != nil {
		return "", false, err
	}
	if recordSheet != "" {
		return newSession(name, recordSheet), true, nil
	}
	return "", false, nil
}

// GetSession gets the current session's name and sheet ID if valid
func GetSession(id string) (string, string) {
	if _, ok := sessions[id]; ok {
		if sessions[id].expired() {
			delete(sessions, id)
			return "", ""
		}
		return sessions[id].Name, sessions[id].RecordSheet
	}
	return "", ""
}

// VerifySession checks if the session is valid
func VerifySession(id string) bool {
	if _, ok := sessions[id]; ok {
		if sessions[id].expired() {
			delete(sessions, id)
			return false
		}
		return true
	}
	return false
}
