package linkPortal

import (
	"fmt"
	"net/http"
	"strings"
)

const jsonContentType = "application/json"
const linkPortalURL = "http://localhost:5000/users/"

type LinkServer struct {
	store PlayerStore
	http.Handler
}

type PlayerStore interface {
	GetUserCreds() UserCredentials
	GetUserLinks() []string
}

func (p *LinkServer) showLinks(w http.ResponseWriter, player string) {
	links := p.store.GetUserLinks()

	fmt.Fprint(w, links)
}

func (p *LinkServer) usersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case http.MethodGet:
		p.showLinks(w, player)
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
