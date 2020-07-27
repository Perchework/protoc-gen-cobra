// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pb

import (
	context "context"
	tls "crypto/tls"
	x509 "crypto/x509"
	fmt "fmt"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	proto "github.com/golang/protobuf/proto"
	cobra "github.com/spf13/cobra"
	pflag "github.com/spf13/pflag"
	oauth2 "golang.org/x/oauth2"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
	oauth "google.golang.org/grpc/credentials/oauth"
	io "io"
	ioutil "io/ioutil"
	net "net"
	os "os"
	filepath "path/filepath"
	strings "strings"
	time "time"
)

var CacheClientDefaultConfig = &_CacheClientConfig{
	ServerAddr:     "localhost:8080",
	RequestFormat:  "json",
	ResponseFormat: "json",
	Timeout:        10 * time.Second,
	AuthTokenType:  "Bearer",
}

type _CacheClientConfig struct {
	ServerAddr         string
	RequestFile        string
	RequestFormat      string
	ResponseFormat     string
	Timeout            time.Duration
	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string
	AuthToken          string
	AuthTokenType      string
	JWTKey             string
	JWTKeyFile         string
}

func (o *_CacheClientConfig) addFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.StringVarP(&o.RequestFile, "request-file", "f", o.RequestFile, "client request file (must be json, yaml, or xml); use \"-\" for stdin + json")
	fs.StringVarP(&o.RequestFormat, "request-format", "i", o.RequestFormat, "request format (json, yaml, or xml)")
	fs.StringVarP(&o.ResponseFormat, "response-format", "o", o.ResponseFormat, "response format (json, prettyjson, xml, prettyxml, or yaml)")
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

func CacheClientCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cache",
		Short: "Cache service client",
		Long:  "",
	}
	CacheClientDefaultConfig.addFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_CacheSetCommand(),
		_CacheGetCommand(),
		_CacheMultiSetCommand(),
		_CacheMultiGetCommand(),
	)
	return cmd
}

func _CacheDial(ctx context.Context) (*grpc.ClientConn, CacheClient, error) {
	cfg := CacheClientDefaultConfig
	opts := []grpc.DialOption{grpc.WithBlock()}
	if cfg.TLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify}
		if cfg.CACertFile != "" {
			caCert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, nil, fmt.Errorf("ca cert: %v", err)
			}
			certPool := x509.NewCertPool()
			certPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = certPool
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
	if cfg.Timeout > 0 {
		var done context.CancelFunc
		ctx, done = context.WithTimeout(ctx, cfg.Timeout)
		defer done()
	}
	conn, err := grpc.DialContext(ctx, cfg.ServerAddr, opts...)
	if err != nil {
		return nil, nil, err
	}
	return conn, NewCacheClient(conn), nil
}

type _CacheRoundTripFunc func(cli CacheClient, in iocodec.Decoder, out iocodec.Encoder) error

func _CacheRoundTrip(ctx context.Context, fn _CacheRoundTripFunc) error {
	cfg := CacheClientDefaultConfig
	if cfg.ResponseFormat == "" {
		cfg.RequestFormat = "json"
	}
	var dm iocodec.DecoderMaker
	r := os.Stdin
	if stat, _ := os.Stdin.Stat(); (stat.Mode()&os.ModeCharDevice) == 0 || cfg.RequestFile == "-" {
		dm = iocodec.DefaultDecoders[cfg.RequestFormat]
	} else if cfg.RequestFile != "" {
		f, err := os.Open(cfg.RequestFile)
		if err != nil {
			return fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		if ext := strings.TrimLeft(filepath.Ext(cfg.RequestFile), "."); ext != "" && ext != cfg.ResponseFormat {
			dm = iocodec.DefaultDecoders[ext]
		}
		r = f
	} else {
		dm = iocodec.DefaultDecoders["noop"]
	}
	if dm == nil {
		return fmt.Errorf("invalid request format: %q", cfg.RequestFormat)
	}
	in := dm.NewDecoder(r)
	if cfg.ResponseFormat == "" {
		cfg.ResponseFormat = "json"
	}
	var em iocodec.EncoderMaker
	if em = iocodec.DefaultEncoders[cfg.ResponseFormat]; em == nil {
		return fmt.Errorf("invalid response format: %q", cfg.ResponseFormat)
	}
	out := em.NewEncoder(os.Stdout)
	conn, client, err := _CacheDial(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	return fn(client, in, out)
}

func _CacheSetCommand() *cobra.Command {
	req := &SetRequest{}

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CacheRoundTrip(cmd.Context(), func(cli CacheClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &SetRequest{}

				if err := in.Decode(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Set(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out.Encode(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, "key", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}

func _CacheGetCommand() *cobra.Command {
	req := &GetRequest{}

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CacheRoundTrip(cmd.Context(), func(cli CacheClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &GetRequest{}

				if err := in.Decode(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.Get(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out.Encode(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, "key", "", "")

	return cmd
}

func _CacheMultiSetCommand() *cobra.Command {
	req := &SetRequest{}

	cmd := &cobra.Command{
		Use:   "multiset",
		Short: "MultiSet RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CacheRoundTrip(cmd.Context(), func(cli CacheClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &SetRequest{}

				stm, err := cli.MultiSet(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in.Decode(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}

				res, err := stm.CloseAndRecv()
				if err != nil {
					return err
				}

				return out.Encode(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, "key", "", "")
	cmd.PersistentFlags().StringVar(&req.Value, "value", "", "")

	return cmd
}

func _CacheMultiGetCommand() *cobra.Command {
	req := &GetRequest{}

	cmd := &cobra.Command{
		Use:   "multiget",
		Short: "MultiGet RPC client",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return _CacheRoundTrip(cmd.Context(), func(cli CacheClient, in iocodec.Decoder, out iocodec.Encoder) error {
				v := &GetRequest{}

				stm, err := cli.MultiGet(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in.Decode(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}

				for {
					res, err := stm.Recv()
					if err != nil {
						if err == io.EOF {
							break
						}
						return err
					}
					if err = out.Encode(res); err != nil {
						return err
					}
				}
				return nil

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.Key, "key", "", "")

	return cmd
}
