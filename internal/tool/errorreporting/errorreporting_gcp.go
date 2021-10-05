package tool

import (
	"log"

	"cloud.google.com/go/errorreporting"
)

var ErrorClient *errorreporting.Client

func LogAndPrintError(err error) {
	if ErrorClient != nil {
		ErrorClient.Report(errorreporting.Entry{
			Error: err,
		})
		log.Print(err)
	}
}
