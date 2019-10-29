package minipool

import (
    "bytes"
    "context"
    "errors"
    "math/big"

    "github.com/ethereum/go-ethereum/common"

    "github.com/rocket-pool/smartnode/shared/services"
    "github.com/rocket-pool/smartnode/shared/services/rocketpool/minipool"
    "github.com/rocket-pool/smartnode/shared/utils/eth"
)


// RocketMinipool NodeWithdrawal event
type NodeWithdrawal struct {
    To common.Address
    EtherAmount *big.Int
    RethAmount *big.Int
    RplAmount *big.Int
    Created *big.Int
}


// Minipool withdraw response type
type MinipoolWithdrawResponse struct {

    // Status
    Success bool                `json:"success"`

    // Withdrawal info
    EtherWithdrawnWei *big.Int  `json:"etherWithdrawnWei"`
    RethWithdrawnWei *big.Int   `json:"rethWithdrawnWei"`
    RplWithdrawnWei *big.Int    `json:"rplWithdrawnWei"`

    // Failure info
    MinipoolDidNotExist bool    `json:"minipoolDidNotExist"`
    WithdrawalsDisabled bool    `json:"withdrawalsDisabled"`
    InvalidNodeOwner bool       `json:"invalidNodeOwner"`
    NodeOwner common.Address    `json:"nodeOwner"`
    InvalidStatus bool          `json:"invalidStatus"`
    Status uint8                `json:"status"`
    NodeDepositDidNotExist bool `json:"nodeDepositDidNotExist"`

}


// Withdraw node deposit from minipool
func WithdrawMinipool(p *services.Provider, minipoolAddress common.Address) (*MinipoolWithdrawResponse, error) {

    // Response
    response := &MinipoolWithdrawResponse{}

    // Get node account
    nodeAccount, _ := p.AM.GetNodeAccount()

    // Check contract code at minipool address
    if code, err := p.Client.CodeAt(context.Background(), minipoolAddress, nil); err != nil {
        return nil, errors.New("Error retrieving contract code at minipool address: " + err.Error())
    } else {
        response.MinipoolDidNotExist = (len(code) == 0)
    }

    // Check minipool exists
    if response.MinipoolDidNotExist {
        return response, nil
    }

    // Initialise minipool contract
    minipoolContract, err := p.CM.NewContract(&minipoolAddress, "rocketMinipool")
    if err != nil {
        return nil, errors.New("Error initialising minipool contract: " + err.Error())
    }

    // Status channels
    withdrawalsDisabledChannel := make(chan bool)
    nodeOwnerChannel := make(chan common.Address)
    statusChannel := make(chan uint8)
    depositNotExistsChannel := make(chan bool)
    errorChannel := make(chan error)

    // Check withdrawals are allowed
    go (func() {
        withdrawalsAllowed := new(bool)
        if err := p.CM.Contracts["rocketNodeSettings"].Call(nil, withdrawalsAllowed, "getWithdrawalAllowed"); err != nil {
            errorChannel <- errors.New("Error checking node withdrawals enabled status: " + err.Error())
        } else {
            withdrawalsDisabledChannel <- !*withdrawalsAllowed
        }
    })()

    // Get minipool node owner
    go (func() {
        nodeOwner := new(common.Address)
        if err := minipoolContract.Call(nil, nodeOwner, "getNodeOwner"); err != nil {
           errorChannel <- errors.New("Error retrieving minipool node owner: " + err.Error())
        } else {
            nodeOwnerChannel <- *nodeOwner
        }
    })()

    // Get minipool status
    go (func() {
        status := new(uint8)
        if err := minipoolContract.Call(nil, status, "getStatus"); err != nil {
            errorChannel <- errors.New("Error retrieving minipool status: " + err.Error())
        } else {
            statusChannel <- *status
        }
    })()

    // Get node deposit status
    go (func() {
        nodeDepositExists := new(bool)
        if err := minipoolContract.Call(nil, nodeDepositExists, "getNodeDepositExists"); err != nil {
            errorChannel <- errors.New("Error retrieving minipool node deposit status: " + err.Error())
        } else {
            depositNotExistsChannel <- !*nodeDepositExists
        }
    })()

    // Receive status
    for received := 0; received < 4; {
        select {
            case response.WithdrawalsDisabled = <-withdrawalsDisabledChannel:
                received++
            case response.NodeOwner = <-nodeOwnerChannel:
                received++
            case response.Status = <-statusChannel:
                received++
            case response.NodeDepositDidNotExist = <-depositNotExistsChannel:
                received++
            case err := <-errorChannel:
                return nil, err
        }
    }

    // Update response
    response.InvalidNodeOwner = !bytes.Equal(response.NodeOwner.Bytes(), nodeAccount.Address.Bytes())
    response.InvalidStatus = !(response.Status == minipool.INITIALIZED || response.Status == minipool.WITHDRAWN || response.Status == minipool.TIMED_OUT)

    // Check minipool status
    if response.WithdrawalsDisabled || response.InvalidNodeOwner || response.InvalidStatus || response.NodeDepositDidNotExist {
        return response, nil
    }

    // Send withdrawal transaction
    txor, err := p.AM.GetNodeAccountTransactor()
    if err != nil { return nil, err }
    txReceipt, err := eth.ExecuteContractTransaction(p.Client, txor, p.NodeContractAddress, p.CM.Abis["rocketNodeContract"], "withdrawMinipoolDeposit", minipoolAddress)
    if err != nil {
        return nil, errors.New("Error withdrawing deposit: " + err.Error())
    } else {
        response.Success = true
    }

    // Get withdrawal event
    if nodeWithdrawalEvents, err := eth.GetTransactionEvents(p.Client, txReceipt, &minipoolAddress, p.CM.Abis["rocketMinipoolDelegateNode"], "NodeWithdrawal", NodeWithdrawal{}); err != nil {
        return nil, errors.New("Error retrieving node deposit withdrawal event: " + err.Error())
    } else if len(nodeWithdrawalEvents) == 0 {
        return nil, errors.New("Could not retrieve node deposit withdrawal event")
    } else {
        nodeWithdrawalEvent := (nodeWithdrawalEvents[0]).(*NodeWithdrawal)
        response.EtherWithdrawnWei = nodeWithdrawalEvent.EtherAmount
        response.RethWithdrawnWei = nodeWithdrawalEvent.RethAmount
        response.RplWithdrawnWei = nodeWithdrawalEvent.RplAmount
    }

    // Return response
    return response, nil

}

