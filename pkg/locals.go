package pkg

import (
	"github.com/plantoncloud/project-planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/atlas/mongodbatlas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	MongodbAtlas *mongodbatlas.MongodbAtlas
}

func initializeLocals(ctx *pulumi.Context, stackInput *mongodbatlas.MongodbAtlasStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.MongodbAtlas = stackInput.Target

	return locals
}
