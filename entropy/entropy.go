package entropy

import (
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"github.com/intelsdi-x/snap/core"
	"os"
)

const (
	vendor = "janczer"

	//pluginName namespace part
	pluginName = "entropy"

	//fs namespace part
	fs = "fs"

	// veersion of entropy plugin
	version = 1

	//pluginType type of plugin
	pluginType = plugin.CollectorPluginType

	//entropy metric from /proc/sys/kernel/random/entropy_avail
	entropy = "entropy"
)

type Plugin struct {
	initialized  bool
	entropy_path string
	host         string
	stat         int
}

var entropyInfo = "/proc/sys/kernel/random/entropy_avail"

func (p *Plugin) GetMetricTypes(cfg plugin.ConfigType) ([]plugin.MetricType, error) {
	metricTypes := []plugin.MetricType{}

	m := plugin.MetricType{
		Namespace_:   core.NewNamespace(vendor, fs, "test", entropy),
		Description_: "Entropy metric",
	}
	metricTypes = append(metricTypes, m)

	return metricTypes, nil
}

func (p *Plugin) CollectMetrics(metricTypes []plugin.MetricType) ([]plugin.MetricType, error) {
	return []plugin.MetricType{}, nil
}

// GetConfigPolicy returns config policy
// It returns error in case retrieval was not successful
func (p *Plugin) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	cp := cpolicy.New()
	//rule, _ := cpolicy.NewStringRule("entropy_path", false)
	//node := cpolicy.NewPolicyNode()
	//node.Add(rule)
	//cp.Add([]string{vendor, fs, pluginName}, node)

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

func New() *Plugin {
	host, err := os.Hostname()

	if err != nil {
		host = "localhost"
	}

	return &Plugin{
		host:         host,
		entropy_path: entropyInfo,
	}
}
