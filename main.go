package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func el(step string, err error) {
	if err != nil {
		log.Fatalf("%v: %v", step, err)
		//panic(err)
	}
}

func main() {
	r := httprouter.New()
	r.GET("/", Index)
	r.GET("/col/:col/:id", Lists)
	r.GET("/db", TestDb)
	http.ListenAndServe(":8080", r)
}

func Test(w http.ResponseWriter, r *http.Response, p httprouter.Params) {
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, "hello from GO")
}

func Lists(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "col: %v\nid iss: %v", p.ByName("col"), p.ByName("id"))
}

func TestDb(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	session, err := mgo.Dial("127.0.0.1, 198.10.10.71")
	el("connect", err)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(
		&Person{Name: "Ball", Phone: "+66 33 123 456", Timestamp: time.Now()},
		&Person{Name: "Fah", Phone: "+66 33 123 799", Timestamp: time.Now()})
	el("insert", err)

	var result []Person
	err = c.Find(bson.M{"name": "Fah"}).Sort("-timestamp").All(&result)
	el("find", err)

	fmt.Fprintf(w, "%v", result)
}
