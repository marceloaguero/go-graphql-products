// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Provider struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive,omitempty"`
}

type ProviderInput struct {
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive,omitempty"`
}