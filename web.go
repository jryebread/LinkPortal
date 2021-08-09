package linkPortal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"

//const linkPortalURL = "http://localhost:5000/users/"

type LinkServer struct {
	store PlayerStore
	http.Handler
}

type PlayerStore interface {
	GetUserCreds() UserCredentials
	GetUserLinks(player string) []UserLinks
	RecordLink(user string, body UserLinks)
}

func (p *LinkServer) showLinks(w http.ResponseWriter, player string) {
	links := p.store.GetUserLinks(player)
	fmt.Println(links)
	jsonInfo, _ := json.Marshal(links)
	log.Printf("jsonInfo: %s\n", jsonInfo)

	fmt.Fprint(w, string(jsonInfo))
}

func (p *LinkServer) recordLink(w http.ResponseWriter, user string,
	body UserLinks) {

	p.store.RecordLink(user, body)
	w.WriteHeader(http.StatusAccepted)
}

func (p *LinkServer) usersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case http.MethodGet:
		p.showLinks(w, player)
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var body UserLinks
		err := decoder.Decode(&body)
		if err != nil {
			log.Fatal(err)
		}
		p.recordLink(w, player, body)
	}
}

func NewLinkServer(store PlayerStore) *LinkServer {
	p := new(LinkServer)

	p.store = store
	router := http.NewServeMux()

	router.Handle("/users/", http.HandlerFunc(p.usersHandler))

	p.Handler = router
	return p
}
