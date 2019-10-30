[![Build Status](https://travis-ci.org/dedis/drand.svg?branch=master)](https://travis-ci.org/dedis/drand)

# Drand - A Distributed Randomness Beacon Daemon

Drand (pronounced "dee-rand") is a distributed randomness beacon daemon written in [Golang](https://golang.org/). Servers running drand can be linked with each other to produce collective, publicly verifiable, unbiasable, unpredictable random values at fixed intervals using bilinear pairings and threshold cryptography. Drand nodes can also serve locally-generated private randomness to clients.

### Disclaimer

**This software is considered experimental and has NOT received a third-party audit yet. Therefore, DO NOT USE it in production or for anything security critical at this point.**

# Table of Contents
1. [Installation](#Installation)
2. [Overview](#Overview)
3. [Operator Guide](#Operator-Guide)
4. [User Guide](#User-Guide)
5. [DrandJS](#DrandJS)
6. [Cryptography Background](#Cryptography-Background)
7. [Contributors and Acknowledgments](#Contributors)
8. [What’s Next?](#What’s-Next?)
9. [License](#License)
10. [Coverage](#Coverage)

## Installation
Drand can be installed via [Golang](https://golang.org/) or [Docker](https://www.docker.com/). By default, drand saves the configuration files such as the long-term key pair, the group file, and the collective public key in the directory `$HOME/.drand/`.
### Via Golang
Make sure that you have a working [Golang installation](https://golang.org/doc/install) and that your [GOPATH](https://golang.org/doc/code.html#GOPATH) is set.  
Then install drand via:
```bash
go get -u github.com/dedis/drand
cd drand
make install
```
Then you can run the command-line application with `drand`
### Via Docker
The setup is explained in [README_docker.md](README_docker.md).

### TLS setup
Running drand behind a reverse proxy is the **default** method of deploying. We provide in [TLS_setup.md](TLS_setup.md) the minimum setup using [nginx](https://www.nginx.com/) and [certbot](https://certbot.eff.org/lets-encrypt/).

## Overview

### Public Randomness
Generating public randomness is the primary functionality of drand. Public
randomness is generated collectively by drand nodes and publicly available. The
main challenge in generating good randomness is that no party involved in the
randomness generation process should be able to predict or bias the final
output. Additionally, the final result has to be third-party verifiable to make
it actually useful for applications like lotteries, sharding, or parameter
generation in security protocols.  

A drand randomness beacon is composed of a distributed set of nodes and has two
phases:

- **Setup:** Each node first generates a *long-term public/private key pair*.
    Then all of the public keys are written to a *group file* together with some
    further metadata required to operate the beacon. After this group file has
    been distributed, the nodes perform a *distributed key generation* (DKG) protocol
    to create the collective public key and one private key share per server. The
    participants NEVER see/use the actual (distributed) private key explicitly but
    instead utilize their respective private key shares for the generation of public
    randomness.

- **Generation:** After the setup, the nodes switch to the randomness
    generation mode. Any of the nodes can initiate a randomness generation round
    by broadcasting a message which all the other participants sign using a t-of-n
    threshold version of the *Boneh-Lynn-Shacham* (BLS) signature scheme and their
    respective private key shares. Once any node (or third-party observer) has
    gathered t partial signatures, it can reconstruct the full BLS
    signature (using Lagrange interpolation). The signature is then hashed using
    SHA-512 to ensure that there is no bias in the byte representation of the final output.
    This hash corresponds to the collective random value and can be verified against
    the collective public key.

### Private Randomness

Private randomness generation is the secondary functionality of drand. Clients
can request private randomness from some or all of the drand nodes which extract
it locally from their entropy pools and send it back in encrypted form. This
can be useful to gather randomness from different entropy sources, for example
in embedded devices.

In this mode we assume that a client has a private/public key pair and
encapsulates its public key towards the server's public key using the ECIES
encryption scheme. After receiving a request, the drand node produces 32 random
bytes locally (using Go's `crypto/rand` interface), encrypts them using the
received public key and sends it back to the client.

**Note:** Assuming that clients without good local entropy sources (such
as embedded devices) use this process to gather high entropy randomness to
bootstrap their local PRNGs, we emphasize that the initial client key pair has
to be provided by a trusted source (such as the device manufacturer). Otherwise
we run into the chicken-and-egg problem of how to produce on the client's side a
secure ephemeral key pair for ECIES encryption without a good (local) source of
randomness.

## Operator Guide
[drand operator guide](https://hackmd.io/@nikkolasg/Hkz2XFWa4)

## User Guide
[client side API](https://hackmd.io/@nikkolasg/HJ9lg5ZTE)

## Cryptography Background

You can learn more about drand, its motivations and how does it work on these
public [slides](https://docs.google.com/presentation/d/1t2ysit78w0lsySwVbQOyWcSDnYxdOBPzY7K2P9UE1Ac/edit?usp=sharing) and this [document](https://hackmd.io/@nikkolasg/HyUAgm234).

## DrandJS
To facilitate the use of drand's randomness in JavaScript-based applications, we provide [DrandJS](https://github.com/PizzaWhisperer/drandjs). The main method `fetchAndVerify` of this JavaScript library fetches from a drand node the latest random beacon generated and then verifies it against the distributed key. For more details on the procedure and instructions on how to use it, refer to the [readme](https://github.com/PizzaWhisperer/drandjs/blob/master/README.md).

## Contributors
Here's the list of people that contributed to drand:
- Nicolas Gailly ([@nikkolasg1](https://twitter.com/nikkolasg1))
- Philipp Jovanovic ([@daeinar](https://twitter.com/daeinar))
- Mathilde Raynal ([@PizzaWhisperer](https://github.com/PizzaWhisperer))
- Gabbi Fisher ([@gabbifish](https://github.com/gabbifish))
- Linus Gasser ([@ineiti](https://github.com/ineiti))
- Jeff Allen ([@jeffallen](https://github.com/jeffallen))
##### Acknowledgments
Thanks to [@herumi](https://github.com/herumi) for providing support on his optimized pairing-based cryptographic library used in the first version.
Thanks to Apostol Vassilev for its interest in drand and the extensive and helpful discussions on the drand design.
Thanks to [@Bren2010](https://github.com/Bren2010) and [@grittygrease](https://github.com/grittygrease) for providing the native Golang bn256 implementation and for their help in the design of drand and future ideas.

## What's Next?
Although being already functional, drand is still at an early development stage and there is a lot left to be done. The list of opened [issues](https://github.com/dedis/drand/issues) is a good place to start. On top of this, drand would benefit from higher-level enhancements such as the following:

+ Move to the BL12-381 curve
+ Add more unit tests
+ Reduce size of Docker
+ Add a systemd unit file
+ Support multiple drand instances within one node
+ Implement a more [failure-resilient DKG protocol](https://eprint.iacr.org/2012/377.pdf) or an approach based on verifiable succint computations (zk-SNARKs, etc).

Feel free to submit feature requests or, even better, pull requests. ;)

## License
The drand source code is released under MIT license, see the file [LICENSE](https://github.com/dedis/drand/blob/master/LICENSE) for the full text.

## Coverage
- EPFL blog [post](https://actu.epfl.ch/news/epfl-helps-launch-globally-distributed-randomness-/)
- Cloudflare crypto week [introduction post](https://new.blog.cloudflare.com/league-of-entropy/) and the more [technical post](https://new.blog.cloudflare.com/inside-the-entropy/).
- Kudelski Security blog [post](https://research.kudelskisecurity.com/2019/06/17/league-of-entropy/)
- OneZero [post](https://onezero.medium.com/the-league-of-entropy-is-making-randomness-truly-random-522f22ce93ce) on the league of entropy
- SlashDot [post](https://science.slashdot.org/story/19/06/17/1921224/the-league-of-entropy-forms-to-offer-acts-of-public-randomness)
- Duo [post](https://duo.com/decipher/the-league-of-entropy-forms-to-offer-acts-of-public-randomness)
