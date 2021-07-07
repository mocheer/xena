package proj4

// Defs
// Defs(32650,"+proj=utm +zone=50 +datum=WGS84 +units=m +no_defs")
func Defs(code int, proj4text string) {
	projStrings[EPSGCode(code)] = proj4text
}
