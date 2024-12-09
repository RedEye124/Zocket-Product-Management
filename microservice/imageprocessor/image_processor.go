package imageprocessor

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nfnt/resize"
)

// DownloadImage downloads an image from the URL
func DownloadImage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	filename := "downloaded_image.jpg"
	outFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// CompressImage compresses the downloaded image
func CompressImage(inputPath string) (string, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return "", err
	}

	compressedImg := resize.Resize(800, 0, img, resize.Lanczos3)
	compressedPath := "compressed_" + inputPath
	outFile, err := os.Create(compressedPath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, compressedImg, nil)
	if err != nil {
		return "", err
	}

	return compressedPath, nil
}

// UploadToS3 uploads the compressed image to AWS S3 and returns the URL
func UploadToS3(compressedPath string) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		return "", err
	}

	file, err := os.Open(compressedPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	s3Client := s3.New(sess)
	bucket := "your-s3-bucket-name"
	key := "compressed_images/" + compressedPath

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    aws.String(s3.ObjectCannedACLPublicRead),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, key), nil
}
