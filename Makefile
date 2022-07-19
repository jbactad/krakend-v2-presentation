dev-create-s3-bucket:
	aws --endpoint-url=http://localhost:4566 s3 mb s3://test-bucket
dev-sync-s3-files:
	aws --endpoint-url=http://localhost:4566 s3 sync .docker/s3/ s3://test-bucket