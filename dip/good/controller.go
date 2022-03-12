package good

import "github.com/revenue-hack/solid-sample/dip/good/gooddomain"

func New() {
	impl := &BookRepositoryImpl{}

	b := gooddomain.Book{}
	// DI
	b.Register(impl)
}
