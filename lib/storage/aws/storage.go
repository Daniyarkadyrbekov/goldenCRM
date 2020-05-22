package aws

import (
	"io/ioutil"
	"strings"

	"go.uber.org/zap"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage struct {
	conn   *s3.S3
	logger *zap.Logger
}

func New() Storage {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("ru-msk"),
		Endpoint: aws.String("http://hb.bizmrg.com"),
	}))

	return Storage{
		conn: s3.New(sess),
	}
}

func (s Storage) GetObject(bucket, key string) ([]byte, error) {
	resp, err := s.conn.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			s.logger.Error("err closing resp body", zap.Error(err))
		}
	}()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func (s Storage) PutObject(bucket, key, content string) error {
	_, err := s.conn.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(content),
	})

	return err
}
