// SPDX-License-Identifier: MIT
pragma solidity >=0.8.12;

contract MySmartContract {
    function Hello() public pure returns (string memory) {
        return "Hello World";
    }
    function Greet(string memory str) public pure returns (string memory) {
        return str;
    }
}