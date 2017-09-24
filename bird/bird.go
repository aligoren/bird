package bird

import (
	"net/http"
	"html/template"
	"../anka"
	"fmt"
	"strings"
	"time"
	"strconv"
)

var mux = http.NewServeMux()


var RealData interface{}
var RouteArray []string
var QueryList map[string][]string

var Url = "127.0.0.1"
var Port = "8081"
var Path = ""
var Type = ""
var HiddenMsg = ""
var StatusCode = 200
var Protocol = ""

var TemplateName = ""
var UseErrorTemplate = false
var ErrorTemplate = ""
var ErrorMessage = ""

func Template(Name string, Data interface{}) {
	TemplateName = Name
	RealData = Data
}

func Message(Msg string) {
	HiddenMsg = Msg 
}

func MessageHidden(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HiddenMsg)
}

func TemplateHidden(w http.ResponseWriter, r *http.Request) {
	tmplt := template.New(TemplateName + ".html")
	tmplt, _ = tmplt.ParseFiles(TemplateName + ".html")

	tmplt.Execute(w, RealData)
}

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}

func NotFound(params ...string) {
	if params[1] != "" && params[2] != "no" {
		ErrorTemplate = params[1]
		UseErrorTemplate = true
		ErrorMessage = params[0]
	} else {
		ErrorMessage = params[0]
	}
}


func Query(val string) string {
	return QueryList[val][0]
}

func Crow(Route string, f func()) {
	
	RouteArray = append(RouteArray, Route)

	fun := func(w http.ResponseWriter, r *http.Request) {

		Path = r.URL.Path
		QueryList = r.URL.Query()
		Type = r.Method
		Protocol = r.Proto

		if contains(RouteArray, r.URL.Path) {
			StatusCode = 200
			w.WriteHeader(StatusCode)
			f()

			if "/" + TemplateName == Route && TemplateName != "" {
				TemplateHidden(w, r)
			} else {
				TemplateName = ""
				if HiddenMsg != "" && TemplateName == "" && TemplateName != Route {
					MessageHidden(w, r)
				}
			}

		} else {
			StatusCode = 404
			w.WriteHeader(StatusCode)

			if ErrorTemplate != "" && UseErrorTemplate {
				type ErrorMessageS struct {
					Message string
				}

				var Mesg = ErrorMessageS{Message: ErrorMessage}

				var templates = template.Must(template.ParseFiles(ErrorTemplate))

				ErrTpl := strings.Split(ErrorTemplate, "/")

				templates.ExecuteTemplate(w, ErrTpl[1], Mesg)

			} else {
				fmt.Fprintf(w, ErrorMessage)
			}
		}

		ShowConnectionInfo()
	}

	mux.HandleFunc(Route, fun)
}

func StaticFiles() {
	fs := http.FileServer(http.Dir(anka.StaticFiles()))
  	mux.Handle("/"+ anka.StaticFiles() +"/", http.StripPrefix("/"+ anka.StaticFiles() +"/", fs))
}


func ShowConnectionInfo() {
	current_time := time.Now().Local().Format("02/Jan/2006 15:04:05")
	status_code := strconv.Itoa(StatusCode)

	fmt.Printf("\n%s - - [%s] \"%s %s %s\" %s -", Url, current_time, Type, Path, Protocol, status_code)
}

func Serve() {

	StaticFiles()

	ServeBird := Url + ":" + Port
	
	fmt.Printf("\n * Bird running on http://%s\n", ServeBird)


	http.ListenAndServe(ServeBird, mux)

}