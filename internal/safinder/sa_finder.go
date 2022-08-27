package safinder

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func SAFinder() {
	fmt.Println("Service Account Finder")
}

func resultOutput(sgName, sgId, desc, vpcId string) {
	tw := table.NewWriter()
	headerTransformer := text.Transformer(func(val interface{}) string {
		return text.Bold.Sprint(val)
	})
	tw.AppendHeader(table.Row{"SG Name", "SG Id", "Description", "VpcId"})
	tw.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:              "#",
			WidthMin:          2,
			WidthMax:          4,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Account Id",
			WidthMax:          12,
			WidthMin:          12,
			AutoMerge:         true,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Cluster",
			WidthMax:          40,
			WidthMin:          40,
			TransformerHeader: headerTransformer,
		},
		{
			Name:              "Description",
			WidthMax:          40,
			WidthMin:          40,
			AutoMerge:         true,
			TransformerHeader: headerTransformer,
		},
	})
	tables := table.Row{
		sgName, sgId, desc, vpcId,
	}
	tw.SetStyle(table.StyleLight)
	tw.Style().Options.SeparateRows = true
	tw.AppendRow(tables)
	tw.SetIndexColumn(1)
	tw.SetTitle("Security Group")
	fmt.Println(tw.Render())
}
