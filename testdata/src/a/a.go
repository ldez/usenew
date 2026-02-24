package a

type User struct{}

func Foo(u User) *User {
	return &u
}

func Bar(u *User) *User {
	return u
}

func _() {
	u := User{}

	_ = Foo(u) // want "The function can be replaced with 'new'"

	_ = Bar(&u)

	_ = func(a bool) *bool { return &a }(true) // want "The function can be replaced with 'new'"

	_ = Pointer(0) // want "The function can be replaced with 'new'"
}

func Pointer[T any](v T) *T { return &v }
