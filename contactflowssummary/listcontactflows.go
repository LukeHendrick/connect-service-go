package contactflowssummary

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukehendrick/connectService/util"
)

const NumberOfFlows = 25

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
	instanceID := chi.URLParam(r, "instanceID")
	contactFlows := generateContactFlows(instanceID)
	util.MarshalAndWriteResponse(w, r, contactFlows)
}

// {
// 	Id:              "abc23",
// 	Arn:             "arn:aws",
// 	Name:            "Abc 123",
// 	ContactFlowType: "STANDARD",
// },
func generateContactFlows(instanceID string) (contactFlows []ContactFlowSummary) {
	contactFlows = []ContactFlowSummary{}
	for i := 0; i < NumberOfFlows; i++ {
		contactFlow := generateContactFlow(instanceID)
		contactFlows = append(contactFlows, contactFlow)
	}
	return
}

func generateContactFlow(instanceID string) (contactFlow ContactFlowSummary) {
	id, arn := generateContactFlowARNs(instanceID)
	name := util.GenerateName()
	return ContactFlowSummary{
		Id:              id,
		Arn:             arn,
		Name:            name,
		ContactFlowType: "STANDARD",
	}
}

func generateContactFlowARNs(instanceID string) (id string, arn string) {
	id = util.GenerateID()
	arn = util.FormatARN(instanceID, "contactflow", id)
	return
}
