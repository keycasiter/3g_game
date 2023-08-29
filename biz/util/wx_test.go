package util

import (
	"context"
	"fmt"
	"testing"
)

func TestDecryptWXOpenData(t *testing.T) {
	m, err := DecryptWXOpenData(
		"wx6d8a318f20ba48f1",
		"8ABV7W23PhR05ukKFWXooA==",
		"Z2W9LYcubVxXZvrzh0nUIDuHV0IOc0p7ZCXZJ34xeOSGCXNk8GbQNVejdqicMwdS1Ns/tCZiwL96jRqW/HubHwni2rHfqPDKeOriWZ59WFfzTm8uwlxbZxGkygptIBYX/fCdDKG6DQLlU8/qNMddiVcghuRo2jZiFohOJIs0GJ9dbsGY5Uryfd8dB7Tx9rmQNcEokDQ3x06pr7SHgCSvmgupTKEZnItY/owjnQnvg/rursSUz7zrPL8uSdKUWsBH+aUFFbDg5K5xYdWsz1+eFG86Hqr1LVBNGIu0kT0Sop+LKu1NcxstrdBDe+SRZKwkxsulSqkq13SempfSlq6txgpZ0QFdBn+tt6/rDT5nZnbf0LVmSW5fwbsUGMTowLTikin7Y8GkBvSYjgwBulvhfjpEyaHhDI3TlNR4MqQ8HbVIqqrE0OIpjWKwzBMzbsRt+Svp12LuNPhLo/6igN4w5Q==",
		"/QH3cb/+V20FxQWHg0rd4g==")
	if err != nil {
		fmt.Errorf("err:%v", err)
		t.Fail()
	}
	fmt.Printf("m:%s", ToJsonString(context.Background(), m))
}
