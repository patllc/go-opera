package lachesis

import (
	"github.com/Fantom-foundation/go-lachesis/src/kvdb"
	"testing"
	"time"

	"github.com/Fantom-foundation/go-lachesis/src/crypto"
	"github.com/Fantom-foundation/go-lachesis/src/hash"
	"github.com/Fantom-foundation/go-lachesis/src/logger"
	"github.com/Fantom-foundation/go-lachesis/src/network"
	"github.com/Fantom-foundation/go-lachesis/src/posnode"
)

func TestRing(t *testing.T) {
	logger.SetTestMode(t)

	ll := LachesisNetworkRing(5)

	if !testing.Short() {
		time.Sleep(10 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
	}

	for _, l := range ll {
		cp := l.consensusStore.GetCheckpoint()
		ep := l.consensusStore.GetSuperFrame()
		t.Logf("%s: SFrame %d, Block %d", l.node.Host(), ep.SuperFrameN, cp.LastBlockN)
		l.Stop()
	}
}

func TestStar(t *testing.T) {
	logger.SetTestMode(t)

	ll := LachesisNetworkStar(5)

	if !testing.Short() {
		time.Sleep(10 * time.Second)
	} else {
		time.Sleep(1 * time.Second)
	}

	for _, l := range ll {
		cp := l.consensusStore.GetCheckpoint()
		ep := l.consensusStore.GetSuperFrame()
		t.Logf("%s: SFrame %d, Block %d", l.node.Host(), ep.SuperFrameN, cp.LastBlockN)
		l.Stop()
	}
}

/*
 * Utils:
 */

// NewForTests makes lachesis node with fake network.
// It does not start any process.
func NewForTests(
	newDb func() kvdb.Database,
	host string,
	key *crypto.PrivateKey,
	conf *Config,
) *Lachesis {
	l := makeLachesis(newDb, host, key, conf, network.FakeListener, posnode.FakeClient(host))

	hash.SetNodeName(l.node.ID, host)
	l.node.SetName(host)
	l.nodeStore.SetName(host)
	l.consensus.SetName(host)
	l.consensusStore.SetName(host)

	return l
}
