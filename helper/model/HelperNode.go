package model

type HelperNode interface {
	Convert(parameter string) (HelperNode, error)
}
