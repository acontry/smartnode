package minipool

import (
	"github.com/urfave/cli"

	"github.com/rocket-pool/smartnode/shared/utils/api"
	cliutils "github.com/rocket-pool/smartnode/shared/utils/cli"
)

// Register subcommands
func RegisterSubcommands(command *cli.Command, name string, aliases []string) {
    command.Subcommands = append(command.Subcommands, cli.Command{
        Name:      name,
        Aliases:   aliases,
        Usage:     "Manage the node's minipools",
        Subcommands: []cli.Command{

            cli.Command{
                Name:      "status",
                Aliases:   []string{"s"},
                Usage:     "Get a list of the node's minipools",
                UsageText: "rocketpool api minipool status",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 0); err != nil { return err }

                    // Run
                    api.PrintResponse(getStatus(c))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-refund",
                Usage:     "Check whether the node can refund ETH from the minipool",
                UsageText: "rocketpool api minipool can-refund minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canRefundMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "refund",
                Aliases:   []string{"r"},
                Usage:     "Refund ETH belonging to the node from a minipool",
                UsageText: "rocketpool api minipool refund minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(refundMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-dissolve",
                Usage:     "Check whether the minipool can be dissolved",
                UsageText: "rocketpool api minipool can-dissolve minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canDissolveMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "dissolve",
                Aliases:   []string{"d"},
                Usage:     "Dissolve an initialized or prelaunch minipool",
                UsageText: "rocketpool api minipool dissolve minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(dissolveMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-exit",
                Usage:     "Check whether the minipool can be exited from the beacon chain",
                UsageText: "rocketpool api minipool can-exit minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canExitMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "exit",
                Aliases:   []string{"e"},
                Usage:     "Exit a staking minipool from the beacon chain",
                UsageText: "rocketpool api minipool exit minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(exitMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-close",
                Usage:     "Check whether the minipool can be closed",
                UsageText: "rocketpool api minipool can-close minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canCloseMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "close",
                Aliases:   []string{"c"},
                Usage:     "Withdraw balance from a dissolved minipool and close it",
                UsageText: "rocketpool api minipool close minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(closeMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-destroy",
                Usage:     "Check whether the minipool can be destroyed",
                UsageText: "rocketpool api minipool can-destroy minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canDestroyMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "destroy",
                Aliases:   []string{"t"},
                Usage:     "Destroy a minipool after it has been withdrawn from, returning its RPL stake",
                UsageText: "rocketpool api minipool destroy minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(destroyMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-delegate-upgrade",
                Usage:     "Check whether the minipool delegate can be upgraded",
                UsageText: "rocketpool api minipool can-delegate-upgrade minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canDelegateUpgrade(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "delegate-upgrade",
                Usage:     "Upgrade this minipool to the latest network delegate contract",
                UsageText: "rocketpool api minipool delegate-upgrade minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(delegateUpgrade(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-delegate-rollback",
                Usage:     "Check whether the minipool delegate can be rolled back",
                UsageText: "rocketpool api minipool can-delegate-rollback minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canDelegateRollback(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "delegate-rollback",
                Usage:     "Rollback the minipool to the previous delegate contract",
                UsageText: "rocketpool api minipool delegate-rollback minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(delegateRollback(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-set-use-latest-delegate",
                Usage:     "Check whether the automatic upgrading setting can be toggled for the minipool",
                UsageText: "rocketpool api minipool can-set-use-latest-delegate minipool-address setting",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 2); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }
                    setting, err := cliutils.ValidateBool("setting", c.Args().Get(1))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canSetUseLatestDelegate(c, minipoolAddress, setting))
                    return nil

                },
            },
            cli.Command{
                Name:      "set-use-latest-delegate",
                Usage:     "Toggle automatic upgrading of minipool delegates to the latest version",
                UsageText: "rocketpool api minipool set-use-latest-delegate minipool-address setting",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 2); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }
                    setting, err := cliutils.ValidateBool("setting", c.Args().Get(1))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(setUseLatestDelegate(c, minipoolAddress, setting))
                    return nil

                },
            },

            cli.Command{
                Name:      "get-use-latest-delegate",
                Usage:     "Gets the current setting of the automatic minipool delegate upgrade toggle",
                UsageText: "rocketpool api minipool get-use-latest-delegate minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(getUseLatestDelegate(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "get-delegate",
                Usage:     "Gets the address of the current delegate contract used by the minipool",
                UsageText: "rocketpool api minipool get-delegate minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(getDelegate(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "get-previous-delegate",
                Usage:     "Gets the address of the previous delegate contract that the minipool will use during a rollback",
                UsageText: "rocketpool api minipool get-previous-delegate minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(getPreviousDelegate(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "get-effective-delegate",
                Usage:     "Gets the address of the effective delegate contract used by the minipool, which takes the UseLatestDelegate setting into account",
                UsageText: "rocketpool api minipool get-effective-delegate minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(getEffectiveDelegate(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-process-withdrawal",
                Usage:     "Check if a withdrawal can be processed on the minipool",
                UsageText: "rocketpool api minipool can-process-withdrawal minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canProcessWithdrawalMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "process-withdrawal",
                Usage:     "Process a withdrawal on the minipool, distributing ETH to the node operator and the staking pool",
                UsageText: "rocketpool api minipool process-withdrawal minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(processWithdrawalMinipool(c, minipoolAddress))
                    return nil

                },
            },

            cli.Command{
                Name:      "can-process-withdrawal-and-destroy",
                Usage:     "Check if a withdrawal and destroy can be processed on the minipool",
                UsageText: "rocketpool api minipool can-process-withdrawal-and-destroy minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(canProcessWithdrawalAndDestroyMinipool(c, minipoolAddress))
                    return nil

                },
            },
            cli.Command{
                Name:      "process-withdrawal-and-destroy",
                Usage:     "Process a withdrawal on the minipool, distributing ETH to the node operator and the staking pool, then destroy it",
                UsageText: "rocketpool api minipool process-withdrawal-and-destroy minipool-address",
                Action: func(c *cli.Context) error {

                    // Validate args
                    if err := cliutils.ValidateArgCount(c, 1); err != nil { return err }
                    minipoolAddress, err := cliutils.ValidateAddress("minipool address", c.Args().Get(0))
                    if err != nil { return err }

                    // Run
                    api.PrintResponse(processWithdrawalAndDestroyMinipool(c, minipoolAddress))
                    return nil

                },
            },

        },
    })
}

