package config

import (
	"LemmyBeProxy/router"
	"LemmyBeProxy/service"
	"LemmyBeProxy/service/backend"
	"LemmyBeProxy/service/frontend"
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
// to — see service/backend/Backend.go. Post, Comment, and Community
// controllers are migrated onto it so far. Every other controller
// (user, search, site, upload) still talks to the raw *piefed.Piefed
// client above directly, and will keep doing so — PieFed-shaped,
// regardless of BACKEND_TYPE — until they get the same migration in
// follow-up work. Setting BACKEND_TYPE=lemmy right now only changes
// Post/Comment/Community endpoint behavior; everything else still
// assumes a Piefed backend until the rest of the controllers move onto
// this interface too.
var activeBackend backend.Backend

// activeFrontend is the wire-format counterpart to activeBackend — see
// service/frontend/Frontend.go. Post and Comment controllers are
// migrated onto it so far; everything else still assumes the current
// (0.19.x) wire format directly regardless of FRONTEND_VERSION.
var activeFrontend frontend.Frontend

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

	// FRONTEND_VERSION is required too, same "no preference" philosophy.
	// "0.19" for the current Lemmy wire format (what every deployment of
	// this proxy has used until now), "0.17" for the older format
	// lemmyBB and similar older clients actually speak.
	frontendVersion, hasFrontendVersion := os.LookupEnv("FRONTEND_VERSION")
	if !hasFrontendVersion {
		panic("FRONTEND_VERSION environment variable not set — must be \"0.19\" or \"0.17\"")
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

	switch frontendVersion {
	case "0.19":
		activeFrontend = frontend.NewFrontend019()
	case "0.17":
		activeFrontend = frontend.NewFrontend017()
	default:
		panic(fmt.Sprintf("unknown FRONTEND_VERSION %q — expected \"0.19\" or \"0.17\"", frontendVersion))
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
