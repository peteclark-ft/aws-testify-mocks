# AWS SDK Testify Mocks

[Testify](https://github.com/stretchr/testify) is a widely used testing framework. S3 is a widely used file store. This go library mocks the Amazon [S3 SDK](https://github.com/aws/aws-sdk-go) so you don't have to.

## Usage

Same as any testify mock:

```
import (
   s3 "github.com/peteclark-ft/aws-testify-mocks"
)

func TestS3(t *testing.T){
   mockS3 := new(MockS3API)

   input := &GetObjectInput{...}
   output := &GetObjectOutput{...}

   mockS3.On("GetObject", input).Return(output, nil)

   // use mockS3!
}
```
