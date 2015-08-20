package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "net/http"
)

 type GvpMeta struct {
    Timestamp string `json:"timestamp"`
    Feed struct {
       FeedAbbrv string `json:"feedAbbrv"`
       FeedName string `json:"feedName"`
      FeedDirectory string `json:"feedDirectory"`
       FeedFilename string `json:"feedFilename"`
    } `json:"feed"`
    Build struct {
        Since int `json:"since"`
        BuildBy string `json:"buildby"`
        ModifiedOnly bool `json:"modifiedonly"`
        DeveloperMode bool `json:"developerMode"`
        Test bool `json:"test"`
    } `json:"build"`
    JiraTickets []string `json:"jiraTickets"`
    Items []JiraTicket `json:"items"`
    BuildTime int `json:"buildtime"`
    BuildTimeSeconds string `json:"buildTimeSeconds"`
    Success bool `json:"success"`
}

type JiraTicket struct {
    AdminDisplay string `json:"adminDisplay"`
    Id string `json:"id"`
    JiraTickets []string `json:"jiraTickets"`
    Path string `json:"path"`
    Success bool `json:"success"`
    Type string `json:"type"`
}


func main() {
    file, err := os.Open("gvp_meta.json")
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    defer file.Close()
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }
    
    var v GvpMeta
    err = json.Unmarshal([]byte(data), &v)
    if err != nil {
        fmt.Printf("error: %v", err)
        return
    }

    output, err := json.MarshalIndent(&v, " ", "    ")
         if err != nil {
                 fmt.Printf("error: %v\n", err)
         }

    // fmt.Println(string(output))
    fmt.Println("Listening on http://localhost:8080/gvp_meta/")

    http.HandleFunc("/gvp_meta/", func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.Write(output)
    })

    http.ListenAndServe(":8080", nil)

}