// Package models provides models for the application.
package models

// Filter is a filter for a query.
type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}
