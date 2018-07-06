package schema

import (
	"log"
	"net/http"

	"fmt"
	rf "github.com/ServiceComb/go-chassis/server/restful"
	"math/rand"
)

var num = rand.Intn(100)

//RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

//Sayhello is a method used to reply user with hello
func (r *RestFulHello) Sayhello(b *rf.Context) {
	id := b.ReadPathParameter("userid")
	log.Printf("get user id: " + id)
	b.Write([]byte(fmt.Sprintf("user %s from %d", id, num)))
}

//Sayhi is a method used to reply user with hello world text
func (r *RestFulHello) Sayhi(b *rf.Context) {
	result := struct {
		Name string
	}{}
	err := b.ReadEntity(&result)
	if err != nil {
		b.Write([]byte(err.Error() + ":Hello Guest, this is an error"))
		return
	}
	b.Write([]byte(result.Name + ":Hello Guest"))
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/sayhello/{userid}", ResourceFuncName: "Sayhello"},
		{Method: http.MethodPost, Path: "/sayhi", ResourceFuncName: "Sayhi"},
	}
}
