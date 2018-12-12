package handlers

import (
	"../commands"
	"../localfiles"
	"../untils"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (cons *Connections) HandlerClients(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, client := range cons.Clients {
		if client.IsConnect && commands.IsLiveClient(client.Connection) {
			client.IsConnect = true
		} else {
			if client.IsConnect {
				client.IsConnect = false
				untils.WriteMsgLog(fmt.Sprintf("disconnected client (%s)", client.Ip))
				client.Connection.Close()
			}
		}
	}

	templ, err := template.ParseFiles("template/index.html")
	CheckError(err)

	err = templ.Execute(w, *cons)
	CheckError(err)
}

func (cons *Connections) HandlerFunctionClient(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err)
		return
	}

	id, err := untils.ConvStringToId(r.FormValue("id"))

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	cl := cons.Clients[id]

	functions, err := commands.GetFunctions(cl.Connection)

	cl.Functions = functions

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	var templ *template.Template
	templ, err = template.ParseFiles("template/client.html")
	CheckError(err)

	err = templ.Execute(w, cl)
	CheckError(err)
}

func (cons *Connections) HandlerExecuteFunction(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, err)
		return
	}

	id, err := untils.ConvStringToId(r.FormValue("id"))

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	function, err := untils.ConvStringToInt(r.FormValue("f"))

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	cl := cons.Clients[id]
	result, err := commands.ExecuteFunction(cl.Connection, function)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	localfiles.SaveLocalFiles(result.Res)

	var templ *template.Template
	templ, err = template.ParseFiles("template/data.html")

	CheckError(err)

	err = templ.Execute(w, result)
	CheckError(err)
}
