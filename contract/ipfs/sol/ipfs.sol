// SPDX-License-Identifier: GNU-3
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

pragma solidity >=0.8.0;

import "./math.sol";
import "./auth.sol";

contract ipfs is DSMath, DSAuth {
    bool                                              public  stopped;
    uint256                                           public  totalSupply;
    mapping(address => uint256)                      public  balanceOf;
    mapping(address => mapping(address => uint256)) public  allowance;
    string                                            public  symbol;
    uint8                                             public  decimals = 18; // standard token precision. override to customize
    string                                            public  name = "";     // Optional token name

    string                                            private  challenge;
    string                                            private  storeChallenge;
    string                                            private  winner;

    uint256                                           private  challengeStartBlock;
    uint256                                           private  challengeCyclePeriod;

    uint256                                           private  storeSeq;
    uint256                                           private  challengeSeq;
    uint256                                           private  rewardSeq;

    uint256                                           private  ipfsMintNum;


    function resetChallengeRule(uint256 _startBlock, uint256 _cyclePeriod) public auth {
        require(_cyclePeriod > 10, "require cyclePeriod > 10");
        if (_startBlock > block.number) {
            challengeStartBlock = _startBlock;
        } else {
            challengeStartBlock = block.number;
        }

        if (_cyclePeriod > 0) {
            challengeCyclePeriod = _cyclePeriod;
        }
    }

    // 0:gen 1:collect 2:announce & check 3:update 4:wait
    function getChallengeStage() public view returns (uint256, string memory, uint256) {
        if (challengeStartBlock <= 0 || challengeCyclePeriod == 0) {
            return (4, "", 0);
            // 未设置起始block 激励尚未开始
        }

        uint256 t1 = (block.number - challengeStartBlock) % (challengeCyclePeriod * 24 * 6);
        if (t1 / (6 * challengeCyclePeriod) < 1) {
            // 验证storeChallenge是否有效
            if (storeSeq * (challengeCyclePeriod * 24 * 6) > (block.number - challengeStartBlock)) {
                return (0, storeChallenge, 6 * challengeCyclePeriod - t1 - 1);
            }
            return (0, "", 6 * challengeCyclePeriod - t1 - 1);
            // gen 阶段，返回 0，存储挑战值
        }
        uint256 temp = t1 % (6 * challengeCyclePeriod);
        uint256 t2 = temp / challengeCyclePeriod;
        if (t2 < 4) {
            // 检查激励值是否有效
            if ((storeSeq + challengeSeq) * (challengeCyclePeriod * 6) > (block.number - challengeStartBlock)) {
                return (1, challenge, 4 * challengeCyclePeriod - temp - 1);
            }
            return (1, "", 4 * challengeCyclePeriod - temp - 1);
            // collect 阶段，返回 1，激励挑战值
        }
        if (t2 < 5) {
            // 检查获奖者是否有效
            if ((rewardSeq + storeSeq) * (challengeCyclePeriod * 6) > (block.number - challengeStartBlock)) {
                return (2, winner, 5 * challengeCyclePeriod - temp - 1);
            }
            return (2, "", 5 * challengeCyclePeriod - temp - 1);
            // announce 阶段，返回 2，获奖者pid
        }
        // 检查激励值是否有效
        if (storeSeq * (challengeCyclePeriod * 24 * 6) > (block.number - challengeStartBlock)) {
            return (3, storeChallenge, 6 * challengeCyclePeriod - temp - 1);
        }
        return (3, "", 6 * challengeCyclePeriod - temp - 1);
        // update 阶段，返回 1，激励挑战值
    }

    function setStoreChallenge(string calldata _storeChallenge) external auth {
        uint256 stage = 4;
        string memory s = "";
        (stage, s,) = getChallengeStage();
        require(stage == 0 && bytes(s).length == 0, "storeChallenge for this period already exist");
        storeChallenge = _storeChallenge;
        storeSeq = (block.number - challengeStartBlock) / (challengeCyclePeriod * 24 * 6) + 1;
        emit SetChallenge(storeChallenge, storeSeq);
    }

    function setChallenge(string calldata _challenge) external auth storeChallengeExist {
        uint256 stage = 4;
        string memory s = "";
        (stage, s,) = getChallengeStage();
        require(stage == 1, "not a proper stage to set challenge");
        require(bytes(s).length == 0, "challenge for this period already exist");
        challenge = _challenge;
        challengeSeq = (block.number - challengeStartBlock) / (challengeCyclePeriod * 6) - storeSeq + 1;
        emit SetChallenge(challenge, challengeSeq);
    }

    function setWinner(string calldata _winner) external auth storeChallengeExist challengeExist {
        uint256 stage = 4;
        string memory s = "";
        (stage, s,) = getChallengeStage();
        require(stage == 2, "not a proper stage to set winner");
        require(bytes(s).length == 0, "winner for this period already exist");
        winner = _winner;
        rewardSeq = (block.number - challengeStartBlock) / (challengeCyclePeriod * 6) - storeSeq + 1;
        emit SetWinner(winner, rewardSeq);
    }

    function getStoreChallenge() public view storeChallengeExist returns (string memory, uint256)  {
        return (storeChallenge, storeSeq * (challengeCyclePeriod * 24 * 6) - (block.number - challengeStartBlock));
    }

    function getChallenge() public view challengeExist returns (string memory, uint256) {
        return (challenge, (storeSeq + challengeSeq) * (challengeCyclePeriod * 6) - (block.number - challengeStartBlock));
    }

    function getWinner() public view winnerExist returns (string memory, uint256) {
        return (winner, (rewardSeq + storeSeq) * (challengeCyclePeriod * 6) - (block.number - challengeStartBlock));
    }



    modifier storeChallengeExist {
        require(storeSeq * (challengeCyclePeriod * 24 * 6) > (block.number - challengeStartBlock), "no proper storeChallenge exist");
        _;
    }

    modifier challengeExist {
        require((storeSeq + challengeSeq) * (challengeCyclePeriod * 6) > (block.number - challengeStartBlock), "no proper challenge exist");
        _;
    }

    modifier winnerExist {
        require((rewardSeq + storeSeq) * (challengeCyclePeriod * 6) > (block.number - challengeStartBlock), "no proper winner exist");
        _;
    }





    constructor(string memory symbol_, uint256 _price) {
        symbol = symbol_;
        challengeCyclePeriod = 10;
        balanceOf[msg.sender] = 10 ** 27;
        ipfsMintNum = 10 ** 20;
        price = _price;
    }

    event Approval(address indexed src, address indexed guy, uint wad);
    event Transfer(address indexed src, address indexed dst, uint wad);
    event Mint(address indexed guy, uint wad);
    event Burn(address indexed guy, uint wad);
    event Stop();
    event Start();
    event SetChallenge(string cge, uint256 cgeSeq);
    event SetStoreChallenge(string cge, uint256 cgeSeq);
    event SetWinner(string cge, uint256 cgeSeq);

    modifier stoppable {
        require(!stopped, "ds-stop-is-stopped");
        _;
    }

    /*function approveSelf(address guy) external returns (bool) {
        return approve(guy, uint(-1));
    }*/

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
        if (src != msg.sender /*&& allowance[src][msg.sender] != uint(-1)*/) {
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
        if (guy != msg.sender /*&& allowance[guy][msg.sender] != uint(-1)*/) {
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

    function setIpfsMintNum(uint256 num) public auth {
        ipfsMintNum = num;
    }

    function getIpfsMintNum() public view returns (uint256) {
        return ipfsMintNum;
    }

    function setName(string memory name_) public auth {
        name = name_;
    }


    struct File {
        // 文件拥有者,文件存储付费者
        address owner;
        uint256 sliceNum;
        uint256 size;
        uint256 expireBlock;
        // 上传文件的节点的链上身份,可以和owner相同
        address uploader;
        uint256 id;
    }

    struct Peer {
        string PeerId;
        string[] AddressList;
        address peerAddress;
        bool valid;
    }

    struct PeerStore {
        mapping(address => Peer) addrPeerMap;
        mapping(address => address) peerList;
        address head;
        uint peerNum;
        mapping(string => address) pidAddrMap;
    }

    struct FileStore {
        mapping(string => File) fileMap;
        string[] fileList;
        uint totalSize;
        uint totalSlice;
    }


    // 文件存储单价  SCToken/kb*链块数
    uint256 public price;

    /*mapping(address => Peer) public addrPeerMap;
    mapping(address => address) public peerList;
    address public head;
    uint public peerNum;
    mapping(string => address) public pidAddrMap;*/

    PeerStore  public peerStore;
    FileStore  private fileStore;


    event Success(string indexed uid);

    modifier IsFileOwner(string memory _cid){
        require(fileStore.fileMap[_cid].owner == address(0) || fileStore.fileMap[_cid].owner == msg.sender, "file is already exist and you are not owner");
        _;
    }

    modifier FileNotExist(string memory _cid){
        require(fileStore.fileMap[_cid].owner == address(0), "file is already exist");
        _;
    }

    modifier FileExist(string memory _cid){
        require(fileStore.fileMap[_cid].owner != address(0), "file not exist");
        _;
    }

    modifier PeerExist{
        require(peerStore.addrPeerMap[msg.sender].valid, "ipfs peer no exist");
        _;
    }

    modifier PeerNotExist{
        require(!peerStore.addrPeerMap[msg.sender].valid, "ipfs peer already exist");
        _;
    }

    function SetPrice(uint256 _price) auth public {
        price = _price;
    }

    function withdraw(string calldata uid, uint wad) auth external {
        bool pay = transferFrom(address(this), owner, wad);
        require(pay, "pay failed");
        emit Success(uid);
    }

    function getPeerList(uint num) public view returns (Peer[] memory) {
        if (num == 0 || num > peerStore.peerNum) {
            num = peerStore.peerNum;
        }
        Peer[] memory res = new Peer[](num);
        address temp = peerStore.head;
        for (uint i = 0; i < num; i++) {
            res[i] = peerStore.addrPeerMap[temp];
            temp = peerStore.peerList[temp];
        }
        return res;
    }

    function getPeerByPid(string memory pid) public view returns (Peer memory) {
        return peerStore.addrPeerMap[peerStore.pidAddrMap[pid]];
    }

    // 更新地址
    function updateAddress(string calldata uid, string[] memory addressList) PeerExist external {
        peerStore.addrPeerMap[msg.sender].AddressList = addressList;
        // 改变peer在peerList的位置
        heartbeat();
        emit Success(uid);
    }

    // 心跳 更新节点顺序活跃度，越活越的节点在列表的位置越靠前
    function heartbeat() PeerExist private {
        if (peerStore.head == address(0)) {
            peerStore.head = msg.sender;
            return;
        }
        if (msg.sender == peerStore.head) {
            return;
        }
        address pre = peerStore.head;
        for (uint i = 1; i < peerStore.peerNum; i++) {
            address temp = peerStore.peerList[pre];
            if (temp == msg.sender) {
                peerStore.peerList[pre] = peerStore.peerList[temp];
                peerStore.peerList[temp] = peerStore.head;
                peerStore.head = temp;
                break;
            }
            pre = temp;
        }
    }

    // 外部心跳调用
    function peerHeartbeat(string calldata uid) PeerExist external {
        heartbeat();
        emit Success(uid);
    }

    function addPeer(string calldata uid, string calldata peerId, string[] memory addressList) PeerNotExist external {
        peerStore.peerList[msg.sender] = peerStore.head;
        peerStore.head = msg.sender;

        peerStore.addrPeerMap[msg.sender] = Peer(peerId, addressList, msg.sender, true);
        peerStore.pidAddrMap[peerId] = msg.sender;
        peerStore.peerNum++;

        emit Success(uid);
    }

    function removePeer(string calldata uid) PeerExist external {
        if (msg.sender == peerStore.head) {
            peerStore.head = peerStore.peerList[peerStore.head];
        } else {
            address pre = peerStore.head;
            for (uint i = 1; i < peerStore.peerNum; i++) {
                address temp = peerStore.peerList[pre];
                if (temp == msg.sender) {
                    peerStore.peerList[pre] = peerStore.peerList[temp];
                    break;
                }
                pre = temp;
            }
        }

        peerStore.peerNum--;
        delete peerStore.pidAddrMap[peerStore.addrPeerMap[msg.sender].PeerId];
        delete peerStore.addrPeerMap[msg.sender];
        delete peerStore.peerList[msg.sender];
        emit Success(uid);
    }

    function addFile(string calldata uid, string calldata cid, uint256 size, uint256 sliceNum, uint256 blockNum, address _owner) FileNotExist(cid) external {
        require(blockNum > 0, "require blockNum > 0");
        require(size > 0, "require size > 0");
        uint256 wad = blockNum * price * size;

        // 上传节点支付给本合约
        bool pay = transferFrom(msg.sender, address(this), wad);
        require(pay, "pay failed");

        uint256 _expire = block.number + blockNum;

        fileStore.fileMap[cid] = File(_owner, sliceNum, size, _expire, msg.sender, fileStore.fileList.length);
        fileStore.fileList.push(cid);

        fileStore.totalSize += size;
        fileStore.totalSlice += sliceNum;

        emit Success(uid);
    }

    function rechargeFile(string calldata uid, string calldata cid, uint256 blockNum) FileExist(cid) external {
        require(blockNum > 0, "require blockNum > 0");
        uint256 wad = price * fileStore.fileMap[cid].size * blockNum;
        bool pay = transferFrom(msg.sender, address(this), wad);
        require(pay, "pay failed");

        if (fileStore.fileMap[cid].expireBlock > block.number) {
            fileStore.fileMap[cid].expireBlock += blockNum;
        } else {
            fileStore.fileMap[cid].expireBlock = blockNum + block.number;
        }
        emit Success(uid);
    }

    function removeFile(string calldata uid, string calldata cid) IsFileOwner(cid) external {
        if (fileStore.fileMap[cid].expireBlock > block.number) {
            uint256 wad = (fileStore.fileMap[cid].expireBlock - block.number) * fileStore.fileMap[cid].size * price;
            bool pay = transferFrom(address(this), msg.sender, wad);
            require(pay, "pay failed");
        }

        fileStore.fileMap[cid].expireBlock = block.number;
        // TODO 是否应该设置此项？暂时保留此项，以避免文件删除后无法再次添加
        fileStore.fileMap[cid].owner = address(0);

        fileStore.totalSize -= fileStore.fileMap[cid].size;
        fileStore.totalSlice -= fileStore.fileMap[cid].sliceNum;

        // 列表删除
        uint256 _id = fileStore.fileMap[cid].id;
        uint256 l = fileStore.fileList.length - 1;

        delete fileStore.fileList[_id];
        fileStore.fileMap[fileStore.fileList[l]].id = _id;
        fileStore.fileList[_id] = fileStore.fileList[l];
        fileStore.fileList.pop();



        emit Success(uid);
    }


    function getFile(string calldata cid) view external returns (File memory){
        return fileStore.fileMap[cid];
    }

    function getFileStatistic() view external returns (uint256, uint256){
        uint256 num = fileStore.fileList.length;
        uint256 size = 0;
        uint256 sliceNum = 0;
        for (uint i = 0; i < num; i++) {
            size += fileStore.fileMap[fileStore.fileList[i]].size;
            sliceNum += fileStore.fileMap[fileStore.fileList[i]].sliceNum;
        }
        return (size, sliceNum);
    }

    function getFileList(uint256 num) view external returns (string[] memory){
        if (num == 0 || num > fileStore.fileList.length) {
            num = fileStore.fileList.length;
        }
        string [] memory res = new string[](num);
        for (uint i = 0; i < num; i++) {
            res[i] = fileStore.fileList[i];
        }
        return res;
    }

    function checkFile(string memory cid) view public returns (bool) {
        return fileStore.fileMap[cid].owner == address(0);
    }
}
