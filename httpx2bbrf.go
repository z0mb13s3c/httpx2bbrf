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
		urlContentType := json["content-type"]
		urlContentLength := json["content-length"]
		urlStatusCode := json["status-code"]
		
		urlUrlString := fmt.Sprintf("%v", urlUrl)
		urlTitleString := fmt.Sprintf("%v", urlTitle)
		urlWebserverString := fmt.Sprintf("%v", urlWebserver)
		urlContentTypeString := fmt.Sprintf("%v", urlContentType)
		urlContentLengthString := fmt.Sprintf("%v", urlContentLength)
		urlStatusCodeString := fmt.Sprintf("%v", urlStatusCode)
		
		if urlUrlString != "<nil>" {
			
			fullCommand := "bbrf url add '" + urlUrlString +"' -t 'title:" + urlTitleString + "'" + " -t 'webserver:" + urlWebserverString + "'" + " -t 'contenttype:" + urlContentTypeString + "'" + " -t 'contentlength:" + urlContentLengthString + "'" + " -t 'statuscode:" + urlStatusCodeString + "'"
			showOutput :=  "echo \"bbrf url add '" + urlUrlString +"' -t 'title:" + urlTitleString + "'" + " -t 'webserver:" + urlWebserverString + "'" + " -t 'contenttype:" + urlContentTypeString + "'" + " -t 'contentlength:" + urlContentLengthString + "'" + " -t 'statuscode:" + urlStatusCodeString + "'\""
			out, _ :=exec.Command("sh","-c",fullCommand).Output()
			out2, _ :=exec.Command("sh","-c",showOutput).Output()
			fmt.Printf("%s", out)
			fmt.Printf("%s", out2)

		} else {
			break
		}
	}
}
