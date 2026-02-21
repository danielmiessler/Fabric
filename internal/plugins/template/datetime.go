// Package template provides datetime operations for the template system
package template

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/danielmiessler/fabric/internal/i18n"
)

// DateTimePlugin handles time and date operations
type DateTimePlugin struct{}

// Apply executes datetime operations with the following formats:
// Time: now (RFC3339), time (HH:MM:SS), unix (timestamp)
// Hour: startofhour, endofhour
// Date: today (YYYY-MM-DD), full (Monday, January 2, 2006)
// Period: startofweek, endofweek, startofmonth, endofmonth
// Relative: rel:-1h, rel:-2d, rel:1w, rel:3m, rel:1y
func (p *DateTimePlugin) Apply(operation string, value string) (string, error) {
	debugf("DateTime: operation=%q value=%q", operation, value)

	now := time.Now()
	debugf("DateTime: reference time=%v", now)

	switch operation {
	// Time operations
	case "now":
		result := now.Format(time.RFC3339)
		debugf("DateTime: now=%q", result)
		return result, nil

	case "time":
		result := now.Format("15:04:05")
		debugf("DateTime: time=%q", result)
		return result, nil

	case "unix":
		result := fmt.Sprintf("%d", now.Unix())
		debugf("DateTime: unix=%q", result)
		return result, nil

	case "startofhour":
		result := now.Truncate(time.Hour).Format(time.RFC3339)
		debugf("DateTime: startofhour=%q", result)
		return result, nil

	case "endofhour":
		result := now.Truncate(time.Hour).Add(time.Hour - time.Second).Format(time.RFC3339)
		debugf("DateTime: endofhour=%q", result)
		return result, nil

	// Date operations
	case "today":
		result := now.Format("2006-01-02")
		debugf("DateTime: today=%q", result)
		return result, nil

	case "full":
		result := now.Format("Monday, January 2, 2006")
		debugf("DateTime: full=%q", result)
		return result, nil

	case "month":
		result := now.Format("January")
		debugf("DateTime: month=%q", result)
		return result, nil

	case "year":
		result := now.Format("2006")
		debugf("DateTime: year=%q", result)
		return result, nil

	case "startofweek":
		result := now.AddDate(0, 0, -int(now.Weekday())).Format("2006-01-02")
		debugf("DateTime: startofweek=%q", result)
		return result, nil

	case "endofweek":
		result := now.AddDate(0, 0, 7-int(now.Weekday())).Format("2006-01-02")
		debugf("DateTime: endofweek=%q", result)
		return result, nil

	case "startofmonth":
		result := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		debugf("DateTime: startofmonth=%q", result)
		return result, nil

	case "endofmonth":
		result := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
		debugf("DateTime: endofmonth=%q", result)
		return result, nil

	case "rel":
		return p.handleRelative(now, value)

	default:
		return "", fmt.Errorf(i18n.T("template_datetime_error_unknown_operation"), operation)
	}
}

func (p *DateTimePlugin) handleRelative(now time.Time, value string) (string, error) {
	debugf("DateTime: handling relative time value=%q", value)

	if value == "" {
		return "", errors.New(i18n.T("template_datetime_error_relative_requires_value"))
	}

	// Try standard duration first (hours, minutes)
	if duration, err := time.ParseDuration(value); err == nil {
		result := now.Add(duration).Format(time.RFC3339)
		debugf("DateTime: relative duration=%q result=%q", duration, result)
		return result, nil
	}

	// Handle date units
	if len(value) < 2 {
		return "", errors.New(i18n.T("template_datetime_error_invalid_relative_format"))
	}

	unit := value[len(value)-1:]
	numStr := value[:len(value)-1]

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return "", fmt.Errorf(i18n.T("template_datetime_error_invalid_number"), value)
	}

	var result string
	switch unit {
	case "d":
		result = now.AddDate(0, 0, num).Format("2006-01-02")
	case "w":
		result = now.AddDate(0, 0, num*7).Format("2006-01-02")
	case "m":
		result = now.AddDate(0, num, 0).Format("2006-01-02")
	case "y":
		result = now.AddDate(num, 0, 0).Format("2006-01-02")
	default:
		return "", fmt.Errorf(i18n.T("template_datetime_error_invalid_unit"), unit)
	}

	debugf("DateTime: relative unit=%q num=%d result=%q", unit, num, result)
	return result, nil
}
