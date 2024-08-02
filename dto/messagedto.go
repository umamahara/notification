package dto

type Message struct {
    // Id      string    `json:"Id"`
    //Title   string `json:"Title"`
    //Desc    string `json:"desc"`

    Channel    string `json:"channel"`
    Receiver    string `json:"receiver"`
    Content string `json:"content"`
}