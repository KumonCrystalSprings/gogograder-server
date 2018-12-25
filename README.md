# GoGoGrader Server

Backend for the GoGoGrader website, built with Go.

## Quickstart

### Setup
1. Create a Google Cloud service account with Google Drive API edit credentials
2. Download the service account's client secret file, and put it in the `config/` folder with the name `client_secret.json`
3. Share the appropriate sheets with the service account
4. Copy the IDs of the above sheets into their respective places in `config/sheet.json`

### Running
1. Make sure you have Golang installed on the server
2. `go build main.go`
3. `./main`