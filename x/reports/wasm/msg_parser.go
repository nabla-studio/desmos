package wasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/cosmwasm"
	"github.com/desmos-labs/desmos/v5/x/commons"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

var _ cosmwasm.MsgParserInterface = MsgsParser{}

type MsgsParser struct {
	cdc codec.Codec
}

func NewWasmMsgParser(cdc codec.Codec) MsgsParser {
	return MsgsParser{cdc: cdc}
}

func (parser MsgsParser) ParseCustomMsgs(contractAddr sdk.AccAddress, data json.RawMessage) ([]sdk.Msg, error) {
	var msg types.ReportsMsg
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse x/reports message from contract %s", contractAddr.String())
	}
	switch {
	case msg.CreateReport != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.CreateReport, &types.MsgCreateReport{})
	case msg.DeleteReport != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.DeleteReport, &types.MsgDeleteReport{})
	case msg.SupportStandardReason != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.SupportStandardReason, &types.MsgSupportStandardReason{})
	case msg.AddReason != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.AddReason, &types.MsgAddReason{})
	case msg.RemoveReason != nil:
		return commons.HandleWasmMsg(parser.cdc, *msg.RemoveReason, &types.MsgRemoveReason{})
	default:
		return nil, errors.Wrap(wasm.ErrInvalidMsg, "cosmwasm-reports-msg-parser: message not supported")
	}
}
