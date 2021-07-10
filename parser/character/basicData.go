package character

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
	"strings"
)

type BasicData struct {
	Name        string `pagser:".frame__chara__name"`
	Title       string `pagser:".frame__chara__title"`
	Server      string `pagser:".frame__chara__world"`
	Datacenter  string `pagser:".frame__chara__world"`
	PortraitUrl string `pagser:".frame__chara__face img->attr(src)"`
	AvatarUrl   string `pagser:".character__detail__image a img->attr(src)"`
	Bio         string `pagser:".character__selfintroduction"`
	Attributes  []struct {
		Name  string `pagser:"span"`
		Value string `pagser:"td"`
	} `pagser:".character__param__list tr"`
}

func (b BasicData) GetServerFromWorld(node *goquery.Selection, args ...string) (out interface{}, err error) {
	worldArray := strings.Fields(node.Text())
	return worldArray[0], nil
}

func (b BasicData) GetDatacenterFromWorld(node *goquery.Selection, args ...string) (out interface{}, err error) {
	worldArray := strings.Fields(node.Text())
	datacenter := worldArray[1]
	datacenter = strings.Replace(datacenter, "(", "", 1)
	datacenter = strings.Replace(datacenter, ")", "", 1)
	return datacenter, nil
}

func (b *BasicData) parseBasicData(data string) error {
	if data == "" {
		return errors.New("no input given")
	}
	// New default config
	p := pagser.New()
	if err := p.Parse(b, data); err != nil {
		return err
	}

	// Print data
	// log.Printf("Page data json: \n-------------\n%v\n-------------\n", func(v interface{}) string {
	// 	data, _ := json.MarshalIndent(v, "", "\t")
	// 	return string(data)
	// }(b))
	return nil
}
