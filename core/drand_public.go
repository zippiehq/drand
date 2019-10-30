package core

import (
	"context"
<<<<<<< HEAD
	"errors"
	"fmt"
	"time"

	"github.com/dedis/drand/beacon"
	"github.com/dedis/drand/ecies"
	"github.com/dedis/drand/entropy"
	"github.com/dedis/drand/key"
	"github.com/dedis/drand/protobuf/drand"
	"google.golang.org/grpc/peer"
)

// Setup is the public method to call during a DKG protocol.
func (d *Drand) Setup(c context.Context, in *drand.SetupPacket) (*drand.Empty, error) {
=======
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/dedis/drand/beacon"
	"github.com/dedis/drand/ecies"
	"github.com/dedis/drand/key"
	"github.com/dedis/drand/protobuf/crypto"
	dkg_proto "github.com/dedis/drand/protobuf/dkg"
	"github.com/dedis/drand/protobuf/drand"
	"github.com/dedis/kyber"
	"github.com/nikkolasg/slog"
)

// Setup is the public method to call during a DKG protocol.
func (d *Drand) Setup(c context.Context, in *dkg_proto.DKGPacket) (*dkg_proto.DKGResponse, error) {
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	d.state.Lock()
	defer d.state.Unlock()
	if d.dkgDone {
		return nil, errors.New("drand: dkg finished already")
	}
	if d.dkg == nil {
		return nil, errors.New("drand: no dkg running")
	}
<<<<<<< HEAD
	d.dkg.Process(c, in.Dkg)
	return new(drand.Empty), nil
}

// Reshare is called when a resharing protocol is in progress
func (d *Drand) Reshare(c context.Context, in *drand.ResharePacket) (*drand.Empty, error) {
=======
	d.dkg.Process(c, in)
	return &dkg_proto.DKGResponse{}, nil
}

