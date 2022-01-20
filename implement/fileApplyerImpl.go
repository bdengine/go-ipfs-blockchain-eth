package implement

import (
	"github.com/bdengine/go-ipfs-blockchain-standard/model"
)

func (p *peerImpl) ApplyLocal(cid string) error {
	return nil
}

func (p *peerImpl) ApplyRemote(cid string) error {
	return nil
}

func (p *peerImpl) GetFile(cid string) (*model.IpfsFile, error) {
	file, err := p.ipfsContract.GetFile(nil, cid)
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
