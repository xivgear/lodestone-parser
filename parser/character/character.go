package character

type Character struct {
	LodestoneId      string
	rawCharacterData string
	rawClassJobsData string
	BasicData        BasicData
	JobData          JobData
	GearData         GearData
}

func (c *Character) ParseCharacterData() error {
	// Todo: Make async
	if err := c.requestLodestone(); err != nil {
		return err
	}
	if err := c.BasicData.parseBasicData(c.rawCharacterData); err != nil {
		return err
	}
	if err := c.GearData.parseGearData(c.rawCharacterData); err != nil {
		return err
	}
	if err := c.JobData.parseJobData(c.rawClassJobsData); err != nil {
		return err
	}
	// Print data
	// log.Printf("Page data json: \n-------------\n%v\n-------------\n", func(v interface{}) string {
	// 	data, _ := json.MarshalIndent(v, "", "\t")
	// 	return string(data)
	// }(c))
	return nil
}

func NewCharacter(lodestoneId string) *Character {
	return &Character{
		LodestoneId: lodestoneId,
		BasicData:   BasicData{},
		JobData:     JobData{},
		GearData:    GearData{},
	}
}
