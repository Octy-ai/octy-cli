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
			t.Errorf("ParseProfiles() returned mismatched columns, (-want +got):\n%s", cmp.Diff(table.columns, columns))
		}

	}
}

func TestParseItems(t *testing.T) {

	// input values
	content := [][]string{
		{
			"item_id",
			"item_category",
			"item_name",
			"item_description",
			"item_price",
		},
		{
			"CY47 0245 5793 4PBG CFBG MKL4 UH7J",
			"Chevrolet",
			"Blazer117",
			"eget eleifend",
			"509",
		},
		{
			"AD20 9343 2499 TRIY ZUR6 RK0J",
			"Mercedes-Benz",
			"S-Class694",
			"suscipit nulla",
			"203",
		},
	}

	// dummy output values
	contentC := [][]string{
		{
			"CY47 0245 5793 4PBG CFBG MKL4 UH7J",
			"Chevrolet",
			"Blazer117",
			"eget eleifend",
			"509",
		},
		{
			"AD20 9343 2499 TRIY ZUR6 RK0J",
			"Mercedes-Benz",
			"S-Class694",
			"suscipit nulla",
			"203",
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
				"item_id",
				"item_category",
				"item_name",
				"item_description",
				"item_price",
			},
		},
	}

	u := NewUpload()

	for _, table := range tables {
		contentC, columns := u.ParseItems(table.content)

		// assess contentC
		if !cmp.Equal(table.contentC, contentC) {
			t.Errorf("ParseItems() returned mismatched contentC value, (-want +got):\n%s", cmp.Diff(table.contentC, contentC))
		}

		// assess columns
		if !cmp.Equal(table.columns, columns) {
			t.Errorf("ParseItems() returned mismatched columns, (-want +got):\n%s", cmp.Diff(table.columns, columns))
		}

	}
}

func TestParseEvents(t *testing.T) {

	// input values
	content := [][]string{
		{
			"event_type",
			"profile_id",
			"created_at",
			"event_properties>>item_id",
			"event_properties>>payment_method",
		},
		{
			"charged",
			"profile-123453",
			"2021-07-10 04:59:58",
			"7226981580969-41368328798377",
			"apple_pay",
		},
		{
			"charged",
			"profile-12345224",
			"2021-07-10 05:59:58",
			"7226979188905-41368318705833",
			"bitcoin",
		},
	}

	// dummy output values
	contentC := [][]string{
		{
			"charged",
			"profile-123453",
			"2021-07-10 04:59:58",
			"{\"item_id\":\"7226981580969-41368328798377\",\"payment_method\":\"apple_pay\"}",
		},
		{
			"charged",
			"profile-12345224",
			"2021-07-10 05:59:58",
			"{\"item_id\":\"7226979188905-41368318705833\",\"payment_method\":\"bitcoin\"}",
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
				"event_type",
				"profile_id",
				"created_at",
				"event_properties>>item_id",
				"event_properties>>payment_method",
			},
		},
	}

	u := NewUpload()

	for _, table := range tables {
		contentC, columns := u.ParseEvents(table.content)

		// assess contentC
		if !cmp.Equal(table.contentC, contentC) {
			t.Errorf("ParseEvents() returned mismatched contentC value, (-want +got):\n%s", cmp.Diff(table.contentC, contentC))
		}

		// assess columns
		if !cmp.Equal(table.columns, columns) {
			t.Errorf("ParseEvents() returned mismatched columns, (-want +got):\n%s", cmp.Diff(table.columns, columns))
		}

	}
}
