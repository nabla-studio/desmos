package v6

import (
	"bytes"
	"encoding/hex"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"

	"github.com/desmos-labs/desmos/v5/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v8 to v9.
// The migration includes:
//
// - fixing the chain links so that their types are correct
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, amino *codec.LegacyAmino) error {
	store := ctx.KVStore(storeKey)

	// Migrate all the chain links
	err := migrateChainLinks(store, cdc, amino)
	if err != nil {
		return err
	}

	return nil
}

// migrateChainLinks migrates the chain links from v8 to v9 by updating their signature value type accordingly.
// During the migration from v6 to v7 there was an error which ended up setting all chain link value types to
// SIGNATURE_VALUE_TYPE_COSMOS_DIRECT. This script fixes that by setting the proper value type based on the plain
// text encode that has been used when creating the signature.
func migrateChainLinks(store sdk.KVStore, cdc codec.BinaryCodec, amino *codec.LegacyAmino) error {
	chainLinksStore := prefix.NewStore(store, types.ChainLinksPrefix)
	iterator := chainLinksStore.Iterator(nil, nil)

	var chainLinks []types.ChainLink
	for ; iterator.Valid(); iterator.Next() {
		var chainLink types.ChainLink
		err := cdc.Unmarshal(iterator.Value(), &chainLink)
		if err != nil {
			return err
		}
		chainLinks = append(chainLinks, chainLink)
	}

	for _, chainLink := range chainLinks {
		// Avoid implicit memory aliasing
		chainLink := chainLink

		var signature types.Signature
		err := cdc.UnpackAny(chainLink.Proof.Signature, &signature)
		if err != nil {
			return err
		}

		value, err := hex.DecodeString(chainLink.Proof.PlainText)
		if err != nil {
			return err
		}

		// Fix the signature
		fixedSig, err := fixSignatureValue(signature, value, cdc, amino)
		if err != nil {
			return err
		}

		// Update the signature Any
		sigAny, err := codectypes.NewAnyWithValue(fixedSig)
		if err != nil {
			return err
		}
		chainLink.Proof.Signature = sigAny

		// Set the link inside the store to update it
		store.Set(
			types.ChainLinksStoreKey(chainLink.User, chainLink.ChainConfig.Name, chainLink.GetAddressData().GetValue()),
			cdc.MustMarshal(&chainLink),
		)
	}

	return nil
}

func fixSignatureValue(signature types.Signature, plainText []byte, cdc codec.BinaryCodec, amino *codec.LegacyAmino) (types.Signature, error) {
	if sig, ok := signature.(*types.SingleSignature); ok {
		return types.NewSingleSignature(getSignatureTypeFromPlainText(plainText, cdc, amino), sig.Signature), nil
	} else if sig, ok := signature.(*types.CosmosMultiSignature); ok {
		// Convert the signatures
		signatures := make([]types.Signature, len(sig.Signatures))
		for i, sigAny := range sig.Signatures {
			var sig types.Signature
			err := cdc.UnpackAny(sigAny, &sig)
			if err != nil {
				return nil, err
			}

			fixedSig, err := fixSignatureValue(sig, plainText, cdc, amino)
			if err != nil {
				return nil, err
			}
			signatures[i] = fixedSig
		}

		// Return the multi sig with the fixed signatures
		return types.NewCosmosMultiSignature(sig.BitArray, signatures), nil
	}

	return nil, fmt.Errorf("invalid signature type: %T", signature)
}

func getSignatureTypeFromPlainText(plainText []byte, cdc codec.BinaryCodec, amino *codec.LegacyAmino) types.SignatureValueType {
	// Check Amino value
	var legacySignDoc legacytx.StdSignDoc
	err := amino.UnmarshalJSON(plainText, &legacySignDoc)
	if err == nil {
		return types.SIGNATURE_VALUE_TYPE_COSMOS_AMINO
	}

	// Check direct value
	var directSignDoc tx.SignDoc
	err = cdc.Unmarshal(plainText, &directSignDoc)

	// Check to make sure the value was a SignDoc. If that's not the case, the two arrays will not match
	if err == nil && bytes.Equal(plainText, cdc.MustMarshal(&directSignDoc)) {
		return types.SIGNATURE_VALUE_TYPE_COSMOS_DIRECT
	}

	return types.SIGNATURE_VALUE_TYPE_RAW
}
