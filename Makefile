API_DIR = /Users/michaeld/src/api-repo
GOOGLE_PROTOS = /Users/michaeld/src/api-common-protos
GOPATH = /Users/michaeld/go

all: go gateway swaggerjson swaggerhtml markdown html

go:
	for proto in `find $(API_DIR) -name '*.proto'`; do protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --go_out=plugins=grpc,paths=source_relative:. $$proto; done

lint:
	for proto in `find $(API_DIR) -name '*.proto'`; do protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --l\
int_out=sort_imports:. $$proto; done

gateway:
	for proto in `find $(API_DIR) -name '*.proto'`; do protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=paths=source_relative:. $$proto; done

swaggerjson:
	for proto in `find $(API_DIR) -name '*.proto'`; do protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --swagger_out=logtostderr=true:. $$proto; done

swaggerhtml:
	for json in `find . -name '*.swagger.json'`; do pushd `dirname $$json`; /Users/michaeld/apps/nodejs/latest/bin/node_modules/.bin/redoc-cli bundle `basename $$json`; popd; done

markdownsuppress:	
	for protodir in `find $(API_DIR) -type d`; do if [ -e $(protodir)/*.proto ] ; then protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --doc_out=.$${protodir#"$(API_DIR)"} --doc_opt=markdown,api.md $$protodir/*.proto ; fi ; done

markdown:	
	for protodir in `find $(API_DIR) -type d`; do protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --doc_out=.$${protodir#"$(API_DIR)"} --doc_opt=markdown,api.md $$protodir/*.proto ; done

html:
	find $(API_DIR) -name '*.proto' | xargs protoc --proto_path=$(API_DIR) -I$(GOOGLE_PROTOS) --doc_out=. --doc_opt=html,index.html 

clean:
	rm -rf ./api
