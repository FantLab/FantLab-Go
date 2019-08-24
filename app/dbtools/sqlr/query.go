package sqlr

import "fmt"

type Query struct {
	Text string
	Args []interface{}
}

func NewQuery(text string) Query {
	return Query{Text: text}
}

func (q Query) WithArgs(args ...interface{}) Query {
	return Query{
		Text: q.Text,
		Args: args,
	}
}

func (q Query) Format(values ...interface{}) Query {
	return Query{
		Text: fmt.Sprintf(q.Text, values...),
		Args: q.Args,
	}
}

func (q Query) Rebind() Query {
	newArgs, counts := flatArgs(q.Args...)
	newQuery := expandQuery(q.Text, BindVarChar, counts)

	return Query{
		Text: newQuery,
		Args: newArgs,
	}
}

func (q Query) String() string {
	return formatQuery(q.Text, BindVarChar, q.Args...)
}
