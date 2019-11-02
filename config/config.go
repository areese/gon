package config

// Config is the configuration structure for gon.
type Config struct {
	// Source is the list of binary files to sign.
	Source []string `hcl:"source"`

	// BundleId is the bundle ID to use for the package that is created.
	// This should be in a format such as "com.example.app". The value can
	// be anything, this is required by Apple.
	BundleId string `hcl:"bundle_id"`

	// Sign are the settings for code-signing the binaries.
	Sign Sign `hcl:"sign,block"`

	// AppleId are the credentials to use to talk to Apple.
	AppleId AppleId `hcl:"apple_id,block"`

	// Zip, if present, creates a notarized zip file as the output. Note
	// that zip files do not support stapling, so the final result will
	// require an internet connection on first use to validate the notarization.
	Zip *Zip `hcl:"zip,block"`
}

// AppleId are the authentication settings for Apple systems.
type AppleId struct {
	// Username is your AC username, typically an email.
	Username string `hcl:"username"`

	// Password is the password for your AC account. This also accepts
	// two additional forms: '@keychain:<name>' which reads the password from
	// the keychain and '@env:<name>' which reads the password from an
	// an environmental variable named <name>.
	Password string `hcl:"password"`

	// Provider is the AC provider. This is optional and only needs to be
	// specified if you're using an Apple ID account that has multiple
	// teams.
	Provider string `hcl:"provider,optional"`
}

// Sign are the options for codesigning the binaries.
type Sign struct {
	// ApplicationIdentity is the ID or name of the certificate to
	// use for signing binaries. This is used for all binaries in "source".
	ApplicationIdentity string `hcl:"application_identity"`
}

// Zip are the options for a zip file as output.
type Zip struct {
	// OutputPath is the path where the final zip file will be saved.
	OutputPath string `hcl:"output_path"`
}
