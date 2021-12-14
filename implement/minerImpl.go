package implement

import "github.com/ipfs/go-ipfs-auth/standard/model"

func (a *peerImpl) GetChallenge() (string, error) {
	cli, contr, err := GetSCTokenContract(a.Client.HttpUrl, a.ContractMap[contractToken].ContractAddr)
	if err != nil {
		return "", err
	}
	defer cli.Close()
	challenge, _, err := contr.GetChallenge(nil)
	return challenge, err
}

func (a *peerImpl) Mining([]model.IpfsMining) error {
	return nil
}
