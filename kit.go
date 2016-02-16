package main

import "strings"
import "time"

func now(fmt string) string {
	return time.Now().Format(fmt)
}

func mapTeam(t string) string {
	t = strings.ToLower(t)

	if t == "s" {
		return "1"
	}

	if t == "联合" || t == "join" {
		return "99"
	}

	return "2"
}
