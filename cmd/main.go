package main

import (
	"bufio"
	"os"

	"github.com/smagulmyrzakhmet/cli_TODO/internal/handler"
	"github.com/smagulmyrzakhmet/cli_TODO/internal/service"
	"github.com/smagulmyrzakhmet/cli_TODO/internal/storage/cache"
)

func main() {
	repo := cache.NewRepositoryImpl()
	cliService := service.NewCLIServiceImpl(repo)
	cliHandler := handler.NewCLIHandler(cliService, bufio.NewReader(os.Stdin))
	cliHandler.Run()
}
