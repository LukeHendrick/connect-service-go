package queuesummary

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lukehendrick/connectService/util"
)

const NumberOfQueues = 25

type QueueSummaryData struct {
	QueueSummaryList []QueueSummary
	NextToken        string
}

type QueueSummary struct {
	Id        string
	Arn       string
	Name      string
	QueueType string
}

func ListQueues(w http.ResponseWriter, r *http.Request) {
	instanceID := chi.URLParam(r, "instanceID")
	fmt.Printf(instanceID)
	queues := generateQueues(instanceID)
	util.MarshalAndWriteResponse(w, r, queues)
}

func generateQueues(instanceID string) (queues []QueueSummary) {
	queues = []QueueSummary{}
	for i := 0; i < NumberOfQueues; i++ {
		queue := generateQueue(instanceID)
		queues = append(queues, queue)
	}
	return
}

func generateQueue(instanceID string) (queue QueueSummary) {
	id, arn := generateQueueARNs(instanceID)
	name := util.GenerateName()
	return QueueSummary{
		Id:        id,
		Arn:       arn,
		Name:      name,
		QueueType: "STANDARD",
	}
}

func generateQueueARNs(instanceID string) (id string, arn string) {
	id = util.GenerateID()
	arn = util.FormatARN(instanceID, "queue", id)
	return
}
