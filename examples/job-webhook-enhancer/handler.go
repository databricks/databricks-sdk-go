package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/service/jobs"
)

type databricksWebhookContent struct {
	EventType   string `json:"event_type"`
	WorkspaceID int64  `json:"workspace_id"`
	Run         struct {
		RunID int64 `json:"run_id"`
	}
	Job struct {
		JobID   int64  `json:"job_id"`
		JobName string `json:"name"`
	}
}

func databricksWebhookHandler(w http.ResponseWriter, r *http.Request) {
	var jobMetadata jobs.Run
	var err error = nil
	JobWebhook := webhookParser(r)
	switch EventType := JobWebhook.EventType; EventType {
	case "jobs.on_start":
		log.Printf("workspace_id: %v | job_name: %s | job_id: %v | run_id: %v | event: %s | current_action: None", JobWebhook.WorkspaceID, JobWebhook.Job.JobName, JobWebhook.Job.JobID, JobWebhook.Run.RunID, JobWebhook.EventType)
		jobMetadata, err = databricksGetJobRunMetadata(JobWebhook)
	case "jobs.on_failure":
		log.Printf("workspace_id: %v | job_name: %s | job_id: %v | run_id: %v | event: %s | current_action: enrich_and_forward", JobWebhook.WorkspaceID, JobWebhook.Job.JobName, JobWebhook.Job.JobID, JobWebhook.Run.RunID, JobWebhook.EventType)
		jobMetadata, err = databricksGetJobRunMetadata(JobWebhook)
	case "jobs.on_success":
		log.Printf("workspace_id: %v | job_name: %s | job_id: %v | run_id: %v | event: %s | current_action: None", JobWebhook.WorkspaceID, JobWebhook.Job.JobName, JobWebhook.Job.JobID, JobWebhook.Run.RunID, JobWebhook.EventType)
		jobMetadata, err = databricksGetJobRunMetadata(JobWebhook)
	default:
		log.Printf("Unexpected event")
	}
	if err != nil {
		panic(err)
	}
	SendDetailedWebhook(jobMetadata)
}

func SendDetailedWebhook(jobMetadata jobs.Run) {
	destinationURL := os.Getenv("WEBHOOK_DESTINATION")
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(jobMetadata)
	req, err := http.NewRequest("POST", destinationURL, payload)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Printf("Enriched webhook response status: %s", resp.Status)
}

func databricksGetJobRunMetadata(webhookContent databricksWebhookContent) (jobs.Run, error) {
	w := databricks.Must(databricks.NewWorkspaceClient(&databricks.Config{
		Host:        os.Getenv("DATABRICKS_HOST"),
		Token:       os.Getenv("DATABRICKS_TOKEN"),
		Credentials: config.PatCredentials{},
	}))
	ctx := context.Background()

	runMetadata, err := w.Jobs.GetRun(ctx, jobs.GetRun{
		RunId:          webhookContent.Run.RunID,
		IncludeHistory: true,
	})
	if err != nil {
		panic(err)
	}
	return *runMetadata, nil
}

func webhookParser(req *http.Request) (webhookContent databricksWebhookContent) {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&webhookContent)
	if err != nil {
		panic(err)
	}
	log.Printf("run_id: %v", webhookContent.Run.RunID)
	return
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/Databricks_Webhook_Event_Handler", databricksWebhookHandler)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
