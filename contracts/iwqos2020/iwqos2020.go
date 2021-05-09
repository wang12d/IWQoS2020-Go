// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iwqos2020

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Iwqos2020ABI is the input ABI used to generate the binding from.
const Iwqos2020ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorization\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cipher\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"name\":\"concat\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"g\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"p\",\"type\":\"bytes\"}],\"name\":\"expmod\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"get_authorize\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_returnC\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fbpie\",\"type\":\"uint256\"}],\"name\":\"get_searchtoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"tok\",\"type\":\"bytes32\"}],\"name\":\"get_taskindex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchfbpie\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"searchtok\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"p\",\"type\":\"bytes\"}],\"name\":\"setP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"token\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"value\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"name\":\"set_taskindex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"authori\",\"type\":\"bytes\"}],\"name\":\"setauthorize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ctoken\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dhash\",\"type\":\"bytes32\"}],\"name\":\"settask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"task_index\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"toBytes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"b\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tok\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fbpie\",\"type\":\"uint256\"}],\"name\":\"try1\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Iwqos2020Bin is the compiled bytecode used for deploying new contracts.
var Iwqos2020Bin = "0x608060405234801561001057600080fd5b50611920806100206000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c80638061d325116100a2578063985407b411610071578063985407b414610868578063b38eacf314610a3d578063cb24ba5814610a5b578063e716dcda14610a9d578063ee4fd78414610afc57610116565b80638061d3251461065157806391327ec614610689578063936000871461070c5780639663deb11461084a57610116565b8063523d2055116100e9578063523d2055146104155780635a75c81114610457578063775a8f5e146104995780637b0e4143146104db5780637b6a6fd8146105a057610116565b80631041f00d1461011b57806311291df81461015357806336b9f5b31461020e5780634fe95b6a146102bf575b600080fd5b6101516004803603604081101561013157600080fd5b810190808035906020019092919080359060200190929190505050610b48565b005b61020c6004803603602081101561016957600080fd5b810190808035906020019064010000000081111561018657600080fd5b82018360208201111561019857600080fd5b803590602001918460018302840111640100000000831117156101ba57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610b64565b005b6102446004803603604081101561022457600080fd5b810190808035906020019092919080359060200190929190505050610b7e565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610284578082015181840152602081019050610269565b50505050905090810190601f1680156102b15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610413600480360360608110156102d557600080fd5b81019080803590602001906401000000008111156102f257600080fd5b82018360208201111561030457600080fd5b8035906020019184602083028401116401000000008311171561032657600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561038657600080fd5b82018360208201111561039857600080fd5b803590602001918460208302840111640100000000831117156103ba57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929080359060200190929190505050610c48565b005b6104416004803603602081101561042b57600080fd5b8101908080359060200190929190505050610cb3565b6040518082815260200191505060405180910390f35b6104836004803603602081101561046d57600080fd5b8101908080359060200190929190505050610cd4565b6040518082815260200191505060405180910390f35b6104c5600480360360208110156104af57600080fd5b8101908080359060200190929190505050610cf1565b6040518082815260200191505060405180910390f35b61059e600480360360408110156104f157600080fd5b81019080803590602001909291908035906020019064010000000081111561051857600080fd5b82018360208201111561052a57600080fd5b8035906020019184600183028401116401000000008311171561054c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610cfe565b005b6105d6600480360360408110156105b657600080fd5b810190808035906020019092919080359060200190929190505050610d4e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106165780820151818401526020810190506105fb565b50505050905090810190601f1680156106435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106876004803603604081101561066757600080fd5b810190808035906020019092919080359060200190929190505050610e14565b005b610691611133565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106d15780820151818401526020810190506106b6565b50505050905090810190601f1680156106fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6107cf6004803603604081101561072257600080fd5b810190808035906020019064010000000081111561073f57600080fd5b82018360208201111561075157600080fd5b8035906020019184600183028401116401000000008311171561077357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291905050506111d1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561080f5780820151818401526020810190506107f4565b50505050905090810190601f16801561083c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61085261124e565b6040518082815260200191505060405180910390f35b6109c26004803603606081101561087e57600080fd5b810190808035906020019064010000000081111561089b57600080fd5b8201836020820111156108ad57600080fd5b803590602001918460018302840111640100000000831117156108cf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803590602001909291908035906020019064010000000081111561093c57600080fd5b82018360208201111561094e57600080fd5b8035906020019184600183028401116401000000008311171561097057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611254565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a025780820151818401526020810190506109e7565b50505050905090810190601f168015610a2f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610a45611572565b6040518082815260200191505060405180910390f35b610a8760048036036020811015610a7157600080fd5b8101908080359060200190929190505050611578565b6040518082815260200191505060405180910390f35b610aa5611590565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610ae8578082015181840152602081019050610acd565b505050509050019250505060405180910390f35b610b3260048036036040811015610b1257600080fd5b8101908080359060200190929190803590602001909291905050506115e8565b6040518082815260200191505060405180910390f35b8060016000848152602001908152602001600020819055505050565b8060059080519060200190610b7a929190611845565b5050565b60606000808481526020019081526020016000208281548110610b9d57fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c3b5780601f10610c1057610100808354040283529160200191610c3b565b820191906000526020600020905b815481529060010190602001808311610c1e57829003601f168201915b5050505050905092915050565b60008090505b81811015610cad576000848281518110610c6457fe5b602002602001015190506000848381518110610c7c57fe5b6020026020010151905080600160008481526020019081526020016000208190555050508080600101915050610c4e565b50505050565b60038181548110610cc057fe5b906000526020600020016000915090505481565b600060016000838152602001908152602001600020549050919050565b60008160001b9050919050565b60008083815260200190815260200160002081908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610d49929190611845565b505050565b60006020528160005260406000208181548110610d6757fe5b90600052602060002001600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e0c5780601f10610de157610100808354040283529160200191610e0c565b820191906000526020600020905b815481529060010190602001808311610def57829003601f168201915b505050505081565b60008090505b600681101561112e5760606000808481526020019081526020016000208281548110610e4257fe5b906000526020600020018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ee05780601f10610eb557610100808354040283529160200191610ee0565b820191906000526020600020905b815481529060010190602001808311610ec357829003601f168201915b505050505090506060610f8e828660058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f845780601f10610f5957610100808354040283529160200191610f84565b820191906000526020600020905b815481529060010190602001808311610f6757829003601f168201915b5050505050611254565b905060008090506060610fa983610fa484610cf1565b6111d1565b90506000816040516020018082805190602001908083835b60208310610fe45780518252602082019150602081019050602083039250610fc1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120905060008060001b90505b8060016000848152602001908152602001600020541461111b576000826001600085815260200190815260200160002054189050600681908060018154018082558091505060019003906000526020600020016000909190919091505560018501945061109f8661109a87610cf1565b6111d1565b9350836040516020018082805190602001908083835b602083106110d857805182526020820191506020810190506020830392506110b5565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012092505061102a565b5050505050508080600101915050610e1a565b505050565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111c95780601f1061119e576101008083540402835291602001916111c9565b820191906000526020600020905b8154815290600101906020018083116111ac57829003601f168201915b505050505081565b606082826040516020018083805190602001908083835b6020831061120b57805182526020820191506020810190506020830392506111e8565b6001836020036101000a03801982511681845116808217855250505050505090500182815260200192505050604051602081830303815290604052905092915050565b60045481565b60606101808251146112ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f756e7175616c6966696564206c656e677468206f66207000000000000000000081525060200191505060405180910390fd5b610180845114611346576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f756e7175616c6966696564206c656e677468206f66206700000000000000000081525060200191505060405180910390fd5b6060845160001b602060001b845160001b878760001b876040516020018087815260200186815260200185815260200184805190602001908083835b602083106113a55780518252602082019150602081019050602083039250611382565b6001836020036101000a03801982511681845116808217855250505050505090500183815260200182805190602001908083835b602083106113fc57805182526020820191506020810190506020830392506113d9565b6001836020036101000a0380198251168184511680821785525050505050509050019650505050505050604051602081830303815290604052905060606101806040519080825280601f01601f19166020018201604052801561146e5781602001600182028036833780820191505090505b50905060606101806040519080825280601f01601f1916602001820160405280156114a85781602001600182028036833780820191505090505b509050610180602082016103806020860160056107d05a03fa6114ca57600080fd5b60008090505b600c8110156115645760008082602002905080602085010151915060008090505b60208110156115545782816020811061150657fe5b1a60f81b868284018151811061151857fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506114f1565b50505080806001019150506114d0565b508193505050509392505050565b60025481565b60016020528060005260406000206000915090505481565b606060068054806020026020016040519081016040528092919081815260200182805480156115de57602002820191906000526020600020905b8154815260200190600101908083116115ca575b5050505050905090565b6000606060008084815260200190815260200160002060008154811061160a57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116a85780601f1061167d576101008083540402835291602001916116a8565b820191906000526020600020905b81548152906001019060200180831161168b57829003601f168201915b5050505050905080600a90805190602001906116c5929190611845565b50606061176d828660058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156117635780601f1061173857610100808354040283529160200191611763565b820191906000526020600020905b81548152906001019060200180831161174657829003601f168201915b5050505050611254565b90508060089080519060200190611785929190611845565b506000809050606061179f8361179a84610cf1565b6111d1565b905080600990805190602001906117b7929190611845565b506000816040516020018082805190602001908083835b602083106117f157805182526020820191506020810190506020830392506117ce565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001209050806007819055506007549550505050505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061188657805160ff19168380011785556118b4565b828001600101855582156118b4579182015b828111156118b3578251825591602001919060010190611898565b5b5090506118c191906118c5565b5090565b6118e791905b808211156118e35760008160009055506001016118cb565b5090565b9056fea2646970667358221220bcbb71c7c277e9d6b60e6252c948e0811e417dd7c244ae88eab6ec9a1f9b666364736f6c63430006030033"

