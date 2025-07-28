package domain

type LoaderSale interface {
	Load() (s []Sale, err error)
}
