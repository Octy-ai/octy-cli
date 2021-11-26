package api

type Upload interface {
	GetReferenceMaps(content *[][]string, resourceType string) (map[string]string, map[int]string, *map[string]int, []error)
	ParseProfiles(content *[][]string) (*[][]string, []string)
	ParseItems(content *[][]string) (*[][]string, []string)
	ParseEvents(content *[][]string) (*[][]string, []string)
}
