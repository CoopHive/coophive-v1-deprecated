// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

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

// PaymentsMetaData contains all meta data concerning the Payments contract.
var PaymentsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCoopHivePayments.PaymentReason\",\"name\":\"reason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumCoopHivePayments.PaymentDirection\",\"name\":\"direction\",\"type\":\"uint8\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"acceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"addResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"agreeResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"checkResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableChangeTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"forfeit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getControllerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"jobCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationAcceptResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"mediationRejectResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controllerAddress\",\"type\":\"address\"}],\"name\":\"setControllerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"setTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundJobCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutAgreeRefundResourceProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutJudgeResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultsCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mediationFee\",\"type\":\"uint256\"}],\"name\":\"timeoutMediateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"resourceProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"jobCreator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"paymentCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeoutCollateral\",\"type\":\"uint256\"}],\"name\":\"timeoutSubmitResults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260018060146101000a81548160ff0219169083151502179055506001600360146101000a81548160ff02191690831515021790555034801562000045575f80fd5b50620000666200005a6200006c60201b60201c565b6200007360201b60201c565b62000134565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61291180620001425f395ff3fe608060405234801561000f575f80fd5b506004361061014b575f3560e01c80638da5cb5b116100c1578063b91880351161007a578063b9188035146102f9578063c4d66de814610315578063c57380a214610331578063d2485cce1461034f578063f2fde38b1461036b578063f3d3d448146103875761014b565b80638da5cb5b146102615780639e3868dc1461027f578063a47029581461029b578063aea38251146102a5578063afe1dff7146102c1578063b1356714146102dd5761014b565b80632a1f9072116101135780632a1f9072146101dd57806338698529146101f9578063407b85b1146102155780634bc28da114610231578063715018a61461023b578063823f3de1146102455761014b565b806302fd8f801461014f57806309cab5101461016b5780630ef0d89e1461018757806310fe9ae8146101a357806326a4e8d2146101c1575b5f80fd5b610169600480360381019061016491906119a2565b6103a3565b005b61018560048036038101906101809190611a35565b610447565b005b6101a1600480360381019061019c9190611ab5565b6104dd565b005b6101ab610566565b6040516101b89190611b30565b60405180910390f35b6101db60048036038101906101d69190611b49565b61058e565b005b6101f760048036038101906101f29190611b74565b6106d6565b005b610213600480360381019061020e91906119a2565b61074c565b005b61022f600480360381019061022a91906119a2565b6107e4565b005b61023961080e565b005b610243610832565b005b61025f600480360381019061025a9190611c2d565b610845565b005b610269610920565b6040516102769190611b30565b60405180910390f35b61029960048036038101906102949190611ab5565b610947565b005b6102a36109cf565b005b6102bf60048036038101906102ba9190611a35565b6109f3565b005b6102db60048036038101906102d69190611a35565b610a89565b005b6102f760048036038101906102f29190611b74565b610b1f565b005b610313600480360381019061030e9190611a35565b610c02565b005b61032f600480360381019061032a9190611b49565b610c96565b005b610339610dd7565b6040516103469190611b30565b60405180910390f35b61036960048036038101906103649190611c2d565b610dff565b005b61038560048036038101906103809190611b49565b610e37565b005b6103a1600480360381019061039c9190611b49565b610eb9565b005b6103ab610fc1565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161461041a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041190611d52565b60405180910390fd5b6104268584845f6110ee565b61043385848360036110ee565b6104408585836003611214565b5050505050565b61044f610fc1565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146104be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b590611de0565b60405180910390fd5b6104ca8483600261133a565b6104d784848360036110ee565b50505050565b6104e5610fc1565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610554576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054b90611de0565b60405180910390fd5b61056183838360036110ee565b505050565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610596611537565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610604576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fb90611e6e565b60405180910390fd5b600360149054906101000a900460ff16610653576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161064a90611efc565b60405180910390fd5b8060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6106de610fc1565b505f8490505f848611156106f457849150610703565b85856107009190611f47565b90505b61071189888a8560046115b5565b61071f8988328660056115b5565b5f811115610734576107338988835f6110ee565b5b61074189898660026110ee565b505050505050505050565b610754610fc1565b508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146107c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107ba90611de0565b60405180910390fd5b6107d085858460026110ee565b6107dd8584836003611214565b5050505050565b6107ec610fc1565b506107fa85848360016110ee565b61080785858460016110ee565b5050505050565b610816611537565b5f600360146101000a81548160ff021916908315150217905550565b61083a611537565b6108435f6116de565b565b61084d610fc1565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614806108b357508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b6108f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e990611fea565b60405180910390fd5b6108ff86868460026110ee565b61090b8685855f6110ee565b61091886858360056110ee565b505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61094f610fc1565b508173ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16146109be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109b590611de0565b60405180910390fd5b6109ca8382600361133a565b505050565b6109d7611537565b5f600160146101000a81548160ff021916908315150217905550565b6109fb610fc1565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610a6a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6190611d52565b60405180910390fd5b610a7784848460036110ee565b610a838482600561133a565b50505050565b610a91610fc1565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610b00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610af790611d52565b60405180910390fd5b610b0c8484845f6110ee565b610b1984848360036110ee565b50505050565b610b27610fc1565b508473ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610b96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8d90611d52565b60405180910390fd5b5f8490505f84861115610bab57849150610bba565b8585610bb79190611f47565b90505b610bc889888a8560046115b5565b5f811115610bdd57610bdc8988835f6110ee565b5b610bea89888560036110ee565b610bf789898660026110ee565b505050505050505050565b610c0a610fc1565b508273ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff1614610c79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7090611d52565b60405180910390fd5b610c8484835f61133a565b610c908482600361133a565b50505050565b5f600160169054906101000a900460ff16159050808015610cc8575060018060159054906101000a900460ff1660ff16105b80610cf65750610cd73061179f565b158015610cf5575060018060159054906101000a900460ff1660ff16145b5b610d35576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2c90612078565b60405180910390fd5b60018060156101000a81548160ff021916908360ff1602179055508015610d715760018060166101000a81548160ff0219169083151502179055505b610d7a8261058e565b8015610dd3575f600160166101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610dca91906120e4565b60405180910390a15b5050565b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610e07610fc1565b50610e148685855f6110ee565b610e228685328460056115b5565b610e2f8686846002611214565b505050505050565b610e3f611537565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ead576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea49061216d565b60405180910390fd5b610eb6816116de565b50565b610ec1611537565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f26906121fb565b60405180910390fd5b600160149054906101000a900460ff16610f7e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7590612289565b60405180910390fd5b8060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b5f8073ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611051576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611048906121fb565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166110916117c1565b73ffffffffffffffffffffffffffffffffffffffff16146110e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110de90612317565b60405180910390fd5b6001905090565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663599efa6b85856040518363ffffffff1660e01b815260040161114b929190612344565b6020604051808303815f875af1158015611167573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061118b91906123a0565b9050806111cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c49061243b565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a984285858585600260405161120595949392919061257c565b60405180910390a15050505050565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166388c2bdfe85856040518363ffffffff1660e01b8152600401611271929190612344565b6020604051808303815f875af115801561128d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112b191906123a0565b9050806112f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ea90612644565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a984285858585600360405161132b95949392919061257c565b60405180910390a15050505050565b8160035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231326040518263ffffffff1660e01b81526004016113959190611b30565b602060405180830381865afa1580156113b0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113d49190612676565b1015611415576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161140c90612711565b60405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635407e93c846040518263ffffffff1660e01b8152600401611470919061272f565b6020604051808303815f875af115801561148c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114b091906123a0565b9050806114f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114e9906127b8565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842843285855f60405161152995949392919061257c565b60405180910390a150505050565b61153f6117c1565b73ffffffffffffffffffffffffffffffffffffffff1661155d610920565b73ffffffffffffffffffffffffffffffffffffffff16146115b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115aa90612820565b60405180910390fd5b565b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663065e86c88686866040518463ffffffff1660e01b81526004016116149392919061283e565b6020604051808303815f875af1158015611630573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061165491906123a0565b905080611696576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161168d906128bd565b60405180910390fd5b7f64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a98428685858560016040516116ce95949392919061257c565b60405180910390a1505050505050565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b5f33905090565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611827826117e1565b810181811067ffffffffffffffff82111715611846576118456117f1565b5b80604052505050565b5f6118586117c8565b9050611864828261181e565b919050565b5f67ffffffffffffffff821115611883576118826117f1565b5b61188c826117e1565b9050602081019050919050565b828183375f83830152505050565b5f6118b96118b484611869565b61184f565b9050828152602081018484840111156118d5576118d46117dd565b5b6118e0848285611899565b509392505050565b5f82601f8301126118fc576118fb6117d9565b5b813561190c8482602086016118a7565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61193e82611915565b9050919050565b61194e81611934565b8114611958575f80fd5b50565b5f8135905061196981611945565b92915050565b5f819050919050565b6119818161196f565b811461198b575f80fd5b50565b5f8135905061199c81611978565b92915050565b5f805f805f60a086880312156119bb576119ba6117d1565b5b5f86013567ffffffffffffffff8111156119d8576119d76117d5565b5b6119e4888289016118e8565b95505060206119f58882890161195b565b9450506040611a068882890161195b565b9350506060611a178882890161198e565b9250506080611a288882890161198e565b9150509295509295909350565b5f805f8060808587031215611a4d57611a4c6117d1565b5b5f85013567ffffffffffffffff811115611a6a57611a696117d5565b5b611a76878288016118e8565b9450506020611a878782880161195b565b9350506040611a988782880161198e565b9250506060611aa98782880161198e565b91505092959194509250565b5f805f60608486031215611acc57611acb6117d1565b5b5f84013567ffffffffffffffff811115611ae957611ae86117d5565b5b611af5868287016118e8565b9350506020611b068682870161195b565b9250506040611b178682870161198e565b9150509250925092565b611b2a81611934565b82525050565b5f602082019050611b435f830184611b21565b92915050565b5f60208284031215611b5e57611b5d6117d1565b5b5f611b6b8482850161195b565b91505092915050565b5f805f805f805f60e0888a031215611b8f57611b8e6117d1565b5b5f88013567ffffffffffffffff811115611bac57611bab6117d5565b5b611bb88a828b016118e8565b9750506020611bc98a828b0161195b565b9650506040611bda8a828b0161195b565b9550506060611beb8a828b0161198e565b9450506080611bfc8a828b0161198e565b93505060a0611c0d8a828b0161198e565b92505060c0611c1e8a828b0161198e565b91505092959891949750929550565b5f805f805f8060c08789031215611c4757611c466117d1565b5b5f87013567ffffffffffffffff811115611c6457611c636117d5565b5b611c7089828a016118e8565b9650506020611c8189828a0161195b565b9550506040611c9289828a0161195b565b9450506060611ca389828a0161198e565b9350506080611cb489828a0161198e565b92505060a0611cc589828a0161198e565b9150509295509295509295565b5f82825260208201905092915050565b7f436f6f70486976655061796d656e74733a2043616e206f6e6c792062652063615f8201527f6c6c656420627920746865204a43000000000000000000000000000000000000602082015250565b5f611d3c602e83611cd2565b9150611d4782611ce2565b604082019050919050565b5f6020820190508181035f830152611d6981611d30565b9050919050565b7f436f6f70486976655061796d656e74733a2043616e206f6e6c792062652063615f8201527f6c6c656420627920746865205250000000000000000000000000000000000000602082015250565b5f611dca602e83611cd2565b9150611dd582611d70565b604082019050919050565b5f6020820190508181035f830152611df781611dbe565b9050919050565b7f4c696c657061645061796d656e74733a20546f6b656e2061646472657373206d5f8201527f75737420626520646566696e6564000000000000000000000000000000000000602082015250565b5f611e58602e83611cd2565b9150611e6382611dfe565b604082019050919050565b5f6020820190508181035f830152611e8581611e4c565b9050919050565b7f436f6f7048697665546f6b656e3a2063616e4368616e6765546f6b656e4164645f8201527f726573732069732064697361626c656400000000000000000000000000000000602082015250565b5f611ee6603083611cd2565b9150611ef182611e8c565b604082019050919050565b5f6020820190508181035f830152611f1381611eda565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611f518261196f565b9150611f5c8361196f565b9250828203905081811115611f7457611f73611f1a565b5b92915050565b7f436f6f70486976655061796d656e74733a2043616e206f6e6c792062652063615f8201527f6c6c656420627920746865205250206f72204a43000000000000000000000000602082015250565b5f611fd4603483611cd2565b9150611fdf82611f7a565b604082019050919050565b5f6020820190508181035f83015261200181611fc8565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c7265615f8201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b5f612062602e83611cd2565b915061206d82612008565b604082019050919050565b5f6020820190508181035f83015261208f81612056565b9050919050565b5f819050919050565b5f60ff82169050919050565b5f819050919050565b5f6120ce6120c96120c484612096565b6120ab565b61209f565b9050919050565b6120de816120b4565b82525050565b5f6020820190506120f75f8301846120d5565b92915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f20615f8201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b5f612157602683611cd2565b9150612162826120fd565b604082019050919050565b5f6020820190508181035f8301526121848161214b565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a20436f6e74726f6c6c65722061645f8201527f6472657373206d75737420626520646566696e65640000000000000000000000602082015250565b5f6121e5603583611cd2565b91506121f08261218b565b604082019050919050565b5f6020820190508181035f830152612212816121d9565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a2063616e4368616e6765436f6e745f8201527f726f6c6c6572416464726573732069732064697361626c656400000000000000602082015250565b5f612273603983611cd2565b915061227e82612219565b604082019050919050565b5f6020820190508181035f8301526122a081612267565b9050919050565b7f436f6e74726f6c6c65724f776e61626c653a204f6e6c792074686520636f6e745f8201527f726f6c6c65722063616e2063616c6c2074686973206d6574686f640000000000602082015250565b5f612301603b83611cd2565b915061230c826122a7565b604082019050919050565b5f6020820190508181035f83015261232e816122f5565b9050919050565b61233e8161196f565b82525050565b5f6040820190506123575f830185611b21565b6123646020830184612335565b9392505050565b5f8115159050919050565b61237f8161236b565b8114612389575f80fd5b50565b5f8151905061239a81612376565b92915050565b5f602082840312156123b5576123b46117d1565b5b5f6123c28482850161238c565b91505092915050565b7f436f6f70486976655061796d656e74733a20526566756e6420657363726f77205f8201527f6661696c65640000000000000000000000000000000000000000000000000000602082015250565b5f612425602683611cd2565b9150612430826123cb565b604082019050919050565b5f6020820190508181035f83015261245281612419565b9050919050565b5f81519050919050565b5f5b83811015612480578082015181840152602081019050612465565b5f8484015250505050565b5f61249582612459565b61249f8185611cd2565b93506124af818560208601612463565b6124b8816117e1565b840191505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60068110612501576125006124c3565b5b50565b5f819050612511826124f0565b919050565b5f61252082612504565b9050919050565b61253081612516565b82525050565b60048110612547576125466124c3565b5b50565b5f81905061255782612536565b919050565b5f6125668261254a565b9050919050565b6125768161255c565b82525050565b5f60a0820190508181035f830152612594818861248b565b90506125a36020830187611b21565b6125b06040830186612335565b6125bd6060830185612527565b6125ca608083018461256d565b9695505050505050565b7f436f6f70486976655061796d656e74733a20536c61736820657363726f7720665f8201527f61696c6564000000000000000000000000000000000000000000000000000000602082015250565b5f61262e602583611cd2565b9150612639826125d4565b604082019050919050565b5f6020820190508181035f83015261265b81612622565b9050919050565b5f8151905061267081611978565b92915050565b5f6020828403121561268b5761268a6117d1565b5b5f61269884828501612662565b91505092915050565b7f436f6f70486976655061796d656e74733a20496e73756666696369656e7420625f8201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b5f6126fb602683611cd2565b9150612706826126a1565b604082019050919050565b5f6020820190508181035f830152612728816126ef565b9050919050565b5f6020820190506127425f830184612335565b92915050565b7f436f6f70486976655061796d656e74733a2050617920657363726f77206661695f8201527f6c65640000000000000000000000000000000000000000000000000000000000602082015250565b5f6127a2602383611cd2565b91506127ad82612748565b604082019050919050565b5f6020820190508181035f8301526127cf81612796565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65725f82015250565b5f61280a602083611cd2565b9150612815826127d6565b602082019050919050565b5f6020820190508181035f830152612837816127fe565b9050919050565b5f6060820190506128515f830186611b21565b61285e6020830185611b21565b61286b6040830184612335565b949350505050565b7f436f6f70486976655061796d656e74733a20506179206a6f62206661696c65645f82015250565b5f6128a7602083611cd2565b91506128b282612873565b602082019050919050565b5f6020820190508181035f8301526128d48161289b565b905091905056fea26469706673582212200fbd64bb30714a16fd91f775939baa1fb84b6f2ba2f2e6f68c9ec13aaa28b24464736f6c63430008150033",
}

// PaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentsMetaData.ABI instead.
var PaymentsABI = PaymentsMetaData.ABI

// PaymentsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PaymentsMetaData.Bin instead.
var PaymentsBin = PaymentsMetaData.Bin

// DeployPayments deploys a new Ethereum contract, binding an instance of Payments to it.
func DeployPayments(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Payments, error) {
	parsed, err := PaymentsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PaymentsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// Payments is an auto generated Go binding around an Ethereum contract.
type Payments struct {
	PaymentsCaller     // Read-only binding to the contract
	PaymentsTransactor // Write-only binding to the contract
	PaymentsFilterer   // Log filterer for contract events
}

// PaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentsSession struct {
	Contract     *Payments         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentsCallerSession struct {
	Contract *PaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentsTransactorSession struct {
	Contract     *PaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentsRaw struct {
	Contract *Payments // Generic contract binding to access the raw methods on
}

// PaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentsCallerRaw struct {
	Contract *PaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentsTransactorRaw struct {
	Contract *PaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayments creates a new instance of Payments, bound to a specific deployed contract.
func NewPayments(address common.Address, backend bind.ContractBackend) (*Payments, error) {
	contract, err := bindPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payments{PaymentsCaller: PaymentsCaller{contract: contract}, PaymentsTransactor: PaymentsTransactor{contract: contract}, PaymentsFilterer: PaymentsFilterer{contract: contract}}, nil
}

// NewPaymentsCaller creates a new read-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsCaller(address common.Address, caller bind.ContractCaller) (*PaymentsCaller, error) {
	contract, err := bindPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsCaller{contract: contract}, nil
}

// NewPaymentsTransactor creates a new write-only instance of Payments, bound to a specific deployed contract.
func NewPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentsTransactor, error) {
	contract, err := bindPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentsTransactor{contract: contract}, nil
}

// NewPaymentsFilterer creates a new log filterer instance of Payments, bound to a specific deployed contract.
func NewPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentsFilterer, error) {
	contract, err := bindPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentsFilterer{contract: contract}, nil
}

// bindPayments binds a generic wrapper to an already deployed contract.
func bindPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.PaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.PaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payments *PaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payments *PaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payments *PaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payments.Contract.contract.Transact(opts, method, params...)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsCaller) GetControllerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "getControllerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsSession) GetControllerAddress() (common.Address, error) {
	return _Payments.Contract.GetControllerAddress(&_Payments.CallOpts)
}

