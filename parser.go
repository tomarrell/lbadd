package main

import (
	"fmt"
	"log"
	"strings"
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

		fmt.Println("\nstep:", p.step.String())
		fmt.Println("remaining:", p.sql[p.cursor:])
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
			p.step = stepSelectFrom

		case stepSelectFrom:
			val, _ := p.pop()
			if strings.ToLower(val) == "from" {
				p.step = stepSelectTable
				break
			}
			return p.query, fmt.Errorf("at SELECT: expected FROM")

		case stepSelectTable:
			val, _ := p.pop()
			p.query.tableName = val
			return p.query, nil

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
	if p.cursor >= len(p.sql) {
		return "", 0
	}

	// Advance the cursor until we reach the end of the
	// input, or the desired character
	buf := ""
	i := p.cursor
	for {
		// Reached the end
		if i == len(p.sql) {
			fmt.Println("returning buf: ", buf)
			return buf, i - p.cursor
		}

		// Reached our desired character
		if p.sql[i] == ' ' {
			fmt.Println("returning buf: ", buf)
			return buf, i - p.cursor
		}

		buf += string(p.sql[i])
		i++
	}
}

func (p *parser) popWhitespace() {
	for {
		if p.sql[p.cursor] != ' ' {
			break
		}
		p.cursor++
	}
}
