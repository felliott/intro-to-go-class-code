package main

type UserID uint64

func main() {
	var i uint64
	i = 42

	takesUserID(i)
}

func takesUserID(id UserID) UserID {
	return id
}
