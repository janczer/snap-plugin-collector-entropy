/*

http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 janczer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/
package entropy

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	vendor = "janczer"

	//pluginName namespace part
	PluginName = "entropy"

	//fs namespace part
	fs = "procfs"

	// veersion of entropy plugin
	PluginVersion = 1
)

// entropyInfo source of data for metric
var entropyInfo = "/proc/sys/kernel/random/entropy_avail"

type EntropyCollector struct{}

// CollectMetrics returns list of collect metrics
func (EntropyCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := make([]plugin.Metric, 0)
	v, err := getEntropy()
	if err == nil {
		mt := plugin.Metric{
			Data:      v,
			Namespace: plugin.NewNamespace(vendor, fs, PluginName),
			Timestamp: time.Now(),
			Version:   PluginVersion,
		}
		metrics = append(metrics, mt)
	}

	return metrics, nil
}

// getEntropy read file entropyInfo and get value
// if function have error then first value will be 0 and second error
func getEntropy() (uint64, error) {
	data, err := ioutil.ReadFile(entropyInfo)
	if err != nil {
		return 0, err
	}
	value := strings.Replace(string(data), "\n", "", -1)

	entropy, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return entropy, nil
}

// GetMetricTypes returns list with one available metric type
func (EntropyCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	m := plugin.Metric{
		Namespace:   plugin.NewNamespace(vendor, fs, PluginName),
		Description: "Entropy metric",
	}
	metrics = append(metrics, m)

	return metrics, nil
}

// GetConfigPolicy returns ConfigPolicy
// It returns error in case retrieval was not successful
func (EntropyCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	cp := plugin.NewConfigPolicy()
	return *cp, nil
}
