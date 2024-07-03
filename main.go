package main

import (
    "encoding/json"
    "net/http"
    "os"
    "strings"
)

type HostnameInfo struct {
    Hostname string
}

func get_env(w http.ResponseWriter, req *http.Request) {
    var js []byte
    var err error

    var_param := req.URL.Query().Get("var")

    var vars map[string]string = make(map[string]string)

    if var_param == "" {
        for _, element := range os.Environ() {
            variable := strings.Split(element, "=")
            vars[variable[0]] = variable[1]
        }
        js, err = json.Marshal(vars)
    } else {
        vars[var_param] = os.Getenv(var_param)
        js, err = json.Marshal(vars)
    }

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func hostname(w http.ResponseWriter, req *http.Request) {
    h, err := os.Hostname()
    host := HostnameInfo{Hostname: h}
    js, err := json.Marshal(host)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func exit(w http.ResponseWriter, req *http.Request) {
    os.Exit(0)
}

func crash(w http.ResponseWriter, req *http.Request) {
    os.Exit(1)
}

func main() {
    if len(os.Args) == 2 && os.Args[1] == "test" {
        _, err := http.Get("http://localhost/hostname")
        if err != nil {
            os.Exit(1)
        }
        os.Exit(0)
    }
    http.HandleFunc("/env", get_env)
    http.HandleFunc("/exit", exit)
    http.HandleFunc("/crash", crash)
    http.HandleFunc("/hostname", hostname)
    http.ListenAndServe(":80", nil)
}
