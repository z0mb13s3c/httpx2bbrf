package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	decoder := json.NewDecoder(os.Stdin)
	
	for {
		json := make(map[string]interface{})
		
		decoder.Decode(&json)
		
		urlUrl := json["url"]	
		urlTitle := json["title"]
		urlWebserver := json["webserver"]
                urlContentType := json["content_type"]
                urlContentLength := json["content_length"]
                urlStatusCode := json["status_code"]
                urlServerResponse := json["body"]
		
		urlUrlString := fmt.Sprintf("%v", urlUrl)
		urlTitleString := fmt.Sprintf("%v", urlTitle)
		urlWebserverString := fmt.Sprintf("%v", urlWebserver)
		urlContentTypeString := fmt.Sprintf("%v", urlContentType)
		urlContentLengthString := fmt.Sprintf("%v", urlContentLength)
		urlStatusCodeString := fmt.Sprintf("%v", urlStatusCode)
		urlServerResponseString := fmt.Sprintf("%v", urlServerResponse)
		
                if urlUrlString != "<nil>" {

                        fullCommand := "bbrf url add '" + urlUrlString + " " + urlStatusCodeString + " " + urlContentLengthString + "' -t 'title:" + urlTitleString + "'" + " -t 'webserver:" + urlWebserverString + "'" + " -t 'contenttype:" + urlContentTypeString + "'" + " -t 'contentlength:" + urlContentLengthString + "'" + " -t 'statuscode:" + urlStatusCodeString + "'" + " -t 'serverresponse:" + urlServerResponseString + "'" + " -p @INFER"
                        out, _ :=exec.Command("sh","-c",fullCommand).Output()
                        fmt.Printf("%s", out)

                } else {
                        break
                }
        }
}
