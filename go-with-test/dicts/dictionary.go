package dicts

type Dict map[string]string

func (d Dict) Search(s string) string {
	return d[s]
}
