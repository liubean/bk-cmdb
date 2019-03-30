

package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/types"
	"configcenter/src/common/util"
	"configcenter/src/source_controller/hostcontroller/app"
	"configcenter/src/source_controller/hostcontroller/app/options"
	"github.com/spf13/pflag"
)

func main() {
	common.SetIdentification(types.CC_MODULE_HOSTCONTROLLER)
	runtime.GOMAXPROCS(runtime.NumCPU())

	blog.InitLogs()
	defer blog.CloseLogs()

	op := options.NewServerOption()
	op.AddFlags(pflag.CommandLine)

	util.InitFlags()

	if err := common.SavePid(); err != nil {
		blog.Error("fail to save pid: err:%s", err.Error())
	}

	if err := app.Run(context.Background(), op); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		blog.Fatal(err)
	}
}
