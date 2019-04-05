package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/manifoldco/promptui"

	"github.com/infinimesh/infinimesh/pkg/apiserver/apipb"
)

var (
	clientConn      *grpc.ClientConn
	namespaceClient apipb.NamespacesClient
	objectClient    apipb.ObjectsClient
	accountClient   apipb.AccountsClient
	deviceClient    apipb.DevicesClient
	shadowClient    apipb.StatesClient
	ctx             context.Context

	noHeaderFlag bool

	namespaceFlag string

	config *Config
)

var rootCmd = &cobra.Command{
	Use:   "inf",
	Short: "Official commandline client for Infinimesh IoT",
}

func init() {
	if cfg, err := ReadConfig(); err == nil {
		cur, err := cfg.GetCurrentContext()
		if err == nil {
			ctx = metadata.AppendToOutgoingContext(context.Background(), "authorization", "bearer "+cur.Token)
		}
		config = cfg

	} else {
		ctx = context.Background()
	}
	rootCmd.AddCommand(loginCmd)

}

func getNamespace() string {
	if allNamespaces {
		return ""
	}
	if namespaceFlag != "" {
		return namespaceFlag
	}
	if config != nil {
		if ctx, err := config.GetCurrentContext(); err == nil {
			return ctx.DefaultNamespace
		}
	}
	return ""
}

func connectGRPC() error {
	current, err := config.GetCurrentContext()
	if err != nil {
		return errors.New("no context found")
	}
	var option grpc.DialOption
	pool, _ := x509.SystemCertPool()
	option = grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: pool}))

	if !current.TLS {
		option = grpc.WithInsecure()
	}

	conn, err := grpc.Dial(current.Server, option)
	if err != nil {
		return err
	}

	clientConn = conn
	namespaceClient = apipb.NewNamespacesClient(conn)
	accountClient = apipb.NewAccountsClient(conn)
	deviceClient = apipb.NewDevicesClient(conn)
	objectClient = apipb.NewObjectsClient(conn)
	shadowClient = apipb.NewStatesClient(conn)
	return nil
}

func disconnectGRPC() error {
	return clientConn.Close()
}

// What to do if unsuccessful connection / no current state unavailable?
func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to account",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Username: ")
		scanner.Scan()
		username := scanner.Text()

		prompt := promptui.Prompt{
			Label:     "Password",
			Mask:      '*',
			Validate:  nil,
			IsConfirm: false,
			Templates: &promptui.PromptTemplates{
				Prompt:  "ABC",
				Valid:   "Password: ",
				Invalid: "Password: ",
				Success: "Password: ",
			},
		}

		password, err := prompt.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read password: %v\n", err)
		}

		response, err := accountClient.Token(context.Background(), &apipb.TokenRequest{Username: username, Password: password})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Login failed: %v\n", err)
			os.Exit(1)
		}

		cfg, err := config.GetCurrentContext()
		if err != nil {
			panic(err)
		}

		cfg.Token = response.Token
		cfg.DefaultNamespace = username

		err = config.Write()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write config: %v\n", err)
		}

		fmt.Println("Logged in successfully.")
	},
	Args: cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return connectGRPC()
	},
	PostRunE: func(cmd *cobra.Command, args []string) error {
		return disconnectGRPC()
	},
}
