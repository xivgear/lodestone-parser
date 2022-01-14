package parser

import (
	"github.com/xivgear/lodestone-parser/parser/character"
)

type Parser struct {
	Character *character.Character
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseCharacter(id string) error {
	p.Character = character.NewCharacter(id)
	if err := p.Character.ParseCharacterData(); err != nil {
		return err
	}
	return nil
}
