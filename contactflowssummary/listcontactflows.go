package contactflowssummary

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ContactFlowSummaryData is the response structure to the api call
type ContactFlowSummaryData struct {
	ContactFlowSummaryList []ContactFlowSummary
	NextToken              string
}

// ContactFlowSummary contains the information for a single flow
type ContactFlowSummary struct {
	Id              string
	Arn             string
	Name            string
	ContactFlowType string
}

// ListContactFlows returns a list of random contact flow Arns and Names
func ListContactFlows(w http.ResponseWriter, r *http.Request) {
	instance := ContactFlowSummaryData{
		NextToken: "aaa",
		ContactFlowSummaryList: []ContactFlowSummary{
			{
				Id:              "abc23",
				Arn:             "arn:aws",
				Name:            "Abc 123",
				ContactFlowType: "STANDARD",
			},
		},
	}
	fmt.Printf("RECEIVED %+v", r)
	json, err := json.Marshal(instance)
	fmt.Printf("RECEIVED %v", "IT")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "x-amzn-RequestId,x-amzn-ErrorType,x-amzn-ErrorMessage,Date")
	w.Write([]byte(json))
}
