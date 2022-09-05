/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/7/30 18:00
 */
package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	contextPackage = protogen.GoImportPath("context")
	application    = protogen.GoImportPath("github.com/go-warrior/pkg/application")
)

func generateFile(gen *protogen.Plugin, file *protogen.File, hasHttp bool) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}

	filename := file.GeneratedFilenamePrefix + ".go"

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-warrior. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-warrior v", version)
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g, hasHttp)
	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, hasHttp bool) {
	if len(file.Services) == 0 {
		return
	}

	for _, service := range file.Services {
		genService(gen, file, g, service, hasHttp)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, hasHttp bool) {
	// Server interface.
	serverType := "I" + service.GoName + "UseCase"
	g.P("// ", serverType, " is the server API for ", service.GoName, " service.")
	g.P("// for forward compatibility")
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}
	g.Annotate(serverType, service.Location)
	g.P("type ", serverType, " interface {")
	for _, method := range service.Methods {
		g.Annotate(serverType+"."+method.GoName, method.Location)
		if method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated() {
			g.P(deprecationComment)
		}

		g.P(serverSignature(g, method))
	}
	g.P("}")
	g.P()
	g.P("// @di di文件根据注解自动生成")
	g.P("func Register", service.GoName, "Server(app ", application.Ident("Application"), ", srv ", serverType, ") {")
	if hasHttp {
		g.P("if app.HttpServer() != nil {")
		g.P("Register", service.GoName, "HTTPServer(app.HttpServer(), srv)")
		g.P("}")
	}

	g.P("if app.GrpcServer() != nil {")
	g.P("Register", service.GoName, "GrpcServer(app.GrpcServer(), srv)")
	g.P("}")
	g.P("}")
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	var reqArgs []string
	ret := "error"
	if !method.Desc.IsStreamingClient() && !method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, g.QualifiedGoIdent(contextPackage.Ident("Context")))
		ret = "(*" + g.QualifiedGoIdent(method.Output.GoIdent) + ", error)"
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "*"+g.QualifiedGoIdent(method.Input.GoIdent))
	}
	if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
		reqArgs = append(reqArgs, method.Parent.GoName+"_"+method.GoName+"Server")
	}
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

const deprecationComment = "// Deprecated: Do not use."