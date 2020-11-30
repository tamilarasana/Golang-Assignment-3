package handlers

import(
	   "log"
	   "net/http"
	   //"encoding/json"
	   "github.com/tamil/Golang-Assignment-3/working/data"

)

type Products struct {
	l*log.Logger

}


func NewProducts(l *log.Logger)*Products{
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){

	if r. Method ==  http.	MethodGet{
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusInternalServerError)
	// lp := data.GetProducts()
	// err := lp.ToJSON(rw)
	// //d, err := json.Marshal(lp)
	// if err != nil{
	// 	http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	// }
	//rw.Write(d)
}


func (g *Products) getProducts(rw http.ResponseWriter, r *http.Request){
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	//d, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
	//rw.Write(d)
}
