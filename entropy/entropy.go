package entropy

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
	"io/ioutil"
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

var entropyInfo = "/proc/sys/kernel/random/entropy_avail"

type EntropyCollector struct{}

func (EntropyCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {

	metrics := make([]plugin.Metric, 0)
	runTime := time.Now()
	value, err := getEntropy()
	if err == nil {
		mt := plugin.Metric{
			Data:      value,
			Namespace: plugin.NewNamespace(vendor, fs, PluginName),
			Timestamp: runTime,
			Version:   1,
		}
		metrics = append(metrics, mt)
	}

	return metrics, nil
}

func getEntropy() (string, error) {
	value, err := ioutil.ReadFile(entropyInfo)
	if err != nil {
		return "", err
	}

	return strings.Replace(string(value), "\n", "", -1), nil
}

func (EntropyCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{}

	m := plugin.Metric{
		Namespace:   plugin.NewNamespace(vendor, fs, PluginName),
		Description: "Entropy metric",
	}
	metrics = append(metrics, m)

	return metrics, nil
}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (EntropyCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	cp := plugin.NewConfigPolicy()
	return *cp, nil
}
