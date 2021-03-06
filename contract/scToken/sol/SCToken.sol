/// token.sol -- ERC20 implementation with minting and burning

// Copyright (C) 2015, 2016, 2017  DappHub, LLC

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

pragma solidity >=0.5.0 <0.5.18;

import "./math.sol";
import "./auth.sol";

contract scToken is DSMath, DSAuth {
    bool                                              public  stopped;
    uint256                                           public  totalSupply;
    mapping (address => uint256)                      public  balanceOf;
    mapping (address => mapping (address => uint256)) public  allowance;
    string                                            public  symbol;
    uint8                                             public  decimals = 18; // standard token precision. override to customize
    string                                            public  name = "";     // Optional token name

    string                                            private  challenge;
    uint256                                           private  challengeSeq;
    uint256                                           private  challengeTime;

    uint256                                           private  ipfsMintNum;

    constructor(string memory symbol_) public {
        symbol = symbol_;
        challengeTime = 600;
        balanceOf[msg.sender] = 10 ** 27;
        ipfsMintNum = 10 ** 20;
    }

    event Approval(address indexed src, address indexed guy, uint wad);
    event Transfer(address indexed src, address indexed dst, uint wad);
    event Mint(address indexed guy, uint wad);
    event Burn(address indexed guy, uint wad);
    event Stop();
    event Start();
    event SetChallenge(string cge, uint256 cgeSeq);

    modifier stoppable {
        require(!stopped, "ds-stop-is-stopped");
        _;
    }

    modifier hasChallenge {
        require(block.number > challengeSeq * challengeTime, "not has challenge");
        _;
    }

    function setChallenge(string calldata challenge_) external  hasChallenge auth {
        challenge = challenge_;
        challengeSeq = block.number / challengeTime + 1;
        emit SetChallenge(challenge, challengeSeq);
    }

    function getChallenge() public view returns (string memory, uint256) {
        if (challengeSeq > 0 && block.number > (challengeSeq - 1) * challengeTime && block.number <= challengeSeq * challengeTime){
            return (challenge, challengeSeq);
        }
        return ("",challengeSeq);

    }

    function approveSelf(address guy) external returns (bool) {
        return approve(guy, uint(-1));
    }

    function approve(address guy, uint wad) public stoppable returns (bool) {
        allowance[msg.sender][guy] = wad;

        emit Approval(msg.sender, guy, wad);

        return true;
    }

    function transfer(address dst, uint wad) external returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    function transferFrom(address src, address dst, uint wad)
    public
    stoppable
    returns (bool)
    {
        if (src != msg.sender && allowance[src][msg.sender] != uint(-1)) {
            require(allowance[src][msg.sender] >= wad, "ds-token-insufficient-approval");
            allowance[src][msg.sender] = sub(allowance[src][msg.sender], wad);
        }

        require(balanceOf[src] >= wad, "ds-token-insufficient-balance");
        balanceOf[src] = sub(balanceOf[src], wad);
        balanceOf[dst] = add(balanceOf[dst], wad);

        emit Transfer(src, dst, wad);

        return true;
    }

    function push(address dst, uint wad) external {
        transferFrom(msg.sender, dst, wad);
    }

    function pull(address src, uint wad) external {
        transferFrom(src, msg.sender, wad);
    }

    function move(address src, address dst, uint wad) external {
        transferFrom(src, dst, wad);
    }


    function mintSelf(uint wad) external {
        mint(msg.sender, wad);
    }

    function mintForIpfs(address ipfsNode) external {
        mint(ipfsNode, ipfsMintNum);
    }

    function burnSelf(uint wad) external {
        burn(msg.sender, wad);
    }

    function mint(address guy, uint wad) public auth stoppable {
        balanceOf[guy] = add(balanceOf[guy], wad);
        totalSupply = add(totalSupply, wad);
        emit Mint(guy, wad);
    }

    function burn(address guy, uint wad) public auth stoppable {
        if (guy != msg.sender && allowance[guy][msg.sender] != uint(-1)) {
            require(allowance[guy][msg.sender] >= wad, "ds-token-insufficient-approval");
            allowance[guy][msg.sender] = sub(allowance[guy][msg.sender], wad);
        }

        require(balanceOf[guy] >= wad, "ds-token-insufficient-balance");
        balanceOf[guy] = sub(balanceOf[guy], wad);
        totalSupply = sub(totalSupply, wad);
        emit Burn(guy, wad);
    }

    function stop() public auth {
        stopped = true;
        emit Stop();
    }

    function start() public auth {
        stopped = false;
        emit Start();
    }

    function setChallengeTime(uint256 _cha) public auth {
        challengeTime = _cha;
    }

    function setIpfsMintNum(uint256 num) public auth {
        ipfsMintNum = num;
    }

    function getIpfsMintNum() public view returns (uint256) {
        return ipfsMintNum;
    }

    function setName(string memory name_) public auth {
        name = name_;
    }
}
