package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type IgoDeliveryStruct struct {
	DeliveryDate int    `json:"deliveryDate"`
	Request      string `json:"request"`
}

type IgoSampleStruct struct {
	IgoSampleId          string `json:"igoSampleId"`
	IgoComplete          bool   `json:"igocomplete"`
	InvestigatorSampleId string `json:"investigatorSampleId"`
}

type IgoRequestSamplesStruct struct {
	DataAccessEmails   string            `json:"dataAccessEmails"`
	DataAnalystEmail   string            `json:"dataAnalystEmail"`
	DataAnalystName    string            `json:"dataAnalystName"`
	InvestigatorEmail  string            `json:"investigatorEmail"`
	InvestigatorName   string            `json:"investigatorName"`
	LabHeadEmail       string            `json:"labHeadEmail"`
	LabHeadName        string            `json:"labHeadName"`
	LibraryType        string            `json:"libraryType"`
	OtherContactEmails string            `json:"otherContactEmails"`
	PIEmail            string            `json:"piEmail"`
	PooledNormals      []string          `json:"pooledNormals"`
	ProjectManagerName string            `json:"projectManagerName"`
	QCAccessEmails     string            `json:"qcAccessEmails"`
	Recipe             string            `json:"recipe"`
	RequestID          string            `json:"requestId"`
	Samples            []IgoSampleStruct `json:"samples"`
	Strand             string            `json:"strand"`
}

type IgoRunStruct struct {
	FastQs        []string `json:"fastqs"`
	FlowCellId    string   `json:"flowCellId"`
	FlowCellLanes []int    `json:"flowCellLanes"`
	ReadLength    string   `json:"readLength"`
	RunDate       string   `json:"runDate"`
	RunId         string   `json:"runId"`
	RunMode       string   `json:"runMode"`
}

type IgoLibraryStruct struct {
	BarcodeId                string         `json:"barcodeId"`
	BarcodeIndex             string         `json:"barcodeIndex"`
	CaptureConcentrationNm   string         `json:"captureConcentrationNm"`
	CaptureInputNg           string         `json:"captureInputNg"`
	CaptureName              string         `json:"captureName"`
	DNAInputNg               float64        `json:"dnaInputNg"`
	LibraryConcentrationNgul float64        `json:"libraryConcentrationNgul"`
	LibraryIgoId             string         `json:"libraryIgoId"`
	LibraryVolume            float64        `json:"libraryVolume"`
	Runs                     []IgoRunStruct `json:"runs"`
}

type IgoQCReportStruct struct {
	IGORecommendation    string `json:"IGORecommendation"`
	Comments             string `json:"comments"`
	InvestigatorDecision string `json:"investigatorDecision"`
	QCReportType         string `json:"qcReportType"`
}

type IgoSampleManifestStruct struct {
	BaitSet              string              `json:"baitSet"`
	CFDNA2dBarcode       string              `json:"cfDNA2dBarcode"`
	CMOInfoIgoId         string              `json:"cmoInfoIgoId"`
	CMOPatientId         string              `json:"cmoPatientId"`
	CMOSampleClass       string              `json:"cmoSampleClass"`
	CMOSampleName        string              `json:"cmoSampleName"`
	CollectionYear       string              `json:"collectionYear"`
	IgoId                string              `json:"igoId"`
	InvestigatorSampleId string              `json:"investigatorSampleId"`
	Libraries            []IgoLibraryStruct  `json:"libraries"`
	OncoTreeCode         string              `json:"oncoTreeCode"`
	Preservation         string              `json:"preservation"`
	QCReports            []IgoQCReportStruct `json:"qcReports"`
	SampleName           string              `json:"sampleName"`
	SampleOrigin         string              `json:"sampleOrigin"`
	Sex                  string              `json:"sex"`
	Species              string              `json:"species"`
	SpecimenType         string              `json:"specimenType"`
	TissueLocation       string              `json:"tissueLocation"`
	TubeId               string              `json:"tubeId"`
	TumorOrNormal        string              `json:"tumorOrNormal"`
}

