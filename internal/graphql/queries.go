package graphql

import "fmt"

type Query struct {
	name    string
	queries map[string]*Query
	field   []string
	Query   string `json:"query"`
}

func NewQuery(name string) *Query {
	return &Query{
		name:    name,
		queries: make(map[string]*Query),
		field:   make([]string, 0),
	}
}

// AddField adds query field to the current query
func (q *Query) AddField(field string) []string {
	q.field = append(q.field, field)
	return q.field
}

func (q *Query) AddQuery(query *Query) {
	q.queries[query.name] = query
}

func (q *Query) GetField() []string {
	return q.field
}

func (q *Query) GetQueries() map[string]*Query {
	return q.queries
}

func (q *Query) Build() string {
	buildStr := ""
	for _, query := range q.queries {
		buildStr += query.Build()
	}
	for _, field := range q.field {
		buildStr += fmt.Sprintf("\n%s", field)
	}
	return fmt.Sprintf("%s{%s}\n", q.name, buildStr)
}
