package implement

import (
	"github.com/ipfs/go-ipfs-auth/standard/model"
	"math/big"
)

func (a *peerImpl) GetPeerList(num int) ([]model.CorePeer, error) {
	list, err := a.ipfsContract.GetPeerList(nil, big.NewInt(int64(num)))
	if err != nil {
		return nil, err
	}
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
	peer, err := a.ipfsContract.GetPeerByPid(nil, id)
	if err != nil {
		return res, err
	}
	res.Addresses = peer.AddressList
	return res, nil
}
