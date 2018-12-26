package auth

import (
	"math/rand"
	"time"

	"../db"
)

const sessionExpirationTime = int64(60 * 60 * 24)
const sessionIDLength = 32
const sessionCleanupInterval = 60 * 60

// Session is a temporary user session given after a successful login
type Session struct {
	Name    string
	Expires int64
}

func (s Session) expired() bool {
	return s.Expires < time.Now().Unix()
}

var sessions map[string]Session

func init() {
	sessions = make(map[string]Session)

	sessionCleanupTicker := time.NewTicker(sessionCleanupInterval * time.Second)
	go func() {
		for {
			select {
			case <-sessionCleanupTicker.C:
				removeExpiredSessions()
			}
		}
	}()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func newSession(name string) string {
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
		Name:    name,
		Expires: time.Now().Unix() + sessionExpirationTime,
	}
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

func removeExpiredSessions() {
	for id, s := range sessions {
		if s.expired() {
			delete(sessions, id)
		}
	}
}

// Login verifies if the name and id match records, and returns a new session ID if it does
func Login(name string, id string) (string, bool, error) {
	ok, err := db.CheckStudent(name, id)
	if err != nil {
		return "", false, err
	}
	if ok {
		return newSession(name), true, nil
	}
	return "", false, nil
}

// VerifySession checks to see if the current session is valid
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
