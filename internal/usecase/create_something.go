package usecase

type CreateSomething struct {
	// Here we define the fields that the usecase needs
}

func NewCreateSomething() *CreateSomething {
	return &CreateSomething{}
}

func (c *CreateSomething) Execute() error {
	// Here we define the business logic
	return nil
}
