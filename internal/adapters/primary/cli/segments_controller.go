package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Octy-ai/octy-cli/pkg/output"
)

func createSegmentsController(clia Adapter, segments Segments) {
	spinner := output.StartNewSpinner("Creating segment definitions ...", os.Stdout)
	created, failed, errs := clia.api.CreateSegments(&segments.Segments)
	if errs != nil {
		errorFactory(errs, spinner)
	}

	fmt.Println("\n--")
	for _, s := range *created {
		output.SPrint(fmt.Sprintf("Created segment definition: '%v' with ID: %v", s.SegmentName, s.SegmentID))
	}
	for _, s := range *failed {
		output.FPrint(fmt.Sprintf("Failed to created segment definition: '%v' due to Error: %v", s.SegmentName, s.ErrorMessage))
	}

	quit(fmt.Sprintf("Created a total of %v segment definition(s)", len(*created)), 0, spinner)
}

func getSegmentsController(clia Adapter, identifiers []string) {
	spinner := output.StartNewSpinner("Getting segment definitions ...", os.Stdout)
	segments, errs := clia.api.GetSegments(identifiers)
	if errs != nil {
		errorFactory(errs, spinner)
	}

	fmt.Println()
	for _, s := range *segments {
		fmt.Println("--")
		fmt.Printf("Segment ID: %v\n", s.SegmentID)
		fmt.Printf("Segment Name: %v\n", s.SegmentName)
		fmt.Printf("Segment Type: %v\n", s.SegmentType)
		fmt.Printf("Segment Subtype: %v\n", s.SegmentSubType)
		fmt.Printf("Segment Timeframe: %v\n", s.SegmentTimeframe)
		fmt.Printf("Event Sequence:\n")
		for _, e := range s.EventSequence {
			fmt.Printf("	Event Sequence >> Event Type: %v\n", e.EventType)
			fmt.Printf("	Event Sequence >> Exp Timeframe: %v\n", e.ExpTimeframe)
			fmt.Printf("	Event Sequence >> Action Inaction: %v\n", e.ActionInaction)
			jsonStr, err := json.Marshal(e.EventProperties)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("	Event Sequence >> Event Properties: %v\n", string(jsonStr))
			fmt.Println("********************************")
		}
		fmt.Printf("Profile Property Name: %v\n", s.ProfilePropertyName)
		fmt.Printf("Profile Property Value: %v\n", s.ProfilePropertyValue)
		fmt.Printf("Profile Count: %v\n", s.ProfileCount)
		fmt.Printf("Created At: %v\n", s.CreatedAt)
		fmt.Println()
	}

	quit(fmt.Sprintf("Found a total of %v segment(s)", len(*segments)), 0, spinner)
}

func deleteSegmentsController(clia Adapter, identifiers []string) {
	spinner := output.StartNewSpinner("Deleting segment definitions ...", os.Stdout)
	deleted, failed, errs := clia.api.DeleteSegments(identifiers)
	if errs != nil {
		errorFactory(errs, spinner)
	}

	fmt.Println("\n--")
	for _, s := range *deleted {
		output.SPrint(fmt.Sprintf("Deleted segment with ID: %v", s.SegmentID))
	}
	for _, s := range *failed {
		output.FPrint(fmt.Sprintf("Failed to delete segment definition: '%v' due to Error: %v", s.SegmentID, s.ErrorMessage))
	}

	quit(fmt.Sprintf("Deleted a total of %v segment definition(s)", len(*deleted)), 0, spinner)
}