// DeployIwqos2020 deploys a new Ethereum contract, binding an instance of Iwqos2020 to it.
func DeployIwqos2020(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Iwqos2020, error) {
	parsed, err := abi.JSON(strings.NewReader(Iwqos2020ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Iwqos2020Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Iwqos2020{Iwqos2020Caller: Iwqos2020Caller{contract: contract}, Iwqos2020Transactor: Iwqos2020Transactor{contract: contract}, Iwqos2020Filterer: Iwqos2020Filterer{contract: contract}}, nil
}

// Iwqos2020 is an auto generated Go binding around an Ethereum contract.
type Iwqos2020 struct {
	Iwqos2020Caller     // Read-only binding to the contract
	Iwqos2020Transactor // Write-only binding to the contract
	Iwqos2020Filterer   // Log filterer for contract events
}

// Iwqos2020Caller is an auto generated read-only Go binding around an Ethereum contract.
type Iwqos2020Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iwqos2020Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Iwqos2020Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iwqos2020Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Iwqos2020Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Iwqos2020Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Iwqos2020Session struct {
	Contract     *Iwqos2020        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Iwqos2020CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Iwqos2020CallerSession struct {
	Contract *Iwqos2020Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Iwqos2020TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Iwqos2020TransactorSession struct {
	Contract     *Iwqos2020Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Iwqos2020Raw is an auto generated low-level Go binding around an Ethereum contract.
type Iwqos2020Raw struct {
	Contract *Iwqos2020 // Generic contract binding to access the raw methods on
}

// Iwqos2020CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Iwqos2020CallerRaw struct {
	Contract *Iwqos2020Caller // Generic read-only contract binding to access the raw methods on
}

// Iwqos2020TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Iwqos2020TransactorRaw struct {
	Contract *Iwqos2020Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIwqos2020 creates a new instance of Iwqos2020, bound to a specific deployed contract.
func NewIwqos2020(address common.Address, backend bind.ContractBackend) (*Iwqos2020, error) {
	contract, err := bindIwqos2020(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iwqos2020{Iwqos2020Caller: Iwqos2020Caller{contract: contract}, Iwqos2020Transactor: Iwqos2020Transactor{contract: contract}, Iwqos2020Filterer: Iwqos2020Filterer{contract: contract}}, nil
}

// NewIwqos2020Caller creates a new read-only instance of Iwqos2020, bound to a specific deployed contract.
func NewIwqos2020Caller(address common.Address, caller bind.ContractCaller) (*Iwqos2020Caller, error) {
	contract, err := bindIwqos2020(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Iwqos2020Caller{contract: contract}, nil
}

// NewIwqos2020Transactor creates a new write-only instance of Iwqos2020, bound to a specific deployed contract.
func NewIwqos2020Transactor(address common.Address, transactor bind.ContractTransactor) (*Iwqos2020Transactor, error) {
	contract, err := bindIwqos2020(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Iwqos2020Transactor{contract: contract}, nil
}

// NewIwqos2020Filterer creates a new log filterer instance of Iwqos2020, bound to a specific deployed contract.
func NewIwqos2020Filterer(address common.Address, filterer bind.ContractFilterer) (*Iwqos2020Filterer, error) {
	contract, err := bindIwqos2020(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Iwqos2020Filterer{contract: contract}, nil
}

// bindIwqos2020 binds a generic wrapper to an already deployed contract.
func bindIwqos2020(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Iwqos2020ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iwqos2020 *Iwqos2020Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iwqos2020.Contract.Iwqos2020Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iwqos2020 *Iwqos2020Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Iwqos2020Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iwqos2020 *Iwqos2020Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Iwqos2020Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iwqos2020 *Iwqos2020CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iwqos2020.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iwqos2020 *Iwqos2020TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iwqos2020.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iwqos2020 *Iwqos2020TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iwqos2020.Contract.contract.Transact(opts, method, params...)
}

// Authorization is a free data retrieval call binding the contract method 0x7b6a6fd8.
//
// Solidity: function authorization(uint256 , uint256 ) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Caller) Authorization(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "authorization", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Authorization is a free data retrieval call binding the contract method 0x7b6a6fd8.
//
// Solidity: function authorization(uint256 , uint256 ) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Session) Authorization(arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	return _Iwqos2020.Contract.Authorization(&_Iwqos2020.CallOpts, arg0, arg1)
}

// Authorization is a free data retrieval call binding the contract method 0x7b6a6fd8.
//
// Solidity: function authorization(uint256 , uint256 ) view returns(bytes)
func (_Iwqos2020 *Iwqos2020CallerSession) Authorization(arg0 *big.Int, arg1 *big.Int) ([]byte, error) {
	return _Iwqos2020.Contract.Authorization(&_Iwqos2020.CallOpts, arg0, arg1)
}

// Cipher is a free data retrieval call binding the contract method 0x523d2055.
//
// Solidity: function cipher(uint256 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Caller) Cipher(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "cipher", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Cipher is a free data retrieval call binding the contract method 0x523d2055.
//
// Solidity: function cipher(uint256 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Session) Cipher(arg0 *big.Int) ([32]byte, error) {
	return _Iwqos2020.Contract.Cipher(&_Iwqos2020.CallOpts, arg0)
}

// Cipher is a free data retrieval call binding the contract method 0x523d2055.
//
// Solidity: function cipher(uint256 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020CallerSession) Cipher(arg0 *big.Int) ([32]byte, error) {
	return _Iwqos2020.Contract.Cipher(&_Iwqos2020.CallOpts, arg0)
}

// Concat is a free data retrieval call binding the contract method 0x93600087.
//
// Solidity: function concat(bytes a, bytes32 b) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Caller) Concat(opts *bind.CallOpts, a []byte, b [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "concat", a, b)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Concat is a free data retrieval call binding the contract method 0x93600087.
//
// Solidity: function concat(bytes a, bytes32 b) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Session) Concat(a []byte, b [32]byte) ([]byte, error) {
	return _Iwqos2020.Contract.Concat(&_Iwqos2020.CallOpts, a, b)
}

// Concat is a free data retrieval call binding the contract method 0x93600087.
//
// Solidity: function concat(bytes a, bytes32 b) view returns(bytes)
func (_Iwqos2020 *Iwqos2020CallerSession) Concat(a []byte, b [32]byte) ([]byte, error) {
	return _Iwqos2020.Contract.Concat(&_Iwqos2020.CallOpts, a, b)
}

// Expmod is a free data retrieval call binding the contract method 0x985407b4.
//
// Solidity: function expmod(bytes g, uint256 x, bytes p) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Caller) Expmod(opts *bind.CallOpts, g []byte, x *big.Int, p []byte) ([]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "expmod", g, x, p)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Expmod is a free data retrieval call binding the contract method 0x985407b4.
//
// Solidity: function expmod(bytes g, uint256 x, bytes p) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Session) Expmod(g []byte, x *big.Int, p []byte) ([]byte, error) {
	return _Iwqos2020.Contract.Expmod(&_Iwqos2020.CallOpts, g, x, p)
}

// Expmod is a free data retrieval call binding the contract method 0x985407b4.
//
// Solidity: function expmod(bytes g, uint256 x, bytes p) view returns(bytes)
func (_Iwqos2020 *Iwqos2020CallerSession) Expmod(g []byte, x *big.Int, p []byte) ([]byte, error) {
	return _Iwqos2020.Contract.Expmod(&_Iwqos2020.CallOpts, g, x, p)
}

// GetAuthorize is a free data retrieval call binding the contract method 0x36b9f5b3.
//
// Solidity: function get_authorize(uint256 tok, uint256 index) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Caller) GetAuthorize(opts *bind.CallOpts, tok *big.Int, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "get_authorize", tok, index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAuthorize is a free data retrieval call binding the contract method 0x36b9f5b3.
//
// Solidity: function get_authorize(uint256 tok, uint256 index) view returns(bytes)
func (_Iwqos2020 *Iwqos2020Session) GetAuthorize(tok *big.Int, index *big.Int) ([]byte, error) {
	return _Iwqos2020.Contract.GetAuthorize(&_Iwqos2020.CallOpts, tok, index)
}

// GetAuthorize is a free data retrieval call binding the contract method 0x36b9f5b3.
//
// Solidity: function get_authorize(uint256 tok, uint256 index) view returns(bytes)
func (_Iwqos2020 *Iwqos2020CallerSession) GetAuthorize(tok *big.Int, index *big.Int) ([]byte, error) {
	return _Iwqos2020.Contract.GetAuthorize(&_Iwqos2020.CallOpts, tok, index)
}

// GetReturnC is a free data retrieval call binding the contract method 0xe716dcda.
//
// Solidity: function get_returnC() view returns(bytes32[])
func (_Iwqos2020 *Iwqos2020Caller) GetReturnC(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "get_returnC")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetReturnC is a free data retrieval call binding the contract method 0xe716dcda.
//
// Solidity: function get_returnC() view returns(bytes32[])
func (_Iwqos2020 *Iwqos2020Session) GetReturnC() ([][32]byte, error) {
	return _Iwqos2020.Contract.GetReturnC(&_Iwqos2020.CallOpts)
}

// GetReturnC is a free data retrieval call binding the contract method 0xe716dcda.
//
// Solidity: function get_returnC() view returns(bytes32[])
func (_Iwqos2020 *Iwqos2020CallerSession) GetReturnC() ([][32]byte, error) {
	return _Iwqos2020.Contract.GetReturnC(&_Iwqos2020.CallOpts)
}

// GetTaskindex is a free data retrieval call binding the contract method 0x5a75c811.
//
// Solidity: function get_taskindex(bytes32 tok) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Caller) GetTaskindex(opts *bind.CallOpts, tok [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "get_taskindex", tok)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTaskindex is a free data retrieval call binding the contract method 0x5a75c811.
//
// Solidity: function get_taskindex(bytes32 tok) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Session) GetTaskindex(tok [32]byte) ([32]byte, error) {
	return _Iwqos2020.Contract.GetTaskindex(&_Iwqos2020.CallOpts, tok)
}

// GetTaskindex is a free data retrieval call binding the contract method 0x5a75c811.
//
// Solidity: function get_taskindex(bytes32 tok) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020CallerSession) GetTaskindex(tok [32]byte) ([32]byte, error) {
	return _Iwqos2020.Contract.GetTaskindex(&_Iwqos2020.CallOpts, tok)
}

// Pp is a free data retrieval call binding the contract method 0x91327ec6.
//
// Solidity: function pp() view returns(bytes)
func (_Iwqos2020 *Iwqos2020Caller) Pp(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "pp")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Pp is a free data retrieval call binding the contract method 0x91327ec6.
//
// Solidity: function pp() view returns(bytes)
func (_Iwqos2020 *Iwqos2020Session) Pp() ([]byte, error) {
	return _Iwqos2020.Contract.Pp(&_Iwqos2020.CallOpts)
}

// Pp is a free data retrieval call binding the contract method 0x91327ec6.
//
// Solidity: function pp() view returns(bytes)
func (_Iwqos2020 *Iwqos2020CallerSession) Pp() ([]byte, error) {
	return _Iwqos2020.Contract.Pp(&_Iwqos2020.CallOpts)
}

// Searchfbpie is a free data retrieval call binding the contract method 0x9663deb1.
//
// Solidity: function searchfbpie() view returns(uint256)
func (_Iwqos2020 *Iwqos2020Caller) Searchfbpie(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "searchfbpie")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Searchfbpie is a free data retrieval call binding the contract method 0x9663deb1.
//
// Solidity: function searchfbpie() view returns(uint256)
func (_Iwqos2020 *Iwqos2020Session) Searchfbpie() (*big.Int, error) {
	return _Iwqos2020.Contract.Searchfbpie(&_Iwqos2020.CallOpts)
}

// Searchfbpie is a free data retrieval call binding the contract method 0x9663deb1.
//
// Solidity: function searchfbpie() view returns(uint256)
func (_Iwqos2020 *Iwqos2020CallerSession) Searchfbpie() (*big.Int, error) {
	return _Iwqos2020.Contract.Searchfbpie(&_Iwqos2020.CallOpts)
}

// Searchtok is a free data retrieval call binding the contract method 0xb38eacf3.
//
// Solidity: function searchtok() view returns(uint256)
func (_Iwqos2020 *Iwqos2020Caller) Searchtok(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "searchtok")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Searchtok is a free data retrieval call binding the contract method 0xb38eacf3.
//
// Solidity: function searchtok() view returns(uint256)
func (_Iwqos2020 *Iwqos2020Session) Searchtok() (*big.Int, error) {
	return _Iwqos2020.Contract.Searchtok(&_Iwqos2020.CallOpts)
}

// Searchtok is a free data retrieval call binding the contract method 0xb38eacf3.
//
// Solidity: function searchtok() view returns(uint256)
func (_Iwqos2020 *Iwqos2020CallerSession) Searchtok() (*big.Int, error) {
	return _Iwqos2020.Contract.Searchtok(&_Iwqos2020.CallOpts)
}

// TaskIndex is a free data retrieval call binding the contract method 0xcb24ba58.
//
// Solidity: function task_index(bytes32 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Caller) TaskIndex(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "task_index", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TaskIndex is a free data retrieval call binding the contract method 0xcb24ba58.
//
// Solidity: function task_index(bytes32 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020Session) TaskIndex(arg0 [32]byte) ([32]byte, error) {
	return _Iwqos2020.Contract.TaskIndex(&_Iwqos2020.CallOpts, arg0)
}

