package main

import (
	"net/http"

	"github.com/product_page/pkg/handler"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		handler := handler.NewHandler()
		// we need to setup the router inside spin handler
		router := chi.NewRouter()
		router.Get("/product-page", handler.ProductPage)
		router.Get("/api/v1/products", handler.ProductsRoute)
		router.Get("/api/v1//products/{id}", handler.ProductRoute)

		// hand the request/response off to chi
		router.ServeHTTP(res, req)
	})
}

func main() {}
