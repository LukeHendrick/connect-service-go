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
	queueSummary := generateQueueSummary(queues)
	util.MarshalAndWriteResponse(w, r, queueSummary)
}

func generateQueueSummary(queues []QueueSummary) (queueSummary QueueSummaryData) {
	queueSummary = QueueSummaryData{
		NextToken:        "",
		QueueSummaryList: queues,
	}
	return
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
	name := util.GenerateName(10)
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
