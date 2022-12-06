TEST_BUCKET := test-bucket
dev-list-s3-files:
	aws --endpoint-url=http://localhost:4566 s3 ls s3://${TEST_BUCKET}
dev-create-s3-bucket:
	aws --endpoint-url=http://localhost:4566 s3 mb s3://${TEST_BUCKET}
dev-sync-s3-files:
	aws --endpoint-url=http://localhost:4566 s3 sync .docker/s3/ s3://${TEST_BUCKET}
build-plugins:
	@for d in plugins/* ; do go build -buildmode=plugin -o "./bin/$$(basename $${d}).so" "./$${d}"; done
