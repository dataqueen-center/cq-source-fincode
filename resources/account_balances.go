package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/dataqueen-center/cq-source-fincode/client"
)

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
			{Name: "amount", Type: arrow.BinaryTypes.String},
			{Name: "tax", Type: arrow.BinaryTypes.String},
			{Name: "total_amount", Type: arrow.BinaryTypes.String},
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
			{Name: "created", Type: arrow.BinaryTypes.String},
			{Name: "updated", Type: arrow.BinaryTypes.String},
		},
	}
}

func fetchPayments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	jsonBytes, statusCode, err := client.Execute("GET", "/v1/payments")
	if err != nil {
		return err
	}
	if statusCode != 200 {
		return fmt.Errorf("unexpected status code: %d", statusCode)
	}

	var paymentsMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &paymentsMap)
	if err != nil {
		return err
	}

	res <- paymentsMap

	return nil
}
