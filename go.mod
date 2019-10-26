module github.com/alexeldeib/cerberus

go 1.13

require (
	github.com/alexeldeib/incendiary-iguana v0.0.13
	github.com/prometheus/common v0.7.0
	github.com/sanity-io/litter v1.2.0
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
