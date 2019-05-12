// Code generated by protoc-gen-cobra.
// source: app/interface/rpc/v1.0/protocol/user.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	cobra "github.com/spf13/cobra"
	context "golang.org/x/net/context"
	credentials "google.golang.org/grpc/credentials"
	envconfig "github.com/kelseyhightower/envconfig"
	filepath "path/filepath"
	grpc "google.golang.org/grpc"
	io "io"
	iocodec "github.com/fiorix/protoc-gen-cobra/iocodec"
	ioutil "io/ioutil"
	json "encoding/json"
	log "log"
	net "net"
	oauth "google.golang.org/grpc/credentials/oauth"
	oauth2 "golang.org/x/oauth2"
	os "os"
	pflag "github.com/spf13/pflag"
	template "text/template"
	time "time"
	tls "crypto/tls"
	x509 "crypto/x509"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ cobra.Command
var _ context.Context
var _ credentials.AuthInfo
var _ envconfig.Decoder
var _ filepath.WalkFunc
var _ grpc.ClientConn
var _ io.Reader
var _ iocodec.Encoder
var _ = ioutil.Discard
var _ json.Encoder
var _ log.Logger
var _ net.IP
var _ oauth.TokenSource
var _ oauth2.Token
var _ os.File
var _ pflag.FlagSet
var _ template.Template
var _ time.Time
var _ tls.Config
var _ x509.Certificate

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

var _DefaultUserServiceClientCommandConfig = _NewUserServiceClientCommandConfig()

type _UserServiceClientCommandConfig struct {
	ServerAddr         string        `envconfig:"SERVER_ADDR" default:"localhost:8080"`
	RequestFile        string        `envconfig:"REQUEST_FILE"`
	PrintSampleRequest bool          `envconfig:"PRINT_SAMPLE_REQUEST"`
	ResponseFormat     string        `envconfig:"RESPONSE_FORMAT" default:"json"`
	Timeout            time.Duration `envconfig:"TIMEOUT" default:"10s"`
	TLS                bool          `envconfig:"TLS"`
	ServerName         string        `envconfig:"TLS_SERVER_NAME"`
	InsecureSkipVerify bool          `envconfig:"TLS_INSECURE_SKIP_VERIFY"`
	CACertFile         string        `envconfig:"TLS_CA_CERT_FILE"`
	CertFile           string        `envconfig:"TLS_CERT_FILE"`
	KeyFile            string        `envconfig:"TLS_KEY_FILE"`
	AuthToken          string        `envconfig:"AUTH_TOKEN"`
	AuthTokenType      string        `envconfig:"AUTH_TOKEN_TYPE" default:"Bearer"`
	JWTKey             string        `envconfig:"JWT_KEY"`
	JWTKeyFile         string        `envconfig:"JWT_KEY_FILE"`
}

func _NewUserServiceClientCommandConfig() *_UserServiceClientCommandConfig {
	c := &_UserServiceClientCommandConfig{}
	envconfig.Process("", c)
	return c
}

func (o *_UserServiceClientCommandConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (must be json, yaml, or xml); use \"-\" for stdin + json")
	fs.BoolVarP(&o.PrintSampleRequest, "print-sample-request", "p", o.PrintSampleRequest, "print sample request file and exit")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (json, prettyjson, yaml, or xml)")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.StringVar(&o.ServerName, "tls-server-name", o.ServerName, "tls server name override")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
	fs.StringVar(&o.AuthToken, "auth-token", o.AuthToken, "authorization token")
	fs.StringVar(&o.AuthTokenType, "auth-token-type", o.AuthTokenType, "authorization token type")
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "jwt key")
	fs.StringVar(&o.JWTKeyFile, "jwt-key-file", o.JWTKeyFile, "jwt key file")
}

var UserServiceClientCommand = &cobra.Command{
	Use: "userservice",
}

