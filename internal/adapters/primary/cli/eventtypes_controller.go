package cli

import (
	"fmt"
	"os"

	"github.com/Octy-ai/octy-cli/pkg/output"
)

func createEventTypesController(clia Adapter, eventTypes EventTypes) {
	spinner := output.StartNewSpinner("Creating event types ...", os.Stdout)
	created, failed, errs := clia.api.CreateEventTypes(&eventTypes.EventTypes)
	if errs != nil {
		errorFactory(errs, spinner)
	}
	fmt.Println("\n--")
	for _, e := range *created {
		output.SPrint(fmt.Sprintf("Created event type: '%v' with ID: %v", e.EventType, e.EventTypeID))
	}
	for _, e := range *failed {
		output.FPrint(fmt.Sprintf("Failed to created event type: '%v' due to Error: %v", e.EventType, e.ErrorMessage))
	}

	quit(fmt.Sprintf("Created a total of %v event type(s)", len(*created)), 0, spinner)
}

func getEventTypesController(clia Adapter, identifiers []string) {
	spinner := output.StartNewSpinner("Getting event types ...", os.Stdout)
	eventTypes, errs := clia.api.GetEventTypes(identifiers)
	if errs != nil {
		errorFactory(errs, spinner)
	}
	fmt.Println()
	for _, e := range *eventTypes {
		fmt.Println("--")
		fmt.Printf("Event type ID: %v\n", e.EventTypeID)
		fmt.Printf("Event type: %v\n", e.EventType)
		fmt.Printf("Event Properties: %v\n", e.EventProperties)
		fmt.Printf("Created At: %v\n", e.CreatedAt)
		fmt.Println()
	}

	quit(fmt.Sprintf("Found a total of %v event type(s)", len(*eventTypes)), 0, spinner)
}

func deleteEventTypesController(clia Adapter, identifiers []string) {
	spinner := output.StartNewSpinner("Deleting event types ...", os.Stdout)
	deleted, failed, errs := clia.api.DeleteEventTypes(identifiers)
	if errs != nil {
		errorFactory(errs, spinner)
	}

	fmt.Println("\n--")
	for _, e := range *deleted {
		output.SPrint(fmt.Sprintf("Deleted event type with ID: %v", e.EventTypeID))
	}
	for _, e := range *failed {
		output.FPrint(fmt.Sprintf("Failed to delete event type: '%v' due to Error: %v", e.EventTypeID, e.ErrorMessage))
	}

	quit(fmt.Sprintf("Deleted a total of %v event type(s)", len(*deleted)), 0, spinner)
}
