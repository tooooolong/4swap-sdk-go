package fswap

import (
	"context"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadAsset(t *testing.T) {
	const (
		assetID = "4d8c508b-91c5-375b-92b0-ee702ed2dac5"
	)

	ctx := context.Background()
	asset, err := ReadAsset(ctx, assetID)
	require.Nil(t, err, "request should be ok")
	require.True(t, asset.Price.GreaterThan(decimal.Zero), "price should be greater than zero")
	require.True(t, len(asset.Chain.Symbol) > 0, "require chain symbol exists")

	t.Log(asset.Symbol, asset.Price)
}

func TestListAssets(t *testing.T) {
	ctx := context.Background()
	assets, err := ListAssets(ctx)
	require.Nil(t, err, "request should be ok")
	assert.NotEmpty(t, assets, "assets should not be empty")
}
