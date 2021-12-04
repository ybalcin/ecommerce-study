package cli

import (
	"bufio"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createcampaign"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createorder"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createproduct"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getcampaigninfo"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getproductinfo"
	"github.com/ybalcin/ecommerce-study/internal/common"
	"github.com/ybalcin/ecommerce-study/internal/ports"
	"log"
	"os"
	"strings"
)

const (
	createProduct  string = "create_product"
	createCampaign string = "create_campaign"
	createOrder    string = "create_order"
	increaseTime   string = "increase_time"

	getCampaignInfo string = "get_campaign_info"
	getProductInfo  string = "get_product_info"
)

func Run() {
	for i := 1; i <= 5; i++ {
		cliPort := ports.NewCLI()

		pwd, _ := os.Getwd()

		file, err := os.Open(fmt.Sprintf("%s/internal/cmd/cli/cases/c_%d.txt", pwd, i))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("\n> > -------------- READING c_%d.txt -------------- < < \n", i)

		scanner := bufio.NewScanner(file)

		scanner.Split(bufio.ScanLines)
		var cmds []string

		for scanner.Scan() {
			cmds = append(cmds, scanner.Text())
		}

		file.Close()

		for _, cmd := range cmds {
			cmdTokens := strings.Split(cmd, " ")

			entry := common.ValueOfSlice(0, cmdTokens)
			if entry == "" {
				continue
			}

			switch entry {
			case createCampaign:
				command := createcampaign.Build(cmd)
				cliPort.CreateCampaign(command)
			case createOrder:
				command := createorder.Build(cmd)
				cliPort.CreateOrder(command)
			case createProduct:
				command := createproduct.Build(cmd)
				cliPort.CreateProduct(command)
			case increaseTime:
				command := increasetime.Build(cmd)
				cliPort.IncreaseTime(command)
			case getCampaignInfo:
				query := getcampaigninfo.Build(cmd)
				cliPort.GetCampaignInfo(query)
			case getProductInfo:
				query := getproductinfo.Build(cmd)
				cliPort.GetProductInfo(query)
			}
		}

		fmt.Printf("> > -------------- END c_%d.txt -------------- < < \n", i)
	}
}
