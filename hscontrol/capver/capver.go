package capver

import (
	"sort"
	"strings"

	xmaps "golang.org/x/exp/maps"
	"tailscale.com/tailcfg"
	"tailscale.com/util/set"
)

func tailscaleVersSorted() []string {
	vers := xmaps.Keys(tailscaleToCapVer)
	sort.Strings(vers)
	return vers
}

func capVersSorted() []tailcfg.CapabilityVersion {
	capVers := xmaps.Keys(capVerToTailscaleVer)
	sort.Slice(capVers, func(i, j int) bool {
		return capVers[i] < capVers[j]
	})
	return capVers
}

// TailscaleVersion returns the Tailscale version for the given CapabilityVersion.
func TailscaleVersion(ver tailcfg.CapabilityVersion) string {
	return capVerToTailscaleVer[ver]
}

// CapabilityVersion returns the CapabilityVersion for the given Tailscale version.
func CapabilityVersion(ver string) tailcfg.CapabilityVersion {
	if !strings.HasPrefix(ver, "v") {
		ver = "v" + ver
	}
	return tailscaleToCapVer[ver]
}

// TailscaleLatest returns the n latest Tailscale versions.
func TailscaleLatest(n int) []string {
	if n <= 0 {
		return nil
	}

	tsSorted := tailscaleVersSorted()

	if n > len(tsSorted) {
		return tsSorted
	}

	return tsSorted[len(tsSorted)-n:]
}

// TailscaleLatestMajorMinor returns the n latest Tailscale versions (e.g. 1.80).
func TailscaleLatestMajorMinor(n int, stripV bool) []string {
	if n <= 0 {
		return nil
	}

	majors := set.Set[string]{}
	for _, vers := range tailscaleVersSorted() {
		if stripV {
			vers = strings.TrimPrefix(vers, "v")
		}
		v := strings.Split(vers, ".")
		majors.Add(v[0] + "." + v[1])
	}

	majorSl := majors.Slice()
	sort.Strings(majorSl)

	if n > len(majorSl) {
		return majorSl
	}

	return majorSl[len(majorSl)-n:]
}

// CapVerLatest returns the n latest CapabilityVersions.
func CapVerLatest(n int) []tailcfg.CapabilityVersion {
	if n <= 0 {
		return nil
	}

	s := capVersSorted()

	if n > len(s) {
		return s
	}

	return s[len(s)-n:]
}
