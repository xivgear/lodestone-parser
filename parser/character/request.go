package character

import (
	"io/ioutil"
	"net/http"
	"sync"
)

var LodestoneUrl = map[string]string{
	"character": "https://na.finalfantasyxiv.com/lodestone/character/",
}

func (c *Character) requestCharacterData() error {
	data, err := request(LodestoneUrl["character"] + c.LodestoneId)
	if err != nil {
		return err
	}
	c.rawCharacterData = data
	return nil
}

func (c *Character) requestJobClassData() error {
	data, err := request(LodestoneUrl["character"] + c.LodestoneId + "/class_job/")
	if err != nil {
		return err
	}
	c.rawClassJobsData = data
	return nil
}

func (c *Character) requestLodestone() error {
	var wg sync.WaitGroup

	var err error
	wg.Add(1)
	go func() {
		err = c.requestCharacterData()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		err = c.requestJobClassData()
		wg.Done()
	}()

	wg.Wait()
	if err != nil {
		return err
	}
	return nil
}

func request(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(html), nil
}
