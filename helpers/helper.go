package helpers

// GetServiceFileByLanguage will get the correct json file
// according to appropriate language
// If language doesn't exists, returns default (english)
func GetServiceFileByLanguage(language string) string {
	switch language {
	case "en":
		return "services_en.json"
	case "fr":
		return "services_fr.json"
	default:
		return "services_en.json"
	}
}
