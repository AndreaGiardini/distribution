package torrentclient

import "github.com/anacrolix/torrent"

var client *torrent.Client

var clientConfig = &torrent.ClientConfig{
	Seed: true,
}

// Create a Torrent Client
func Create() *torrent.Client {
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
