package provider

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/geekcell/terraform-provider-tools/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataMapSanitizer() *schema.Resource {
	return &schema.Resource{
		Description: "This data source takes a map and removes all null values in the output.",
		ReadContext: dataMapSanitizerRead,
		Schema: map[string]*schema.Schema{
			"input": {
				Description: "The map to sanitize.",
				Type:        schema.TypeMap,
				Required:    true,
			},
			"output": {
				Description: "The sanitized map.",
				Type:        schema.TypeMap,
				Computed:    true,
			},
		},
	}
}

func dataMapSanitizerRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	inputMap := d.Get("input").(map[string]interface{})
	utils.RemoveNullValues(inputMap)

	hash := md5.Sum([]byte(fmt.Sprint(inputMap)))
	d.SetId(hex.EncodeToString(hash[:]))
	d.Set("output", inputMap)

	return nil
}
