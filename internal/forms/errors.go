package forms

type errors map[string][]string

// Add adds an error message for given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get return the first message
func (e errors) Get(filed string) string {
	es := e[filed]

	if len(es) == 0 {
		return ""
	}
	return es[0]
}
