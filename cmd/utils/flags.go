Name:  "sepolia",
		Usage: "Sepolia network: pre-configured proof-of-work test network",
	}
	KilnFlag = cli.BoolFlag{
		Name:  "kiln",
		Usage: "Kiln network: pre-configured proof-of-work to proof-of-stake test network",
	}
	DeveloperFlag = cli.BoolFlag{
		Name:  "dev",
		Usage: "Ephemeral proof-of-authority network with a pre-funded developer account, mining enabled",
@@ -839,6 +843,9 @@ func MakeDataDir(ctx *cli.Context) string {
		if ctx.GlobalBool(SepoliaFlag.Name) {
			return filepath.Join(path, "sepolia")
		}
		if ctx.GlobalBool(KilnFlag.Name) {
			return filepath.Join(path, "kiln")
		}
		return path
	}
	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
@@ -893,6 +900,8 @@ func setBootstrapNodes(ctx *cli.Context, cfg *p2p.Config) {
		urls = params.RinkebyBootnodes
	case ctx.GlobalBool(GoerliFlag.Name):
		urls = params.GoerliBootnodes
	case ctx.GlobalBool(KilnFlag.Name):
		urls = params.KilnBootnodes
	case cfg.BootstrapNodes != nil:
		return // already set, don't apply defaults.
	}