// GetControllerAddress is a free data retrieval call binding the contract method 0xc57380a2.
//
// Solidity: function getControllerAddress() view returns(address)
func (_Payments *PaymentsCallerSession) GetControllerAddress() (common.Address, error) {
	return _Payments.Contract.GetControllerAddress(&_Payments.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsSession) GetTokenAddress() (common.Address, error) {
	return _Payments.Contract.GetTokenAddress(&_Payments.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Payments *PaymentsCallerSession) GetTokenAddress() (common.Address, error) {
	return _Payments.Contract.GetTokenAddress(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payments.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payments *PaymentsCallerSession) Owner() (common.Address, error) {
	return _Payments.Contract.Owner(&_Payments.CallOpts)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AcceptResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "acceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AcceptResult is a paid mutator transaction binding the contract method 0xb1356714.
//
// Solidity: function acceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AddResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "addResult", dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AddResult(dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AddResult is a paid mutator transaction binding the contract method 0x09cab510.
//
// Solidity: function addResult(string dealId, address resourceProvider, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AddResult(dealId string, resourceProvider common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AddResult(&_Payments.TransactOpts, dealId, resourceProvider, resultsCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeJobCreator(opts *bind.TransactOpts, dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeJobCreator is a paid mutator transaction binding the contract method 0xb9188035.
//
// Solidity: function agreeJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) AgreeResourceProvider(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "agreeResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) AgreeResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// AgreeResourceProvider is a paid mutator transaction binding the contract method 0x9e3868dc.
//
// Solidity: function agreeResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) AgreeResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.AgreeResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) CheckResult(opts *bind.TransactOpts, dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "checkResult", dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) CheckResult(dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckResult(&_Payments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// CheckResult is a paid mutator transaction binding the contract method 0xaea38251.
//
// Solidity: function checkResult(string dealId, address jobCreator, uint256 timeoutCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) CheckResult(dealId string, jobCreator common.Address, timeoutCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.CheckResult(&_Payments.TransactOpts, dealId, jobCreator, timeoutCollateral, mediationFee)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsTransactor) DisableChangeControllerAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "disableChangeControllerAddress")
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeControllerAddress(&_Payments.TransactOpts)
}

// DisableChangeControllerAddress is a paid mutator transaction binding the contract method 0xa4702958.
//
// Solidity: function disableChangeControllerAddress() returns()
func (_Payments *PaymentsTransactorSession) DisableChangeControllerAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeControllerAddress(&_Payments.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsTransactor) DisableChangeTokenAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "disableChangeTokenAddress")
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeTokenAddress(&_Payments.TransactOpts)
}

// DisableChangeTokenAddress is a paid mutator transaction binding the contract method 0x4bc28da1.
//
// Solidity: function disableChangeTokenAddress() returns()
func (_Payments *PaymentsTransactorSession) DisableChangeTokenAddress() (*types.Transaction, error) {
	return _Payments.Contract.DisableChangeTokenAddress(&_Payments.TransactOpts)
}

// Forfeit is a paid mutator transaction binding the contract method 0x407b85b1.
//
// Solidity: function forfeit(string dealId, address jobCreator, address resourceProvider, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) Forfeit(opts *bind.TransactOpts, dealId string, jobCreator common.Address, resourceProvider common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "forfeit", dealId, jobCreator, resourceProvider, paymentCollateral, timeoutCollateral)
}

