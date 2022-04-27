package locations

import (
	"os"
	"path/filepath"

	"github.com/aripalo/vegas-credentials/internal/utils"

	"github.com/adrg/xdg"
)

// OS/User Cache Data Directory.
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
var CacheDir = mustEnsureAppDir(xdg.CacheHome)

// OS/User State Data Directory.
// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
var StateDir = mustEnsureAppDir(xdg.StateHome)

// The directory where vegas-credentials executable is located.
var ExecDir = filepath.Dir(utils.Must(os.Executable))

// AWS config file location. Usually in $HOME/.aws/config or %USERPROFILE%/.aws/config
// unless user has set $AWS_CONFIG_FILE environment variable.
// https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-where
var AwsConfig = utils.Must(getAwsConfigFile)

// Location of the Yubikey Manager CLI (ykman) executable.
// Empty if not available.
var YkmanPath = getYkmanPath()
