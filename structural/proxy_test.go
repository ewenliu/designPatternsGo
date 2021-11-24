package structural

import "testing"

func TestProxy(t *testing.T){
	db := &DB{}
	proxy := &DBProxy{Auth: true, DB: db}
	proxy.Get()
	proxy.Set()

	proxy.Auth = false
	proxy.Set()
}