package character

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
	"log"
	"strings"
)

type GearData struct {
	Items []*struct {
		Name       string `pagser:".db-tooltip__item__name"`
		Category   string `pagser:".db-tooltip__item__category"`
		ILevel     int    `pagser:".db-tooltip__item__level->ILevelParse()"`
		EquipClass string `pagser:".db-tooltip__item_equipment__class"`
		EquipLevel int    `pagser:".db-tooltip__item_equipment__level->EquipLevelParse()"`
		Stats      []*struct {
			Name  string `pagser:""` // Is getting fixed by postParse()
			Value string `pagser:""` // Is getting fixed by postParse()
		} `pagser:".db-tooltip__basic_bonus li"`
		Materia []*struct {
			Name  string `pagser:".db-tooltip__materia__txt->GetNameFromMateriaBlob()"`
			Stats []*struct {
				Name  string `pagser:""` // Is getting fixed by postParse()
				Value string `pagser:""` // Is getting fixed by postParse()
			} `pagser:".db-tooltip__materia__txt--base"`
		} `pagser:".db-tooltip__materia li"`
	} `pagser:".character__view .db-tooltip.db-tooltip__wrapper.item_detail_box"`
	// find('.db-tooltip.db-tooltip__wrapper.item_detail_box')
}

func (g GearData) GetNameFromMateriaBlob(node *goquery.Selection, args ...string) (out interface{}, err error) {
	html, _ := node.Html()
	name := strings.Split(html, "<br/>")[0]
	//log.Println(node.Find(".db-tooltip__materia__txt--base").Text())
	return name, nil
}

func (g GearData) GetStatFromMateriaBlob(node *goquery.Selection, args ...string) (out interface{}, err error) {
	html, _ := node.Html()
	name := strings.Split(html, "<br/>")[0]
	//log.Println(node.Find(".db-tooltip__materia__txt--base").Text())
	return name, nil
}

func (g GearData) ILevelParse(node *goquery.Selection, args ...string) (out interface{}, err error) {
	iLevel := strings.Split(node.Text(), " ")[len(strings.Split(node.Text(), " "))-1]
	return iLevel, nil
}

func (g GearData) EquipLevelParse(node *goquery.Selection, args ...string) (out interface{}, err error) {
	equipLevel := strings.Split(node.Text(), " ")[1]
	return equipLevel, nil
}

func (g GearData) StatsParse(node *goquery.Selection, args ...string) (out interface{}, err error) {
	node.Each(func(i int, selection *goquery.Selection) {
		log.Println(selection.Text())
	})
	return node, nil
}

func (g *GearData) parseGearData(data string) error {
	if data == "" {
		return errors.New("no input given")
	}
	// New default config
	p := pagser.New()
	if err := p.Parse(g, data); err != nil {
		return err
	}
	g.postParse()

	// Print data
	// log.Printf("Page data json: \n-------------\n%v\n-------------\n", func(v interface{}) string {
	// 	data, _ := json.MarshalIndent(v, "", "\t")
	// 	return string(data)
	// }(g))
	return nil
}

func (g *GearData) postParse() {
	for _, item := range g.Items {
		for _, stat := range item.Stats {
			list := strings.Split(stat.Name, " +")
			stat.Name = list[0]
			stat.Value = list[1]
		}
		for _, materia := range item.Materia {
			for _, stat := range materia.Stats {
				list := strings.Split(stat.Name, " +")
				stat.Name = list[0]
				stat.Value = list[1]
			}
		}
	}
}
