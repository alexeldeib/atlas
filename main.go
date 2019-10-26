/*
Copyright 2019 Alexander Eldeib.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/alexeldeib/cerberus/pkg/imds"
	"github.com/sanity-io/litter"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	// +kubebuilder:scaffold:scheme
}

type Configuration struct {
	Cloud             string            `json:"cloud"`
	Location          Location          `json:"location"`
	ResourceGroupName string            `json:"resourceGroupName"`
	Tags              map[string]string `json:"tags"`
}

type Location struct {
	Canonical string `json:"canonical"`
	Short     string `json:"short"`
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(func(o *zap.Options) {
		o.Development = true
	}))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
		Port:               9443,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)
	req.Header.Add("Metadata", "True")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2019-03-11")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		setupLog.Error(err, "failed reading json")
	}

	var data imds.Metadata
	if err := json.Unmarshal(resp_body, &data); err != nil {
		setupLog.Error(err, "failed unmarshalling json")
	}

	litter.Dump(data)

	log := ctrl.Log.WithName("cerberus")
	log.Info("passed data dump")

	tags := map[string]string{}
	tagString := strings.Split(data.Compute.Tags, ";")
	for _, tag := range tagString {
		tuple := strings.Split(tag, ":")
		tags[tuple[0]] = tuple[1]
	}

	litter.Dump(tags)

	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
