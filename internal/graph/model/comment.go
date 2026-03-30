package model

type Comment struct {
    ID     string
    TaskID string
    Text   string
    Author string // user ID
}