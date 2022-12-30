package handlers

import (
	"golumbus/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// handler POST
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// handle PUT
	if r.Method == http.MethodPut {
		// expect id in URI parameters
		re := regexp.MustCompile(`/([0-9]+)`)
		ids := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(ids) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(ids[0][1])
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.logger.Println("PUT productId:", id)

		p.updateProduct(id, w, r)
		return
	}

	// handle DELETE
	if r.Method == http.MethodDelete {
		// expect id in URI parameters
		re := regexp.MustCompile(`/([0-9]+)`)
		ids := re.FindAllStringSubmatch(r.URL.Path, -1)

		if len(ids) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(ids[0][1])
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.logger.Println("DELETE productId:", id)

		p.deleteProduct(id, w, r)
		return

	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle GET request")
	prods := data.GetProducts()
	err := prods.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to encode json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle POST request")
	prod := data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		p.logger.Println(err)
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
	}
	data.AddProduct(&prod)

}

func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle PUT request")
	prod := data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		p.logger.Println(err)
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
	p.logger.Printf("Update product %d successfully!\n", id)
}

func (p *Products) deleteProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle DELETE request")
	err := data.DeleteProduct(id)
	if err != nil {
		p.logger.Println(err)
		http.Error(w, "Unable to delete", http.StatusBadRequest)
	}
}
