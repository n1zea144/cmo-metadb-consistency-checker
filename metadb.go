package main

import (
	"encoding/json"
)

type MetaDBRequestStruct struct {
	DataAccessEmails   string   `json:"dataAccessEmails"`
	DataAnalystEmail   string   `json:"dataAnalystEmail"`
	DataAnalystName    string   `json:"dataAnalystName"`
	InvestigatorEmail  string   `json:"investigatorEmail"`
	InvestigatorName   string   `json:"investigatorName"`
	LabHeadEmail       string   `json:"labHeadEmail"`
	LabHeadName        string   `json:"labHeadName"`
	LibraryType        string   `json:"libraryType"`
	OtherContactEmails string   `json:"otherContactEmails"`
	PIEmail            string   `json:"piEmail"`
	PooledNormals      []string `json:"pooledNormals"`
	ProjectManagerName string   `json:"projectManagerName"`
	QCAccessEmails     string   `json:"qcAccessEmails"`
	Recipe             string   `json:"recipe"`
	RequestID          string   `json:"requestId"`
	//Samples            []SampleStruct `json:"samples"`
	Strand string `json:"strand"`
}
