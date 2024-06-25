package util

type Nil struct {
	Path []string
}

func (n *Nil) ID() string {
	return JoinPath(n.Path)
}

func (n *Nil) JSON() string {
	return "null"
}

func (n *Nil) Type() string {
	return "string"
}

func (n *Nil) Ptr() any {
	return nil
}

func (n *Nil) Dig(p string) (Object, bool) {
	return nil, false
}
