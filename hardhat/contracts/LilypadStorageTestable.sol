// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.6;

import "./CoopHiveStorage.sol";

contract CoopHiveStorageTestable is CoopHiveStorage {
  function _checkControllerAccess() internal pure override returns (bool) {
    return true;
  }
}
