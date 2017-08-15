// Code generated by 'github.com/sacloud/usacloud/tools/gen-command-completion'; DO NOT EDIT

package completion

import (
	"fmt"
	"github.com/sacloud/usacloud/command"
	"github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/schema"
)

func BridgeListCompleteFlags(ctx command.Context, params *params.ListBridgeParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Bridge"].Commands["list"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Bridge"].Commands["list"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "from", "offset":
		param := define.Resources["Bridge"].Commands["list"].BuildedParams().Get("from")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "max", "limit":
		param := define.Resources["Bridge"].Commands["list"].BuildedParams().Get("max")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "sort":
		param := define.Resources["Bridge"].Commands["list"].BuildedParams().Get("sort")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func BridgeCreateCompleteFlags(ctx command.Context, params *params.CreateBridgeParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Bridge"].Commands["create"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["Bridge"].Commands["create"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func BridgeReadCompleteFlags(ctx command.Context, params *params.ReadBridgeParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["Bridge"].Commands["read"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func BridgeUpdateCompleteFlags(ctx command.Context, params *params.UpdateBridgeParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "name":
		param := define.Resources["Bridge"].Commands["update"].BuildedParams().Get("name")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "description", "desc":
		param := define.Resources["Bridge"].Commands["update"].BuildedParams().Get("description")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "id":
		param := define.Resources["Bridge"].Commands["update"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}

func BridgeDeleteCompleteFlags(ctx command.Context, params *params.DeleteBridgeParam, flagName string, currentValue string) {
	var comp schema.CompletionFunc

	switch flagName {
	case "id":
		param := define.Resources["Bridge"].Commands["delete"].BuildedParams().Get("id")
		if param != nil {
			comp = param.Param.CompleteFunc
		}
	case "output-type", "out":
		comp = schema.CompleteInStrValues("json", "csv", "tsv")
	}

	if comp != nil {
		words := comp(ctx, currentValue)
		for _, w := range words {
			fmt.Println(w)
		}
	}
}