package clparams

type StateVersion uint8

const (
	Phase0Version    StateVersion = 0
	AltairVersion    StateVersion = 1
	BellatrixVersion StateVersion = 2
	CapellaVersion   StateVersion = 3 // Unimplemented!
)

// stringToClVersion converts the string to the current state version.
func StringToClVersion(s string) StateVersion {
	switch s {
	case "phase0":
		return Phase0Version
	case "altair":
		return AltairVersion
	case "bellatrix":
		return BellatrixVersion
	case "capella":
		return CapellaVersion
	default:
		panic("unsupported fork version: " + s)
	}
}
