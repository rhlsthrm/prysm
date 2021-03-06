// Package flags contains all configuration runtime flags for
// the validator service.
package flags

import (
	"github.com/urfave/cli/v2"
)

var (
	// DisableAccountMetricsFlag defines the graffiti value included in proposed blocks, default false.
	DisableAccountMetricsFlag = &cli.BoolFlag{
		Name: "disable-account-metrics",
		Usage: "Disable prometheus metrics for validator accounts. Operators with high volumes " +
			"of validating keys may wish to disable granular prometheus metrics as it increases " +
			"the data cardinality.",
	}
	// BeaconRPCProviderFlag defines a beacon node RPC endpoint.
	BeaconRPCProviderFlag = &cli.StringFlag{
		Name:  "beacon-rpc-provider",
		Usage: "Beacon node RPC provider endpoint",
		Value: "localhost:4000",
	}
	// CertFlag defines a flag for the node's TLS certificate.
	CertFlag = &cli.StringFlag{
		Name:  "tls-cert",
		Usage: "Certificate for secure gRPC. Pass this and the tls-key flag in order to use gRPC securely.",
	}
	// SlasherRPCProviderFlag defines a slasher node RPC endpoint.
	SlasherRPCProviderFlag = &cli.StringFlag{
		Name:  "slasher-rpc-provider",
		Usage: "Slasher node RPC provider endpoint",
		Value: "localhost:4002",
	}
	// SlasherCertFlag defines a flag for the slasher node's TLS certificate.
	SlasherCertFlag = &cli.StringFlag{
		Name:  "slasher-tls-cert",
		Usage: "Certificate for secure slasher gRPC. Pass this and the tls-key flag in order to use gRPC securely.",
	}
	// DisablePenaltyRewardLogFlag defines the ability to not log reward/penalty information during deployment
	DisablePenaltyRewardLogFlag = &cli.BoolFlag{
		Name:  "disable-rewards-penalties-logging",
		Usage: "Disable reward/penalty logging during cluster deployment",
	}
	// GraffitiFlag defines the graffiti value included in proposed blocks
	GraffitiFlag = &cli.StringFlag{
		Name:  "graffiti",
		Usage: "String to include in proposed blocks",
	}
	// GrpcRetriesFlag defines the number of times to retry a failed gRPC request.
	GrpcRetriesFlag = &cli.UintFlag{
		Name:  "grpc-retries",
		Usage: "Number of attempts to retry gRPC requests",
		Value: 5,
	}
	// GrpcHeadersFlag defines a list of headers to send with all gRPC requests.
	GrpcHeadersFlag = &cli.StringFlag{
		Name: "grpc-headers",
		Usage: "A comma separated list of key value pairs to pass as gRPC headers for all gRPC " +
			"calls. Example: --grpc-headers=key=value",
	}
	// KeyManager specifies the key manager to use.
	KeyManager = &cli.StringFlag{
		Name:  "keymanager",
		Usage: "The keymanger to use (unencrypted, interop, keystore, wallet)",
		Value: "",
	}
	// KeyManagerOpts specifies the key manager options.
	KeyManagerOpts = &cli.StringFlag{
		Name:  "keymanageropts",
		Usage: "The options for the keymanger, either a JSON string or path to same",
		Value: "",
	}
	// KeystorePathFlag defines the location of the keystore directory for a validator's account.
	KeystorePathFlag = &cli.StringFlag{
		Name:  "keystore-path",
		Usage: "Path to the desired keystore directory",
	}
	// MergeSourceDirectories defines the locations of source validator databases that will be merged.
	MergeSourceDirectories = &cli.StringFlag{
		Name:  "merge-source-dirs",
		Usage: "A comma-separated list of directories containing validator databases that will be merged.",
	}
	// MergeTargetDirectory defines the locations where the merged database will be stored.
	MergeTargetDirectory = &cli.StringFlag{
		Name:  "merge-target-dir",
		Usage: "The directory where the merged database will be stored.",
	}
	// MonitoringPortFlag defines the http port used to serve prometheus metrics.
	MonitoringPortFlag = &cli.Int64Flag{
		Name:  "monitoring-port",
		Usage: "Port used to listening and respond metrics for prometheus.",
		Value: 8081,
	}

	// PasswordFlag defines the password value for storing and retrieving validator private keys from the keystore.
	PasswordFlag = &cli.StringFlag{
		Name:  "password",
		Usage: "String value of the password for your validator private keys",
	}
	// SplitSourceDirectory defines the location of the validator database that will be split.
	SplitSourceDirectory = &cli.StringFlag{
		Name: "split-source-dir",
		Usage: "The directory containing the validator database that will be split.",
	}
	// SplitTargetDirectory defines the location where the split outcome will be stored.
	SplitTargetDirectory = &cli.StringFlag{
		Name: "split-target-dir",
		Usage: "The directory where the split outcome will be stored.",
	}
	// UnencryptedKeysFlag specifies a file path of a JSON file of unencrypted validator keys as an
	// alternative from launching the validator client from decrypting a keystore directory.
	UnencryptedKeysFlag = &cli.StringFlag{
		Name:  "unencrypted-keys",
		Usage: "Filepath to a JSON file of unencrypted validator keys for easier launching of the validator client",
		Value: "",
	}
)
