package storage

import (
	"github.com/TRON-US/go-btfs/core/commands/storage/announce"
	"github.com/TRON-US/go-btfs/core/commands/storage/challenge"
	"github.com/TRON-US/go-btfs/core/commands/storage/contracts"
	"github.com/TRON-US/go-btfs/core/commands/storage/hosts"
	"github.com/TRON-US/go-btfs/core/commands/storage/info"
	"github.com/TRON-US/go-btfs/core/commands/storage/path"
	"github.com/TRON-US/go-btfs/core/commands/storage/stats"
	"github.com/TRON-US/go-btfs/core/commands/storage/upload/upload"

	cmds "github.com/TRON-US/go-btfs-cmds"
)

var StorageCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Interact with storage services on BTFS.",
		ShortDescription: `
Storage services include client upload operations, host storage operations,
host information sync/display operations, and BTT payment-related routines.`,
	},
	Subcommands: map[string]*cmds.Command{
		"upload":    upload.StorageUploadCmd,
		"hosts":     hosts.StorageHostsCmd,
		"info":      info.StorageInfoCmd,
		"announce":  announce.StorageAnnounceCmd,
		"challenge": challenge.StorageChallengeCmd,
		"stats":     stats.StorageStatsCmd,
		"contracts": contracts.StorageContractsCmd,
		"path":      path.PathCmd,
		"dcrepair":  upload.StorageDcRepairRouterCmd,
	},
}
