// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jobcreator

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// JobcreatorMetaData contains all meta data concerning the Jobcreator contract.
var JobcreatorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"calling_contract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"}],\"name\":\"JobAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"JobForfeited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"forfeit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextJobID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"module\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"inputs\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"}],\"name\":\"runJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"name\":\"setRequiredDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dataId\",\"type\":\"string\"}],\"name\":\"submitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff021916908315150217905550348015610029575f80fd5b5061004661003b61004b60201b60201c565b61005260201b60201c565b610113565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611f4b806101205f395ff3fe608060405234801561000f575f80fd5b50600436106100fe575f3560e01c8063a470295811610095578063d2a715c011610064578063d2a715c01461024a578063f2fde38b14610268578063f3d3d44814610284578063fb7cfdd7146102a0576100fe565b8063a4702958146101d6578063c4d66de8146101e0578063c57380a2146101fc578063c75555fa1461021a576100fe565b806358e56db4116100d157806358e56db4146101745780636c0f1f5814610192578063715018a6146101ae5780638da5cb5b146101b8576100fe565b806310fe9ae81461010257806326a4e8d21461012057806336aef9671461013c5780634c526c7614610158575b5f80fd5b61010a6102be565b6040516101179190611008565b60405180910390f35b61013a6004803603810190610135919061105c565b6102e6565b005b610156600480360381019061015191906111f6565b610400565b005b610172600480360381019061016d919061127e565b610530565b005b61017c610585565b60405161018991906112b8565b60405180910390f35b6101ac60048036038101906101a791906111f6565b61058e565b005b6101b6610683565b005b6101c0610696565b6040516101cd9190611008565b60405180910390f35b6101de6106bd565b005b6101fa60048036038101906101f5919061105c565b6106e1565b005b610204610829565b6040516102119190611008565b60405180910390f35b610234600480360381019061022f91906113b3565b610851565b60405161024191906112b8565b60405180910390f35b610252610acc565b60405161025f91906112b8565b60405180910390f35b610282600480360381019061027d919061105c565b610ad2565b005b61029e6004803603810190610299919061105c565b610b54565b005b6102a8610c5c565b6040516102b591906112b8565b60405180910390f35b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102ee610c62565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361035c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035390611495565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610408610ce0565b505f60065f8581526020019081526020015f2090505f815f015403610462576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610459906114fd565b60405180910390fd5b806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166336aef9678585856040518463ffffffff1660e01b81526004016104c293929190611585565b5f604051808303815f87803b1580156104d9575f80fd5b505af11580156104eb573d5f803e3d5ffd5b505050507f365cc3878e524c2da89b4cad993fe1d16f0aa1698b90ecc6ac111c47732b72e384848460405161052293929190611585565b60405180910390a150505050565b610538610ce0565b505f811161057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057290611612565b60405180910390fd5b8060048190555050565b5f600454905090565b610596610ce0565b505f60065f8581526020019081526020015f2090505f815f0154036105f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e7906114fd565b60405180910390fd5b806001015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636c0f1f588585856040518463ffffffff1660e01b815260040161065093929190611585565b5f604051808303815f87803b158015610667575f80fd5b505af1158015610679573d5f803e3d5ffd5b5050505050505050565b61068b610c62565b6106945f610e0d565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6106c5610c62565b5f600160146101000a81548160ff021916908315150217905550565b5f600160169054906101000a900460ff16159050808015610713575060018060159054906101000a900460ff1660ff16105b80610741575061072230610ece565b158015610740575060018060159054906101000a900460ff1660ff16145b5b610780576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610777906116a0565b60405180910390fd5b60018060156101000a81548160ff021916908360ff16021790555080156107bc5760018060166101000a81548160ff0219169083151502179055505b6107c5826102e6565b5f6005819055508015610825575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498600160405161081c919061170c565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60045460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e8461089b610829565b6040518363ffffffff1660e01b81526004016108b8929190611725565b602060405180830381865afa1580156108d3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108f79190611760565b1015610938576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092f906117d5565b60405180910390fd5b60016005546109479190611820565b6005819055506040518060a0016040528060055481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481525060065f60055481526020019081526020015f205f820151815f01556020820151816001015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019081610a5e9190611a44565b506080820151816004019080519060200190610a7b929190610ef7565b509050507faa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a60055433848787604051610ab8959493929190611c16565b60405180910390a160055490509392505050565b60055481565b610ada610c62565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3f90611ce5565b60405180910390fd5b610b5181610e0d565b50565b610b5c610c62565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc190611d73565b60405180910390fd5b600160149054906101000a900460ff16610c19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1090611e01565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60045481565b610c6a610ef0565b73ffffffffffffffffffffffffffffffffffffffff16610c88610696565b73ffffffffffffffffffffffffffffffffffffffff1614610cde576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd590611e69565b60405180910390fd5b565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610d70576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6790611d73565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610db0610ef0565b73ffffffffffffffffffffffffffffffffffffffff1614610e06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dfd90611ef7565b60405180910390fd5b6001905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b828054828255905f5260205f20908101928215610f3d579160200282015b82811115610f3c578251829081610f2c9190611a44565b5091602001919060010190610f15565b5b509050610f4a9190610f4e565b5090565b5b80821115610f6d575f8181610f649190610f71565b50600101610f4f565b5090565b508054610f7d90611880565b5f825580601f10610f8e5750610fab565b601f0160209004905f5260205f2090810190610faa9190610fae565b5b50565b5b80821115610fc5575f815f905550600101610faf565b5090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ff282610fc9565b9050919050565b61100281610fe8565b82525050565b5f60208201905061101b5f830184610ff9565b92915050565b5f604051905090565b5f80fd5b5f80fd5b61103b81610fe8565b8114611045575f80fd5b50565b5f8135905061105681611032565b92915050565b5f602082840312156110715761107061102a565b5b5f61107e84828501611048565b91505092915050565b5f819050919050565b61109981611087565b81146110a3575f80fd5b50565b5f813590506110b481611090565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611108826110c2565b810181811067ffffffffffffffff82111715611127576111266110d2565b5b80604052505050565b5f611139611021565b905061114582826110ff565b919050565b5f67ffffffffffffffff821115611164576111636110d2565b5b61116d826110c2565b9050602081019050919050565b828183375f83830152505050565b5f61119a6111958461114a565b611130565b9050828152602081018484840111156111b6576111b56110be565b5b6111c184828561117a565b509392505050565b5f82601f8301126111dd576111dc6110ba565b5b81356111ed848260208601611188565b91505092915050565b5f805f6060848603121561120d5761120c61102a565b5b5f61121a868287016110a6565b935050602084013567ffffffffffffffff81111561123b5761123a61102e565b5b611247868287016111c9565b925050604084013567ffffffffffffffff8111156112685761126761102e565b5b611274868287016111c9565b9150509250925092565b5f602082840312156112935761129261102a565b5b5f6112a0848285016110a6565b91505092915050565b6112b281611087565b82525050565b5f6020820190506112cb5f8301846112a9565b92915050565b5f67ffffffffffffffff8211156112eb576112ea6110d2565b5b602082029050602081019050919050565b5f80fd5b5f61131261130d846112d1565b611130565b90508083825260208201905060208402830185811115611335576113346112fc565b5b835b8181101561137c57803567ffffffffffffffff81111561135a576113596110ba565b5b80860161136789826111c9565b85526020850194505050602081019050611337565b5050509392505050565b5f82601f83011261139a576113996110ba565b5b81356113aa848260208601611300565b91505092915050565b5f805f606084860312156113ca576113c961102a565b5b5f84013567ffffffffffffffff8111156113e7576113e661102e565b5b6113f3868287016111c9565b935050602084013567ffffffffffffffff8111156114145761141361102e565b5b61142086828701611386565b925050604061143186828701611048565b9150509250925092565b5f82825260208201905092915050565b7f546f6b656e2061646472657373000000000000000000000000000000000000005f82015250565b5f61147f600d8361143b565b915061148a8261144b565b602082019050919050565b5f6020820190508181035f8301526114ac81611473565b9050919050565b7f4a6f62206e6f7420666f756e64000000000000000000000000000000000000005f82015250565b5f6114e7600d8361143b565b91506114f2826114b3565b602082019050919050565b5f6020820190508181035f830152611514816114db565b9050919050565b5f81519050919050565b5f5b83811015611542578082015181840152602081019050611527565b5f8484015250505050565b5f6115578261151b565b611561818561143b565b9350611571818560208601611525565b61157a816110c2565b840191505092915050565b5f6060820190506115985f8301866112a9565b81810360208301526115aa818561154d565b905081810360408301526115be818461154d565b9050949350505050565b7f4d696e206465706f7369740000000000000000000000000000000000000000005f82015250565b5f6115fc600b8361143b565b9150611607826115c8565b602082019050919050565b5f6020820190508181035f830152611629816115f0565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f61168a602e8361143b565b915061169582611630565b604082019050919050565b5f6020820190508181035f8301526116b78161167e565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6116f66116f16116ec846116be565b6116d3565b6116c7565b9050919050565b611706816116dc565b82525050565b5f60208201905061171f5f8301846116fd565b92915050565b5f6040820190506117385f830185610ff9565b6117456020830184610ff9565b9392505050565b5f8151905061175a81611090565b92915050565b5f602082840312156117755761177461102a565b5b5f6117828482850161174c565b91505092915050565b7f546f6b656e20616c6c6f77616e6365206e6f7420656e6f7567680000000000005f82015250565b5f6117bf601a8361143b565b91506117ca8261178b565b602082019050919050565b5f6020820190508181035f8301526117ec816117b3565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61182a82611087565b915061183583611087565b925082820190508082111561184d5761184c6117f3565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061189757607f821691505b6020821081036118aa576118a9611853565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261190c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118d1565b61191686836118d1565b95508019841693508086168417925050509392505050565b5f61194861194361193e84611087565b6116d3565b611087565b9050919050565b5f819050919050565b6119618361192e565b61197561196d8261194f565b8484546118dd565b825550505050565b5f90565b61198961197d565b611994818484611958565b505050565b5b818110156119b7576119ac5f82611981565b60018101905061199a565b5050565b601f8211156119fc576119cd816118b0565b6119d6846118c2565b810160208510156119e5578190505b6119f96119f1856118c2565b830182611999565b50505b505050565b5f82821c905092915050565b5f611a1c5f1984600802611a01565b1980831691505092915050565b5f611a348383611a0d565b9150826002028217905092915050565b611a4d8261151b565b67ffffffffffffffff811115611a6657611a656110d2565b5b611a708254611880565b611a7b8282856119bb565b5f60209050601f831160018114611aac575f8415611a9a578287015190505b611aa48582611a29565b865550611b0b565b601f198416611aba866118b0565b5f5b82811015611ae157848901518255600182019150602085019450602081019050611abc565b86831015611afe5784890151611afa601f891682611a0d565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f611b568261151b565b611b608185611b3c565b9350611b70818560208601611525565b611b79816110c2565b840191505092915050565b5f611b8f8383611b4c565b905092915050565b5f602082019050919050565b5f611bad82611b13565b611bb78185611b1d565b935083602082028501611bc985611b2d565b805f5b85811015611c045784840389528151611be58582611b84565b9450611bf083611b97565b925060208a01995050600181019050611bcc565b50829750879550505050505092915050565b5f60a082019050611c295f8301886112a9565b611c366020830187610ff9565b611c436040830186610ff9565b8181036060830152611c55818561154d565b90508181036080830152611c698184611ba3565b90509695505050505050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f611ccf60268361143b565b9150611cda82611c75565b604082019050919050565b5f6020820190508181035f830152611cfc81611cc3565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f611d5d60358361143b565b9150611d6882611d03565b604082019050919050565b5f6020820190508181035f830152611d8a81611d51565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f611deb60398361143b565b9150611df682611d91565b604082019050919050565b5f6020820190508181035f830152611e1881611ddf565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f611e5360208361143b565b9150611e5e82611e1f565b602082019050919050565b5f6020820190508181035f830152611e8081611e47565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f611ee1603b8361143b565b9150611eec82611e87565b604082019050919050565b5f6020820190508181035f830152611f0e81611ed5565b905091905056fea26469706673582212201b80b6736325ac50ded152ea3c7f6de1b08a14e50ce5812dd79eed8f6ebe744964736f6c63430008150033",
}

// JobcreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use JobcreatorMetaData.ABI instead.
var JobcreatorABI = JobcreatorMetaData.ABI

// JobcreatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobcreatorMetaData.Bin instead.
var JobcreatorBin = JobcreatorMetaData.Bin

// DeployJobcreator deploys a new Ethereum contract, binding an instance of Jobcreator to it.
func DeployJobcreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Jobcreator, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobcreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// Jobcreator is an auto generated Go binding around an Ethereum contract.
type Jobcreator struct {
	JobcreatorCaller     // Read-only binding to the contract
	JobcreatorTransactor // Write-only binding to the contract
	JobcreatorFilterer   // Log filterer for contract events
}

// JobcreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobcreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobcreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobcreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobcreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobcreatorSession struct {
	Contract     *Jobcreator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobcreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobcreatorCallerSession struct {
	Contract *JobcreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// JobcreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobcreatorTransactorSession struct {
	Contract     *JobcreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// JobcreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobcreatorRaw struct {
	Contract *Jobcreator // Generic contract binding to access the raw methods on
}

// JobcreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobcreatorCallerRaw struct {
	Contract *JobcreatorCaller // Generic read-only contract binding to access the raw methods on
}

// JobcreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobcreatorTransactorRaw struct {
	Contract *JobcreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobcreator creates a new instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreator(address common.Address, backend bind.ContractBackend) (*Jobcreator, error) {
	contract, err := bindJobcreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobcreator{JobcreatorCaller: JobcreatorCaller{contract: contract}, JobcreatorTransactor: JobcreatorTransactor{contract: contract}, JobcreatorFilterer: JobcreatorFilterer{contract: contract}}, nil
}

