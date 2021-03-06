# go-adstxt-crawler
[Ads.txt](https://iabtechlab.com/ads-txt-about/) crawler and parser based on [IAB Ads.txt Specification Version 1.0.1](https://iabtechlab.com/wp-content/uploads/2017/09/IABOpenRTB_Ads.txt_Public_Spec_V1-0-1.pdf) implemented in Go

This library provides a mechanism for obtaining and parsing Ads.txt file from websites, or parse your local copy of Ads.txt file

# Motivation
There are some nice online tools for crawling and validating Ads.txt files (for example [Ads.txt validator](https://adstxt.adnxs.com) from AppNexus or another [Ads.txt Validator](https://www.adstxtvalidator.com) by AdReform) that follows [IAB Ads.txt Specification Version 1.0.1](https://iabtechlab.com/wp-content/uploads/2017/09/IABOpenRTB_Ads.txt_Public_Spec_V1-0-1.pdf). 
However, you cannot easily use those tools for massive site scanning since they do not provide free API.

There are also few open source projects I've found Github for scanning Ads.txt files, but at least the ones that I've tried were not fully competible with latest [Ads.txt Spec](https://iabtechlab.com/wp-content/uploads/2017/09/IABOpenRTB_Ads.txt_Public_Spec_V1-0-1.pdf). You can use them, of course, and they would do a decent job, but they are lacking a good validation mechanism to ensure that Ads.txt format is correct and follows the official spec.

This Ads.txt library allows massive sites crawling, and follows [IAB Ads.txt Specification Version 1.0.1](https://iabtechlab.com/wp-content/uploads/2017/09/IABOpenRTB_Ads.txt_Public_Spec_V1-0-1.pdf) to validate that the Ads.txt file is valid

# Examples
You can see [main.go](https://github.com/tzafrirben/go-adstxt-crawler/blob/master/main.go) file for a short example of the adstxt library 2 main methods: adstxt.Get to fetch and parse Ads.txt file from the remote host, or adstxt.ParseBody that can be used to parse the content of a local Ads.txt file

```
req, err := adstxt.NewRequest("example.com")
if err != nil {
  log.Fatal(err)
}
res, err := adstxt.Get(req)
if err != nil {
  log.Fatal(err)
}
// res now holds Ads.txt file DataRecords, Variables and Warnings for Ads.txt parse warnings
for _, r := range res.DataRecords { ... }
for _, v := range res.Variables { ... }
for _, w := range res.Warnings { ... }
```

You can also parse local Ads.txt file in a similar way
```
body, err := ioutil.ReadFile("/<path_to>/ads.txt")
if err != nil {
  log.Fatal(err)
}
rec, err := adstxt.ParseBody(body)
if err != nil {
  log.Fatal(err)
}
// rec now holds Ads.txt file DataRecords, Variables and Warnings for Ads.txt parse warnings
for _, r := range rec.DataRecords { ... }
for _, v := range rec.Variables { ... }
for _, w := range rec.Warnings { ... } 
```

# Import as a Library
import "github.com/tzafrirben/go-adstxt-crawler/adstxt" and you can use adstxt library in your code

# ToDo
Adstxt library still does not support Async crawling. The next version should have a method for scanning multiple domains simultaneously (vs the adstxt.Get method that only scan single Ads.txt file

## LICENSE

MIT

## Author
Tzafrir Ben Ami (a.k.a. tzafrirben)
