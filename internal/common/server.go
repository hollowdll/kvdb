package common

import "math"

const (
	// ServerDefaultHost is the default address of the server.
	ServerDefaultHost string = "localhost"
	// ServerDefaultPort is the default port of the server.
	ServerDefaultPort uint16 = 12345
	// DbMaxKeyCount is the maximum number of keys a database can hold.
	DbMaxKeyCount uint32 = math.MaxUint32
)
