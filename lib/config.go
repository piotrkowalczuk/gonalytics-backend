package lib

const (
	APIConfigConsumer           = "api"
	TrackerConfigConsumer       = "tracker"
	ActionsWorkerConfigConsumer = "worker/actions"
)

type Config struct {
	Server    ServerConfig    `xml:"server"`
	Logger    LoggerConfig    `xml:"logger"`
	Cassandra CassandraConfig `xml:"cassandra"`
	GeoIP     GeoIPConfig     `xml:"geo-ip"`
	RabbitMQ  RabbitMQConfig  `xml:"rabbit-mq"`
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
	NbOfChannels int64  `xml:"nb-of-channels"`
	Adapter      string `xml:"adapter"`
	Settings     string `xml:"settings"`
}

// CassandraConfig ...
type CassandraConfig struct {
	Hosts []struct {
		Host string `xml:"host"`
	} `xml:"hosts"`
	Keyspace string `xml:"keyspace"`
}

// GetHosts ...
func (cc *CassandraConfig) GetHosts() []string {
	hosts := make([]string, len(cc.Hosts))

	for _, host := range cc.Hosts {
		hosts = append(hosts, host.Host)
	}

	return hosts
}

// RabbitMQConfig ...
type RabbitMQConfig struct {
	ConnectionString string `xml:"connection-string"`
}

// GeoIPConfig ...
type GeoIPConfig struct {
	Path string `xml:"path"`
}
