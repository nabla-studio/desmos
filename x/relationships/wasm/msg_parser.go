package wasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/cosmwasm"
	"github.com/desmos-labs/desmos/v5/x/commons"
	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{
		cdc: cdc,
	}
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.RelationshipsMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse x/relationships message from contract %s", contractAddr.String())
	}

	switch {
	case msg.CreateRelationship != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CreateRelationship, &types.MsgCreateRelationship{})
	case msg.DeleteRelationship != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeleteRelationship, &types.MsgDeleteRelationship{})
	case msg.BlockUser != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.BlockUser, &types.MsgBlockUser{})
	case msg.UnblockUser != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.UnblockUser, &types.MsgUnblockUser{})

	default:
		return nil, errors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-relationships-msg-parser: message not supported")
	}
}
