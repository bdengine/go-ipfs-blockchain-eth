// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.4.23;


abstract contract SCToken {
    function transferFrom(address src, address dst, uint wad) virtual public  returns (bool);
}


contract IPFS {
    struct Authority {
        address owner;
        uint256 size;
        uint256  expireBlock;
    }

    struct Peer {
        string PeerId;
        string[] AddressList;
        bool valid;
    }



    // 合约拥有者
    address public owner;
    // erc20合约
    SCToken tokenContract;
    // 文件存储单价  SCToken/kb*链块数
    uint256 public price;

    mapping(address => Peer) public addrPeerMap;
    mapping(address => address) public peerList;
    address public head;
    uint public peerNum;
    mapping(string => address) public pidAddrMap;

    modifier PeerExist{
        require(addrPeerMap[msg.sender].valid,"ipfs peer no exist");
        _;
    }

    modifier PeerNotExist{
        require(!addrPeerMap[msg.sender].valid,"ipfs peer already exist");
        _;
    }

    function getPeerList(uint num) public view returns (Peer[] memory) {
        if (num == 0 || num > peerNum) {
            num = peerNum;
        }
        Peer[] memory res = new Peer[](num);
        address  temp = head;
        for (uint i = 0; i < num; i++) {
            res[i] = addrPeerMap[temp];
            temp = peerList[temp];
        }
        return res;
    }

    function getPeerByPid(string calldata pid) public view returns (Peer memory) {
        return addrPeerMap[pidAddrMap[pid]];
    }


    // 更新地址
    function updateAddress(string calldata uid, string[] memory addressList) PeerExist  external {
        addrPeerMap[msg.sender].AddressList = addressList;
        // 改变peer在peerList的位置
        heartbeat();
        emit Success(uid);
    }

    // 心跳 更新节点顺序活跃度，越活越的节点在列表的位置越靠前
    function heartbeat()PeerExist private {
        if (head == address(0)){
            head = msg.sender;
            return;
        }
        if (msg.sender == head) {
            return;
        }
        address pre = head;
        for(uint i = 1;i < peerNum;i++){
            address  temp = peerList[pre];
            if (temp == msg.sender) {
                peerList[pre] = peerList[temp];
                peerList[temp] = head;
                head = temp;
                break;
            }
            pre = temp;
        }
    }

    // 外部心跳调用
    function peerHeartbeat(string calldata uid)PeerExist external{
        heartbeat();
        emit Success(uid);
    }

    function addPeer(string calldata uid,Peer memory _peer) PeerNotExist external {
        _peer.valid = true;
        peerList[msg.sender] = head;
        head = msg.sender;

        addrPeerMap[msg.sender] = _peer;
        pidAddrMap[_peer.PeerId] = msg.sender;
        peerNum++;

        emit Success(uid);
    }

    function removePeer(string calldata uid) PeerExist external {
        if (msg.sender == head) {
            head = address(0);
            return;
        }
        address pre = head;
        for(uint i = 1;i < peerNum;i++){
            address  temp = peerList[pre];
            if (temp == msg.sender) {
                peerList[pre] = peerList[temp];
                break;
            }
            pre = temp;
        }

        peerNum--;
        delete pidAddrMap[addrPeerMap[msg.sender].PeerId];
        delete addrPeerMap[msg.sender];
        delete peerList[msg.sender];
        emit Success(uid);
    }



    mapping(string => Authority) public fileMap;

    function addFile(string calldata uid, string calldata cid, uint256 size, uint256 wad) IsFileOwner(cid) external payable {
        require(wad > 0,"require wad > 0");

        bool pay = tokenContract.transferFrom(msg.sender,address(this),wad);
        require(pay,"pay failed");


        uint256 bNum =  block.number;
        if (fileMap[cid].expireBlock > block.number){
            bNum = fileMap[cid].expireBlock;
        }

        if (fileMap[cid].size > 0) {
            size = fileMap[cid].size;
        }

        uint256  _expire = bNum+wad/(size*price);
        fileMap[cid] = Authority(msg.sender,size,_expire);

        emit Success(uid);
    }

    function removeFile(string calldata uid, string calldata cid) IsFileOwner(cid) external payable {
        if ((fileMap[cid].expireBlock - block.number) > 0) {
            uint256 wad = (fileMap[cid].expireBlock - block.number)*fileMap[cid].size*price;
            bool pay = tokenContract.transferFrom(address(this),msg.sender,wad);
            require(pay,"pay failed");
        }

        fileMap[cid].expireBlock = block.number;
        // TODO 是否应该设置此项？
        fileMap[cid].owner = address(0);

        emit Success(uid);
    }




    event Success(string indexed uid);


    modifier IsOwner(){
        require(msg.sender == owner);
        _;
    }

    modifier IsFileOwner(string calldata _cid){
        require(fileMap[_cid].owner==address(0) || fileMap[_cid].owner == msg.sender,"file is already exist and you are not owner");
        _;
    }

    constructor(uint256 _price,address _token){
        owner = msg.sender;
        price = _price;
        tokenContract = SCToken(_token);
    }

    function SetOwner(address _to) IsOwner public {
        owner = _to;
    }

    function SetToken(address _token) IsOwner public {
        tokenContract = SCToken(_token);
    }

    function SetPrice(uint256 _price) IsOwner public {
        price = _price;
    }

    function withdraw(string calldata uid,uint wad) IsOwner external {
        bool pay = tokenContract.transferFrom(address(this),owner,wad);
        require(pay,"pay failed");
        emit Success(uid);
    }
}