// NewJobcreatorCaller creates a new read-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorCaller(address common.Address, caller bind.ContractCaller) (*JobcreatorCaller, error) {
	contract, err := bindJobcreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorCaller{contract: contract}, nil
}

// NewJobcreatorTransactor creates a new write-only instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*JobcreatorTransactor, error) {
	contract, err := bindJobcreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobcreatorTransactor{contract: contract}, nil
}

// NewJobcreatorFilterer creates a new log filterer instance of Jobcreator, bound to a specific deployed contract.
func NewJobcreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*JobcreatorFilterer, error) {
	contract, err := bindJobcreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobcreatorFilterer{contract: contract}, nil
}

// bindJobcreator binds a generic wrapper to an already deployed contract.
func bindJobcreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JobcreatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.JobcreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.JobcreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobcreator *JobcreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobcreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobcreator *JobcreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobcreator *JobcreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobcreator.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetControllerAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetControllerAddress(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) GetRequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getRequiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetRequiredDeposit is a free data retrieval call binding the contract method 0x58e56db4.
//
// Solidity: function getRequiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) GetRequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.GetRequiredDeposit(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) GetTokenAddress() (common.Address, error) {
	return _Jobcreator.Contract.GetTokenAddress(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) NextJobID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "nextJobID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// NextJobID is a free data retrieval call binding the contract method 0xd2a715c0.
