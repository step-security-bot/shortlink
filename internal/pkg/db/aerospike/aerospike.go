package aerospike

import (
	"context"
	"net/url"
	"strconv"

	aero "github.com/aerospike/aerospike-client-go"
	"github.com/spf13/viper"
)

// Config - config
type Config struct {
	host string
	port int
}

// Store implementation of db interface
type Store struct {
	client *aero.Client
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return err
	}

	// Connect to Aerospike
	s.client, err = aero.NewClient(s.config.host, s.config.port)
	if err != nil {
		return err
	}

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) Close() error {
	s.client.Close()
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_AEROSPIKE_URI", "tcp://localhost:3000") // Aerospike URI

	conf, err := url.Parse(viper.GetString("STORE_AEROSPIKE_URI"))
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(conf.Port())
	if err != nil {
		return err
	}

	s.config = Config{
		host: conf.Hostname(),
		port: port,
	}

	return nil
}
