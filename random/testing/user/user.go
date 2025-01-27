package myuser

type User struct {
	FirstName string
	LastName  string
}

func (u User) FullName() string {
	switch {
	case u.FirstName != "" && u.LastName != "":
		return u.FirstName + " " + u.LastName
	case u.FirstName != "":
		return u.FirstName
	case u.LastName != "":
		return "mr. " + u.LastName
	default:
		return "Anonymous"
	}
}