// TaskIndex is a free data retrieval call binding the contract method 0xcb24ba58.
//
// Solidity: function task_index(bytes32 ) view returns(bytes32)
func (_Iwqos2020 *Iwqos2020CallerSession) TaskIndex(arg0 [32]byte) ([32]byte, error) {
	return _Iwqos2020.Contract.TaskIndex(&_Iwqos2020.CallOpts, arg0)
}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 x) view returns(bytes32 b)
func (_Iwqos2020 *Iwqos2020Caller) ToBytes(opts *bind.CallOpts, x *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Iwqos2020.contract.Call(opts, &out, "toBytes", x)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 x) view returns(bytes32 b)
func (_Iwqos2020 *Iwqos2020Session) ToBytes(x *big.Int) ([32]byte, error) {
	return _Iwqos2020.Contract.ToBytes(&_Iwqos2020.CallOpts, x)
}

// ToBytes is a free data retrieval call binding the contract method 0x775a8f5e.
//
// Solidity: function toBytes(uint256 x) view returns(bytes32 b)
func (_Iwqos2020 *Iwqos2020CallerSession) ToBytes(x *big.Int) ([32]byte, error) {
	return _Iwqos2020.Contract.ToBytes(&_Iwqos2020.CallOpts, x)
}

// GetSearchtoke is a paid mutator transaction binding the contract method 0x8061d325.
//
// Solidity: function get_searchtoke(uint256 tok, uint256 fbpie) returns()
func (_Iwqos2020 *Iwqos2020Transactor) GetSearchtoke(opts *bind.TransactOpts, tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "get_searchtoke", tok, fbpie)
}

