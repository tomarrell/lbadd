package main

import (
	"fmt"
	"log"
	"strings"
)

type step int

const (
	stepInit step = iota
	stepSelectField
	stepSelectComma
)

func parse(sql string) (query, error) {
	p := parser{
		cursor:          0,
		sql:             sql,
		step:            stepInit,
		query:           query{},
		err:             nil,
		nextUpdateField: "",
	}

	return p.parse()
}

type parser struct {
	cursor          int
	sql             string
	step            step
	query           query
	err             error
	nextUpdateField string
}

func (p *parser) parse() (query, error) {
	q, err := p.doParse()
	if err != nil {
		log.Printf("Err: failed to parse sql: %v", err)
		return query{}, err
	}

	return q, nil
}

func (p *parser) doParse() (query, error) {
	for {
		// Check if we've hit the end of the query
		if p.cursor >= len(p.sql) {
			return p.query, p.err
		}

		switch p.step {
		case stepInit:
			switch strings.ToLower(p.peek()) {
			case selectQuery.String():
				p.query.queryType = selectQuery
				p.step = stepSelectField
				p.pop()
			default:
				return p.query, fmt.Errorf("unrecognised query type")
			}

		case stepSelectField:
			val, _ := p.pop()
			p.query.fields = append(p.query.fields, val)
			p.step = stepSelectComma
		default:
			return p.query, nil
		}
	}
}

// TODO
var reservedWords = []string{
	"(", ")", ">=", "<=", "!=", ",", "=", ">", "<", "SELECT", "INSERT INTO", "VALUES", "UPDATE", "DELETE FROM",
	"WHERE", "FROM", "SET",
}

func (p *parser) peek() string {
	val, _ := p.peekWithCount()
	return val
}

func (p *parser) pop() (string, int) {
	p.popWhitespace()
	val, adv := p.peekWithCount()
	p.cursor += adv

	return val, adv
}

func (p *parser) peekWithCount() (string, int) {
	buf := ""

	i := p.cursor
	for ; p.sql[i] != ' '; i++ {
		buf += string(p.sql[i])
	}

	return buf, i - p.cursor
}

func (p *parser) popWhitespace() string {
	for ; p.sql[p.cursor] == ' '; p.cursor++ {
	}

	return ""
}
