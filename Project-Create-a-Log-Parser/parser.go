package main

import (
	"fmt"
	"strconv"
	"strings"
)

func newParser() parser {

	return parser{sum:make(map[string]result) }
}

type result struct {
	domain string
	visits int

	metrics map[string]int
}

type parser struct {
	sum     map[string]result //total visits per domain
	domains []string       //number of parsed lines
	total   int            //total visits to all domains
	lines   int            //number of parsed lines
}

func parse(p parser, line string) (parsed result,err error) {

	//var (
	//	parsed result
	//	err error
	//	)

	fields := strings.Fields(line)
	//fmt.Printf("domain: %s - visits: %s\n", fields[0], fields[1])

	if len(fields) != 2 {
		err = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		//return parsed, err
		return
	}

	parsed.domain = fields[0]

	//sum the total visits per domain
	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		err = fmt.Errorf("wrong input: %q(line #%d)", fields[1], p.lines)
		//return parsed,err
		return
	}

	//return parsed, err
	return
}

func update(p parser,parsed result) parser  {
	domain, visits := parsed.domain, parsed.visits


	//collect the unique domains
	if _, ok := p.sum[domain]; ok {
		p.domains = append(p.domains, domain)
	}

	//keep track of total and per domain visits
	p.total += visits



	r := result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
	p.sum[domain] = r
	return p
}