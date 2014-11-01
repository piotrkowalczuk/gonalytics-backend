package lib

// Config ...
type Config struct {
	Server    ServerConfig    `xml:"server"`
	Logger    LoggerConfig    `xml:"logger"`
	Cassandra CassandraConfig `xml:"cassandra"`
	GeoIP     GeoIPConfig     `xml:"geo-ip"`
	RabbitMQ  RabbitMQConfig  `xml:"rabbit-mq"`
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
