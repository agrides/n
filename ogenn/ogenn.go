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

func From[T any, G getter[T], S setter[T]](in G, out S) S {
	To(in, out)
	return out
}

type ptrSetter[S any, T any] interface {
	setter[T]
	*S
}

// Into creates a new S, sets value from getter, and returns S.
func Into[S any, T any, PT ptrSetter[S, T]](in getter[T]) S {
	out := new(S)
	v, ok := in.Get()
	if ok {
		PT(out).SetTo(v)
	} else {
		PT(out).SetToNull()
	}
	return *out
}
