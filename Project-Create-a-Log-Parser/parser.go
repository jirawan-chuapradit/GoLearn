package main

import (
	"fmt"
	"strconv"
	"strings"
)

// result stores the parsed result for a domain
type result struct {
	domain string
	visits int

	metrics map[string]int
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result //total visits per domain
	domains []string          //number of parsed lines
	total   int               //total visits to all domains
	lines   int               //number of parsed lines
	lerr    error
}

//newPasser contructs, initializes and returns a new parser
func newParser() *parser {

	return &parser{sum: make(map[string]result)}
}

// parse parses a log line and returns the parsed result with an error
func parse(p *parser, line string) (r result) {

	if p.lerr != nil {
		return
	}


	p.lines++

	fields := strings.Fields(line)


	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		//return parsed, err
		return
	}

	var err error

	r.domain = fields[0]

	//sum the total visits per domain
	r.visits, err = strconv.Atoi(fields[1])
	if r.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q(line #%d)", fields[1], p.lines)
		//return parsed,err

	}

	//return parsed, err
	return
}

// update updates the parser for given parsing resu lt
func update(p *parser, r result) {
	if p.lerr != nil {
		return
	}

	//collect the unique domains
	if _, ok := p.sum[r.domain]; ok {
		p.domains = append(p.domains, r.domain)
	}

	//keep track of total and per domain visits
	p.total += r.visits

	//create and assign a new copy of `visit`
	p.sum[r.domain] =  result{
		domain: r.domain,
		visits: r.visits + p.sum[r.domain].visits,
	}

}

func err(p *parser) error {
	return p.lerr
}