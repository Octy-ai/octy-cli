package cli

import (
	"fmt"
	"os"
	"strconv"

	c "github.com/Octy-ai/octy-cli/internal/application/domain/configurations"
	"github.com/Octy-ai/octy-cli/pkg/output"
	"github.com/olekukonko/tablewriter"
)

func getAccConfigController(clia Adapter) {
	spinner := output.StartNewSpinner("Getting account configurations ...", os.Stdout)
	accConf, errs := clia.api.GetAccountConfigs()
	if errs != nil {
		errorFactory(errs, spinner)
	}
	renderAccConfigTable(accConf)
	quit("", 0, spinner)
}

func setAccConfigController(clia Adapter, accountConf AccountConf) {
	spinner := output.StartNewSpinner("Updating account configurations ...", os.Stdout)
	if errs := clia.api.SetAccountConfigs(&accountConf.Configurations); errs != nil {
		errorFactory(errs, spinner)
	}
	renderAccConfigTable(&accountConf.Configurations)
	quit("Updated account configurations", 0, spinner)
}

func getAlgoConfigController(clia Adapter, configType string) {
	spinner := output.StartNewSpinner(fmt.Sprintf("Getting %s algorithm configurations ...", configType), os.Stdout)
	algoConf, errs := clia.api.GetAlgorithmConfigs()
	if errs != nil {
		errorFactory(errs, spinner)
	}
	renderAlgoConfigTable(configType, algoConf)
	quit("", 0, spinner)
}

func setAlgoConfigController(clia Adapter, algoConf AlgoConf) {
	spinner := output.StartNewSpinner("Updating algoritm configurations ...", os.Stdout)
	if errs := clia.api.SetAlgorithmConfigs(&algoConf.Configurations); errs != nil {
		errorFactory(errs, spinner)
	}

	renderAlgoConfigTable("all", &algoConf.Configurations)
	quit("Processing algorithm configuration updates", 0, spinner)
}

//
// Private functions
//

func renderAccConfigTable(configs *c.OctyAccountConfigurations) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Config", "Values"})
	tableData := [][]string{}
	tableData = append(tableData, []string{"Contact name", configs.ContactName})
	tableData = append(tableData, []string{"Contact surname", configs.ContactSurname})
	tableData = append(tableData, []string{"Contact email address", configs.ContactEmailAddress})
	tableData = append(tableData, []string{"Notification Webhook", configs.WebhookURL})
	tableData = append(tableData, []string{"Authenticated ID Key", configs.AuthenticatedIDKey})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(tableData)
	fmt.Println("\n--")
	table.Render()
}

func renderAlgoConfigTable(configType string, configs *[]c.OctyAlgorithmConfiguration) {

	for _, c := range *configs {
		switch c.AlgorithmName {

		case "rec":

			if configType != "rec" && configType != "all" {
				continue
			}
			// init new table writer
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Config", "Values"})

			//init empty table data arr
			tableData := [][]string{}
			fmt.Println("\n--")
			output.BPrint("Recommendations configurations:\n")
			tableData = append(tableData, []string{"Recommend interacted items", strconv.FormatBool(c.Configurations.RecommendInteractedItems)})

			// set default value if nothing in stop list
			if len(c.Configurations.ItemIDStopList) < 1 {
				tableData = append(tableData, []string{"Item id stop list", "(none set)"})
			}
			//iter over stop list
			for _, v := range c.Configurations.ItemIDStopList {
				tableData = append(tableData, []string{"Item id stop list (pending validation)", v})
			}

			// set default value if profile features not set
			if len(c.Configurations.ProfileFeatures) < 1 {
				tableData = append(tableData, []string{"Profile features", "(none set)"})
			}
			//iter over profile features
			for _, v := range c.Configurations.ProfileFeatures {
				tableData = append(tableData, []string{"Profile features", v})
			}

			// populate table and render
			table.SetAutoMergeCells(true)
			table.SetRowLine(true)
			table.AppendBulk(tableData)
			table.Render()

		case "churn":

			if configType != "churn" && configType != "all" {
				continue
			}
			// init new table writer
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Config", "Values"})

			//init empty table data arr
			tableData := [][]string{}
			fmt.Println("\n--")
			output.BPrint("Churn prediction configurations:\n")
			// set default value if profile features not set
			if len(c.Configurations.ProfileFeatures) < 1 {
				tableData = append(tableData, []string{"Profile features", "(none set)"})
			}
			//iter over profile features
			for _, v := range c.Configurations.ProfileFeatures {
				tableData = append(tableData, []string{"Profile features", v})
			}

			// populate table and render
			table.SetAutoMergeCells(true)
			table.SetRowLine(true)
			table.AppendBulk(tableData)
			table.Render()

		}

	}

}
