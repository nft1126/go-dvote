package types

type MessageContext interface {
	ConnectionType() string
	Send(Message)
}

// Message is a wrapper for messages from various net transport modules
type Message struct {
	Data      []byte
	TimeStamp int32
	Namespace string

	Context MessageContext
}

// Connection describes the settings for any of the transports defined in the net module, note that not all
// fields are used for all transport types.
type Connection struct {
	Topic        string // channel/topic for topic based messaging such as PubSub
	Encryption   string // what type of encryption to use
	TransportKey string // transport layer key for encrypting messages
	Key          string // this node's key
	Address      string // this node's address
	Path         string // specific path on which a transport should listen
	SSLDomain    string // ssl domain
	SSLCertDir   string // ssl certificates directory
	Port         int    // specific port on which a transport should listen
}

type DataStore struct {
	Datadir string
}