// GetSearchtoke is a paid mutator transaction binding the contract method 0x8061d325.
//
// Solidity: function get_searchtoke(uint256 tok, uint256 fbpie) returns()
func (_Iwqos2020 *Iwqos2020Session) GetSearchtoke(tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.GetSearchtoke(&_Iwqos2020.TransactOpts, tok, fbpie)
}

// GetSearchtoke is a paid mutator transaction binding the contract method 0x8061d325.
//
// Solidity: function get_searchtoke(uint256 tok, uint256 fbpie) returns()
func (_Iwqos2020 *Iwqos2020TransactorSession) GetSearchtoke(tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.GetSearchtoke(&_Iwqos2020.TransactOpts, tok, fbpie)
}

// SetP is a paid mutator transaction binding the contract method 0x11291df8.
//
// Solidity: function setP(bytes p) returns()
func (_Iwqos2020 *Iwqos2020Transactor) SetP(opts *bind.TransactOpts, p []byte) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "setP", p)
}

// SetP is a paid mutator transaction binding the contract method 0x11291df8.
//
// Solidity: function setP(bytes p) returns()
func (_Iwqos2020 *Iwqos2020Session) SetP(p []byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.SetP(&_Iwqos2020.TransactOpts, p)
}

// SetP is a paid mutator transaction binding the contract method 0x11291df8.
//
// Solidity: function setP(bytes p) returns()
func (_Iwqos2020 *Iwqos2020TransactorSession) SetP(p []byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.SetP(&_Iwqos2020.TransactOpts, p)
}

