package index

import (
	"ycoem/lib/db"
	model "ycoem/lib/models"
	"ycoem/lib/service/abstract"
	"ycoem/lib/service/cases"
	"ycoem/lib/service/firm"
	"ycoem/lib/service/product"
	"ycoem/lib/service/service"
	"ycoem/lib/service/solution"

	jsoniter "github.com/json-iterator/go"
)

type IndexService struct {
}

type Node2 struct {
	URL  string `json:"url"`
	Text string `json:"text"`
}

type Nodes struct {
	URL  string  `json:"url"`
	Node []Node2 `json:"node"`
}

func (this IndexService) CreateMenuBar() (err error) {
	//先把产品 拼出来
	casesService := cases.CasesService{}
	firmService := firm.FirmService{}
	serviceService := service.ServiceService{}
	productService := product.ProductService{}
	solutionService := solution.SolutionService{}

	tableList := []abstract.Service{casesService, firmService, serviceService, productService, solutionService}

	nameMap := map[string]string{
		casesService.TableName():    "案例",
		firmService.TableName():     "关于我们",
		serviceService.TableName():  "服务支持",
		productService.TableName():  "产品",
		solutionService.TableName(): "解决方案",
	}

	parentUrl := "/"
	menubar := make(map[string]Nodes)
	for k := range tableList {
		table := tableList[k]
		list, err := abstract.Options("name", table)
		if err != nil {
			continue
		}
		nodes := []Node2{}

		for k2 := range list {
			var node Node2
			node.Text = list[k2]
			node.URL = "/" + table.TableName() + "/" + node.Text
			nodes = append(nodes, node)
		}

		if table.TableName() == "firm" {
			var node Node2
			node.Text = "公司介绍"
			node.URL = "/" + table.TableName() + "/" + node.Text
			nodes = append(nodes, node)
		}

		if len(nodes) > 0 {
			parentUrl = nodes[0].URL
		}

		nodes2 := Nodes{
			URL:  parentUrl,
			Node: nodes,
		}

		menubar[nameMap[tableList[k].TableName()]] = nodes2
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b, err := json.Marshal(menubar)
	if err != nil {
		return
	}
	err = db.Set(model.NavbarModel, string(b))
	return
}
