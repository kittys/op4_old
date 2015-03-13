package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/couchbase/go-couchbase"
	"github.com/julienschmidt/httprouter"
)

func el(step string, err error) {
	if err != nil {
		log.Fatalf("%v: %v", step, err)
	}
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/:col/:id", Lists)
	http.ListenAndServe(":8080", r)
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "hello from GO")
}

func Lists(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "col: %v\nid: %v", p.ByName("col"), p.ByName("id"))
}

func getDb() {
	c, err := couchbase.Connect("http://localhost:8091/")
	el("connect", err)

	pool, err := c.GetPool("default")
	el("get pool", err)

	bucket, err := pool.GetBucket("default")
	el("get bucket", err)

	bucket.Set("someKey", 0, []string{"an", "example", "list"})
}
