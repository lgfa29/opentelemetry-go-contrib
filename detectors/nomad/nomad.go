package nomad

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

type resourceDetector struct{}

const (
	NomadRegionKey     = attribute.Key("nomad.region")
	NomadDatacenterKey = attribute.Key("nomad.dc")
	NomadNamespaceKey  = attribute.Key("nomad.namespace")

	NomadJobParentIDKey = attribute.Key("nomad.job_parent_id")
	NomadJobIDKey       = attribute.Key("nomad.job_id")
	NomadJobNameKey     = attribute.Key("nomad.job_name")
	NomadGroupNameKey   = attribute.Key("nomad.group_name")
	NomadTaskNameKey    = attribute.Key("nomad.task_name")

	NomadAllocIDKey    = attribute.Key("nomad.alloc_id")
	NomadAllocNameKey  = attribute.Key("nomad.alloc_name")
	NomadAllocIndexKey = attribute.Key("nomad.alloc_index")
)

var _ resource.Detector = (*resourceDetector)(nil)

func NewResourceDetector() resource.Detector {
	return &resourceDetector{}
}

func (d *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	attributes := []attribute.KeyValue{
		NomadRegionKey.String(os.Getenv("NOMAD_REGION")),
		NomadDatacenterKey.String(os.Getenv("NOMAD_DC")),
		NomadNamespaceKey.String(os.Getenv("NOMAD_NAMESPACE")),
		NomadJobParentIDKey.String(os.Getenv("NOMAD_JOB_PARENT_ID")),
		NomadJobIDKey.String(os.Getenv("NOMAD_JOB_ID")),
		NomadJobNameKey.String(os.Getenv("NOMAD_JOB_NAME")),
		NomadGroupNameKey.String(os.Getenv("NOMAD_GROUP_NAME")),
		NomadTaskNameKey.String(os.Getenv("NOMAD_TASK_NAME")),
		NomadAllocIDKey.String(os.Getenv("NOMAD_ALLOC_ID")),
		NomadAllocNameKey.String(os.Getenv("NOMAD_ALLOC_NAME")),
		NomadAllocIndexKey.String(os.Getenv("NOMAD_ALLOC_INDEX")),
	}

	return resource.NewWithAttributes(semconv.SchemaURL, attributes...), nil
}
