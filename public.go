package main

import (
<<<<<<< HEAD
	"github.com/dedis/drand/core"
	"github.com/dedis/drand/net"
	"github.com/dedis/drand/protobuf/drand"
=======
	"encoding/hex"

	"github.com/dedis/drand/core"
	"github.com/dedis/drand/key"
	"github.com/dedis/drand/net"
	"github.com/dedis/drand/protobuf/drand"
	"github.com/dedis/kyber"
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	"github.com/nikkolasg/slog"
	"github.com/urfave/cli"
)

func getPrivateCmd(c *cli.Context) error {
	if !c.Args().Present() {
		slog.Fatal("Get private takes a group file as argument.")
	}
	defaultManager := net.NewCertManager()
	if c.IsSet("tls-cert") {
		defaultManager.Add(c.String("tls-cert"))
	}
	ids := getNodes(c)
	client := core.NewGrpcClientFromCert(defaultManager)
	var resp []byte
	var err error
	for _, public := range ids {
		resp, err = client.Private(public)
		if err == nil {
			slog.Infof("drand: successfully retrieved private randomness "+
				"from %s", public.Addr)
			break
		}
		slog.Infof("drand: error contacting node %s: %s", public.Addr, err)
	}
	if resp == nil {
		slog.Fatalf("drand: zero successful contacts with nodes")
	}

	type private struct {
<<<<<<< HEAD
		Randomness []byte
	}

	printJSON(&private{resp})
	return nil
}

func getPublicRandomness(c *cli.Context) error {
=======
		Randomness string
	}

	printJSON(&private{hex.EncodeToString(resp)})
	return nil
}

func getPublicCmd(c *cli.Context) error {
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	if !c.Args().Present() {
		slog.Fatal("Get public command takes a group file as argument.")
	}
	defaultManager := net.NewCertManager()
	if c.IsSet("tls-cert") {
		defaultManager.Add(c.String("tls-cert"))
	}

	ids := getNodes(c)
	group := getGroup(c)
	if group.PublicKey == nil {
		slog.Fatalf("drand: group file must contain the distributed public key!")
	}

	public := group.PublicKey
	client := core.NewGrpcClientFromCert(defaultManager)
	isTLS := !c.Bool("tls-disable")
	var resp *drand.PublicRandResponse
	var err error
	for _, id := range ids {
		if c.IsSet("round") {
			resp, err = client.Public(id.Addr, public, isTLS, c.Int("round"))
		} else {
			resp, err = client.LastPublic(id.Addr, public, isTLS)
		}
		if err == nil {
			slog.Infof("drand: public randomness retrieved from %s", id.Addr)
			break
		}
		slog.Printf("drand: could not get public randomness from %s: %s", id.Addr, err)
	}
<<<<<<< HEAD

	printJSON(resp)
=======
	type publicRand struct {
		Round      uint64
		Previous   string
		Randomness string
	}
	s := &publicRand{
		Round:      resp.Round,
		Previous:   hex.EncodeToString(resp.Previous),
		Randomness: hex.EncodeToString(resp.Randomness),
	}

	printJSON(s)
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	return nil
}

func getCokeyCmd(c *cli.Context) error {
	defaultManager := net.NewCertManager()
	if c.IsSet("tls-cert") {
		defaultManager.Add(c.String("tls-cert"))
	}
	ids := getNodes(c)
	client := core.NewGrpcClientFromCert(defaultManager)
<<<<<<< HEAD
	var dkey *drand.DistKeyResponse
=======
	var dkey kyber.Point
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	var err error
	for _, id := range ids {
		dkey, err = client.DistKey(id.Addr, !c.Bool("tls-disable"))
		if err == nil {
			break
		}
		slog.Printf("drand: error fetching distributed key from %s : %s",
			id.Addr, err)
	}
	if dkey == nil {
		slog.Fatalf("drand: can't retrieve dist. key from all nodes")
	}
<<<<<<< HEAD
	printJSON(dkey)
=======
	str := key.PointToString(dkey)
	type distkey struct {
		CollectiveKey string
	}
	printJSON(&distkey{str})
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	return nil
}
