//go run ./main.go ./limsrest.go -limsrest-url="https://igolims.mskcc.org:8443" -username= -password= -delivery-date="2021/02/25"

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

func checkArgs(osArgs []string) (*url.URL, string, string, time.Time) {

	if len(osArgs) < 4 {
		fmt.Println("Usage: cmo-metadb-consistency-checker -limsrest-url=[URL] -username=[USERNAME] -password=[PASSWORD] -delivery-date=[DATE]")
		os.Exit(1)
	}

	limsURLStr := flag.String("limsrest-url", "", "Base URL to LimsRest")
	user := flag.String("username", "", "LimsRest username")
	pwd := flag.String("password", "", "LimsRest password")
	delDteStr := flag.String("delivery-date", "", "Delivery Date")

	flag.Parse()

	limsURL, err := url.Parse(*limsURLStr)
	if err != nil {
		log.Fatal(err)
	}

	delDte, err := time.Parse("2006/01/02", *delDteStr)
	if err != nil {
		log.Fatal(err)
	}

	return limsURL, *user, *pwd, delDte
}

func main() {
	limsURL, user, pwd, delDte := checkArgs(os.Args)
	igoRequests := GetLimsRestRequests(limsURL, user, pwd, delDte)
	for _, ir := range igoRequests {
		b, err := json.Marshal(ir)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%+v\n", ir)
		fmt.Println(string(b))
	}
}
