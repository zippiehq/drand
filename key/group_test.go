package key

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

<<<<<<< HEAD
	"github.com/stretchr/testify/require"
	kyber "go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/util/random"
=======
	kyber "github.com/dedis/kyber"
	"github.com/dedis/kyber/util/random"
	"github.com/stretchr/testify/require"
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
)

func TestGroupSaveLoad(t *testing.T) {
	n := 3
	ids := make([]*Identity, n)
	dpub := make([]kyber.Point, n)
	for i := 0; i < n; i++ {
		ids[i] = &Identity{
			Key:  G2.Point().Mul(G2.Scalar().Pick(random.New()), nil),
			Addr: "--",
		}
		dpub[i] = ids[i].Key
	}

	group := LoadGroup(ids, &DistPublic{dpub}, DefaultThreshold(n))
	group.Period = time.Second * 4

	gtoml := group.TOML().(*GroupTOML)
	require.NotNil(t, gtoml.PublicKey)

	// faking distributed public key coefficients
	groupFile, err := ioutil.TempFile("", "group.toml")
	require.NoError(t, err)
	groupPath := groupFile.Name()
	groupFile.Close()
	defer os.RemoveAll(groupPath)

	require.NoError(t, Save(groupPath, group, false))

	loaded := &Group{}
	require.NoError(t, Load(groupPath, loaded))

<<<<<<< HEAD
	require.Equal(t, len(loaded.Nodes), len(group.Nodes))
	require.Equal(t, loaded.Threshold, group.Threshold)
	require.True(t, loaded.PublicKey.Equal(group.PublicKey))
=======
	require.Equal(t, loaded.Nodes, group.Nodes)
	require.Equal(t, loaded.Threshold, group.Threshold)
	require.Equal(t, loaded.PublicKey, group.PublicKey)
>>>>>>> 246580c89478d335ddfbe1c84b8e3afc01153128
	require.Equal(t, loaded.Period, group.Period)
}
