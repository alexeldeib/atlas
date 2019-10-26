// Copyright 2019 Alexander Eldeib.

package configmap

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/alexeldeib/atlas/pkg/imds"
)

func TestConfigMap(t *testing.T) {
	wd, err := os.Getwd()
	failIf(t, err)

	contents, err := ioutil.ReadFile(filepath.Join(wd, "../../testdata/fake.json"))
	failIf(t, err)

	data := &imds.Metadata{}
	if err = json.Unmarshal(contents, data); err != nil {
		t.Errorf(err.Error())
	}

	got, err := New(data, "default")
	failIf(t, err)

	want := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "region-config",
			Namespace: "default",
		},
		Data: map[string]string{
			"cloud":         "AzurePublicCloud",
			"resourceGroup": "bar",
			"env":           "int",
			"location":      "eastus2",
			"shortName":     "eus2",
			"geography":     "unitedstates",
		},
	}

	diff := cmp.Diff(want, got)
	if diff != "" {
		t.Errorf(diff)
	}
}

func TestGetEnv(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"should match int env if present": {"foo-int-bar", "int"},
		"should match int env by default": {"foo-bar", "int"},
		"should match prod env":           {"foo-bar-prod", "prod"},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got := getEnv(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

//nolint:dupl
func TestGetGeo(t *testing.T) {
	tests := map[string]struct {
		input string
		fail  bool
		want  string
	}{
		"should return us for eus":   {"eastus2", false, "unitedstates"},
		"should fail to find foobar": {"foobar", true, ""},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got, err := getGeography(tc.input)
			if tc.fail && err == nil {
				t.Fatalf("expected failure, but succeeded")
			}
			if !tc.fail && err != nil {
				t.Fatalf("expected success, but failed")
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

//nolint:dupl
func TestGetShortName(t *testing.T) {
	tests := map[string]struct {
		input string
		fail  bool
		want  string
	}{
		"should return eus2 for eastus2": {"eastus2", false, "eus2"},
		"should fail for foobar":         {"foobar", true, ""},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			got, err := getShortName(tc.input)
			if tc.fail && err == nil {
				t.Fatalf("expected failure, but succeeded")
			}
			if !tc.fail && err != nil {
				t.Fatalf("expected success, but failed")
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func failIf(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}
