module github.com/alexeldeib/atlas

go 1.13

require (
	github.com/go-logr/zapr v0.1.1 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.7.0
	github.com/prometheus/client_golang v1.2.1 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.2.0 // indirect
	go.uber.org/zap v1.11.0 // indirect
	k8s.io/api v0.20.0-alpha.2
	k8s.io/apimachinery v0.20.0-alpha.2
	k8s.io/client-go v0.20.0-alpha.2
	sigs.k8s.io/controller-runtime v0.3.0
)

// `go mod tidy` often breaks kubernetes dependencies. If that happens,
// you need to pin stable combinations of the following four repositories.
// you need to take whatever version controller runtime uses, which is here for v0.3.0
//
// replace (
// 	k8s.io/api => k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
// 	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
// 	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
// 	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
// )
