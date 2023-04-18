package mymod

type Person struct {
	Name     string
	password string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetPassword(password string) {
	p.password = password
}

func (p *Person) GetPassword() string {
	return p.password
}
