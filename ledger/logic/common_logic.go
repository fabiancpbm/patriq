package logic

type ILogic[Model any] interface {
	Validate(model *Model) (*Model, error)
}