// Forfeit is a paid mutator transaction binding the contract method 0x407b85b1.
//
// Solidity: function forfeit(string dealId, address jobCreator, address resourceProvider, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) Forfeit(dealId string, jobCreator common.Address, resourceProvider common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Forfeit(&_Payments.TransactOpts, dealId, jobCreator, resourceProvider, paymentCollateral, timeoutCollateral)
}

// Forfeit is a paid mutator transaction binding the contract method 0x407b85b1.
//
// Solidity: function forfeit(string dealId, address jobCreator, address resourceProvider, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) Forfeit(dealId string, jobCreator common.Address, resourceProvider common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.Forfeit(&_Payments.TransactOpts, dealId, jobCreator, resourceProvider, paymentCollateral, timeoutCollateral)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsTransactor) Initialize(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "initialize", _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.Initialize(&_Payments.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _tokenAddress) returns()
func (_Payments *PaymentsTransactorSession) Initialize(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.Initialize(&_Payments.TransactOpts, _tokenAddress)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationAcceptResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationAcceptResult", dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationAcceptResult is a paid mutator transaction binding the contract method 0x2a1f9072.
//
// Solidity: function mediationAcceptResult(string dealId, address resourceProvider, address jobCreator, uint256 jobCost, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationAcceptResult(dealId string, resourceProvider common.Address, jobCreator common.Address, jobCost *big.Int, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationAcceptResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, jobCost, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) MediationRejectResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "mediationRejectResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// MediationRejectResult is a paid mutator transaction binding the contract method 0xd2485cce.
//
// Solidity: function mediationRejectResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) MediationRejectResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.MediationRejectResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payments.Contract.RenounceOwnership(&_Payments.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Payments *PaymentsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Payments.Contract.RenounceOwnership(&_Payments.TransactOpts)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsTransactor) SetControllerAddress(opts *bind.TransactOpts, _controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "setControllerAddress", _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetControllerAddress(&_Payments.TransactOpts, _controllerAddress)
}