func _DialUserService() (*grpc.ClientConn, UserServiceClient, error) {
	cfg := _DefaultUserServiceClientCommandConfig
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Timeout),
	}
	if cfg.TLS {
		tlsConfig := &tls.Config{}
		if cfg.InsecureSkipVerify {
			tlsConfig.InsecureSkipVerify = true
		}
		if cfg.CACertFile != "" {
			cacert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certpool := x509.NewCertPool()
			certpool.AppendCertsFromPEM(cacert)
			tlsConfig.RootCAs = certpool
		}
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		if cfg.ServerName != "" {
			tlsConfig.ServerName = cfg.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(cfg.ServerAddr)
			tlsConfig.ServerName = addr
		}
		//tlsConfig.BuildNameToCertificate()
		cred := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	if cfg.AuthToken != "" {
		cred := oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: cfg.AuthToken,
			TokenType:   cfg.AuthTokenType,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKey != "" {
		cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	if cfg.JWTKeyFile != "" {
		cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
		if err != nil {
			return nil, nil, fmt.Errorf("jwt key file: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}
	conn, err := grpc.Dial(cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewUserServiceClient(conn), nil
}

type _UserServiceRoundTripFunc func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error

func _UserServiceRoundTrip(sample interface{}, fn _UserServiceRoundTripFunc) error {
	cfg := _DefaultUserServiceClientCommandConfig
	var em iocodec.EncoderMaker
	var ok bool
	if cfg.ResponseFormat == "" {
		em = iocodec.DefaultEncoders["json"]
	} else {
		em, ok = iocodec.DefaultEncoders[cfg.ResponseFormat]
		if !ok {
			return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
		}
	}
	if cfg.PrintSampleRequest {
		return em.NewEncoder(os.Stdout).Encode(sample)
	}
	var d iocodec.Decoder
	if cfg.RequestFile == "" || cfg.RequestFile == "-" {
		d = iocodec.DefaultDecoders["json"].NewDecoder(os.Stdin)
	} else {
		f, err := os.Open(cfg.RequestFile)
		if err != nil {
			return fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		ext := filepath.Ext(cfg.RequestFile)
		if len(ext) > 0 && ext[0] == '.' {
			ext = ext[1:]
		}
		dm, ok := iocodec.DefaultDecoders[ext]
		if !ok {
			return fmt.Errorf("invalid request file format: %q", ext)
		}
		d = dm.NewDecoder(f)
	}
	conn, client, err := _DialUserService()
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, d, em.NewEncoder(os.Stdout))
}

var _UserServiceCreateClientCommand = &cobra.Command{
	Use:  "create",
	Long: "Create client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	create -p > req.json

Submit request using file:
	create -f req.json

Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | create --tls`,
	Run: func(cmd *cobra.Command, args []string) {
		var v CreateUserRequest
		err := _UserServiceRoundTrip(v, func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error {

			err := in.Decode(&v)
			if err != nil {
				return err
			}

			resp, err := cli.Create(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UserServiceClientCommand.AddCommand(_UserServiceCreateClientCommand)
	_DefaultUserServiceClientCommandConfig.AddFlags(_UserServiceCreateClientCommand.Flags())
}

var _UserServiceReadClientCommand = &cobra.Command{
	Use:  "read",
	Long: "Read client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	read -p > req.json

Submit request using file:
	read -f req.json

Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | read --tls`,
	Run: func(cmd *cobra.Command, args []string) {
		var v ReadUserRequest
		err := _UserServiceRoundTrip(v, func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error {

			err := in.Decode(&v)
			if err != nil {
				return err
			}

			resp, err := cli.Read(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UserServiceClientCommand.AddCommand(_UserServiceReadClientCommand)
	_DefaultUserServiceClientCommandConfig.AddFlags(_UserServiceReadClientCommand.Flags())
}

var _UserServiceUpdateClientCommand = &cobra.Command{
	Use:  "update",
	Long: "Update client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	update -p > req.json

Submit request using file:
	update -f req.json

Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | update --tls`,
	Run: func(cmd *cobra.Command, args []string) {
		var v UpdateUserRequest
		err := _UserServiceRoundTrip(v, func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error {

			err := in.Decode(&v)
			if err != nil {
				return err
			}

			resp, err := cli.Update(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UserServiceClientCommand.AddCommand(_UserServiceUpdateClientCommand)
	_DefaultUserServiceClientCommandConfig.AddFlags(_UserServiceUpdateClientCommand.Flags())
}

var _UserServiceListClientCommand = &cobra.Command{
	Use:  "list",
	Long: "List client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	list -p > req.json

Submit request using file:
	list -f req.json

Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | list --tls`,
	Run: func(cmd *cobra.Command, args []string) {
		var v ListUserRequest
		err := _UserServiceRoundTrip(v, func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error {

			err := in.Decode(&v)
			if err != nil {
				return err
			}

			resp, err := cli.List(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UserServiceClientCommand.AddCommand(_UserServiceListClientCommand)
	_DefaultUserServiceClientCommandConfig.AddFlags(_UserServiceListClientCommand.Flags())
}

var _UserServiceDeleteClientCommand = &cobra.Command{
	Use:  "delete",
	Long: "Delete client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	delete -p > req.json

Submit request using file:
	delete -f req.json

Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | delete --tls`,
	Run: func(cmd *cobra.Command, args []string) {
		var v DeleteUserRequest
		err := _UserServiceRoundTrip(v, func(cli UserServiceClient, in iocodec.Decoder, out iocodec.Encoder) error {

			err := in.Decode(&v)
			if err != nil {
				return err
			}

			resp, err := cli.Delete(context.Background(), &v)

			if err != nil {
				return err
			}

			return out.Encode(resp)

		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	UserServiceClientCommand.AddCommand(_UserServiceDeleteClientCommand)
	_DefaultUserServiceClientCommandConfig.AddFlags(_UserServiceDeleteClientCommand.Flags())
}
