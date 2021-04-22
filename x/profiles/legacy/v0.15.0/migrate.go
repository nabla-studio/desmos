package v0150

import (
	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
)

const (
	ModuleName = "profiles"
)

// Migrate accepts exported genesis state from v0.13.0 and migrates it to v0.15.0
// genesis state.
func Migrate(oldGenState v0130profiles.GenesisState) GenesisState {
	return GenesisState{
		Profiles:             ConvertProfiles(oldGenState.Profiles),
		DTagTransferRequests: ConvertDTagTransferRequest(oldGenState.DTagTransferRequests),
		Params:               ConvertParams(oldGenState.Params),
	}
}

// ------------------------------------------------------------------------------------------------------------------

func ConvertProfiles(oldProfiles []v0130profiles.Profile) []Profile {
	profiles := make([]Profile, len(oldProfiles))

	for index, profile := range oldProfiles {
		profiles[index] = newProfile(profile)
	}

	return profiles
}

// ------------------------------------------------------------------------------------------------------------------

func ConvertDTagTransferRequest(oldDTagTransferRequests []v0130profiles.DTagTransferRequest) []DTagTransferRequest {
	dTagTransferRequests := make([]DTagTransferRequest, len(oldDTagTransferRequests))

	for index, dTagTransferRequest := range oldDTagTransferRequests {
		dTagTransferRequests[index] = DTagTransferRequest{
			DTagToTrade: dTagTransferRequest.DTagToTrade,
			Sender:      dTagTransferRequest.Sender.String(),
			Receiver:    dTagTransferRequest.Receiver.String(),
		}
	}

	return dTagTransferRequests
}

// ------------------------------------------------------------------------------------------------------------------

func ConvertParams(old v0130profiles.Params) Params {
	return Params{
		MonikerParams: MonikerParams{
			MinMonikerLength: old.MonikerParams.MinMonikerLen,
			MaxMonikerLength: old.MonikerParams.MaxMonikerLen,
		},
		DTagParams: DTagParams{
			RegEx:         old.DTagParams.RegEx,
			MinDTagLength: old.DTagParams.MinDTagLen,
			MaxDTagLength: old.DTagParams.MaxDTagLen,
		},
		MaxBioLength: old.MaxBioLen,
	}
}
