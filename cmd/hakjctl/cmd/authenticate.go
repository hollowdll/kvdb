package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/hollowdll/hakjdb/api/v1/authpb"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/client"
	"github.com/hollowdll/hakjdb/cmd/hakjctl/config"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var (
	cmdAuthenticate = &cobra.Command{
		Use:   "authenticate",
		Short: "Authenticate to the server",
		Long:  "Authenticate to the server using password. If no options provided, prompts the user to enter password.",
		Example: `# Prompt password
hakjctl authenticate

# Pass password as argument
hakjctl authenticate -p your-password

# Read password from environment variable HAKJCTL_PASSWORD
hakjctl authenticate --password-from-env`,
		Run: func(cmd *cobra.Command, args []string) {
			password := ""
			if isReadPasswordFromEnv {
				password = readPasswordFromEnv()
			} else if passedPassword != "" {
				password = passedPassword
			} else {
				password = promptPassword("Password: ")
			}
			authenticate(password)
		},
	}
	passedPassword        string
	isReadPasswordFromEnv bool
)

func init() {
	cmdAuthenticate.Flags().StringVarP(&passedPassword, "password", "p", "", "The password to use")
	cmdAuthenticate.Flags().BoolVar(&isReadPasswordFromEnv, "password-from-env", false, "Read password from environment variable HAKJCTL_PASSWORD")
}

func promptPassword(prompt string) string {
	fmt.Print(prompt)
	fd := int(os.Stdin.Fd())
	bytes, err := term.ReadPassword(fd)
	if err != nil {
		cobra.CheckErr(err)
	}
	fmt.Println()
	return string(bytes)
}

func authenticate(password string) {
	ctx, cancel := context.WithTimeout(context.Background(), config.GetCmdTimeout())
	defer cancel()

	req := &authpb.AuthenticateRequest{Password: password}
	res, err := client.GrpcAuthClient.Authenticate(ctx, req)
	client.CheckGrpcError(err)

	tokenCachePath, err := client.GetTokenCacheFilePath()
	cobra.CheckErr(err)
	err = client.WriteTokenCache(tokenCachePath, res.AuthToken)
	cobra.CheckErr(err)
	fmt.Println("OK")
}

func readPasswordFromEnv() string {
	password, ok := os.LookupEnv(config.EnvVarPassword)
	if !ok {
		return ""
	}
	return password
}
