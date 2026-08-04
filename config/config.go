package config

import (
	"LemmyBeProxy/router"
	"LemmyBeProxy/service"
	"LemmyBeProxy/service/backend"
	lemmyService "LemmyBeProxy/service/lemmy"
	piefedService "LemmyBeProxy/service/piefed"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var AppHttpPort = 8080
var AppRouter *router.Router
var CertificateFile string
var CertificateKey string
var CorsRegex *regexp.Regexp

var simulateLemmy bool
var piefed *piefedService.Piefed
var activityPub *service.ActivityPub

// activeBackend is the shared instance every migrated controller talks
// to — see service/backend/Backend.go. Post and Comment controllers are
// migrated onto it so far. Every other controller (community, user,
// search, site, upload) still talks to the raw *piefed.Piefed client
// above directly, and will keep doing so — PieFed-shaped, regardless of
// BACKEND_TYPE — until they get the same migration in follow-up work.
// Setting BACKEND_TYPE=lemmy right now only changes Post/Comment
// endpoint behavior; everything else still assumes a Piefed backend
// until the rest of the controllers move onto this interface too.
var activeBackend backend.Backend

func init() {
	if port, exists := os.LookupEnv("PORT"); exists {
		parsed, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}

		AppHttpPort = parsed
	}
	if simulate, exists := os.LookupEnv("SIMULATE"); exists {
		var err error
		simulateLemmy, err = strconv.ParseBool(simulate)
		if err != nil {
			panic(err)
		}
	}

	// BACKEND_TYPE and BACKEND_INSTANCE are both required, with no
	// default — this proxy takes no position on which backend you're
	// running, so every deployment must say explicitly. "piefed" or
	// "lemmy" for BACKEND_TYPE; BACKEND_INSTANCE is that backend's
	// hostname (e.g. retrofed.com).
	backendType, hasBackendType := os.LookupEnv("BACKEND_TYPE")
	if !hasBackendType {
		panic("BACKEND_TYPE environment variable not set — must be \"piefed\" or \"lemmy\"")
	}
	backendInstance, hasBackendInstance := os.LookupEnv("BACKEND_INSTANCE")
	if !hasBackendInstance {
		panic("BACKEND_INSTANCE environment variable not set")
	}

	piefed = piefedService.NewPiefed(backendInstance)
	activityPub = service.NewActivityPub()

	switch backendType {
	case "piefed":
		activeBackend = piefedService.NewPiefedBackend(piefedService.NewPiefed(backendInstance))
	case "lemmy":
		activeBackend = lemmyService.NewLemmyBackend(lemmyService.NewLemmy(backendInstance))
	default:
		panic(fmt.Sprintf("unknown BACKEND_TYPE %q — expected \"piefed\" or \"lemmy\"", backendType))
	}

	AppRouter = router.NewRouter()

	if certFile, exists := os.LookupEnv("CERTIFICATE_FILE"); exists {
		CertificateFile = certFile
	}
	if certKey, exists := os.LookupEnv("CERTIFICATE_KEY"); exists {
		CertificateKey = certKey
	}
	if regexStr, exists := os.LookupEnv("CORS_REGEX"); exists {
		CorsRegex = regexp.MustCompile(regexStr)
	}
}
