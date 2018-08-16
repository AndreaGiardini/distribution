package torrentclient

import "github.com/anacrolix/torrent"

var client *torrent.Client

var clientConfig = torrent.NewDefaultClientConfig()

// Create a Torrent Client
func Create() *torrent.Client {
	clientConfig.Seed = true
	client, _ = torrent.NewClient(clientConfig)
	return client
}

// Close the running Torrent Client
func Close() {
	client.Close()
}

// GetClient returns the running Torrent Client
func GetClient() *torrent.Client {
	return client
}
