package net

import (
	"time"

	"google.golang.org/grpc"

<<<<<<< HEAD
	//"github.com/dedis/drand/protobuf/control"

=======
	"github.com/dedis/drand/protobuf/control"
	"github.com/dedis/drand/protobuf/dkg"
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	"github.com/dedis/drand/protobuf/drand"
)

//var DefaultTimeout = time.Duration(30) * time.Second

// Gateway is the main interface to communicate to other drand nodes. It
// acts as a listener to receive incoming requests and acts a client connecting
// to drand particpants.
// The gateway fixes all drand functionalities offered by drand.
type Gateway struct {
	Listener
	ProtocolClient
}

// CallOption is simply a wrapper around the grpc options
type CallOption = grpc.CallOption

<<<<<<< HEAD
=======
// InternalClient represents all methods callable on drand nodes which are
// internal to the system. See relevant protobuf files in `/protobuf` for more
// informations.
type InternalClient interface {
	NewBeacon(p Peer, in *drand.BeaconRequest, opts ...CallOption) (*drand.BeaconResponse, error)
	Setup(p Peer, in *dkg.DKGPacket, opts ...CallOption) (*dkg.DKGResponse, error)
	Reshare(p Peer, in *dkg.ResharePacket, opts ...CallOption) (*dkg.ReshareResponse, error)
}

>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
// Listener is the active listener for incoming requests.
type Listener interface {
	Service
	Start()
	Stop()
}

// Service holds all functionalities that a drand node should implement
type Service interface {
	drand.PublicServer
	drand.ControlServer
	drand.ProtocolServer
}

// NewGrpcGatewayInsecure returns a grpc Gateway listening on "listen" for the
// public methods, listening on "port" for the control methods, using the given
// Service s with the given options.
func NewGrpcGatewayInsecure(listen string, s Service, opts ...grpc.DialOption) Gateway {
	return Gateway{
<<<<<<< HEAD
		ProtocolClient: NewGrpcClient(opts...),
		Listener:       NewTCPGrpcListener(listen, s),
=======
		InternalClient:  NewGrpcClient(opts...),
		Listener:        NewTCPGrpcListener(listen, s),
		ControlListener: NewTCPGrpcControlListener(cs, port),
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	}
}

// NewGrpcGatewayFromCertManager returns a grpc gateway using the TLS
// certificate manager
func NewGrpcGatewayFromCertManager(listen string, certPath, keyPath string, certs *CertManager, s Service, opts ...grpc.DialOption) Gateway {
	l, err := NewTLSGrpcListener(listen, certPath, keyPath, s, grpc.ConnectionTimeout(500*time.Millisecond))
	if err != nil {
		panic(err)
	}
	return Gateway{
<<<<<<< HEAD
		ProtocolClient: NewGrpcClientFromCertManager(certs, opts...),
		Listener:       l,
=======
		InternalClient:  NewGrpcClientFromCertManager(certs, opts...),
		Listener:        l,
		ControlListener: NewTCPGrpcControlListener(cs, port),
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	}
}

// StartAll starts the control and public functionalities of the node
func (g Gateway) StartAll() {
	go g.Listener.Start()
}

// StopAll stops the control and public functionalities of the node
func (g Gateway) StopAll() {
	g.Listener.Stop()
}
