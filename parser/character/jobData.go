package character

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/foolin/pagser"
)

type JobData struct {
	Tanks []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->eq(0)"`

	Healers []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->eq(2)"`

	Melee []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->eq(1)"`

	PhysicalRange []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->PhysicalRangeSelector()"`

	MagicalRange []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->MagicalRangeSelector()"`

	Gatherer []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->eq(4)"`

	Crafter []struct {
		Job []struct {
			Name       string   `pagser:".character__job__name->text()"`
			Level      int      `pagser:".character__job__level->text()"`
			Experience []string `pagser:".character__job__exp->textSplit('/', true)"`
		} `pagser:".character__job li"`
	} `pagser:".character__job__role->eq(5)"`
}

// Todo: Make one function out of those Range selectors
func (j JobData) PhysicalRangeSelector(node *goquery.Selection, args ...string) (out interface{}, err error) {
	var physicalRange *goquery.Selection
	node.Siblings().Each(func(i int, selection *goquery.Selection) {
		if i == 2 {
			selection.Find(".character__job").Each(func(i int, selection *goquery.Selection) {
				if i == 0 {
					physicalRange = selection
				}
			})
		}
	})
	return physicalRange, nil
}

func (j JobData) MagicalRangeSelector(node *goquery.Selection, args ...string) (out interface{}, err error) {
	var magicalRange *goquery.Selection
	node.Siblings().Each(func(i int, selection *goquery.Selection) {
		if i == 2 {
			selection.Find(".character__job").Each(func(i int, selection *goquery.Selection) {
				if i == 1 {
					magicalRange = selection
				}
			})
		}
	})
	return magicalRange, nil
}

func (j *JobData) parseJobData(data string) error {
	if data == "" {
		return errors.New("no input given")
	}

	p := pagser.New()
	if err := p.Parse(j, data); err != nil {
		return err
	}

	// Print data
	// log.Printf("Page data json: \n-------------\n%v\n-------------\n", func(v interface{}) string {
	// 	data, _ := json.MarshalIndent(v, "", "\t")
	// 	return string(data)
	// }(j))
	return nil
}
