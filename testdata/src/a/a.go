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

	_ = Foo(u) // want "This call can be replaced with the built-in 'new' function."

	_ = Bar(&u)

	_ = func(a bool) *bool { return &a }(true) // want "This call can be replaced with the built-in 'new' function."

	_ = Pointer(0) // want "This call can be replaced with the built-in 'new' function."
}

func Pointer[T any](v T) *T { return &v }
