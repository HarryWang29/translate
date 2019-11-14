package translate

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"translate/model"
	"translate/util"
)

type web struct {
	apiBase
	SchDecoder *schema.Decoder
	SchEncoder *schema.Encoder
}

func NewWeb(typ string, args model.CliArgs) api {
	return &web{
		apiBase: apiBase{
			typ:  typ,
			args: args,
		},
		SchDecoder: schema.NewDecoder(),
		SchEncoder: schema.NewEncoder(),
	}
}

func (t *web) Run() error {
	r := mux.NewRouter()
	r.HandleFunc("/translate", t.handler).Methods(http.MethodGet)

	http.Handle("/", context.ClearHandler(r))
	log.Printf("http is listening on %d", t.args.Port)
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", t.args.Port), nil))
	return nil
}

func (t *web) handler(w http.ResponseWriter, r *http.Request) {
	//处理参数
	err := r.ParseForm()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	//每次请求产生一次备份
	args := t.args
	t.typ = r.Form.Get("from")
	args.Target = r.Form.Get("to")
	subLinks := r.Form["subLink"]
	for _, value := range subLinks {
		//防止两处参数重复
		if util.StringInSlice(value, t.args.SubLinks) {
			continue
		}
		args.SubLinks = append(args.SubLinks, value)
	}
	rule, err := t.translate(t.typ, args)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write(rule)
}
