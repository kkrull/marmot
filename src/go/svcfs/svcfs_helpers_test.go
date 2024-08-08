package svcfs_test

import "github.com/kkrull/marmot/svcfs"

func jsonMetaRepoAdmin(args *jsonMetaRepoAdminArgs) *svcfs.JsonMetaRepoAdmin {
	if args == nil {
		args = &jsonMetaRepoAdminArgs{}
	}

	return svcfs.NewJsonMetaRepoAdmin(args.Version())
}

type jsonMetaRepoAdminArgs struct {
	version string
}

func (args jsonMetaRepoAdminArgs) Version() string {
	if args.version == "" {
		return "42"
	} else {
		return args.version
	}
}
