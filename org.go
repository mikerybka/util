package util

type Org struct {
	Members []string
}

func (o *Org) HasMember(id string) bool {
	for _, m := range o.Members {
		if m == id {
			return true
		}
	}
	return false
}
