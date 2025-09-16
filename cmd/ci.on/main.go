package main

import (
	"fmt"
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/assert"
	"github.com/ci-on-dev/ci.on-cli/internal/logs"
	"github.com/spf13/cobra"
)

var logLevel string
var testFile string
var logService logs.LogsService
var assertService assert.AssertService
var service MainInterface

type mainService struct {
	logService    logs.LogsService
	assertService assert.AssertService
}

type MainInterface interface {
	RunInit() error
	RunTest(testFile string)
}

func NewMainService(logService logs.LogsService, assertService assert.AssertService) MainInterface {
	return mainService{logService: logService, assertService: assertService}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "ci.on",
		Short: "CLI para inicialização e testes do CI",
		Long:  `Uma CLI para rodar inicializações (init) e testes (test) no pipeline.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Agora o logLevel já foi setado pelo Cobra
			fmt.Printf("Inicializando serviços com log-level: %s\n", logLevel)
			logService = logs.NewLogsService(logLevel)
			assertService = assert.NewAssertService(logService)
			service = NewMainService(logService, assertService)
			return nil
		},
	}

	// Flag global
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Define o modo de log (ex: debug, info, warn, error)")

	// Subcomando init
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Executa a inicialização",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Rodando init com log-level:", logLevel)
			if err := service.RunInit(); err != nil {
				return fmt.Errorf("init falhou: %w", err)
			}
			return nil
		},
	}

	// Subcomando test
	testCmd := &cobra.Command{
		Use:   "test",
		Short: "Executa os testes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Rodando test com log-level:", logLevel)
			fmt.Println("Arquivo de teste:", testFile)
			service.RunTest(testFile)
		},
	}

	testCmd.Flags().StringVarP(&testFile, "test-file", "t", "", "Caminho do arquivo de teste")
	testCmd.MarkFlagRequired("test-file")

	rootCmd.AddCommand(initCmd, testCmd)

	// Executa o comando
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.RemoveAll("tmp")
}
