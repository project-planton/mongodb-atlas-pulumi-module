package main

import (
	mongodbatlasv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/atlas/mongodbatlas/v1"
	"github.com/pkg/errors"
	"github.com/project-planton/mongodb-atlas-pulumi-module/pkg"
	"github.com/project-planton/pulumi-module-golang-commons/pkg/stackinput"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		stackInput := &mongodbatlasv1.MongodbAtlasStackInput{}

		if err := stackinput.LoadStackInput(ctx, stackInput); err != nil {
			return errors.Wrap(err, "failed to load stack-input")
		}

		return pkg.Resources(ctx, stackInput)
	})
}
