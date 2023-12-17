package gyocharo

import (
	"fmt"
	"net/http"
	"strings"
)

type ReqURLPath = string
type ReqHandlerFunc func(res http.ResponseWriter, req *http.Request)

type Router struct {
	GetRoutes    map[ReqURLPath]ReqHandlerFunc
	PostRoutes   map[ReqURLPath]ReqHandlerFunc
	PatchRoutes  map[ReqURLPath]ReqHandlerFunc
	PutRoutes    map[ReqURLPath]ReqHandlerFunc
	DeleteRoutes map[ReqURLPath]ReqHandlerFunc

	StaticTypes []string
}

func NewRouter() Router {
	return Router{
		StaticTypes:  []string{"css", "png", "jpg", "mp4"},
		GetRoutes:    map[ReqURLPath]ReqHandlerFunc{},
		PostRoutes:   map[ReqURLPath]ReqHandlerFunc{},
		PatchRoutes:  map[ReqURLPath]ReqHandlerFunc{},
		PutRoutes:    map[ReqURLPath]ReqHandlerFunc{},
		DeleteRoutes: map[ReqURLPath]ReqHandlerFunc{},
	}
}

func (r Router) Get(url ReqURLPath, handler ReqHandlerFunc) {
	r.GetRoutes[url] = handler
}

func (r Router) Post(url ReqURLPath, handler ReqHandlerFunc) {
	r.PostRoutes[url] = handler
}

func (r Router) Patch(url ReqURLPath, handler ReqHandlerFunc) {
	r.PatchRoutes[url] = handler
}

func (r Router) Put(url ReqURLPath, handler ReqHandlerFunc) {
	r.PutRoutes[url] = handler
}

func (r Router) Delete(url ReqURLPath, handler ReqHandlerFunc) {
	r.DeleteRoutes[url] = handler
}

func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	URL := req.URL.Path

	// Try to serve accepted Assets first
	for _, elem := range r.StaticTypes {
		if strings.Contains(URL, elem) {
			assetPath := string(URL[strings.Index(req.URL.Path, "/"):])
			http.ServeFile(res, req, fmt.Sprintf("public/%v", assetPath))
			return
		}
	}

	// Then serve HTML
	if strings.Contains(req.Header.Get("Accept"), "text/html") {
		var m ReqHandlerFunc = nil
		switch req.Method {
		case "POST":
			m = r.PostRoutes[req.URL.Path]
		case "PUT":
			m = r.PutRoutes[req.URL.Path]
		case "DELETE":
			m = r.DeleteRoutes[req.URL.Path]
		case "PATCH":
			m = r.PatchRoutes[req.URL.Path]
		default:
			m = r.GetRoutes[req.URL.Path]
		}
		if m != nil {
			m(res, req)
			return
		}
	}
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte("404 Page or assets not found"))
}
