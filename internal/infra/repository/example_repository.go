package repository

type ExampleRepository struct {
}

func NewExampleRepository() *ExampleRepository {
	return &ExampleRepository{}
}

func (r *ExampleRepository) ExampleMethod() {
	// Do something
}
