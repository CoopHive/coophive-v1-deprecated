// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./CoopHiveToken.sol";

// a version of CoopHiveToken that can be called by any address
// so we can run unit tests
contract CoopHiveTokenTestable is CoopHiveToken {
  constructor(
    string memory name,
    string memory symbol,
    uint256 initialSupply
  ) CoopHiveToken(name, symbol, initialSupply) {}

  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
