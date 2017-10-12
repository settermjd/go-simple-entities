package simple_entity

type User struct {
	name string
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) getName() string {
  return u.name
}