// Reshare is called when a resharing protocol is in progress
func (d *Drand) Reshare(c context.Context, in *dkg_proto.ResharePacket) (*dkg_proto.ReshareResponse, error) {
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	d.state.Lock()
	defer d.state.Unlock()

	if d.nextGroupHash == "" {
<<<<<<< HEAD
		return nil, fmt.Errorf("drand %s: can't reshare because InitReshare has not been called", d.priv.Public.Addr)
=======
		return nil, errors.New("drand: can't reshare because InitReshare has not been called")
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	}

	// check that we are resharing to the new group that we expect
	if in.GroupHash != d.nextGroupHash {
		return nil, errors.New("drand: can't reshare to new group: incompatible hashes")
	}

<<<<<<< HEAD
	if !d.nextFirstReceived && d.nextOldPresent {
		d.nextFirstReceived = true
		// go routine since StartDKG requires the global lock
		go d.StartDKG()
=======
	if in.Packet == nil {
		// indicator that we should start the DKG as we are one node in the old
		// list that should reshare its share
		go d.StartDKG()
		return &dkg_proto.ReshareResponse{}, nil
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	}

	if d.dkg == nil {
		return nil, errors.New("drand: no dkg setup yet")
	}

<<<<<<< HEAD
	d.nextFirstReceived = true
	if in.Dkg != nil {
		// first packet from the "leader" contains a nil packet for
		// nodes that are in the old list that must broadcast their
		// deals.
		d.dkg.Process(c, in.Dkg)
	}
	return new(drand.Empty), nil
=======
	// we just relay to the dkg
	d.dkg.Process(c, in.Packet)
	return &dkg_proto.ReshareResponse{}, nil
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
}

// NewBeacon methods receives a beacon generation requests and answers
// with the partial signature from this drand node.
func (d *Drand) NewBeacon(c context.Context, in *drand.BeaconRequest) (*drand.BeaconResponse, error) {
	d.state.Lock()
	defer d.state.Unlock()
	if d.beacon == nil {
		return nil, errors.New("drand: beacon not setup yet")
	}
	return d.beacon.ProcessBeacon(c, in)
}

<<<<<<< HEAD
// PublicRand returns a public random beacon according to the request. If the Round
// field is 0, then it returns the last one generated.
func (d *Drand) PublicRand(c context.Context, in *drand.PublicRandRequest) (*drand.PublicRandResponse, error) {
	d.state.Lock()
	defer d.state.Unlock()
	if d.beacon == nil {
		return nil, errors.New("drand: beacon generation not started yet")
	}
=======
// Public returns a public random beacon according to the request. If the Round
// field is 0, then it returns the last one generated.
func (d *Drand) Public(c context.Context, in *drand.PublicRandRequest) (*drand.PublicRandResponse, error) {
	d.state.Lock()
	defer d.state.Unlock()
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	var beacon *beacon.Beacon
	var err error
	if in.GetRound() == 0 {
		beacon, err = d.beaconStore.Last()
	} else {
		beacon, err = d.beaconStore.Get(in.GetRound())
	}
	if err != nil {
		return nil, fmt.Errorf("can't retrieve beacon: %s", err)
	}
<<<<<<< HEAD
	peer, ok := peer.FromContext(c)
	if ok {
		d.log.With("module", "public").Info("public_rand", peer.Addr.String(), "round", beacon.Round)
		d.log.Info("public rand", peer.Addr.String(), "round", beacon.Round)
	}
	h := RandomnessHash()
	h.Write(beacon.GetSignature())
	randomness := h.Sum(nil)
	return &drand.PublicRandResponse{
		Previous:   beacon.GetPreviousSig(),
		Round:      beacon.Round,
		Signature:  beacon.GetSignature(),
		Randomness: randomness,
	}, nil
}

// PrivateRand returns an ECIES encrypted random blob of 32 bytes from /dev/urandom
func (d *Drand) PrivateRand(c context.Context, priv *drand.PrivateRandRequest) (*drand.PrivateRandResponse, error) {
	protoPoint := priv.GetRequest().GetEphemeral()
	point := key.G2.Point()
	if err := point.UnmarshalBinary(protoPoint); err != nil {
		return nil, err
	}
	msg, err := ecies.Decrypt(key.G2, ecies.DefaultHash, d.priv.Key, priv.GetRequest())
	if err != nil {
		d.log.With("module", "public").Error("private", "invalid ECIES", "err", err.Error())
=======
	return &drand.PublicRandResponse{
		Previous:   beacon.PreviousRand,
		Round:      beacon.Round,
		Randomness: beacon.Randomness,
	}, nil
}

// Private returns an ECIES encrypted random blob of 32 bytes from /dev/urandom
func (d *Drand) Private(c context.Context, priv *drand.PrivateRandRequest) (*drand.PrivateRandResponse, error) {
	protoPoint := priv.GetRequest().GetEphemeral()
	point, err := crypto.ProtoToKyberPoint(protoPoint)
	if err != nil {
		return nil, err
	}
	groupable, ok := point.(kyber.Groupable)
	if !ok {
		return nil, errors.New("point is not on a registered curve")
	}
	if groupable.Group().String() != key.G2.String() {
		return nil, errors.New("point is not on the supported curve")
	}
	msg, err := ecies.Decrypt(key.G2, ecies.DefaultHash, d.priv.Key, priv.GetRequest())
	if err != nil {
		slog.Debugf("drand: received invalid ECIES private request: %s", err)
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
		return nil, errors.New("invalid ECIES request")
	}

	clientKey := key.G2.Point()
	if err := clientKey.UnmarshalBinary(msg); err != nil {
		return nil, errors.New("invalid client key")
	}
<<<<<<< HEAD
	randomness, err := entropy.GetRandom(nil, 32)
	if err != nil {
		return nil, fmt.Errorf("error gathering randomness: %s", err)
	} else if len(randomness) != 32 {
		return nil, fmt.Errorf("error gathering randomness: expected 32 bytes, got %d", len(randomness))
=======
	var randomness [32]byte
	if n, err := rand.Read(randomness[:]); err != nil {
		return nil, errors.New("error gathering randomness")
	} else if n != 32 {
		return nil, errors.New("error gathering randomness")
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	}

	obj, err := ecies.Encrypt(key.G2, ecies.DefaultHash, clientKey, randomness[:])
	return &drand.PrivateRandResponse{Response: obj}, err
}
<<<<<<< HEAD

// Home ...
func (d *Drand) Home(c context.Context, in *drand.HomeRequest) (*drand.HomeResponse, error) {
	peer, ok := peer.FromContext(c)
	if ok {
		d.log.With("module", "public").Info("home", peer.Addr.String())
	}
	return &drand.HomeResponse{
		Status: fmt.Sprintf("drand up and running on %s",
			d.priv.Public.Address()),
	}, nil
}

// Group replies with the current group of this drand node in a TOML encoded
// format
func (d *Drand) Group(ctx context.Context, in *drand.GroupRequest) (*drand.GroupResponse, error) {
	d.state.Lock()
	defer d.state.Unlock()
	if d.group == nil {
		return nil, errors.New("drand: no dkg group setup yet")
	}
	gtoml := d.group.TOML().(*key.GroupTOML)
	var resp = new(drand.GroupResponse)
	resp.Nodes = make([]*drand.Node, len(gtoml.Nodes))
	for i, n := range gtoml.Nodes {
		resp.Nodes[i] = &drand.Node{
			Address: n.Address,
			Key:     n.Key,
			TLS:     n.TLS,
		}
	}
	resp.Threshold = uint32(gtoml.Threshold)
	// take the period in second -> ms. grouptoml already transforms it to toml
	ms := uint32(d.group.Period / time.Millisecond)
	resp.Period = ms
	if gtoml.PublicKey != nil {
		resp.Distkey = make([]string, len(gtoml.PublicKey.Coefficients))
		copy(resp.Distkey, gtoml.PublicKey.Coefficients)
	}
	return resp, nil
}
=======
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
