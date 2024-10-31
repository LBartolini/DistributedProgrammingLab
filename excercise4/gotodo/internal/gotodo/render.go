package gotodo

type Render interface {
	Render(todos *Todos) error
}