// SetTaskindex is a paid mutator transaction binding the contract method 0x4fe95b6a.
//
// Solidity: function set_taskindex(bytes32[] token, bytes32[] value, uint256 len) returns()
func (_Iwqos2020 *Iwqos2020Transactor) SetTaskindex(opts *bind.TransactOpts, token [][32]byte, value [][32]byte, len *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "set_taskindex", token, value, len)
}

// SetTaskindex is a paid mutator transaction binding the contract method 0x4fe95b6a.
//
// Solidity: function set_taskindex(bytes32[] token, bytes32[] value, uint256 len) returns()
func (_Iwqos2020 *Iwqos2020Session) SetTaskindex(token [][32]byte, value [][32]byte, len *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.SetTaskindex(&_Iwqos2020.TransactOpts, token, value, len)
}

// SetTaskindex is a paid mutator transaction binding the contract method 0x4fe95b6a.
//
// Solidity: function set_taskindex(bytes32[] token, bytes32[] value, uint256 len) returns()
func (_Iwqos2020 *Iwqos2020TransactorSession) SetTaskindex(token [][32]byte, value [][32]byte, len *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.SetTaskindex(&_Iwqos2020.TransactOpts, token, value, len)
}

// Setauthorize is a paid mutator transaction binding the contract method 0x7b0e4143.
//
// Solidity: function setauthorize(uint256 tok, bytes authori) returns()
func (_Iwqos2020 *Iwqos2020Transactor) Setauthorize(opts *bind.TransactOpts, tok *big.Int, authori []byte) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "setauthorize", tok, authori)
}

