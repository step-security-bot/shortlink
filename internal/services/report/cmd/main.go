package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"

	"github.com/shortlink-org/shortlink/internal/services/report/shared"
	"github.com/shortlink-org/shortlink/internal/services/report/workflow"
)

func main() {

	// create a namespace
	c1, err := client.NewNamespaceClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	t := time.Now()
	duration := time.Since(t) + time.Hour*24*1
	err = c1.Register(context.Background(), &workflowservice.RegisterNamespaceRequest{
		Namespace:                        client.DefaultNamespace,
		WorkflowExecutionRetentionPeriod: &duration,
	})
	if err != nil {
		// log.Fatalln("unable to create Temporal client", err)
	}

	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: shared.GreetingTaskQueue,
	}

	// Start the Workflow
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
