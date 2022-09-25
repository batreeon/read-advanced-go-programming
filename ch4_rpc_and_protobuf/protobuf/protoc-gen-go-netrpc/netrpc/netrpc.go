package netrpc

import (
	"bytes"
	"log"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go-netrpc/descriptor"
	"github.com/golang/protobuf/protoc-gen-go-netrpc/generator"
)

type netrpcPlugin struct{ *generator.Generator }

func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}

func (p *netrpcPlugin) Name() string                { return "netrpc" }
func (p *netrpcPlugin) Init(g *generator.Generator) { p.Generator = g }
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "net/rpc"`)
}

type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}
type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

func (p *netrpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}
	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}
	return spec
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)
	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}
	p.P(buf.String())
}
