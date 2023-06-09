package resources

import (
	_ "embed"
	"testing"

	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/dataqueen-center/cq-source-fincode/client"
)

func TestPayments(t *testing.T) {
	var res PageWrapper[PaymentsResponse]
	if err := faker.FakeObject(&res); err != nil {
		t.Fatal(err)
	}

	ts := client.TestServer(t, res)

	defer ts.Close()
	client.TestHelper(t, Payments(), ts)
}
