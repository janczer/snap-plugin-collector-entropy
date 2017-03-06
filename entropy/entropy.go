package entropy

import (
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"io/ioutil"
	"strings"
	"time"
)

const (
	vendor = "janczer"

	//pluginName namespace part
	pluginName = "entropy"

	//fs namespace part
	fs = "procfs"

	// veersion of entropy plugin
	version = 1

	//pluginType type of plugin
	pluginType = plugin.CollectorPluginType
)

var entropyInfo = "/proc/sys/kernel/random/entropy_avail"

type Plugin struct{}

func New() *Plugin {
	return &Plugin{}
}

func (p *Plugin) CollectMetrics(mts []plugin.MetricType) ([]plugin.MetricType, error) {

	metrics := make([]plugin.MetricType, 0)
	runTime := time.Now()
	value, err := getEntropy()
	if err == nil {
		mt := plugin.MetricType{
			Data_:      value,
			Namespace_: core.NewNamespace(vendor, fs, pluginName),
			Timestamp_: runTime,
			Version_:   1,
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

func (p *Plugin) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	metricTypes := []plugin.MetricType{}

	m := plugin.MetricType{
		Namespace_:   core.NewNamespace(vendor, fs, pluginName),
		Description_: "Entropy metric",
	}
	metricTypes = append(metricTypes, m)

	return metricTypes, nil
}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (p *Plugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	return cp, nil
}

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		pluginName,
		version,
		pluginType,
		[]string{},
		[]string{plugin.SnapGOBContentType},
		plugin.ConcurrencyCount(1),
	)
}
