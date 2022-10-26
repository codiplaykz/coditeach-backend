package helpers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jackc/pgx/v4"
	"mime/multipart"
)

func PgSqlRowsToJson(rows pgx.Rows) []map[string]interface{} {
	fieldDescriptions := rows.FieldDescriptions()
	var columns []string
	for _, col := range fieldDescriptions {
		columns = append(columns, string(col.Name))
	}

	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}

	return tableData
}

func UploadImageToSpace(filename string, file multipart.File) {
	key := "DO002GCU89VYXXQ8XV67"
	secret := "qcHdya6wOF/BtcTvpyPhlYsHFRGL/coJz0SxnXXyv6Y"

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String("https://fra1.digitaloceanspaces.com"),
		S3ForcePathStyle: aws.Bool(false),
		Region:           aws.String("fra1"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	object := s3.PutObjectInput{
		Bucket: aws.String("cp-space"),
		Key:    aws.String("projects-images/" + filename),
		Body:   file,
		ACL:    aws.String("public-read"),
	}

	test, err := s3Client.PutObject(&object)

	fmt.Println(test.ETag)

	if err != nil {
		fmt.Println(err.Error())
	}
}
