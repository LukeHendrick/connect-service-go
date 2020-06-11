package contactflowssummary

import (
	"net/http"

	"github.com/lukehendrick/connectService/util"
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
	util.MarshalAndWriteResponse(w, r, instance)
}
