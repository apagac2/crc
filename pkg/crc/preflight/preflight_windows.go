package preflight

var hypervPreflightChecks = [...]Check{
	{
		configKeySuffix:  "check-administrator-user",
		checkDescription: "Checking if running as normal user",
		check:            checkIfRunningAsNormalUser,
		fixDescription:   "crc should be ran as a normal user",
		flags:            NoFix,
	},
	{
		configKeySuffix:  "check-windows-version",
		checkDescription: "Checking Windows 10 release",
		check:            checkVersionOfWindowsUpdate,
		fixDescription:   "Please manually update your Windows 10 installation",
		flags:            NoFix,
	},
	{
		configKeySuffix:  "check-windows-edition",
		checkDescription: "Checking Windows edition",
		check:            checkWindowsEdition,
		fixDescription:   "Your Windows edition is not supported. Consider using Professional or Enterprise editions of Windows",
		flags:            NoFix,
	},
	{
		configKeySuffix:  "check-hyperv-installed",
		checkDescription: "Checking if Hyper-V is installed and operational",
		check:            checkHyperVInstalled,
		fixDescription:   "Installing Hyper-V",
		fix:              fixHyperVInstalled,
	},
	{
		configKeySuffix:  "check-user-in-hyperv-group",
		checkDescription: "Checking if user is a member of the Hyper-V Administrators group",
		check:            checkIfUserPartOfHyperVAdmins,
		fixDescription:   "Adding user to the Hyper-V Administrators group",
		fix:              fixUserPartOfHyperVAdmins,
	},
	{
		configKeySuffix:  "check-hyperv-service-running",
		checkDescription: "Checking if Hyper-V service is enabled",
		check:            checkHyperVServiceRunning,
		fixDescription:   "Enabling Hyper-V service",
		fix:              fixHyperVServiceRunning,
	},
	{
		configKeySuffix:  "check-hyperv-switch",
		checkDescription: "Checking if the Hyper-V virtual switch exist",
		check:            checkIfHyperVVirtualSwitchExists,
		fixDescription:   "Unable to perform Hyper-V administrative commands. Please reboot your system and run 'crc setup' to complete the setup process",
		flags:            NoFix,
	},
	{
		cleanupDescription: "Removing dns server from interface",
		cleanup:            removeDNSServerAddress,
		flags:              CleanUpOnly,
	},
	{
		cleanupDescription: "Removing the crc VM if exists",
		cleanup:            removeCrcVM,
		flags:              CleanUpOnly,
	},
}

var traySetupChecks = [...]Check{
	{
		checkDescription: "Checking if tray binary is present",
		check:            checkTrayBinaryExists,
		fixDescription:   "Caching tray binary",
		fix:              fixTrayBinaryExists,
		flags:            SetupOnly,
	},
	{
		checkDescription:   "Checking if tray is installed",
		check:              checkIfTrayInstalled,
		fixDescription:     "Installing CodeReady Containers tray",
		fix:                fixTrayInstalled,
		cleanupDescription: "Uninstalling tray if installed",
		cleanup:            removeTray,
		flags:              SetupOnly,
	},
}

func getPreflightChecks() []Check {
	checks := []Check{}
	checks = append(checks, genericPreflightChecks[:]...)
	checks = append(checks, hypervPreflightChecks[:]...)

	// Experimental feature
	if EnableExperimentalFeatures {
		checks = append(checks, traySetupChecks[:]...)
	}

	return checks
}

// StartPreflightChecks performs the preflight checks before starting the cluster
func StartPreflightChecks() {
	doPreflightChecks(getPreflightChecks())
}

// SetupHost performs the prerequisite checks and setups the host to run the cluster
func SetupHost() {
	doFixPreflightChecks(getPreflightChecks())
}

func RegisterSettings() {
	doRegisterSettings(getPreflightChecks())
}

func CleanUpHost() {
	// A user can use setup with experiment flag
	// and not use cleanup with same flag, to avoid
	// any extra step/confusion we are just adding the checks
	// which are behind the experiment flag. This way cleanup
	// perform action in a sane way.
	checks := getPreflightChecks()
	checks = append(checks, traySetupChecks[:]...)
	doCleanUpPreflightChecks(checks)
}
