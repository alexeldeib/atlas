module github.com/alexeldeib/cerberus

go 1.13

require (
	github.com/prometheus/common v0.4.1
	github.com/sanity-io/litter v1.2.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.4.0 // indirect
	k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
	k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
	k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
	sigs.k8s.io/controller-runtime v0.3.0
)

// replace (
// 	k8s.io/api => k8s.io/api v0.0.0-20190918195907-bd6ac527cfd2
// 	k8s.io/apiextensions-apiserver => 	k8s.io/apiextensions-apiserver v0.0.0-20190918201827-3de75813f604
// 	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190817020851-f2f3a405f61d
// 	k8s.io/client-go => k8s.io/client-go v0.0.0-20190918200256-06eb1244587a
// )
