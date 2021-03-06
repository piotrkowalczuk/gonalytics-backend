package lib

import "encoding/xml"

const (
	// APIConfigConsumer ...
	APIConfigConsumer = "api"
	// TrackerConfigConsumer ...
	TrackerConfigConsumer = "tracker"
	// ActionsWorkerConfigConsumer ...
	ActionsWorkerConfigConsumer = "worker/actions"
)

// Config ...
type Config struct {
	Server    ServerConfig    `xml:"server"`
	Logger    LoggerConfig    `xml:"logger"`
	Cassandra CassandraConfig `xml:"cassandra"`
	GeoIP     GeoIPConfig     `xml:"geo-ip"`
	Kafka     KafkaConfig     `xml:"kafka"`
}

// APIConfig ...
type APIConfig struct {
	Config
}

// TrackerConfig ...
type TrackerConfig struct {
	Config
}

// ActionsWorkerConfig ...
type ActionsWorkerConfig struct {
	Config
	Worker       WorkerConfig `xml:"worker"`
	MetricsGroup XMLMetrics   `xml:"metrics"`
}

// XMLMetric ...
type XMLMetric struct {
	XMLName         xml.Name      `xml:"metric"`
	DimensionsGroup XMLDimensions `xml:"dimensions"`
}

// XMLMetrics ...
type XMLMetrics struct {
	XMLName xml.Name    `xml:"metrics"`
	Metrics []XMLMetric `xml:"metric"`
}

// XMLDimensions ...
type XMLDimensions struct {
	XMLName    xml.Name    `xml:"dimensions"`
	Dimensions []Dimension `xml:"dimension"`
}

// ServerConfig ...
type ServerConfig struct {
	Host string `xml:"host"`
	Port string `xml:"port"`
}

// GetAddress ...
func (sc *ServerConfig) GetAddress() string {
	return sc.Host + ":" + sc.Port
}

// WorkerConfig ...
type WorkerConfig struct {
	Concurency int64 `xml:"concurency"`
}

// LoggerConfig ...
type LoggerConfig struct {
	Adapter  string `xml:"adapter"`
	Settings string `xml:"settings"`
}

// CassandraConfig ...
type CassandraConfig struct {
	Hosts []struct {
		Host string `xml:"host"`
	} `xml:"hosts"`
	Keyspace string `xml:"keyspace"`
}

// GetHosts ...
func (cc CassandraConfig) GetHosts() []string {
	hosts := make([]string, len(cc.Hosts))

	for index, host := range cc.Hosts {
		hosts[index] = host.Host
	}

	return hosts
}

// KafkaConfig describes configuration of kafka messaging system.
type KafkaConfig struct {
	ConnectionString string `xml:"connection-string"`
	Publishers       struct {
		Action KafkaPublisherConfig `xml:"action"`
	} `xml:"publishers"`
	Consumers struct {
		Action KafkaConsumerConfig `xml:"action"`
	} `xml:"consumers"`
}

// KafkaPublisherConfig describes configuration of kafka publisher.
type KafkaPublisherConfig struct {
	Topic     string `xml:"topic"`
	Partition int    `xml:"partition"`
}

// KafkaConsumerConfig describes configuration of kafka consumer.
type KafkaConsumerConfig struct {
	Topic     string `xml:"topic"`
	Partition int    `xml:"partition"`
	Offset    uint64 `xml:"offset"`
	MaxSize   uint32 `xml:"max-size"`
}

// GeoIPConfig ...
type GeoIPConfig struct {
	Path string `xml:"path"`
}
