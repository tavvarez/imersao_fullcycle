package entity

type Investor struct {
	ID string
	Name string
	AssetsPosition []*InvestorAssetsPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
        ID: id,
		AssetsPosition: []*InvestorAssetsPosition{},
	}	
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetsPosition) {
	i.AssetPosition = append(i.AssetsPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
        i.AssetPosition = append(i.AssetPosition, NewInvestorAssetsPosition(assetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetsPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
            return assetPosition
        }
	}
	return nil
}

type InvestorAssetsPosition struct {
	AssetsID string
	Shares int
}

func NewInvestorAssetsPosition(assetsID string, shares int) *InvestorAssetsPosition {
	return &InvestorAssetsPosition{
			AssetsID: assetsID,
			Shares: shares,
		}
}