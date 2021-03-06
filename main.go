package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"ads.txt/adstxt"
)

func main() {

	d := flag.String("domain", "", "The remote host to fetch and parse Ads.txt file form.")
	f := flag.String("path", "", "Path to Ads.txt file to parse.")

	flag.Parse()

	// fetch and download Ads.txt file from remote host
	if len(*d) > 0 {
		req, err := adstxt.NewRequest(*d)
		if err != nil {
			log.Fatal(err)
		}
		res, err := adstxt.Get(req)
		if err != nil {
			log.Fatal(err)
		}
		showResults(res.Records)
	}

	// parse local file
	if len(*f) > 0 {
		body, err := ioutil.ReadFile(*f)
		if err != nil {
			log.Fatal(err)
		}
		rec, err := adstxt.ParseBody(body)
		if err != nil {
			log.Fatal(err)
		}
		showResults(rec)
	}
}

func showResults(r *adstxt.Records) {
	if len(r.Warnings) > 0 {
		log.Println("Warnings:")
		for _, w := range r.Warnings {
			j, _ := json.Marshal(w)
			log.Println(string(j))
		}
	}

	if len(r.DataRecords) > 0 {
		log.Println("Data Records:")
		for _, r := range r.DataRecords {
			j, _ := json.Marshal(r)
			log.Println(string(j))
		}
	}

	if len(r.Variables) > 0 {
		log.Println("Variables:")
		for _, v := range r.Variables {
			j, _ := json.Marshal(v)
			log.Println(string(j))
		}
	}
}
