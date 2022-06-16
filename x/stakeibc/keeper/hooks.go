package keeper

import (
	"fmt"
	"strconv"

	epochstypes "github.com/Stride-Labs/stride/x/epochs/types"
	"github.com/Stride-Labs/stride/x/stakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	// every epoch
	k.Logger(ctx).Info(fmt.Sprintf("Handling epoch start %s %d", epochIdentifier, epochNumber))
	if epochIdentifier == "stride_epoch" {
		k.Logger(ctx).Info(fmt.Sprintf("Stride Epoch %d", epochNumber))
		depositInterval := int64(k.GetParam(ctx, types.KeyDepositInterval))
		// NOTE: delegateInterval was collapsed into deposit interval - is this ok?
		if epochNumber%depositInterval == 0 {
			// track the current deposit epoch number in stakeibc
			if epochNumber < 0 {
				panic("epochs must be positive")
			}
			epochTracker := types.EpochTracker{
				EpochIdentifier: epochIdentifier,
				EpochNumber: uint64(epochNumber),
			}
			// deposit records *must* exist for this epoch
			k.SetEpochTracker(ctx, epochTracker)
			// TODO TEST-72 move this function to the keeper
			k.Logger(ctx).Info("Triggering deposits")

			// Create one new deposit record / host zone for the next epoch
			createDepositRecords := func(index int64, zoneInfo types.HostZone) (stop bool) {
				// create a deposit record / host zone
				depositRecord := types.NewDepositRecord(0, zoneInfo.HostDenom, zoneInfo.ChainId, types.DepositRecord_TRANSFER, uint64(epochNumber))
				k.AppendDepositRecord(ctx, *depositRecord)
				return false
			}
			// Iterate the zones and apply icaReinvest
			k.IterateHostZones(ctx, createDepositRecords)

			// process previous deposit records
			depositRecords := k.GetAllDepositRecord(ctx)
			addr := k.accountKeeper.GetModuleAccount(ctx, types.ModuleName).GetAddress().String()
			for _, depositRecord := range depositRecords {
				if depositRecord.EpochNumber < uint64(epochNumber) {
					pstr := fmt.Sprintf("\tProcessing deposit {%d} {%s} {%d} {%s}", depositRecord.Id, depositRecord.Denom, depositRecord.Amount)
					k.Logger(ctx).Info(pstr)
					hostZone, hostZoneFound := k.GetHostZone(ctx, depositRecord.HostZoneId)
					if !hostZoneFound {
						k.Logger(ctx).Error("Host zone not found for deposit record {%d}", depositRecord.Id)
						continue
					}
					delegateAccount := hostZone.GetDelegationAccount()
					if delegateAccount == nil || delegateAccount.Address == "" {
						k.Logger(ctx).Error("Zone %s is missing a delegation address!", hostZone.ChainId)
						continue
					}
					delegateAddress := delegateAccount.Address
					// PROCESS TRANSFERS
					// TODO(TEST-89): Set NewHeight relative to the most recent known gaia height (based on the LC)
					// TODO(TEST-90): why do we have two gaia LCs?
					if depositRecord.Status == types.DepositRecord_TRANSFER {
						timeoutHeight := clienttypes.NewHeight(0, 1000000)
						transferCoin := sdk.NewCoin(hostZone.GetIBCDenom(), sdk.NewInt(int64(depositRecord.Amount)))
						goCtx := sdk.WrapSDKContext(ctx)
		
						msg := ibctypes.NewMsgTransfer("transfer", hostZone.TransferChannelId, transferCoin, addr, delegateAddress, timeoutHeight, 0)
						_, err := k.transferKeeper.Transfer(goCtx, msg)
						if err != nil {
							pstr := fmt.Sprintf("\tERROR WITH DEPOSIT RECEIPT {%d}", depositRecord.Id)
							k.Logger(ctx).Info(pstr)
							panic(err)
						}
					// PROCESS STAKING
					// TODO: put this in a function in delegation.go!
					// in general, put this logic in functions
					} else if depositRecord.Status == types.DepositRecord_STAKE {
						k.Logger(ctx).Info(fmt.Sprintf("\tdelegation staking callback on %s", hostZone.HostDenom))
						processAmount := strconv.FormatInt(depositRecord.Amount, 10) + hostZone.HostDenom
						amt, err := sdk.ParseCoinNormalized(processAmount)
						if err != nil {
							k.Logger(ctx).Error(fmt.Sprintf("Could not process coin %s: %s", hostZone.HostDenom, err))
							return
						}
						err = k.DelegateOnHost(ctx, hostZone, amt)
						if err != nil {
							k.Logger(ctx).Error(fmt.Sprintf("Did not stake %s on %s", processAmount, hostZone.ChainId))
							return
						} else {
							k.Logger(ctx).Info(fmt.Sprintf("Successfully staked %s on %s", processAmount, hostZone.ChainId))
						}

						ctx.EventManager().EmitEvents(sdk.Events{
							sdk.NewEvent(
								sdk.EventTypeMessage,
								sdk.NewAttribute("hostZone", hostZone.ChainId),
								sdk.NewAttribute("newAmountStaked", strconv.FormatInt(depositRecord.Amount, 10)),
							),
						})
					}
				}
			}
		}

		// process withdrawals
		// TODO(TEST-88): restructure this to be more efficient, we should only have to loop
		// over host zones once
		reinvestInterval := int64(k.GetParam(ctx, types.KeyReinvestInterval))
		if epochNumber%reinvestInterval == 0 {
			icaReinvest := func(index int64, zoneInfo types.HostZone) (stop bool) {
				// Verify the delegation ICA is registered
				delegationIca := zoneInfo.GetDelegationAccount()
				if delegationIca == nil || delegationIca.Address == "" {
					k.Logger(ctx).Error("Zone %s is missing a delegation address!", zoneInfo.ChainId)
					return true
				}
				withdrawIca := zoneInfo.GetWithdrawalAccount()
				if withdrawIca == nil || withdrawIca.Address == "" {
					k.Logger(ctx).Error("Zone %s is missing a withdrawal address!", zoneInfo.ChainId)
					return true
				}
				// err := k.ReinvestRewards(ctx, zoneInfo)
				// if err != nil {
				// 	k.Logger(ctx).Error("Did not withdraw rewards on %s", zoneInfo.ChainId)
				// 	return true
				// } else {
				// 	k.Logger(ctx).Info(fmt.Sprintf("Successfully withdrew rewards on %s", zoneInfo.ChainId))
				// }
				return false
			}

			// Iterate the zones and apply icaReinvest
			k.IterateHostZones(ctx, icaReinvest)
		}
	}
}

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	// every epoch
	k.Logger(ctx).Info(fmt.Sprintf("Handling epoch end %s %d", epochIdentifier, epochNumber))

}

// Hooks wrapper struct for incentives keeper
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
	h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
