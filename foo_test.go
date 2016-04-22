package vetrepro_test

import "github.com/prashantv/vetrepro"

// This causes a failure
var _ = &vetrepro.Container{
	List: vetrepro.StringList{"hello", "world", "asdasd", "asdasd"},
}

// This works even though it's essentially the same.
var list = vetrepro.StringList{"hello", "world", "asdasd", "asdasd"}
var _ = &vetrepro.Container{
	List: list,
}
