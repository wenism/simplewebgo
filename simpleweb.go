package main

import (
    "html/template"
    "log"
    "log/syslog"
    "os"
    "net/http"
    "strings"
)

// This is injected at build time
var AppVersion = "undefined"
var BuiltOn = "undefined"
var BuiltUsing = "undefined"

// This is injected at runtime
var environment = os.Getenv("ENVIRONMENT") 
var containerEngineVersion = os.Getenv("CONTAINER_ENGINE_VERSION")
var operatingSystem = os.Getenv("OS")
var cloudProvider = os.Getenv("CLOUD_PROVIDER")

var templates = template.Must(template.ParseFiles("hello.template.html"))

type Model struct {
    Environment string
    AppVersion string
    BuiltOn string
    BuiltUsing string
    ContainerEngineVersion string
    OperatingSystem string
    CloudProvider string
}

func renderTemplate(w http.ResponseWriter, tmpl string, m *Model) {
    err := templates.ExecuteTemplate(w, tmpl+".template.html", m)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func getModel() *Model {
    return &Model{Environment: environment,
		AppVersion: AppVersion,
	    BuiltOn: BuiltOn,
	    BuiltUsing: BuiltUsing,
	    ContainerEngineVersion: containerEngineVersion,
	    OperatingSystem: operatingSystem,
      CloudProvider: cloudProvider}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "hello", getModel())
	
    log.Print("responding to / request from " + strings.Split(r.RemoteAddr,":")[0])
}

func setupLog() {
    // Configure logger to write to the syslog. You could do this in init(), too.
    logwriter, e := syslog.New(syslog.LOG_NOTICE, "simplewebgo")
    if e == nil {
        log.SetOutput(logwriter)
    }
}

func main() {
  setupLog()
    
  log.Printf("Starting application version %s on %s environment", AppVersion, environment)
    
	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(":9999", nil)
}