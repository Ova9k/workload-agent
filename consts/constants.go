package consts

import "crypto"

// Define constants to be accessed in other packages
const (
	MTWILSON_API_URL                            = "MTWILSON_API_URL"
	MTWILSON_API_USERNAME                       = "MTWILSON_API_USERNAME"
	MTWILSON_API_PASSWORD                       = "MTWILSON_API_PASSWORD"
	MTWILSON_TLS_SHA256                         = "MTWILSON_TLS_CERT_SHA256"
	WLS_API_URL                                 = "WLS_API_URL"
	WLS_API_USERNAME                            = "WLS_API_USERNAME"
	WLS_API_PASSWORD                            = "WLS_API_PASSWORD"
	WLS_TLS_SHA256                              = "WLS_TLS_SHA256"
	LOG_LEVEL                                   = "LOG_LEVEL"
	AikSecretKeyName                            = "aik.secret"
	TAConfigDirEnvVar                           = "TRUSTAGENT_CONFIGURATION"
	TAConfigAikSecretCmd                        = "tagent config aik.secret"
	TAAikPemFileName                            = "aik.pem"
	TAUserNameEnvVar                            = "TRUSTAGENT_USERNAME"
	BindingKeyFileName                          = "bindingkey.json"
	SigningKeyFileName                          = "signingkey.json"
	BindingKeyPemFileName                       = "bindingkey.pem"
	SigningKeyPemFileName                       = "signingkey.pem"
	ImageVmCountAssociationFileName             = "image_vm_association"
	EnvFileName                                 = "workloadagent.env"
	DevMapperDirPath                            = "/dev/mapper/"
	MountPath                                   = "/mnt/workload-agent/crypto/"
	LogFileName                                 = "workloadagent.log"
	LogDirPath                                  = "/var/log/workload-agent/"
	DaemonLogFileName                           = "daemon.log"
	ConfigFileName                              = "config.yml"
	ConfigDirPath                               = "/etc/workload-agent/"
	OptDirPath                                  = "/opt/workload-agent/"
	BinDirPath                                  = "/opt/workload-agent/bin/"
	RunDirPath                                  = "/var/run/workload-agent/"
	LibvirtHookFilePath                         = "/etc/libvirt/hooks/qemu"
	DaemonFileName                              = "wlagentd"
	PIDFileName                                 = "wlagent.pid"
	RPCSocketFileName                           = "wlagent.sock"
	WlagentSymLink                              = "/usr/local/bin/wlagent"
	PemCertificateHeader                        = "CERTIFICATE"
	HashingAlgorithm                crypto.Hash = crypto.SHA256
	WLABinFilePath                              = "/usr/local/bin/wlagent"
	AIKPemFileName                              = "aik.pem"
)
