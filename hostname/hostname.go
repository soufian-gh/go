// Package hostname parses v1 and v2 hostnames into their constituent parts. It
// is intended to help in the transition from v1 to v2 names on the platform.
// M-Lab go programs that need to parse hostnames should use this package.
package hostname

import (
	"fmt"
	"regexp"
)

// Hostname represents an M-Lab hostname and all of its constituent parts.
type Hostname struct {
	hostname string
	machine  string
	site     string
	project  string
	domain   string
	version  string
}

// Parse parses an M-Lab hostname and breaks it into its constituent parts.
func Parse(name string) (Hostname, error) {
	var parts Hostname

	reInit := regexp.MustCompile(`^mlab[1-4]([.-])`)
	reV1 := regexp.MustCompile(`^(mlab[1-4])\.([a-z]{3}[0-9tc]{2})\.(measurement-lab.org)$`)
	reV2 := regexp.MustCompile(`^(mlab[1-4])-([a-z]{3}[0-9tc]{2})\.(.*?)\.(measurement-lab.org)$`)

	mInit := reInit.FindAllStringSubmatch(name, -1)
	if len(mInit) != 1 || len(mInit[0]) != 2 {
		return parts, fmt.Errorf("Invalid hostname: %s", name)
	}

	switch mInit[0][1] {
	case "-":
		mV2 := reV2.FindAllStringSubmatch(name, -1)
		if len(mV2) != 1 || len(mV2[0]) != 5 {
			return parts, fmt.Errorf("Invalid v2 hostname: %s", name)
		}
		parts = Hostname{
			hostname: mV2[0][0],
			machine:  mV2[0][1],
			site:     mV2[0][2],
			project:  mV2[0][3],
			domain:   mV2[0][4],
			version:  "v2",
		}
	case ".":
		mV1 := reV1.FindAllStringSubmatch(name, -1)
		if len(mV1) != 1 || len(mV1[0]) != 4 {
			return parts, fmt.Errorf("Invalid v1 hostname: %s", name)
		}
		parts = Hostname{
			hostname: mV1[0][0],
			machine:  mV1[0][1],
			site:     mV1[0][2],
			project:  "",
			domain:   mV1[0][3],
			version:  "v1",
		}
	}

	return parts, nil
}
