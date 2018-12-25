package auth

import (
	"fmt"
	"math/rand"
	"time"
)

type Session struct {
	Name    string
	Expires int64
}

var sessions map[string]Session

func init() {
	sessions = make(map[string]Session)
}

var src = rand.NewSource(time.Now().UnixNano())

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
		id = randString(32)
		if _, ok := sessions[id]; !ok {
			unique = true
		}
	}

	s := Session{
		Name:    name,
		Expires: time.Now().Unix() + 86400,
	}
	sessions[id] = s

	fmt.Println(sessions)
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

func Login(name string, id string) (string, bool) {
	return newSession(name), true
}
