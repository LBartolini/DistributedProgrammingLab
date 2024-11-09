package server

type FileServer interface {
	ListenAndServe(address string) error
	Close()
}
