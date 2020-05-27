package generators

type (
	GeneratorNameDoesNotExist struct {
		Name string
		Err  error
	}
)

func (e GeneratorNameDoesNotExist) Error() string {
	return e.Name + ": " + e.Err.Error()
}
