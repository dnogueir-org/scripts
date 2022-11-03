package main
import "fmt"

// https://yantar.prd.gcp.gruposbf.com.br/api/navigation?resultsPerPage=40&page=1&scoringProfile=scoreByRanking&fields=genero:masculino&bu=centauro

// base params
var (
	baseUrl = "https://yantar.prd.gcp.gruposbf.com.br/api"
	fixedParams = [4]string{
		"resultsPerPage=40",
		"page=1",
		"scoringProfile=scoreByRanking",
		"bu=centauro",
	}
)

var filters = make(map[string][]string)

// 


func main() {
	fmt.Println(baseUrl)
	fmt.Println(fixedParams)

	filters["genero"] = genero
	filters["tamanho"] = tamanho
	filters["marca"] = marca

	fmt.Println(genero)

}