package implement

import (
	"github.com/ipfs/go-ipfs-auth/standard/model"
)

func (a *peerImpl) ApplyLocal(cid string) error {
	return nil
}

func (a *peerImpl) ApplyRemote(cid string) error {
	return nil
}

func (a *peerImpl) GetFile(cid string) (*model.IpfsFile, error) {
	file, err := a.ipfsContract.GetFile(nil, cid)
	if err != nil {
		return nil, err
	}
	ipfsFile := model.IpfsFile{
		Cid:         cid,
		Owner:       file.Owner.String(),
		Size:        file.Size.Uint64(),
		ExpireBlock: file.ExpireBlock.Uint64(),
	}
	return &ipfsFile, nil
}
