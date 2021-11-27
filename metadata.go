package dropshipping

import (
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	ApiBaseUrl = "https://developers.cjdropshipping.com"
)

type (
	Client struct {
		sync.Mutex
		Email               string
		Password            string
		Http                *http.Client
		Logger              io.Writer
		AccessToken         *AccessTokenResponse
		AccessTokenExpiryAt time.Time
	}

	AccessTokenResponse struct {
		Code    int    `json:"code"`
		Result  bool   `json:"result"`
		Message string `json:"message"`
		Data    struct {
			AccessToken            string    `json:"accessToken"`
			AccessTokenExpiryDate  time.Time `json:"accessTokenExpiryDate"`
			RefreshToken           string    `json:"refreshToken"`
			RefreshTokenExpiryDate time.Time `json:"refreshTokenExpiryDate"`
			CreateDate             time.Time `json:"createDate"`
		} `json:"data"`
		RequestId string `json:"requestId"`
	}

	ProductQueryResponse struct {
		Code    int    `json:"code"`
		Result  bool   `json:"result"`
		Message string `json:"message"`
		Data    struct {
			Pid            string               `json:"pid"`
			ProductName    string               `json:"productName"`
			ProductNameEn  string               `json:"productNameEn"`
			ProductSku     string               `json:"productSku"`
			ProductImage   string               `json:"productImage"`
			ProductWeight  string               `json:"productWeight"`
			ProductUnit    string               `json:"productUnit"`
			ProductType    string               `json:"productType"`
			CategoryId     string               `json:"categoryId"`
			CategoryName   string               `json:"categoryName"`
			EntryCode      string               `json:"entryCode"`
			EntryName      string               `json:"entryName"`
			EntryNameEn    string               `json:"entryNameEn"`
			MaterialName   string               `json:"materialName"`
			MaterialNameEn string               `json:"materialNameEn"`
			MaterialKey    string               `json:"materialKey"`
			PackingWeight  string               `json:"packingWeight"`
			PackingName    string               `json:"packingName"`
			PackingNameEn  string               `json:"packingNameEn"`
			PackingKey     string               `json:"packingKey"`
			ProductKey     string               `json:"productKey"`
			ProductKeyEn   string               `json:"productKeyEn"`
			ProductPro     string               `json:"productPro"`
			ProductProEn   []string             `json:"productProEn"`
			SellPrice      string               `json:"sellPrice"`
			SourceFrom     int                  `json:"sourceFrom"`
			Description    string               `json:"description"`
			Variants       []ProductVariantData `json:"variants"`
			AddMarkStatus  int                  `json:"addMarkStatus"`
			CreaterTime    time.Time            `json:"createrTime"`
		} `json:"data"`
		RequestId string `json:"requestId"`
	}
	ProductVariantResponse struct {
		Code      int                  `json:"code"`
		Result    bool                 `json:"result"`
		Message   string               `json:"message"`
		Data      []ProductVariantData `json:"data"`
		RequestId string               `json:"requestId"`
	}
	ProductVariantByIdResponse struct {
		Code      int                `json:"code"`
		Result    bool               `json:"result"`
		Message   string             `json:"message"`
		Data      ProductVariantData `json:"data"`
		RequestId string             `json:"requestId"`
	}
	ProductVariantData struct {
		Vid              string    `json:"vid"`
		Pid              string    `json:"pid"`
		VariantName      string    `json:"variantName"`
		VariantNameEn    string    `json:"variantNameEn"`
		VariantImage     string    `json:"variantImage"`
		VariantSku       string    `json:"variantSku"`
		VariantUnit      string    `json:"variantUnit"`
		VariantProperty  string    `json:"variantProperty"`
		VariantKey       string    `json:"variantKey"`
		VariantLength    int       `json:"variantLength"`
		VariantWidth     int       `json:"variantWidth"`
		VariantHeight    int       `json:"variantHeight"`
		VariantVolume    int       `json:"variantVolume"`
		VariantWeight    float64   `json:"variantWeight"`
		VariantSellPrice float64   `json:"variantSellPrice"`
		CreateTime       time.Time `json:"createTime"`
	}

	ProductSearchRequest struct {
		PageNum         int    `json:"pageNum"`
		PageSize        int    `json:"pageSize"`
		CategoryId      string `json:"categoryId"`
		CategoryKeyword string `json:"categoryKeyword"`
		Pid             string `json:"pid"`
		ProductSku      string `json:"productSku"`
		ProductType     string `json:"productType"`
	}
	ProductSearchResponse struct {
		Code    int    `json:"code"`
		Result  bool   `json:"result"`
		Message string `json:"message"`
		Data    struct {
			PageNum  int `json:"pageNum"`
			PageSize int `json:"pageSize"`
			Total    int `json:"total"`
			List     []struct {
				Pid           string `json:"pid"`
				ProductName   string `json:"productName"`
				ProductNameEn string `json:"productNameEn"`
				ProductSku    string `json:"productSku"`
				ProductImage  string `json:"productImage"`
				ProductWeight string `json:"productWeight"`
				ProductType   string `json:"productType"`
				ProductUnit   string `json:"productUnit"`
				SellPrice     string `json:"sellPrice"`
				CategoryId    string `json:"categoryId"`
				CategoryName  string `json:"categoryName"`
				SourceFrom    int    `json:"sourceFrom"`
				Remark        string `json:"remark"`
				CreateTime    string `json:"createTime"`
			} `json:"list"`
		} `json:"data"`
		RequestId string `json:"requestId"`
	}
)