// Setauthorize is a paid mutator transaction binding the contract method 0x7b0e4143.
//
// Solidity: function setauthorize(uint256 tok, bytes authori) returns()
func (_Iwqos2020 *Iwqos2020Session) Setauthorize(tok *big.Int, authori []byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Setauthorize(&_Iwqos2020.TransactOpts, tok, authori)
}

// Setauthorize is a paid mutator transaction binding the contract method 0x7b0e4143.
//
// Solidity: function setauthorize(uint256 tok, bytes authori) returns()
func (_Iwqos2020 *Iwqos2020TransactorSession) Setauthorize(tok *big.Int, authori []byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Setauthorize(&_Iwqos2020.TransactOpts, tok, authori)
}

// Settask is a paid mutator transaction binding the contract method 0x1041f00d.
//
// Solidity: function settask(bytes32 ctoken, bytes32 dhash) returns()
func (_Iwqos2020 *Iwqos2020Transactor) Settask(opts *bind.TransactOpts, ctoken [32]byte, dhash [32]byte) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "settask", ctoken, dhash)
}

// Settask is a paid mutator transaction binding the contract method 0x1041f00d.
//
// Solidity: function settask(bytes32 ctoken, bytes32 dhash) returns()
func (_Iwqos2020 *Iwqos2020Session) Settask(ctoken [32]byte, dhash [32]byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Settask(&_Iwqos2020.TransactOpts, ctoken, dhash)
}

