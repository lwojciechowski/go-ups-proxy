package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

const (
	upsRequestString = `
		{
			"UPSSecurity": {
					"UsernameToken": {
							"Username": "{{.Username}}",
							"Password": "{{.Password}}"
					},
					"ServiceAccessToken": {
							"AccessLicenseNumber": "{{.AccessKey}}"
					}
			},
			"TrackRequest": {
					"Request": {
							"RequestOption": "1",
							"TransactionReference": {
									"CustomerContext": "Your Test Case Summary Description"
							}
					},
					"InquiryNumber": "{{.Tracking}}"
			}
		}
	`
)

var upsRequestTpl = template.Must(template.New("upsRequest").Parse(upsRequestString))

func QueryUPS(tracking string) *http.Response {
	client := &http.Client{}
	pr, pw := io.Pipe()

	go func() {
		vars := map[string]interface{}{
			"Username":  os.Getenv("UPS_USERNAME"),
			"Password":  os.Getenv("UPS_PASSWORD"),
			"AccessKey": os.Getenv("UPS_ACCESS_KEY"),
			"Tracking":  tracking,
		}
		upsRequestTpl.Execute(pw, vars)
		pw.Close()
	}()

	isProd := true
	upsTrackingURL := "https://wwwcie.ups.com/rest/Track"

	if isProd {
		upsTrackingURL = "https://onlinetools.ups.com/webservices/Track"
	}

	resp, _ := client.Post(upsTrackingURL, "application/json", pr)

	return resp
}
