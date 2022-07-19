package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/luraproject/lura/v2/logging"
)

type S3Client struct {
	client *s3.Client
	logger logging.Logger
}

func NewS3Client(l logging.Logger) (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		r := os.Getenv("AWS_DEFAULT_REGION")
		e := os.Getenv("S3_ENDPOINT")

		l.Debug(fmt.Sprintf("using s3 endpoint %s and region %s", e, r))
		opts.Region = r
		opts.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               e,
				SigningRegion:     region,
				HostnameImmutable: true,
			}, nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	c := s3.NewFromConfig(cfg)
	return &S3Client{
		client: c,
		logger: l,
	}, nil
}

func (s S3Client) GetBytes(ctx context.Context, bucket string, key string) ([]byte, error) {
	if errors.Is(ctx.Err(), context.Canceled) {
		return nil, ctx.Err()
	}
	s.logger.Debug(fmt.Sprintf("request to fetch %s from %s received", key, bucket))

	o, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	if errors.Is(ctx.Err(), context.Canceled) {
		return nil, ctx.Err()
	}

	s.logger.Debug(fmt.Sprintf("successfully fetched content"))
	defer o.Body.Close()

	body, err := ioutil.ReadAll(o.Body)
	s.logger.Debug(string(body))
	if err != nil {
		return nil, err
	}

	if errors.Is(ctx.Err(), context.Canceled) {
		return nil, ctx.Err()
	}
	return body, nil
}
