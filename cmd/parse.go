package cmd

import "strings"

func parseStatus(raw string) (state, uptime, health string) {
	switch {
	case strings.HasPrefix(raw, "Up"):
		state = "up"
	case strings.HasPrefix(raw, "Exited"):
		state = "exited"
	case strings.HasPrefix(raw, "Restarting"):
		state = "restarting"
	default:
		state = "unknown"
	}

	rest := raw[strings.Index(raw, " ")+1:]

	openIdx := strings.Index(rest, " (")
	if openIdx == -1 {
		uptime = rest
		return
	}

	uptime = rest[:openIdx]

	closeIdx := strings.Index(rest, ")")
	if closeIdx == -1 {
		return
	}

	inner := rest[openIdx+2 : closeIdx]
	health = strings.TrimPrefix(inner, "health: ")
	return
}

func parseContainers(out string) []Container {
	containers := []Container{}
	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "\t")
		if len(parts) != 4 {
			continue
		}

		state, uptime, health := parseStatus(parts[1])

		c := Container{
			Name:   parts[0],
			Image:  parts[2],
			Stack:  parts[3],
			State:  state,
			Uptime: uptime,
			Health: health,
		}
		containers = append(containers, c)
	}
	return containers
}
