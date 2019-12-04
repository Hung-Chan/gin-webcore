package administrators

type (
	// Administrator .
	Administrator struct {
		Account  string
		Password string
		Name     string
		GroupID  int
		LevelID  int
		Token    string
		Remark   string
		Enable   int
	}
)
