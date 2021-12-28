// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.0;


abstract contract SCToken {
    function transferFrom(address src, address dst, uint wad) virtual public returns (bool);
}


contract IPFS {
    struct Authority {
        // 文件拥有者,文件存储付费者
        address owner;
        uint256 size;
        uint256 expireBlock;
        // 上传文件的节点的链上身份,可以和owner相同
        address uploader;
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

    mapping(string => Authority) public fileMap;


    event Success(string indexed uid);


    modifier IsOwner(){
        require(msg.sender == owner);
        _;
    }

    modifier IsFileOwner(string calldata _cid){
        require(fileMap[_cid].owner == address(0) || fileMap[_cid].owner == msg.sender, "file is already exist and you are not owner");
        _;
    }

    modifier FileNotExist(string calldata _cid){
        require(fileMap[_cid].owner == address(0), "file is already exist");
        _;
    }

    modifier FileExist(string calldata _cid){
        require(fileMap[_cid].owner != address(0), "file not exist");
        _;
    }

    modifier PeerExist{
        require(addrPeerMap[msg.sender].valid, "ipfs peer no exist");
        _;
    }

    modifier PeerNotExist{
        require(!addrPeerMap[msg.sender].valid, "ipfs peer already exist");
        _;
    }


    constructor(uint256 _price, address _token){
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

    function withdraw(string calldata uid, uint wad) IsOwner external {
        bool pay = tokenContract.transferFrom(address(this), owner, wad);
        require(pay, "pay failed");
        emit Success(uid);
    }

    function getPeerList(uint num) public view returns (Peer[] memory) {
        if (num == 0 || num > peerNum) {
            num = peerNum;
        }
        Peer[] memory res = new Peer[](num);
        address temp = head;
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
    function updateAddress(string calldata uid, string[] memory addressList) PeerExist external {
        addrPeerMap[msg.sender].AddressList = addressList;
        // 改变peer在peerList的位置
        heartbeat();
        emit Success(uid);
    }

    // 心跳 更新节点顺序活跃度，越活越的节点在列表的位置越靠前
    function heartbeat() PeerExist private {
        if (head == address(0)) {
            head = msg.sender;
            return;
        }
        if (msg.sender == head) {
            return;
        }
        address pre = head;
        for (uint i = 1; i < peerNum; i++) {
            address temp = peerList[pre];
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
    function peerHeartbeat(string calldata uid) PeerExist external {
        heartbeat();
        emit Success(uid);
    }

    function addPeer(string calldata uid, string calldata peerId, string[] calldata addressList) PeerNotExist external {
        peerList[msg.sender] = head;
        head = msg.sender;

        addrPeerMap[msg.sender] = Peer(peerId, addressList, true);
        pidAddrMap[peerId] = msg.sender;
        peerNum++;

        emit Success(uid);
    }

    function removePeer(string calldata uid) PeerExist external {
        if (msg.sender == head) {
            head = peerList[head];
        } else {
            address pre = head;
            for (uint i = 1; i < peerNum; i++) {
                address temp = peerList[pre];
                if (temp == msg.sender) {
                    peerList[pre] = peerList[temp];
                    break;
                }
                pre = temp;
            }
        }

        peerNum--;
        delete pidAddrMap[addrPeerMap[msg.sender].PeerId];
        delete addrPeerMap[msg.sender];
        delete peerList[msg.sender];
        emit Success(uid);
    }

    function addFile(string calldata uid, string calldata cid, uint256 size, uint256 blockNum, address _owner) FileNotExist(cid) external {
        require(blockNum > 0, "require blockNum > 0");
        require(size > 0, "");
        uint256 wad = blockNum * price * size;

        // 上传节点支付给本合约
        bool pay = tokenContract.transferFrom(msg.sender, address(this), wad);
        require(pay, "pay failed");

        uint256 _expire = block.number + blockNum;
        fileMap[cid] = Authority(_owner, size, _expire, msg.sender);

        emit Success(uid);
    }

    function rechargeFile(string calldata uid, string calldata cid, uint256 blockNum) FileExist(cid) external {
        require(blockNum > 0, "require blockNum > 0");
        uint256 wad = price * fileMap[cid].size * blockNum;
        bool pay = tokenContract.transferFrom(msg.sender, address(this), wad);
        require(pay, "pay failed");

        if (fileMap[cid].expireBlock > block.number) {
            fileMap[cid].expireBlock += blockNum;
        } else {
            fileMap[cid].expireBlock = blockNum + block.number;
        }
        emit Success(uid);
    }

    function removeFile(string calldata uid, string calldata cid) IsFileOwner(cid) external {
        if (fileMap[cid].expireBlock > block.number) {
            uint256 wad = (fileMap[cid].expireBlock - block.number) * fileMap[cid].size * price;
            bool pay = tokenContract.transferFrom(address(this), msg.sender, wad);
            require(pay, "pay failed");
        }

        fileMap[cid].expireBlock = block.number;
        // TODO 是否应该设置此项？暂时保留此项，以避免文件删除后无法再次添加
        fileMap[cid].owner = address(0);

        emit Success(uid);
    }
}
