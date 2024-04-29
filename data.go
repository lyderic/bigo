package main

import "time"

type Database struct {
	Encryted bool     `json: "encrypted"`
	Folders  []Folder `json: "folders"`
	Items    []Item   `json: "items"`
}

type Folder struct {
	Id   string `json: "id"`
	Name string `json: "name"`
}

type Item struct {
	PasswordHistory []PasswordHistory `json: "passwordHistory"`
	RevisionDate    time.Time         `json: "revisionDate"`
	CreationDate    time.Time         `json: "creationDate"`
	DeletedDate     time.Time         `json: "deletedDate"`
	Object          string            `json: "object"`
	Id              string            `json: "id"`
	OrganizationId  string            `json: "organizationId"`
	FolderId        string            `json: "folderId"`
	Type            int               `json: "type"`
	Reprompt        int               `json: "reprompt"`
	Name            string            `json: "name"`
	Notes           string            `json: "notes"`
	Favorite        bool              `json: "favorite"`
	CollectionIds   []string          `json: "collectionIds"`
	Login           Login             `json: "login"`
}

type PasswordHistory struct {
	LastUsedDate time.Time
	password     string
}

type Login struct {
	Fido2Credentials     []string  `json: "fido2Credentials"`
	Uris                 []URI     `json: "uris"`
	Username             string    `json: "username"`
	Password             string    `json: "password"`
	Totp                 string    `json: "totp"`
	PasswordRevisionDate time.Time `json: "passwordRevisionDate"`
}

type URI struct {
	Match int    `json: "match"`
	Uri   string `json: "uri"`
}
