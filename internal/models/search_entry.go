package models

type SearchEntry struct {
	TotalCount int32         `json:"total_count"`
	Items      []interface{} `json:"items"`
}

func (se *SearchEntry) Append(totalCount int32, items []interface{}) {
	se.TotalCount += totalCount
	se.Items = append(se.Items, items...)
}

func NewSearchEntry(totalCount int32, items []interface{}) *SearchEntry {
	return &SearchEntry{TotalCount: totalCount, Items: items}
}

func SearchEntryConcat(sEntries ...*SearchEntry) SearchEntry {
	result := SearchEntry{}
	for _, entry := range sEntries {
		result.Append(entry.TotalCount, entry.Items)
	}

	return result
}
