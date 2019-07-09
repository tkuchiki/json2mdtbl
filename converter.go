package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/olekukonko/tablewriter"
)

type Converter struct {
	reader    *bufio.Reader
	writer    *tablewriter.Table
	dec       *json.Decoder
	firstChar string
	keys      []string
	rows      [][]string
}

func NewConverter(r io.Reader, w io.Writer) *Converter {
	converter := &Converter{
		reader: bufio.NewReader(r),
		writer: tablewriter.NewWriter(w),
		keys:   make([]string, 0),
		rows:   make([][]string, 0),
	}

	b, _ := converter.reader.Peek(1)
	converter.dec = json.NewDecoder(converter.reader)

	converter.firstChar = string(b)

	return converter
}

func (c *Converter) setKeys(obj map[string]interface{}) {
	if len(obj) > 0 {
		keys := make([]string, 0)
		for key, _ := range obj {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		c.keys = keys
	}
}

func (c *Converter) generateRow(obj map[string]interface{}) []string {
	row := make([]string, 0)
	for _, key := range c.keys {
		row = append(row, fmt.Sprintf("%v", obj[key]))
	}

	return row
}

func (c *Converter) Read() error {
	if c.firstChar == "{" {
		for {
			var obj map[string]interface{}
			err := c.dec.Decode(&obj)

			if len(obj) > 0 {
				c.setKeys(obj)
				row := c.generateRow(obj)

				c.rows = append(c.rows, row)
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

		}
	} else if c.firstChar == "[" {
		for {
			var obj []map[string]interface{}
			err := c.dec.Decode(&obj)

			for _, o := range obj {
				c.setKeys(o)
				row := c.generateRow(o)

				c.rows = append(c.rows, row)
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

		}
	} else {
		return fmt.Errorf("invalid JSON")
	}

	return nil
}

func (c *Converter) Write() {
	c.writer.SetHeader(c.keys)
	c.writer.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	c.writer.SetCenterSeparator("|")
	c.writer.AppendBulk(c.rows)
	c.writer.Render()
}
