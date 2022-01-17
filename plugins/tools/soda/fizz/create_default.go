package fizz

import (
	"strings"

	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/flect"
)

type createDefault struct{}

func (ct *createDefault) match(name string) bool {
	return false
}

func (ct *createDefault) GenerateFizz(name string, args []string) (string, string, error) {
	var up, down string
	name = ""

	table := fizz.NewTable(name, map[string]interface{}{
		"timestamps": false,
	})

	for _, arg := range args[0:] {
		slice := strings.Split(arg, ":")
		if len(slice) == 1 {
			slice = append(slice, "string")
		}

		o := fizz.Options{}
		name := flect.Underscore(slice[0])
		colType := columnType(slice[1])

		if name == "id" {
			o["primary"] = true
		}

		if strings.HasPrefix(strings.ToLower(slice[1]), "nulls.") {
			o["null"] = true
		}

		if err := table.Column(name, colType, o); err != nil {
			return up, down, err
		}
	}

	up = table.Fizz()
	down = table.UnFizz()

	return up, down, nil
}
