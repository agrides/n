package ogenn

type getter[T any] interface {
	Get() (v T, ok bool)
}

type setter[T any] interface {
	SetTo(v T)
	SetToNull()
}

func To[T any, G getter[T], S setter[T]](in G, out S) {
	v, ok := in.Get()
	if ok {
		out.SetTo(v)
		return
	}
	out.SetToNull()
}
