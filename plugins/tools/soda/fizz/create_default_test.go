package fizz

import (
	"strings"
	"testing"
)

func Test_CreateDefault_Table(t *testing.T) {
	ac := createDefault{}
	t.Run("with name and no args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_users", []string{})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		upContains := []string{
			"# You can add your migration from this point. For example:",
			`# create_table("users") {`,
			`#   t.Column("id", "uuid", {primary: true})`,
			`#   t.Column("email", "string", {})`,
		}

		for _, c := range upContains {
			if !strings.Contains(up, c) {
				t.Errorf("expected %v but got %v", c, up)
			}
		}

		downContains := []string{
			"# You can add your migration from this point. For example:",
			`# drop_table("users")`,
		}

		for _, c := range downContains {
			if !strings.Contains(down, c) {
				t.Errorf("expected %v but got %v", c, down)
			}
		}
	})

	t.Run("with name and args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_users", []string{"email"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		upContains := []string{
			"# You can add your migration from this point. For example:",
			`# create_table("users") {`,
			`#   t.Column("id", "uuid", {primary: true})`,
			`#   t.Column("email", "string", {})`,
		}

		for _, c := range upContains {
			if !strings.Contains(up, c) {
				t.Errorf("expected %v but got %v", c, up)
			}
		}

		downContains := []string{
			"# You can add your migration from this point. For example:",
			`# drop_table("users")`,
		}

		for _, c := range downContains {
			if !strings.Contains(down, c) {
				t.Errorf("expected %v but got %v", c, down)
			}
		}
	})
}
