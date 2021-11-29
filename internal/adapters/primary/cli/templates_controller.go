package cli

import (
	"encoding/json"
	"fmt"
	"os"

	te "github.com/Octy-ai/octy-cli/internal/application/domain/messaging"
	"github.com/Octy-ai/octy-cli/pkg/output"
)

func createUpdateTemplatesController(clia Adapter, t Templates) {

	createTemplates := *te.NewTemps()
	updateTemplates := *te.NewTemps()

	for _, template := range t.Templates {
		if template.TemplateID == "" {
			createTemplates = append(createTemplates, template)
		} else {
			updateTemplates = append(updateTemplates, template)
		}
	}

	if len(createTemplates) > 0 {
		spinner := output.StartNewSpinner("Creating templates ...", os.Stdout)
		var createdCount int = 0
		created, failed, errs := clia.api.CreateTemplates(&createTemplates)
		if errs != nil {
			errorFactory(errs, false, spinner)
		}
		fmt.Println("\n--")
		if created != nil {
			createdCount = len(*created)
			for _, t := range *created {
				output.SPrint(fmt.Sprintf("Created message template: '%v' with ID: %v", t.FriendlyName, t.TemplateID))
			}
		}
		if failed != nil {
			for _, t := range *failed {
				output.FPrint(fmt.Sprintf("Failed to created message template: '%v' due to Error: %v", t.FriendlyName, t.ErrorMessage))
			}
		}

		output.SPrint(fmt.Sprintf("Created a total of %v message template(s)", createdCount))
		spinner.Stop()
	}

	if len(updateTemplates) > 0 {
		fmt.Println("\n--")
		spinner := output.StartNewSpinner("Updating templates ...", os.Stdout)
		var updatedCount int = 0
		updated, failed, errs := clia.api.UpdateTemplates(&updateTemplates)
		if errs != nil {
			errorFactory(errs, false, spinner)
		}
		fmt.Println("\n--")
		if updated != nil {
			updatedCount = len(*updated)
			for _, t := range *updated {
				output.SPrint(fmt.Sprintf("Updated message template: '%v' with ID: %v", t.FriendlyName, t.TemplateID))
			}
		}

		if failed != nil {
			for _, t := range *failed {
				output.FPrint(fmt.Sprintf("Failed to update message template: '%v' due to Error: %v", t.FriendlyName, t.ErrorMessage))
			}
		}

		output.SPrint(fmt.Sprintf("Updated a total of %v message template(s)", updatedCount))
		spinner.Stop()
	}

	quit("", 0, nil)
}

func getTemplatesController(clia Adapter, identifiers []string, ids bool) {
	spinner := output.StartNewSpinner("Getting message templates ...", os.Stdout)
	templates, errs := clia.api.GetTemplates(identifiers)
	if errs != nil {
		errorFactory(errs, true, spinner)
	}
	fmt.Println()
	for _, t := range *templates {
		if ids {
			fmt.Println("--")
			fmt.Printf("Template ID: %v\n", t.TemplateID)
			continue
		}
		fmt.Println("--")
		fmt.Printf("Template ID: %v\n", t.TemplateID)
		fmt.Printf("Friendly Name: %v\n", t.FriendlyName)
		fmt.Printf("Template Type: %v\n", t.TemplateType)
		fmt.Printf("Title: %v\n", t.Title)
		fmt.Printf("Content: %v\n", t.Content)
		fmt.Printf("Required Data: %v\n", t.RequiredData)
		DefaultValuesJsonStr, err := json.Marshal(t.DefaultValues)
		if err != nil {
			fmt.Printf("Default Values: {} \n")
		} else {
			fmt.Printf("Default Values: %v\n", string(DefaultValuesJsonStr))
		}
		MetadataJsonStr, err := json.Marshal(t.Metadata)
		if err != nil {
			fmt.Printf("Metadata: null \n")
		} else {
			fmt.Printf("Metadata: %v\n", string(MetadataJsonStr))
		}
		fmt.Printf("Created At: %v\n", t.CreatedAt)
		fmt.Printf("Updated At: %v\n", t.UpdatedAt)
		fmt.Println()
	}

	quit(fmt.Sprintf("Found a total of %v message template(s)", len(*templates)), 0, spinner)
}

func deleteTemplatesController(clia Adapter, identifiers []string) {
	spinner := output.StartNewSpinner("Deleting message templates ...", os.Stdout)
	deleted, failed, errs := clia.api.DeleteTemplates(identifiers)
	if errs != nil {
		errorFactory(errs, true, spinner)
	}

	fmt.Println("\n--")
	for _, t := range *deleted {
		output.SPrint(fmt.Sprintf("Deleted message template with ID: %v", t.TemplateID))
	}
	for _, t := range *failed {
		output.FPrint(fmt.Sprintf("Failed to delete message template: '%v' due to Error: %v", t.TemplateID, t.ErrorMessage))
	}

	quit(fmt.Sprintf("Deleted a total of %v message template(s)", len(*deleted)), 0, spinner)
}
