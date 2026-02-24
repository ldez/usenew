package b

type User struct{}

func Foo(u User) *User {
	return &u
}

func Bar(u *User) *User {
	return u
}

func _() {
	u := User{}

	_ = Foo(u)

	_ = Bar(&u)

	_ = func(a bool) *bool { return &a }(true)

	_ = Pointer(0)
}

func Pointer[T any](v T) *T { return &v }
