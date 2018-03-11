package models

type Person struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Age int `json:"age"`
}

type PersonFunc interface {
  GetId() int
  GetName() string
  GetAge() int
}

func (p Person) GetId() int {
  return p.Id
}

func (p Person) GetName() string {
  return p.Name
}

func (p Person) GetAge() int {
  return p.Age
}
