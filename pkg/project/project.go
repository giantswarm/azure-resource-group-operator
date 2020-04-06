package project

var (
	bundleVersion = "0.0.1"
	description   = "The azure-resource-group-operator does something."
	gitSHA        = "n/a"
	name          = "azure-resource-group-operator"
	source        = "https://github.com/giantswarm/azure-resource-group-operator"
	version       = "n/a"
)

func BundleVersion() string {
	return bundleVersion
}

func Description() string {
	return description
}

func GitSHA() string {
	return gitSHA
}

func Name() string {
	return name
}

func Source() string {
	return source
}

func Version() string {
	return version
}
