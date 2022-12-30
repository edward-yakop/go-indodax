package indodax

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GetPairs(t *testing.T) {
	cl, err := NewClient("", "")
	require.NoError(t, err, "NewClient")

	gotPairs, err := cl.GetPairs(context.Background())
	require.NoError(t, err, "GetPairs")
	require.NotNil(t, gotPairs)
	assert.NotEmpty(t, gotPairs.Pairs, "Pairs.Entries")

	btcIdr, ok := gotPairs.Pairs["BTCIDR"]
	require.True(t, ok, "Pairs[BTCIDR]")
	assert.Equal(t, "btcidr", btcIdr.Id, "pair.Id")
	assert.Equal(t, "idr", btcIdr.BaseCurrency, "pair.BaseCurrency")
	assert.Equal(t, "btc", btcIdr.TradedCurrency, "pair.TradedCurrency")
}
