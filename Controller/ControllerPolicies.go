package Controller

var Policies = map[string]map[string][]string{
	"*":       {"*": {}},
	"private": {"*": {"ValidateHeaderToken", "ValidateTest"}},
}