// SetControllerAddress is a paid mutator transaction binding the contract method 0xf3d3d448.
//
// Solidity: function setControllerAddress(address _controllerAddress) returns()
func (_Payments *PaymentsTransactorSession) SetControllerAddress(_controllerAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetControllerAddress(&_Payments.TransactOpts, _controllerAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsTransactor) SetTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "setTokenAddress", _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetTokenAddress(&_Payments.TransactOpts, _tokenAddress)
}

// SetTokenAddress is a paid mutator transaction binding the contract method 0x26a4e8d2.
//
// Solidity: function setTokenAddress(address _tokenAddress) returns()
func (_Payments *PaymentsTransactorSession) SetTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Payments.Contract.SetTokenAddress(&_Payments.TransactOpts, _tokenAddress)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundJobCreator(opts *bind.TransactOpts, dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundJobCreator", dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundJobCreator is a paid mutator transaction binding the contract method 0xafe1dff7.
//
// Solidity: function timeoutAgreeRefundJobCreator(string dealId, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundJobCreator(dealId string, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundJobCreator(&_Payments.TransactOpts, dealId, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutAgreeRefundResourceProvider(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutAgreeRefundResourceProvider", dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutAgreeRefundResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutAgreeRefundResourceProvider is a paid mutator transaction binding the contract method 0x0ef0d89e.
//
// Solidity: function timeoutAgreeRefundResourceProvider(string dealId, address resourceProvider, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutAgreeRefundResourceProvider(dealId string, resourceProvider common.Address, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutAgreeRefundResourceProvider(&_Payments.TransactOpts, dealId, resourceProvider, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutJudgeResults(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutJudgeResults", dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutJudgeResults(dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutJudgeResults is a paid mutator transaction binding the contract method 0x38698529.
//
// Solidity: function timeoutJudgeResults(string dealId, address resourceProvider, address jobCreator, uint256 resultsCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutJudgeResults(dealId string, resourceProvider common.Address, jobCreator common.Address, resultsCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutJudgeResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, resultsCollateral, timeoutCollateral)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactor) TimeoutMediateResult(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutMediateResult", dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsSession) TimeoutMediateResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutMediateResult is a paid mutator transaction binding the contract method 0x823f3de1.
//
// Solidity: function timeoutMediateResult(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 resultsCollateral, uint256 mediationFee) returns()
func (_Payments *PaymentsTransactorSession) TimeoutMediateResult(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, resultsCollateral *big.Int, mediationFee *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutMediateResult(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, resultsCollateral, mediationFee)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactor) TimeoutSubmitResults(opts *bind.TransactOpts, dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "timeoutSubmitResults", dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsSession) TimeoutSubmitResults(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutSubmitResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TimeoutSubmitResults is a paid mutator transaction binding the contract method 0x02fd8f80.
//
// Solidity: function timeoutSubmitResults(string dealId, address resourceProvider, address jobCreator, uint256 paymentCollateral, uint256 timeoutCollateral) returns()
func (_Payments *PaymentsTransactorSession) TimeoutSubmitResults(dealId string, resourceProvider common.Address, jobCreator common.Address, paymentCollateral *big.Int, timeoutCollateral *big.Int) (*types.Transaction, error) {
	return _Payments.Contract.TimeoutSubmitResults(&_Payments.TransactOpts, dealId, resourceProvider, jobCreator, paymentCollateral, timeoutCollateral)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payments.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payments *PaymentsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payments.Contract.TransferOwnership(&_Payments.TransactOpts, newOwner)
}

// PaymentsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Payments contract.
type PaymentsInitializedIterator struct {
	Event *PaymentsInitialized // Event containing the contract specifics and raw log

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
func (it *PaymentsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsInitialized)
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
		it.Event = new(PaymentsInitialized)
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
func (it *PaymentsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsInitialized represents a Initialized event raised by the Payments contract.
type PaymentsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Payments *PaymentsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PaymentsInitializedIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PaymentsInitializedIterator{contract: _Payments.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Payments *PaymentsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PaymentsInitialized) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsInitialized)
				if err := _Payments.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Payments *PaymentsFilterer) ParseInitialized(log types.Log) (*PaymentsInitialized, error) {
	event := new(PaymentsInitialized)
	if err := _Payments.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Payments contract.
type PaymentsOwnershipTransferredIterator struct {
	Event *PaymentsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsOwnershipTransferred)
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
		it.Event = new(PaymentsOwnershipTransferred)
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
func (it *PaymentsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsOwnershipTransferred represents a OwnershipTransferred event raised by the Payments contract.
type PaymentsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payments *PaymentsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payments.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentsOwnershipTransferredIterator{contract: _Payments.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Payments *PaymentsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Payments.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsOwnershipTransferred)
				if err := _Payments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Payments *PaymentsFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentsOwnershipTransferred, error) {
	event := new(PaymentsOwnershipTransferred)
	if err := _Payments.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentsPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the Payments contract.
type PaymentsPaymentIterator struct {
	Event *PaymentsPayment // Event containing the contract specifics and raw log

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
func (it *PaymentsPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentsPayment)
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
		it.Event = new(PaymentsPayment)
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
func (it *PaymentsPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentsPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentsPayment represents a Payment event raised by the Payments contract.
type PaymentsPayment struct {
	DealId    string
	Payee     common.Address
	Amount    *big.Int
	Reason    uint8
	Direction uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) FilterPayment(opts *bind.FilterOpts) (*PaymentsPaymentIterator, error) {

	logs, sub, err := _Payments.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &PaymentsPaymentIterator{contract: _Payments.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *PaymentsPayment) (event.Subscription, error) {

	logs, sub, err := _Payments.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentsPayment)
				if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0x64861f505d0cfce7a0cc3629c70eb54f7de27be35939b48300935694958a9842.
//
// Solidity: event Payment(string dealId, address payee, uint256 amount, uint8 reason, uint8 direction)
func (_Payments *PaymentsFilterer) ParsePayment(log types.Log) (*PaymentsPayment, error) {
	event := new(PaymentsPayment)
	if err := _Payments.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
