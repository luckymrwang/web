package handler

import (
	"net/http"

	"urlooker/web/g"
	"urlooker/web/http/render"
)

func GetHostIpItem(w http.ResponseWriter, r *http.Request) {
	hostname := HostnameRequired(r)
	ipItem, exists := g.DetectedItemMap.Get(hostname)
	if !exists {
		render.Data(w, "")
		return
	}
	render.Data(w, ipItem)
}
