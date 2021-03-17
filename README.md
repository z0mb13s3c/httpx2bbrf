# httpx2bbrf

Simple tool written in GOLANG to send the json output from [Project Discovery's HTTPX](https://github.com/projectdiscovery/httpx) tool to [Honoki's BBRF](
https://github.com/honoki/bbrf-client).

This tool comes with a warning. It's not by any way, shape or form the best code out there. I put it together out of need. If you want to improve it, feel free to do so.

## Prerequisites:

- Both tools need to be installed (duh?)
- HTTPX needs to be executable from every folder (ex: /usr/bin)

## Installation:
```
git clone https://github.com/z0mb13s3c/httpx2bbrf.git
cd httpx2bbrf
go build httpx2bbrf.go
sudo cp httpx2bbrf.go /usr/bin
```

## Usage:
```
bbrf domains | httpx -json -silent | httpx2bbrf
```

If you want to also add the server response:
```
bbrf domains | httpx -response-in-json -json -silent | httpx2bbrf
```
