package prices

import "errors"

type Article struct {
	Name      string
	CostPrice float64
	Tax       float64
}

func GetSellPrice(idArticle string) (float64, error) {
	article, err := GetArticle(idArticle)
	if err != nil {
		return 0, err
	}
	price, err := CalcPrice(article)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func GetArticle(idArticle string) (Article, error) {

	if idArticle == "ASX123" {
		return Article{
			"Test article",
			500,
			5.0,
		}, nil
	}

	return Article{}, errors.New("article not found")
}

func CalcPrice(art Article) (float64, error) {
	if art.Tax == 0 {
		return 0, errors.New("tax undefined")
	}
	return art.Tax + art.CostPrice, nil
}
