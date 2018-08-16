package torrentclient

import (
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
)

var client *torrent.Client

var clientConfig = torrent.NewDefaultClientConfig()

// Create a Torrent Client
func Create() *torrent.Client {
	clientConfig.Seed = true
	clientConfig.NoDefaultPortForwarding = true
	clientConfig.DisableAggressiveUpload = false
	clientConfig.NoDHT = false
	client, _ = torrent.NewClient(clientConfig)
	return client
}

// Close the running Torrent Client
func Close() {
	client.Close()
}

func GetClientStorage(path string) storage.ClientImpl {
	return storage.ClientImpl(storage.NewFile(path))
}

// GetClient returns the running Torrent Client
func GetClient() *torrent.Client {
	return client
}

func TorrentSpecFromMI(mi *metainfo.MetaInfo) *torrent.TorrentSpec {
	return torrent.TorrentSpecFromMetaInfo(mi)
}
