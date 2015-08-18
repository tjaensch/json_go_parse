package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
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
    JiraTickets interface{}
    Items interface{}
    BuildTime int `json:"buildtime"`
    BuildTimeSeconds string `json:"buildTimeSeconds"`
    Success bool `json:"success"`
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

    fmt.Println(string(output))

}