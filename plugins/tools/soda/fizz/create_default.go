package fizz

type createDefault struct{}

func (ct *createDefault) match(name string) bool {
	return false
}

func (ct *createDefault) GenerateFizz(name string, args []string) (string, string, error) {
	up := `<%
# You can add your migration from this point. For example:
# create_table("users") {
#   t.Column("id", "uuid", {primary: true})
#   t.Column("email", "string", {})
# }
%>`

	down := `<%
# You can add your migration from this point. For example:
# drop_table("users")
%>`

	return up, down, nil
}
