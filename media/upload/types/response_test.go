package types

import (
	"testing"

	"github.com/928799934/go-twitter/resources"
	"github.com/stretchr/testify/assert"
)

func Test_InitializeOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *InitializeOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &InitializeOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &InitializeOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &InitializeOutput{},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}

func Test_AppendOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *AppendOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &AppendOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &AppendOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &AppendOutput{},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}

func Test_FinalizeOutput_HasPartialError(t *testing.T) {
	var errorTitle string = "test partial error"
	cases := []struct {
		name   string
		res    *FinalizeOutput
		expect bool
	}{
		{
			name: "has partial error",
			res: &FinalizeOutput{
				Errors: []resources.PartialError{
					{Title: &errorTitle},
				},
			},
			expect: true,
		},
		{
			name: "has no partial error",
			res: &FinalizeOutput{
				Errors: []resources.PartialError{},
			},
			expect: false,
		},
		{
			name:   "partial error is nil",
			res:    &FinalizeOutput{},
			expect: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			hpe := c.res.HasPartialError()
			assert.Equal(tt, c.expect, hpe)
		})
	}
}
