package fizz

import (
	"strings"
	"testing"
)

func Test_CreateDefault_Table(t *testing.T) {
	ac := createDefault{}
	t.Run("with table name and no args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_users", []string{})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP := `create_table("")`
		expectedDown := `drop_table("")`

		if !strings.Contains(up, expectedUP) {
			t.Errorf("expected %v but got %v", expectedUP, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})

	t.Run("with table name and args", func(t *testing.T) {
		up, down, err := ac.GenerateFizz("create_users", []string{"email"})
		if err != nil {
			t.Error("should not be nil but got err")
		}

		expectedUP1 := `create_table("")`
		expectedUP2 := `t.Column("email", "string", {})`
		expectedDown := `drop_table("")`

		if !strings.Contains(up, expectedUP1) {
			t.Errorf("expected %v but got %v", expectedUP1, up)
		}

		if !strings.Contains(up, expectedUP2) {
			t.Errorf("expected %v but got %v", expectedUP2, up)
		}

		if down != expectedDown {
			t.Errorf("expected %v but got %v", expectedDown, down)
		}
	})
}