// Settask is a paid mutator transaction binding the contract method 0x1041f00d.
//
// Solidity: function settask(bytes32 ctoken, bytes32 dhash) returns()
func (_Iwqos2020 *Iwqos2020TransactorSession) Settask(ctoken [32]byte, dhash [32]byte) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Settask(&_Iwqos2020.TransactOpts, ctoken, dhash)
}

// Try1 is a paid mutator transaction binding the contract method 0xee4fd784.
//
// Solidity: function try1(uint256 tok, uint256 fbpie) returns(bytes32)
func (_Iwqos2020 *Iwqos2020Transactor) Try1(opts *bind.TransactOpts, tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.contract.Transact(opts, "try1", tok, fbpie)
}

// Try1 is a paid mutator transaction binding the contract method 0xee4fd784.
//
// Solidity: function try1(uint256 tok, uint256 fbpie) returns(bytes32)
func (_Iwqos2020 *Iwqos2020Session) Try1(tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Try1(&_Iwqos2020.TransactOpts, tok, fbpie)
}

// Try1 is a paid mutator transaction binding the contract method 0xee4fd784.
//
// Solidity: function try1(uint256 tok, uint256 fbpie) returns(bytes32)
func (_Iwqos2020 *Iwqos2020TransactorSession) Try1(tok *big.Int, fbpie *big.Int) (*types.Transaction, error) {
	return _Iwqos2020.Contract.Try1(&_Iwqos2020.TransactOpts, tok, fbpie)
}
