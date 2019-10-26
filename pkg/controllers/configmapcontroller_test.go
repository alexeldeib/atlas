package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/alexeldeib/atlas/pkg/randstring"
)

var _ = Describe("ConfigMap Controller", func() {
	const timeout = time.Second * 30
	const interval = time.Second * 1
	Context("Config map reconciler", func() {
		It("Should create the region-config configmap", func() {
			ctx := context.Background()

			name := randstring.NewLowerCaseAlphaNumeric(10)
			namespace := &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: name,
				},
			}

			err := k8sClient.Create(ctx, namespace)
			Expect(err).ToNot(HaveOccurred())

			key := types.NamespacedName{
				Name:      "region-config",
				Namespace: name,
			}

			nsName := types.NamespacedName{
				Name: name,
			}

			found := &corev1.ConfigMap{}

			Eventually(func() bool {
				return k8sClient.Get(context.Background(), key, found) == nil
			}, timeout, interval).Should(BeTrue())

			err = k8sClient.Delete(context.Background(), namespace)
			Expect(err).ToNot(HaveOccurred())

			Eventually(func() bool {
				return k8sClient.Get(context.Background(), key, namespace) != nil
			}, timeout, interval).Should(BeTrue())

			Eventually(func() bool {
				return k8sClient.Get(context.Background(), nsName, namespace) != nil
			}, timeout, interval).Should(BeTrue())
		})
	})
})
