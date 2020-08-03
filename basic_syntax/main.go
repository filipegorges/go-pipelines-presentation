// START 1 OMIT
package main // HL

import "fmt" // HL

type Presenter interface { // HL
	Introduce()
}

type Language struct { // HL
	Name string
	Maintainer string
	Motivations []string
}

func NewLanguage(name, maintainer string, motivations []string) *Language { // HL
	return &Language{
		Name: name,
		Maintainer: maintainer,
		Motivations: motivations,
	}
}
// END 1 OMIT
// START 2 OMIT
func (l *Language) Introduce() { // HL
	fmt.Printf("%s is maintained by %s\n", l.Name, l.Maintainer)
	fmt.Println("designed to provide:")
	for i, motivation := range l.Motivations {
		fmt.Printf("%d - %s\n", i+1, motivation)
	}
}

func main() { // HL
	var presenter Presenter = NewLanguage( // HL
		"Golang",
		"Google",
		[]string{
			"Native concurrency constructs",
			"Garbage collection",
			"Ease of adoption",
		},
	)
	presenter.Introduce() // HL
}
// END 2 OMIT
