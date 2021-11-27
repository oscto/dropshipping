package main

import (
	"context"
	"fmt"
	"log"

	"github.com/oscto/dropshipping"
)

func main() {
	client, err := dropshipping.NewClient("", "", false)
	if err != nil {
		fmt.Println("NewClient err", err)
	}

	client.SetLogger(log.Writer())

	token, _ := client.GetAccessToken(context.Background())
	fmt.Println("token", token.Data.AccessToken)

	var query dropshipping.ProductSearchRequest
	query.CategoryKeyword = "puzzle"
	query.PageSize = 10
	query.PageNum = 0
	search, err := client.ProductSearch(context.Background(), query)
	fmt.Println(err, "search", search.Data.Total)

	productQuery, err := client.ProductQuery(context.Background(), "", "CJJJJTJT05843")
	fmt.Println(err, "ProductQuery", productQuery.Data.ProductName)

	variantQuery, err := client.ProductVariantQuery(context.Background(), "", "CJJJJTJT05843")
	fmt.Println(err, "ProductVariantQuery", variantQuery.Data[0].VariantNameEn)

	id, _ := client.ProductVariantById(context.Background(), "D4057F56-3F09-4541-8461-9D76D014846D")
	fmt.Println("ProductVariantById", id.Data.VariantNameEn)
}
