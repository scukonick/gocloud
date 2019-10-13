package lib

type Computer interface {
	Run(code string) (string, error)
}
