package data_upload

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetReferenceMaps(t *testing.T) {

	// input values
	content := [][]string{
		{
			"customer_id",
			"has_charged",
			"profile_data>>likes",
			"platform_info>>os",
			"profile_data>>gender",
			"profile_data>>score",
			"platform_info>>using_vpn",
		},
		{
			"2748275927498",
			"True",
			"43",
			"linux",
			"female",
			"38.74587133177674",
			"t",
		},
		{
			"2742642659798",
			"False",
			"43",
			"linux",
			"male",
			"38.74587133177674",
			"true",
		},
	}

	// dummy output values
	objectRowIDXMapDummy := map[string]int{
		"2748275927498": 1,
		"2742642659798": 2,
	}

	tables := []struct {
		content         *[][]string
		resourceType    string
		provided        map[string]string
		providedKeys    map[int]string
		objectRowIDXMap *map[string]int
	}{
		{
			&content,
			"profiles",
			map[string]string{
				"customer_id":   "string",
				"has_charged":   "bool",
				"profile_data":  "nested",
				"platform_info": "nested",
			},
			map[int]string{
				0: "customer_id",
				1: "has_charged",
				2: "profile_data",
				3: "platform_info",
			},
			&objectRowIDXMapDummy,
		},
	}

	u := NewUpload()

	for _, table := range tables {
		provided, providedKeys, objectRowIDXMap, _ := u.GetReferenceMaps(table.content, table.resourceType)

		// assess provided
		if !cmp.Equal(table.provided, provided) {
			t.Errorf("GetReferenceMaps() returned mismatched provided columns map, (-want +got):\n%s", cmp.Diff(table.provided, provided))
		}

		// assess provided keys
		if !cmp.Equal(table.providedKeys, providedKeys) {
			t.Errorf("GetReferenceMaps() returned mismatched provided column keys map, (-want +got):\n%s", cmp.Diff(table.providedKeys, providedKeys))
		}

		// assess objectRowIDXMap
		if !cmp.Equal(table.objectRowIDXMap, objectRowIDXMap) {
			t.Errorf("GetReferenceMaps() returned mismatched object row index map, (-want +got):\n%s", cmp.Diff(table.objectRowIDXMap, objectRowIDXMap))
		}
	}

}

func TestParseProfiles(t *testing.T) {

	// input values
	content := [][]string{
		{
			"customer_id",
			"has_charged",
			"profile_data>>likes",
			"profile_data>>gender",
			"profile_data>>score",
			"platform_info>>os",
			"platform_info>>using_vpn",
		},
		{
			"2748275927498",
			"True",
			"65",
			"female",
			"38.74587133177674",
			"linux",
			"t",
		},
		{
			"2742642659798",
			"False",
			"43",
			"male",
			"38.74587133177674",
			"macOS",
			"true",
		},
	}

	// dummy output values
	contentC := [][]string{
		{
			"2748275927498",
			"True",
			"{\"gender\":\"female\",\"likes\":65,\"score\":38.74587133177674}",
			"{\"os\":\"linux\",\"using_vpn\":true}",
		},
		{
			"2742642659798",
			"False",
			"{\"gender\":\"male\",\"likes\":43,\"score\":38.74587133177674}",
			"{\"os\":\"macOS\",\"using_vpn\":true}",
		},
	}

	tables := []struct {
		content  *[][]string
		contentC *[][]string
		columns  []string
	}{
		{
			&content,
			&contentC,
			[]string{
				"customer_id",
				"has_charged",
				"profile_data>>likes",
				"profile_data>>gender",
				"profile_data>>score",
				"platform_info>>os",
				"platform_info>>using_vpn",
			},
		},
	}

	u := NewUpload()

	for _, table := range tables {
		contentC, columns := u.ParseProfiles(table.content)

		// assess contentC
		if !cmp.Equal(table.contentC, contentC) {
			t.Errorf("ParseProfiles() returned mismatched contentC value, (-want +got):\n%s", cmp.Diff(table.contentC, contentC))
		}

		// assess columns
		if !cmp.Equal(table.columns, columns) {
			t.Errorf("GetReferenceMaps() returned mismatched columns, (-want +got):\n%s", cmp.Diff(table.columns, columns))
		}

	}
}
