// Copyright 2019 Alexander Eldeib.

package controllers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexeldeib/atlas/pkg/imds"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var imdsData imds.Metadata

func TestAPIs(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecsWithDefaultAndCustomReporters(t,
		"Controller Suite",
		[]Reporter{envtest.NewlineReporter{}})
}

var _ = BeforeSuite(func(done Done) {
	logf.SetLogger(zap.LoggerTo(GinkgoWriter, true))

	By("Setting up fake instance metadata")
	wd, err := os.Getwd()
	Expect(err).ToNot(HaveOccurred())

	contents, err := ioutil.ReadFile(filepath.Join(wd, "../../testdata/fake.json"))
	Expect(err).ToNot(HaveOccurred())

	err = json.Unmarshal(contents, &imdsData)
	Expect(err).ToNot(HaveOccurred())

	By("bootstrapping test environment")
	testEnv = &envtest.Environment{}

	cfg, err = testEnv.Start()
	Expect(err).ToNot(HaveOccurred())
	Expect(cfg).ToNot(BeNil())

	// +kubebuilder:scaffold:scheme
	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:             scheme.Scheme,
		MetricsBindAddress: "0",
	})
	Expect(err).ToNot(HaveOccurred())

	k8sClient = k8sManager.GetClient()
	Expect(k8sClient).ToNot(BeNil())

	err = (&ConfigMapController{
		Client:   k8sClient,
		Metadata: &imdsData,
	}).SetupWithManager(k8sManager)
	Expect(err).ToNot(HaveOccurred())

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		Expect(err).ToNot(HaveOccurred())
	}()

	close(done)
}, 60)

var _ = AfterSuite(func() {
	By("tearing down the test environment")
	err := testEnv.Stop()
	Expect(err).ToNot(HaveOccurred())
})
