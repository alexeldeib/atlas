package configmap

import (
	"errors"
	"strings"

	"github.com/alexeldeib/atlas/pkg/imds"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var locationToShortName = map[string]string{
	"southeastasia":      "sea",
	"uksouth":            "suk",
	"eastus":             "eus",
	"eastus2":            "eus2",
	"westus":             "wus",
	"westus2":            "wus2",
	"northeurope":        "neu",
	"westeurope":         "weu",
	"canadacentral":      "ccan",
	"eastasia":           "ea",
	"centralus":          "cus",
	"westcentralus":      "wcus",
	"australiacentral":   "cau",
	"australiaeast":      "eau",
	"australiasoutheast": "seau",
	"eastus2eap":         "canary",
	"japaneast":          "ejp",
	"northcentralus":     "ncus",
	"francecentral":      "par",
	"southcentralus":     "scus",
	"koreacentral":       "se",
}

var locationToGeography = map[string]string{
	"canadacentral":      "canada",
	"eastasia":           "asiapacific",
	"centralus":          "unitedstates",
	"southeastasia":      "asiapacific",
	"uksouth":            "unitedkingdom",
	"eastus":             "unitedstates",
	"eastus2":            "unitedstates",
	"westus":             "unitedstates",
	"westus2":            "unitedstates",
	"northeurope":        "europe",
	"westeurope":         "europe",
	"westcentralus":      "unitedstates",
	"australiacentral":   "australia",
	"australiaeast":      "australia",
	"australiasoutheast": "australia",
	"eastus2eap":         "canary",
	"japaneast":          "japan",
	"northcentralus":     "unitedstates",
	"francecentral":      "france",
	"southcentralus":     "unitedstates",
	"koreacentral":       "korea",
}

func New(data imds.Metadata, namespace string) (*v1.ConfigMap, error) {
	env := getEnv(data.Compute.ResourceGroupName)
	shortName, err := getShortName(data.Compute.Location)
	if err != nil {
		return nil, err
	}

	geo, err := getGeography(data.Compute.Location)
	if err != nil {
		return nil, err
	}

	result := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "region-config",
			Namespace: namespace,
		},
		Data: map[string]string{
			"resourceGroup": data.Compute.ResourceGroupName,
			"env":           env,
			"location":      data.Compute.Location,
			"shortName":     shortName,
			"geography":     geo,
		},
	}
	return result, nil
}

func getGeography(location string) (string, error) {
	geo, ok := locationToGeography[location]
	if !ok {
		return "", errors.New("location geography not found")
	}
	return geo, nil
}

func getEnv(group string) string {
	env := "int"
	if strings.Contains(group, "prod") {
		env = "prod"
	}
	return env
}

func getShortName(location string) (string, error) {
	short, ok := locationToShortName[location]
	if !ok {
		return "", errors.New("location shortname not found")
	}
	return short, nil
}
