package Common

type Command interface {
	Execute() error
}