//
// Solidity: function nextJobID() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) NextJobID() (*big.Int, error) {
	return _Jobcreator.Contract.NextJobID(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Jobcreator *JobcreatorCallerSession) Owner() (common.Address, error) {
	return _Jobcreator.Contract.Owner(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCaller) RequiredDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Jobcreator.contract.Call(opts, &out, "requiredDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// RequiredDeposit is a free data retrieval call binding the contract method 0xfb7cfdd7.
//
// Solidity: function requiredDeposit() view returns(uint256)
func (_Jobcreator *JobcreatorCallerSession) RequiredDeposit() (*big.Int, error) {
	return _Jobcreator.Contract.RequiredDeposit(&_Jobcreator.CallOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Jobcreator *JobcreatorTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Jobcreator.Contract.DisableChangeControllerAddress(&_Jobcreator.TransactOpts)
}

// Forfeit is a paid mutator transaction binding the contract method 0x36aef967.
//
// Solidity: function forfeit(uint256 id, string dealId, string message) returns()
func (_Jobcreator *JobcreatorTransactor) Forfeit(opts *bind.TransactOpts, id *big.Int, dealId string, message string) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "forfeit", id, dealId, message)
}

// Forfeit is a paid mutator transaction binding the contract method 0x36aef967.
//
// Solidity: function forfeit(uint256 id, string dealId, string message) returns()
func (_Jobcreator *JobcreatorSession) Forfeit(id *big.Int, dealId string, message string) (*types.Transaction, error) {
	return _Jobcreator.Contract.Forfeit(&_Jobcreator.TransactOpts, id, dealId, message)
}

// Forfeit is a paid mutator transaction binding the contract method 0x36aef967.
//
// Solidity: function forfeit(uint256 id, string dealId, string message) returns()
func (_Jobcreator *JobcreatorTransactorSession) Forfeit(id *big.Int, dealId string, message string) (*types.Transaction, error) {
	return _Jobcreator.Contract.Forfeit(&_Jobcreator.TransactOpts, id, dealId, message)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.Initialize(&_Jobcreator.TransactOpts, _tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Jobcreator *JobcreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Jobcreator.Contract.RenounceOwnership(&_Jobcreator.TransactOpts)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactor) RunJob(opts *bind.TransactOpts, module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "runJob", module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// RunJob is a paid mutator transaction binding the contract method 0xc75555fa.
//
// Solidity: function runJob(string module, string[] inputs, address payee) returns(uint256)
func (_Jobcreator *JobcreatorTransactorSession) RunJob(module string, inputs []string, payee common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.RunJob(&_Jobcreator.TransactOpts, module, inputs, payee)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetControllerAddress(&_Jobcreator.TransactOpts, _controllerAddress)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactor) SetRequiredDeposit(opts *bind.TransactOpts, cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setRequiredDeposit", cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetRequiredDeposit is a paid mutator transaction binding the contract method 0x4c526c76.
//
// Solidity: function setRequiredDeposit(uint256 cost) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetRequiredDeposit(cost *big.Int) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetRequiredDeposit(&_Jobcreator.TransactOpts, cost)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Jobcreator *JobcreatorTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.SetTokenAddress(&_Jobcreator.TransactOpts, _tokenAddress)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactor) SubmitResults(opts *bind.TransactOpts, id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "submitResults", id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// SubmitResults is a paid mutator transaction binding the contract method 0x6c0f1f58.
//
// Solidity: function submitResults(uint256 id, string dealId, string dataId) returns()
func (_Jobcreator *JobcreatorTransactorSession) SubmitResults(id *big.Int, dealId string, dataId string) (*types.Transaction, error) {
	return _Jobcreator.Contract.SubmitResults(&_Jobcreator.TransactOpts, id, dealId, dataId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Jobcreator *JobcreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Jobcreator.Contract.TransferOwnership(&_Jobcreator.TransactOpts, newOwner)
}

// JobcreatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Jobcreator contract.
type JobcreatorInitializedIterator struct {
	Event *JobcreatorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobcreatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobcreatorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobcreatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorInitialized represents a Initialized event raised by the Jobcreator contract.
type JobcreatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*JobcreatorInitializedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JobcreatorInitializedIterator{contract: _Jobcreator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JobcreatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorInitialized)
				if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Jobcreator *JobcreatorFilterer) ParseInitialized(log types.Log) (*JobcreatorInitialized, error) {
	event := new(JobcreatorInitialized)
	if err := _Jobcreator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorJobAddedIterator is returned from FilterJobAdded and is used to iterate over the raw logs and unpacked data for JobAdded events raised by the Jobcreator contract.
type JobcreatorJobAddedIterator struct {
	Event *JobcreatorJobAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobcreatorJobAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorJobAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobcreatorJobAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobcreatorJobAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorJobAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorJobAdded represents a JobAdded event raised by the Jobcreator contract.
type JobcreatorJobAdded struct {
	Id              *big.Int
	CallingContract common.Address
	Payee           common.Address
	Module          string
	Inputs          []string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterJobAdded is a free log retrieval operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) FilterJobAdded(opts *bind.FilterOpts) (*JobcreatorJobAddedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return &JobcreatorJobAddedIterator{contract: _Jobcreator.contract, event: "JobAdded", logs: logs, sub: sub}, nil
}

// WatchJobAdded is a free log subscription operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) WatchJobAdded(opts *bind.WatchOpts, sink chan<- *JobcreatorJobAdded) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "JobAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorJobAdded)
				if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobAdded is a log parse operation binding the contract event 0xaa171d38d08c39622dbd189ddd0a5c0cf3ea10eeda0c7d4c2c75d13fcc4fe14a.
//
// Solidity: event JobAdded(uint256 id, address calling_contract, address payee, string module, string[] inputs)
func (_Jobcreator *JobcreatorFilterer) ParseJobAdded(log types.Log) (*JobcreatorJobAdded, error) {
	event := new(JobcreatorJobAdded)
	if err := _Jobcreator.contract.UnpackLog(event, "JobAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorJobForfeitedIterator is returned from FilterJobForfeited and is used to iterate over the raw logs and unpacked data for JobForfeited events raised by the Jobcreator contract.
type JobcreatorJobForfeitedIterator struct {
	Event *JobcreatorJobForfeited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobcreatorJobForfeitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorJobForfeited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobcreatorJobForfeited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobcreatorJobForfeitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorJobForfeitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorJobForfeited represents a JobForfeited event raised by the Jobcreator contract.
type JobcreatorJobForfeited struct {
	Id      *big.Int
	DealId  string
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterJobForfeited is a free log retrieval operation binding the contract event 0x365cc3878e524c2da89b4cad993fe1d16f0aa1698b90ecc6ac111c47732b72e3.
//
// Solidity: event JobForfeited(uint256 id, string dealId, string message)
func (_Jobcreator *JobcreatorFilterer) FilterJobForfeited(opts *bind.FilterOpts) (*JobcreatorJobForfeitedIterator, error) {

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "JobForfeited")
	if err != nil {
		return nil, err
	}
	return &JobcreatorJobForfeitedIterator{contract: _Jobcreator.contract, event: "JobForfeited", logs: logs, sub: sub}, nil
}

// WatchJobForfeited is a free log subscription operation binding the contract event 0x365cc3878e524c2da89b4cad993fe1d16f0aa1698b90ecc6ac111c47732b72e3.
//
// Solidity: event JobForfeited(uint256 id, string dealId, string message)
func (_Jobcreator *JobcreatorFilterer) WatchJobForfeited(opts *bind.WatchOpts, sink chan<- *JobcreatorJobForfeited) (event.Subscription, error) {

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "JobForfeited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorJobForfeited)
				if err := _Jobcreator.contract.UnpackLog(event, "JobForfeited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJobForfeited is a log parse operation binding the contract event 0x365cc3878e524c2da89b4cad993fe1d16f0aa1698b90ecc6ac111c47732b72e3.
//
// Solidity: event JobForfeited(uint256 id, string dealId, string message)
func (_Jobcreator *JobcreatorFilterer) ParseJobForfeited(log types.Log) (*JobcreatorJobForfeited, error) {
	event := new(JobcreatorJobForfeited)
	if err := _Jobcreator.contract.UnpackLog(event, "JobForfeited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobcreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Jobcreator contract.
type JobcreatorOwnershipTransferredIterator struct {
	Event *JobcreatorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *JobcreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobcreatorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(JobcreatorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *JobcreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobcreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobcreatorOwnershipTransferred represents a OwnershipTransferred event raised by the Jobcreator contract.
type JobcreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JobcreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JobcreatorOwnershipTransferredIterator{contract: _Jobcreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JobcreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Jobcreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobcreatorOwnershipTransferred)
				if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Jobcreator *JobcreatorFilterer) ParseOwnershipTransferred(log types.Log) (*JobcreatorOwnershipTransferred, error) {
	event := new(JobcreatorOwnershipTransferred)
	if err := _Jobcreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
