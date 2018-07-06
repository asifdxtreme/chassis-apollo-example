package schema

import (
	rf "github.com/ServiceComb/go-chassis/server/restful"
	"math/rand"
	"net/http"
	"github.com/ServiceComb/go-chassis/core"
	"github.com/ServiceComb/go-chassis/client/rest"
	"context"
	"log"
	"io/ioutil"
	"fmt"
)

var num = rand.Intn(100)

//RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

//Sayhello is a method used to reply user with hello
func (r *RestFulHello) Sayhello(b *rf.Context) {
	/*id := b.ReadPathParameter("userid")
	log.Printf("get user id: " + id)
	b.Write([]byte(fmt.Sprintf("user %s from %d", id, num)))*/
	id := b.ReadPathParameter("userid")
	restInvoker := core.NewRestInvoker()
	req, _ := rest.NewRequest("GET", "cse://Server/sayhello/"+id)

	resp , err :=  restInvoker.ContextDo(context.TODO(), req)
	if err != nil {
		log.Printf("Error in request call to Server : ", err)
	}
	respbody, _ := ioutil.ReadAll(resp.Resp.Body)
	log.Printf("Response from Server is : ", string(respbody))
	b.Write([]byte(fmt.Sprintf("user %s from %d", string(respbody), num)))
}

//Sayhi is a method used to reply user with hello world text
func (r *RestFulHello) Sayhi(b *rf.Context) {
	result := struct {
		Name string
	}{}
	err := b.ReadEntity(&result)
	if err != nil {
		b.Write([]byte(err.Error() + ":hello world"))
		return
	}
	b.Write([]byte(result.Name + ":hello world"))
	return
}

//URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: http.MethodGet, Path: "/greet/{userid}", ResourceFuncName: "Sayhello"},
		{Method: http.MethodPost, Path: "/guest", ResourceFuncName: "Sayhi"},
	}
}