type IgoRequestStruct struct {
	BicAnalysis        bool                      `json:"bicAnalysis"`
	CMORequest         bool                      `json:"cmoRequest"`
	DataAccessEmails   string                    `json:"dataAccessEmails"`
	DataAnalystEmail   string                    `json:"dataAnalystEmail"`
	DataAnalystName    string                    `json:"dataAnalystName"`
	InvestigatorEmail  string                    `json:"investigatorEmail"`
	InvestigatorName   string                    `json:"investigatorName"`
	IsCmoRequest       bool                      `json:"isCmoRequest"`
	LabHeadEmail       string                    `json:"labHeadEmail"`
	LabHeadName        string                    `json:"labHeadName"`
	LibraryType        string                    `json:"libraryType"`
	OtherContactEmails string                    `json:"otherContactEmails"`
	PIEmail            string                    `json:"piEmail"`
	PooledNormals      []string                  `json:"pooledNormals"`
	ProjectManagerName string                    `json:"projectManagerName"`
	QCAccessEmails     string                    `json:"qcAccessEmails"`
	Recipe             string                    `json:"recipe"`
	RequestID          string                    `json:"requestId"`
	Samples            []IgoSampleManifestStruct `json:"samples"`
	Strand             string                    `json:"strand"`
}

func getIgoRequestStruct(request IgoRequestSamplesStruct) IgoRequestStruct {
	r := IgoRequestStruct{}
	r.DataAccessEmails = request.DataAccessEmails
	r.DataAnalystEmail = request.DataAnalystEmail
	r.DataAnalystName = request.DataAnalystName
	r.InvestigatorEmail = request.InvestigatorEmail
	r.InvestigatorName = request.InvestigatorName
	r.LabHeadEmail = request.LabHeadEmail
	r.LabHeadName = request.LabHeadName
	r.LibraryType = request.LibraryType
	r.OtherContactEmails = request.OtherContactEmails
	r.PIEmail = request.PIEmail
	r.PooledNormals = request.PooledNormals
	r.ProjectManagerName = request.ProjectManagerName
	r.QCAccessEmails = request.QCAccessEmails
	r.Recipe = request.Recipe
	r.RequestID = request.RequestID
	r.Strand = request.Strand
	return r
}

func getLimsResp(limsRestURL *url.URL, user string, pwd string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, limsRestURL.String(), nil)
	req.SetBasicAuth(user, pwd)
	req.Header.Set("Accept", "application/json")
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: customTransport}
	return client.Do(req)
}

func getDeliveries(limsHostURL *url.URL, user string, pwd string, deliveryDate time.Time) []IgoDeliveryStruct {
	limsRestString := fmt.Sprintf("%s/LimsRest/api/getDeliveries?timestamp=%d", limsHostURL.String(), deliveryDate.Unix()*1000)
	limsRestURL, err := url.Parse(limsRestString)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := getLimsResp(limsRestURL, user, pwd)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var dels []IgoDeliveryStruct
	err = json.Unmarshal(body, &dels)
	if err != nil {
		log.Fatal(err)
	}

	return dels
}

func getRequests(limsHostURL *url.URL, user string, pwd string, deliveryDate time.Time) []IgoRequestSamplesStruct {
	var reqSamples []IgoRequestSamplesStruct
	dels := getDeliveries(limsHostURL, user, pwd, deliveryDate)
	for _, del := range dels {
		limsRestString := fmt.Sprintf("%s/LimsRest/api/getRequestSamples?request=%s", limsHostURL.String(), del.Request)
		limsRestURL, err := url.Parse(limsRestString)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := getLimsResp(limsRestURL, user, pwd)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		rs := IgoRequestSamplesStruct{}
		err = json.Unmarshal(body, &rs)
		if err != nil {
			log.Fatal(err)
		}
		reqSamples = append(reqSamples, rs)
	}
	return reqSamples
}

func getIgoRequests(limsHostURL *url.URL, user string, pwd string, request IgoRequestSamplesStruct, c chan IgoRequestStruct) {
	igoRequestStruct := getIgoRequestStruct(request)
	for _, s := range request.Samples {
		limsRestString := fmt.Sprintf("%s/LimsRest/api/getSampleManifest?igoSampleId=%s", limsHostURL.String(), s.IgoSampleId)
		limsRestURL, err := url.Parse(limsRestString)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := getLimsResp(limsRestURL, user, pwd)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var sms []IgoSampleManifestStruct
		err = json.Unmarshal(body, &sms)
		if err != nil {
			log.Fatal(err)
		}
		igoRequestStruct.Samples = append(igoRequestStruct.Samples, sms[0])
	}
	c <- igoRequestStruct
}

func GetLimsRestRequests(limsHostURL *url.URL, user string, pwd string, deliveryDate time.Time) []IgoRequestStruct {
	var igoRequests []IgoRequestStruct
	requests := getRequests(limsHostURL, user, pwd, deliveryDate)
	c := make(chan IgoRequestStruct)
	for _, r := range requests {
		go getIgoRequests(limsHostURL, user, pwd, r, c)
	}
	for i := 0; i < len(requests); i++ {
		igoRequests = append(igoRequests, <-c)
	}
	return igoRequests
}
