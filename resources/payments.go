package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/dataqueen-center/cq-source-fincode/client"
)

type PaymentsResponse struct {
	ShopID          string  `json:"shop_id"`
	ID              string  `json:"id"`
	PayType         string  `json:"pay_type"`
	Status          string  `json:"status"`
	AccessID        string  `json:"access_id"`
	ProcessDate     string  `json:"process_date"`
	JobCode         string  `json:"job_code"`
	ItemCode        string  `json:"item_code"`
	Amount          float64 `json:"amount"`
	Tax             float64 `json:"tax"`
	TotalAmount     float64 `json:"total_amount"`
	CustomerGroupID string  `json:"customer_group_id"`
	CustomerID      string  `json:"customer_id"`
	CardNo          string  `json:"card_no"`
	CardID          string  `json:"card_id"`
	Expire          string  `json:"expire"`
	HolderName      string  `json:"holder_name"`
	CardNoHash      string  `json:"card_no_hash"`
	Method          string  `json:"method"`
	PayTimes        string  `json:"pay_times"`
	Forward         string  `json:"forward"`
	Issuer          string  `json:"issuer"`
	TransactionID   string  `json:"transaction_id"`
	Approve         string  `json:"approve"`
	AuthMaxDate     string  `json:"auth_max_date"`
	ClientField1    string  `json:"client_field_1"`
	ClientField2    string  `json:"client_field_2"`
	ClientField3    string  `json:"client_field_3"`
	TdsType         string  `json:"tds_type"`
	Tds2Type        string  `json:"tds2_type"`
	Tds2RetUrl      string  `json:"tds2_ret_url"`
	Tds2Status      string  `json:"tds2_status"`
	MerchantName    string  `json:"merchant_name"`
	SendUrl         string  `json:"send_url"`
	SubscriptionID  string  `json:"subscription_id"`
	Brand           string  `json:"brand"`
	ErrorCode       string  `json:"error_code"`
	Created         string  `json:"created"`
	Updated         string  `json:"updated"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

func Payments() *schema.Table {
	return &schema.Table{
		Name:     "fincode_payments",
		Resolver: fetchPayments,
		Columns: []schema.Column{
			{Name: "shop_id", Type: arrow.BinaryTypes.String},
			{Name: "id", Type: arrow.BinaryTypes.String},
			{Name: "pay_type", Type: arrow.BinaryTypes.String},
			{Name: "status", Type: arrow.BinaryTypes.String},
			{Name: "access_id", Type: arrow.BinaryTypes.String},
			{Name: "process_date", Type: arrow.BinaryTypes.String},
			{Name: "job_code", Type: arrow.BinaryTypes.String},
			{Name: "item_code", Type: arrow.BinaryTypes.String},
			{Name: "amount", Type: arrow.PrimitiveTypes.Float64},
			{Name: "tax", Type: arrow.PrimitiveTypes.Float64},
			{Name: "total_amount", Type: arrow.PrimitiveTypes.Float64},
			{Name: "customer_group_id", Type: arrow.BinaryTypes.String},
			{Name: "customer_id", Type: arrow.BinaryTypes.String},
			{Name: "card_no", Type: arrow.BinaryTypes.String},
			{Name: "card_id", Type: arrow.BinaryTypes.String},
			{Name: "expire", Type: arrow.BinaryTypes.String},
			{Name: "holder_name", Type: arrow.BinaryTypes.String},
			{Name: "card_no_hash", Type: arrow.BinaryTypes.String},
			{Name: "method", Type: arrow.BinaryTypes.String},
			{Name: "pay_times", Type: arrow.BinaryTypes.String},
			{Name: "forward", Type: arrow.BinaryTypes.String},
			{Name: "issuer", Type: arrow.BinaryTypes.String},
			{Name: "transaction_id", Type: arrow.BinaryTypes.String},
			{Name: "approve", Type: arrow.BinaryTypes.String},
			{Name: "auth_max_date", Type: arrow.BinaryTypes.String},
			{Name: "client_field_1", Type: arrow.BinaryTypes.String},
			{Name: "client_field_2", Type: arrow.BinaryTypes.String},
			{Name: "client_field_3", Type: arrow.BinaryTypes.String},
			{Name: "tds_type", Type: arrow.BinaryTypes.String},
			{Name: "tds2_type", Type: arrow.BinaryTypes.String},
			{Name: "tds2_ret_url", Type: arrow.BinaryTypes.String},
			{Name: "tds2_status", Type: arrow.BinaryTypes.String},
			{Name: "merchant_name", Type: arrow.BinaryTypes.String},
			{Name: "send_url", Type: arrow.BinaryTypes.String},
			{Name: "subscription_id", Type: arrow.BinaryTypes.String},
			{Name: "brand", Type: arrow.BinaryTypes.String},
			{Name: "error_code", Type: arrow.BinaryTypes.String},
			{Name: "created_at", Type: arrow.BinaryTypes.String},
			{Name: "updated_at", Type: arrow.BinaryTypes.String},
		},
	}
}

func fetchPayments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	const processDateFrom = "2023/01/01"
	const endpoint = "/v1/payments"
	const method = "GET"
	total_resources := 0

	// TODO: Change to []string{"Card", "Konbini", "Paypay"}
	for _, payType := range []string{"Card"} {
		page := 1
		// go func() {
		for {
			paymentsPage, err := fetchPaymentsByPaymentType(client, endpoint, method, payType, processDateFrom, page)
			if err != nil {
				return err
			}
			client.Logger.Info().Msgf("Pages: %d/%d (%s)", paymentsPage.CurrentPage, paymentsPage.LastPage, payType)
			client.Logger.Info().Msgf("Total count: %d (%s)", paymentsPage.TotalCount, payType)
			for _, payment := range paymentsPage.List {
				client.Logger.Debug().Msgf("Payment: %v", payment)
				// parse date: 2023/05/09 20:03:32.528
				created_at, err := time.Parse("2006/01/02 15:04:05.000", payment.Created)
				if err != nil {
					return err
				}
				payment.CreatedAt = created_at.Format(time.RFC3339Nano)
				updated_at, err := time.Parse("2006/01/02 15:04:05.000", payment.Updated)
				if err != nil {
					return err
				}
				payment.UpdatedAt = updated_at.Format(time.RFC3339Nano)
				res <- payment
				total_resources++
			}
			if paymentsPage.LastPage == paymentsPage.CurrentPage || total_resources >= paymentsPage.TotalCount {
				break
			}
			page++
		}
		// }()
	}
	client.Logger.Info().Msg(fmt.Sprintf("Total resources: %d", total_resources))

	return nil
}

func fetchPaymentsByPaymentType(client *client.Client, endpoint string, method string, payType string, processDateFrom string, page int) (*PageWrapper[PaymentsResponse], error) {
	URL := client.BaseUrl + endpoint
	req, err := http.NewRequest(
		method,
		URL,
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	params := req.URL.Query()
	params.Add("page", fmt.Sprintf("%d", page))
	params.Add("limit", "50")
	params.Add("sort", "created")
	params.Add("pay_type", payType)
	params.Add("process_date_from", processDateFrom)
	req.URL.RawQuery = params.Encode()

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		client.Logger.Error().Msg(string(bodyBytes))
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var paymentsMap PageWrapper[PaymentsResponse]
	err = json.Unmarshal(bodyBytes, &paymentsMap)
	if err != nil {
		return nil, err
	}
	return &paymentsMap, nil
}
