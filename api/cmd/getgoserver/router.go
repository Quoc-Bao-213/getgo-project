package main

import (
	"net/http"

	product "github.com/Quoc-Bao-213/getgo-project/api/internal/controller/products"
	"github.com/Quoc-Bao-213/getgo-project/api/internal/handler/rest/v1/products"
	"github.com/go-chi/chi/v5"
)

type router struct {
	productCtrl product.Controller
}

func (rtr router) routes(r chi.Router) {
	r.Group(rtr.public)
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello world!!"))
		})
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		// products
		r.Group(func(r chi.Router) {
			pattern := prefix + "/products/"

			r.Post(pattern, products.New(rtr.productCtrl).Create())
			r.Get(pattern, products.New(rtr.productCtrl).GetAllProducts())
			r.Put(pattern+"{productID}", products.New(rtr.productCtrl).GetProductDetails())
		})
	})
}
