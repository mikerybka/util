package util

type Org struct {
	Readers Set[string]
	Writers Set[string]
}

func (o *Org) IsReader(userID string) bool {
	return o.Readers.Has(userID)
}

func (o *Org) IsWriter(userID string) bool {
	return o.Writers.Has(userID)
}
