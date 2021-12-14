package implement

import (
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"math/big"
)

func (a *peerImpl) GetPeerList() ([]model.CorePeer, error) {
	httpClient, contract, err := GetIpfsContract(a.Client.HttpUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return nil, err
	}
	defer httpClient.Close()

	list, err := contract.GetPeerList(nil, big.NewInt(-1))
	res := make([]model.CorePeer, len(list))
	for i, peer := range list {
		res[i] = model.CorePeer{
			PeerId:    peer.PeerId,
			Addresses: peer.AddressList,
		}
	}

	return res, nil
}

func (a *peerImpl) GetUserCode() (string, error) {
	return a.address.String(), nil
}

func (a *peerImpl) GetPeer(id string) (model.CorePeer, error) {
	res := model.CorePeer{PeerId: id}
	cli, contract, err := GetIpfsContract(a.Client.HttpUrl, a.ContractMap[contractIpfs].ContractAddr)
	if err != nil {
		return res, err
	}
	defer cli.Close()
	peer, err := contract.GetPeerByPid(nil, id)
	if err != nil {
		return res, err
	}
	res.Addresses = peer.AddressList
	return res, nil
}
