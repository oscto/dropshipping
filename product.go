package dropshipping

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ProductSearch
// Endpoint /api2.0/v1/product/list
func (c *Client) ProductSearch(ctx context.Context, search ProductSearchRequest) (*ProductSearchResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", ApiBaseUrl, UriProductSearch), nil)
	if err != nil {
		return nil, err
	}
	payload := url.Values{}
	payload.Add("pageNum", strconv.Itoa(search.PageNum))
	payload.Add("pageSize", strconv.Itoa(search.PageSize))
	payload.Add("categoryId", search.CategoryId)
	payload.Add("categoryKeyword", search.CategoryKeyword)
	payload.Add("pid", search.Pid)
	payload.Add("productSku", search.ProductSku)
	payload.Add("productType", search.ProductType)
	req.URL.RawQuery = payload.Encode()

	//rep := &ProductSearchResponse{}
	rep := &ProductSearchResponse{}
	if err := c.SendWithAuth(req, rep); err != nil {
		fmt.Println("xxxxxxerrrrrr", err)
		return nil, err
	}
	fmt.Println("rrrrrrrrrrrrrrrrr\n\n\n\n\n\n\n\n", rep)

	return rep, nil
}

// ProductQuery
// Endpoint /api2.0/v1/product/query
func (c *Client) ProductQuery(ctx context.Context, pid, sku string) (*ProductQueryResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", ApiBaseUrl, UriProductQuery), nil)
	if err != nil {
		return nil, err
	}
	payload := url.Values{}
	payload.Add("pid", pid)
	payload.Add("productSku", sku)
	req.URL.RawQuery = payload.Encode()

	rep := &ProductQueryResponse{}
	if err := c.SendWithAuth(req, rep); err != nil {
		return nil, err
	}

	return rep, nil
}

// ProductVariantQuery
// Endpoint /api2.0/v1/product/variant/query
func (c *Client) ProductVariantQuery(ctx context.Context, pid, sku string) (*ProductVariantResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", ApiBaseUrl, UriProductVariantQuery), nil)
	if err != nil {
		return nil, err
	}
	payload := url.Values{}
	payload.Add("pid", pid)
	payload.Add("productSku", sku)
	req.URL.RawQuery = payload.Encode()

	rep := &ProductVariantResponse{}
	if err := c.SendWithAuth(req, rep); err != nil {
		return nil, err
	}

	return rep, nil
}

// ProductVariantById
// Endpoint /api2.0/v1/product/variant/queryByVid
func (c *Client) ProductVariantById(ctx context.Context, vid string) (*ProductVariantByIdResponse, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s%s", ApiBaseUrl, UriProductVariantById), nil)
	if err != nil {
		return nil, err
	}
	payload := url.Values{}
	payload.Add("vid", vid)
	req.URL.RawQuery = payload.Encode()

	rep := &ProductVariantByIdResponse{}
	if err := c.SendWithAuth(req, rep); err != nil {
		return nil, err
	}

	return rep, nil
}
