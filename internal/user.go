package internal

type User struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Surname string `json:"surname,omitempty" yaml:"surname,omitempty"`
	Email   string `json:"email,omitempty" yaml:"email,omitempty"`
	Phone   string `json:"phone,omitempty" yaml:"phone,omitempty"`
}
