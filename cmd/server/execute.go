package server

import (
	"fmt"
	"github.com/swartz-k/chatgpt-app/interface/web"
	"github.com/swartz-k/chatgpt-app/pkg/config"
	"github.com/swartz-k/chatgpt-app/pkg/log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type option struct {
	Debug  bool
	Addr   string
	Config string
}

// rootCmd represents the base command when called without any subcommands
var (
	opt = &option{}

	rootCmd = &cobra.Command{
		Use:   "server",
		Short: "server",
		Long:  `Server provides ChatGPT backend server`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return startServer(opt)
		},
	}
)

func init() {
	flags := rootCmd.PersistentFlags()
	flags.BoolVarP(&opt.Debug, "debug", "D", false, "Enable debug messages")
	flags.StringVarP(&opt.Addr, "addr", "L", "127.0.0.1:8092", "Server listening addr")
	flags.StringVarP(&opt.Config, "config", "C", "config.json", "Server Config path")
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}

func startServer(opt *option) error {
	if opt == nil || !opt.Debug {
		gin.SetMode(gin.ReleaseMode)
		log.NewLoggerWithLevel(3)
	} else {
		gin.SetMode(gin.DebugMode)
		log.NewLoggerWithLevel(101)
	}

	r := gin.New()
	// we trade config order option > env > config
	cfg := config.GetConfig(&opt.Config)
	if opt.Addr != "" {
		cfg.Addr = opt.Addr
	}
	// get persistence
	//_, err := persistence.GetByConfig(cfg)
	//if err != nil {
	//	return errors.Wrap(err, "new persistence")
	//}

	// register route
	web.Register(r)
	log.V(1).Info("Starting server at %s", cfg.Addr)
	return r.Run(cfg.Addr)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
