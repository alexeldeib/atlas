/*
Copyright 2019 Alexander Eldeib.
*/

package imds

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func New() (Metadata, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://169.254.169.254/metadata/instance", nil)
	req.Header.Add("Metadata", "True")

	q := req.URL.Query()
	q.Add("format", "json")
	q.Add("api-version", "2019-03-11")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return Metadata{}, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Metadata{}, err
	}

	var data Metadata
	err = json.Unmarshal(resp_body, &data)
	return data, err

	// litter.Dump(data)

	// tags := map[string]string{}
	// tagString := strings.Split(data.Compute.Tags, ";")
	// for _, tag := range tagString {
	// 	tuple := strings.Split(tag, ":")
	// 	tags[tuple[0]] = tuple[1]
	// }

	// litter.Dump(tags)
	// log.Info("passed tags")
}

type Metadata struct {
	Compute Compute `json:"compute"`
	// Network Network `json:"network"`
}

type Compute struct {
	AzEnvironment string `json:"azEnvironment"`
	// CustomData           string      `json:"customData"`
	Location string `json:"location"`
	Name     string `json:"name"`
	// Offer                string      `json:"offer"`
	// OsType               string      `json:"osType"`
	PlacementGroupID string `json:"placementGroupId"`
	// Plan                 Plan        `json:"plan"`
	PlatformFaultDomain  string `json:"platformFaultDomain"`
	PlatformUpdateDomain string `json:"platformUpdateDomain"`
	// Provider             string      `json:"provider"`
	// PublicKeys []PublicKey `json:"publicKeys"`
	// Publisher            string      `json:"publisher"`
	ResourceGroupName string `json:"resourceGroupName"`
	ResourceID        string `json:"resourceId"`
	Sku               string `json:"sku"`
	SubscriptionID    string `json:"subscriptionId"`
	Tags              string `json:"tags"`
	Version           string `json:"version"`
	VMID              string `json:"vmId"`
	// VMScaleSetName    string `json:"vmScaleSetName"`
	VMSize string `json:"vmSize"`
	Zone   string `json:"zone"`
}

// type Plan struct {
// 	Name      string `json:"name"`
// 	Product   string `json:"product"`
// 	Publisher string `json:"publisher"`
// }

// type PublicKey struct {
// 	KeyData string `json:"keyData"`
// 	Path    string `json:"path"`
// }

// type Network struct {
// 	Interface []Interface `json:"interface"`
// // }

// type Interface struct {
// 	IPv4       IPv4   `json:"ipv4"`
// 	IPv6       IPv6   `json:"ipv6"`
// 	MacAddress string `json:"macAddress"`
// }

// type IPv4 struct {
// 	IPAddress []IPAddress `json:"ipAddress"`
// 	Subnet    []Subnet    `json:"subnet"`
// }

// type IPv6 struct {
// 	IPAddress []IPAddress `json:"ipAddress"`
// }

// type Subnet struct {
// 	Address string `json:"address"`
// 	Prefix  string `json:"prefix"`
// }

// type IPAddress struct {
// 	PrivateIPAddress string `json:"privateIpAddress"`
// 	PublicIPAddress  string `json:"publicIpAddress"`
// }
