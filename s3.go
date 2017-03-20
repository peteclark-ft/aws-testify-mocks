package s3

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/mock"
)

type MockS3API struct {
	mock.Mock
}

func NewMockS3API() s3iface.S3API {
	return new(MockS3API)
}

func (m *MockS3API) AbortMultipartUploadRequest(abort *s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput) {
	args := m.Called(abort)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.AbortMultipartUploadOutput)
}

func (m *MockS3API) AbortMultipartUpload(abort *s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error) {
	args := m.Called(abort)
	return args.Get(0).(*s3.AbortMultipartUploadOutput), args.Error(1)
}

func (m *MockS3API) CompleteMultipartUploadRequest(c *s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput) {
	args := m.Called(c)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.CompleteMultipartUploadOutput)
}

func (m *MockS3API) CompleteMultipartUpload(c *s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error) {
	args := m.Called(c)
	return args.Get(0).(*s3.CompleteMultipartUploadOutput), args.Error(1)
}

func (m *MockS3API) CopyObjectRequest(c *s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput) {
	args := m.Called(c)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.CopyObjectOutput)
}

func (m *MockS3API) CopyObject(c *s3.CopyObjectInput) (*s3.CopyObjectOutput, error) {
	args := m.Called(c)
	return args.Get(0).(*s3.CopyObjectOutput), args.Error(1)
}

func (m *MockS3API) CreateBucketRequest(c *s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput) {
	args := m.Called(c)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.CreateBucketOutput)
}

func (m *MockS3API) CreateBucket(c *s3.CreateBucketInput) (*s3.CreateBucketOutput, error) {
	args := m.Called(c)
	return args.Get(0).(*s3.CreateBucketOutput), args.Error(1)
}

func (m *MockS3API) CreateMultipartUploadRequest(c *s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput) {
	args := m.Called(c)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.CreateMultipartUploadOutput)
}

func (m *MockS3API) CreateMultipartUpload(c *s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error) {
	args := m.Called(c)
	return args.Get(0).(*s3.CreateMultipartUploadOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketRequest(d *s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketOutput)
}

func (m *MockS3API) DeleteBucket(d *s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketAnalyticsConfigurationRequest(d *s3.DeleteBucketAnalyticsConfigurationInput) (*request.Request, *s3.DeleteBucketAnalyticsConfigurationOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketAnalyticsConfigurationOutput)
}

func (m *MockS3API) DeleteBucketAnalyticsConfiguration(d *s3.DeleteBucketAnalyticsConfigurationInput) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketAnalyticsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketCorsRequest(d *s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketCorsOutput)
}

func (m *MockS3API) DeleteBucketCors(d *s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketCorsOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketInventoryConfigurationRequest(d *s3.DeleteBucketInventoryConfigurationInput) (*request.Request, *s3.DeleteBucketInventoryConfigurationOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketInventoryConfigurationOutput)
}

func (m *MockS3API) DeleteBucketInventoryConfiguration(d *s3.DeleteBucketInventoryConfigurationInput) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketInventoryConfigurationOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketLifecycleRequest(d *s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketLifecycleOutput)
}

func (m *MockS3API) DeleteBucketLifecycle(d *s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketLifecycleOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketMetricsConfigurationRequest(d *s3.DeleteBucketMetricsConfigurationInput) (*request.Request, *s3.DeleteBucketMetricsConfigurationOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketMetricsConfigurationOutput)
}

func (m *MockS3API) DeleteBucketMetricsConfiguration(d *s3.DeleteBucketMetricsConfigurationInput) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketMetricsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketPolicyRequest(d *s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketPolicyOutput)
}

func (m *MockS3API) DeleteBucketPolicy(d *s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketPolicyOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketReplicationRequest(d *s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketReplicationOutput)
}

func (m *MockS3API) DeleteBucketReplication(d *s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketReplicationOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketTaggingRequest(d *s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketTaggingOutput)
}

func (m *MockS3API) DeleteBucketTagging(d *s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketTaggingOutput), args.Error(1)
}

func (m *MockS3API) DeleteBucketWebsiteRequest(d *s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteBucketWebsiteOutput)
}

func (m *MockS3API) DeleteBucketWebsite(d *s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteBucketWebsiteOutput), args.Error(1)
}

func (m *MockS3API) DeleteObjectRequest(d *s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteObjectOutput)
}

func (m *MockS3API) DeleteObject(d *s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteObjectOutput), args.Error(1)
}

func (m *MockS3API) DeleteObjectTaggingRequest(d *s3.DeleteObjectTaggingInput) (*request.Request, *s3.DeleteObjectTaggingOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteObjectTaggingOutput)
}

func (m *MockS3API) DeleteObjectTagging(d *s3.DeleteObjectTaggingInput) (*s3.DeleteObjectTaggingOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteObjectTaggingOutput), args.Error(1)
}

func (m *MockS3API) DeleteObjectsRequest(d *s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.DeleteObjectsOutput)
}

func (m *MockS3API) DeleteObjects(d *s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.DeleteObjectsOutput), args.Error(1)
}

func (m *MockS3API) GetBucketAccelerateConfigurationRequest(g *s3.GetBucketAccelerateConfigurationInput) (*request.Request, *s3.GetBucketAccelerateConfigurationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketAccelerateConfigurationOutput)
}

func (m *MockS3API) GetBucketAccelerateConfiguration(g *s3.GetBucketAccelerateConfigurationInput) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketAccelerateConfigurationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketAclRequest(d *s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput) {
	args := m.Called(d)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketAclOutput)
}

func (m *MockS3API) GetBucketAcl(d *s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error) {
	args := m.Called(d)
	return args.Get(0).(*s3.GetBucketAclOutput), args.Error(1)
}

func (m *MockS3API) GetBucketAnalyticsConfigurationRequest(g *s3.GetBucketAnalyticsConfigurationInput) (*request.Request, *s3.GetBucketAnalyticsConfigurationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketAnalyticsConfigurationOutput)
}

func (m *MockS3API) GetBucketAnalyticsConfiguration(g *s3.GetBucketAnalyticsConfigurationInput) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketAnalyticsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketCorsRequest(g *s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketCorsOutput)
}

func (m *MockS3API) GetBucketCors(g *s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketCorsOutput), args.Error(1)
}

func (m *MockS3API) GetBucketInventoryConfigurationRequest(g *s3.GetBucketInventoryConfigurationInput) (*request.Request, *s3.GetBucketInventoryConfigurationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketInventoryConfigurationOutput)
}

func (m *MockS3API) GetBucketInventoryConfiguration(g *s3.GetBucketInventoryConfigurationInput) (*s3.GetBucketInventoryConfigurationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketInventoryConfigurationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketLifecycleRequest(g *s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketLifecycleOutput)
}

func (m *MockS3API) GetBucketLifecycle(g *s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketLifecycleOutput), args.Error(1)
}

func (m *MockS3API) GetBucketLifecycleConfigurationRequest(g *s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketLifecycleConfigurationOutput)
}

func (m *MockS3API) GetBucketLifecycleConfiguration(g *s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketLifecycleConfigurationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketLocationRequest(g *s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketLocationOutput)
}

func (m *MockS3API) GetBucketLocation(g *s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketLocationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketLoggingRequest(g *s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketLoggingOutput)
}

func (m *MockS3API) GetBucketLogging(g *s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketLoggingOutput), args.Error(1)
}

func (m *MockS3API) GetBucketMetricsConfigurationRequest(g *s3.GetBucketMetricsConfigurationInput) (*request.Request, *s3.GetBucketMetricsConfigurationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketMetricsConfigurationOutput)
}

func (m *MockS3API) GetBucketMetricsConfiguration(g *s3.GetBucketMetricsConfigurationInput) (*s3.GetBucketMetricsConfigurationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketMetricsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketNotificationRequest(g *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.NotificationConfigurationDeprecated)
}

func (m *MockS3API) GetBucketNotification(g *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.NotificationConfigurationDeprecated), args.Error(1)
}

func (m *MockS3API) GetBucketNotificationConfigurationRequest(g *s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.NotificationConfiguration)
}

func (m *MockS3API) GetBucketNotificationConfiguration(g *s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.NotificationConfiguration), args.Error(1)
}

func (m *MockS3API) GetBucketPolicyRequest(g *s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketPolicyOutput)
}

func (m *MockS3API) GetBucketPolicy(g *s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketPolicyOutput), args.Error(1)
}

func (m *MockS3API) GetBucketReplicationRequest(g *s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketReplicationOutput)
}

func (m *MockS3API) GetBucketReplication(g *s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketReplicationOutput), args.Error(1)
}

func (m *MockS3API) GetBucketRequestPaymentRequest(g *s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketRequestPaymentOutput)
}

func (m *MockS3API) GetBucketRequestPayment(g *s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketRequestPaymentOutput), args.Error(1)
}

func (m *MockS3API) GetBucketTaggingRequest(*s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput) {
	args := m.Called()
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketTaggingOutput)
}

func (m *MockS3API) GetBucketTagging(g *s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketTaggingOutput), args.Error(1)
}

func (m *MockS3API) GetBucketVersioningRequest(g *s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketVersioningOutput)
}

func (m *MockS3API) GetBucketVersioning(g *s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketVersioningOutput), args.Error(1)
}

func (m *MockS3API) GetBucketWebsiteRequest(g *s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetBucketWebsiteOutput)
}

func (m *MockS3API) GetBucketWebsite(g *s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetBucketWebsiteOutput), args.Error(1)
}

func (m *MockS3API) GetObjectRequest(g *s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetObjectOutput)
}

func (m *MockS3API) GetObject(g *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetObjectOutput), args.Error(1)
}

func (m *MockS3API) GetObjectAclRequest(g *s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetObjectAclOutput)
}

func (m *MockS3API) GetObjectAcl(g *s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetObjectAclOutput), args.Error(1)
}

func (m *MockS3API) GetObjectTaggingRequest(g *s3.GetObjectTaggingInput) (*request.Request, *s3.GetObjectTaggingOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetObjectTaggingOutput)
}

func (m *MockS3API) GetObjectTagging(g *s3.GetObjectTaggingInput) (*s3.GetObjectTaggingOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetObjectTaggingOutput), args.Error(1)
}

func (m *MockS3API) GetObjectTorrentRequest(g *s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.GetObjectTorrentOutput)
}

func (m *MockS3API) GetObjectTorrent(g *s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.GetObjectTorrentOutput), args.Error(1)
}

func (m *MockS3API) HeadBucketRequest(g *s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput) {
	args := m.Called(g)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.HeadBucketOutput)
}

func (m *MockS3API) HeadBucket(g *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	args := m.Called(g)
	return args.Get(0).(*s3.HeadBucketOutput), args.Error(1)
}

func (m *MockS3API) HeadObjectRequest(h *s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput) {
	args := m.Called(h)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.HeadObjectOutput)
}

func (m *MockS3API) HeadObject(h *s3.HeadObjectInput) (*s3.HeadObjectOutput, error) {
	args := m.Called(h)
	return args.Get(0).(*s3.HeadObjectOutput), args.Error(1)
}

func (m *MockS3API) ListBucketAnalyticsConfigurationsRequest(l *s3.ListBucketAnalyticsConfigurationsInput) (*request.Request, *s3.ListBucketAnalyticsConfigurationsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListBucketAnalyticsConfigurationsOutput)
}

func (m *MockS3API) ListBucketAnalyticsConfigurations(l *s3.ListBucketAnalyticsConfigurationsInput) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListBucketAnalyticsConfigurationsOutput), args.Error(1)
}

func (m *MockS3API) ListBucketInventoryConfigurationsRequest(l *s3.ListBucketInventoryConfigurationsInput) (*request.Request, *s3.ListBucketInventoryConfigurationsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListBucketInventoryConfigurationsOutput)
}

func (m *MockS3API) ListBucketInventoryConfigurations(l *s3.ListBucketInventoryConfigurationsInput) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListBucketInventoryConfigurationsOutput), args.Error(1)
}

func (m *MockS3API) ListBucketMetricsConfigurationsRequest(l *s3.ListBucketMetricsConfigurationsInput) (*request.Request, *s3.ListBucketMetricsConfigurationsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListBucketMetricsConfigurationsOutput)
}

func (m *MockS3API) ListBucketMetricsConfigurations(l *s3.ListBucketMetricsConfigurationsInput) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListBucketMetricsConfigurationsOutput), args.Error(1)
}

func (m *MockS3API) ListBucketsRequest(l *s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListBucketsOutput)
}

func (m *MockS3API) ListBuckets(l *s3.ListBucketsInput) (*s3.ListBucketsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListBucketsOutput), args.Error(1)
}

func (m *MockS3API) ListMultipartUploadsRequest(l *s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListMultipartUploadsOutput)
}

func (m *MockS3API) ListMultipartUploads(l *s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListMultipartUploadsOutput), args.Error(1)
}

func (m *MockS3API) ListMultipartUploadsPages(l *s3.ListMultipartUploadsInput, f func(*s3.ListMultipartUploadsOutput, bool) bool) error {
	args := m.Called(l, f)
	return args.Error(0)
}

func (m *MockS3API) ListObjectVersionsRequest(l *s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListObjectVersionsOutput)
}

func (m *MockS3API) ListObjectVersions(l *s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListObjectVersionsOutput), args.Error(1)
}

func (m *MockS3API) ListObjectVersionsPages(l *s3.ListObjectVersionsInput, f func(*s3.ListObjectVersionsOutput, bool) bool) error {
	args := m.Called(l, f)
	return args.Error(0)
}

func (m *MockS3API) ListObjectsRequest(l *s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListObjectsOutput)
}

func (m *MockS3API) ListObjects(l *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListObjectsOutput), args.Error(1)
}

func (m *MockS3API) ListObjectsPages(l *s3.ListObjectsInput, f func(*s3.ListObjectsOutput, bool) bool) error {
	args := m.Called(l, f)
	return args.Error(0)
}

func (m *MockS3API) ListObjectsV2Request(l *s3.ListObjectsV2Input) (*request.Request, *s3.ListObjectsV2Output) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListObjectsV2Output)
}

func (m *MockS3API) ListObjectsV2(l *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListObjectsV2Output), args.Error(1)
}

func (m *MockS3API) ListObjectsV2Pages(l *s3.ListObjectsV2Input, f func(*s3.ListObjectsV2Output, bool) bool) error {
	args := m.Called(l, f)
	return args.Error(0)
}

func (m *MockS3API) ListPartsRequest(l *s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput) {
	args := m.Called(l)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.ListPartsOutput)
}

func (m *MockS3API) ListParts(l *s3.ListPartsInput) (*s3.ListPartsOutput, error) {
	args := m.Called(l)
	return args.Get(0).(*s3.ListPartsOutput), args.Error(1)
}

func (m *MockS3API) ListPartsPages(l *s3.ListPartsInput, f func(*s3.ListPartsOutput, bool) bool) error {
	args := m.Called(l, f)
	return args.Error(0)
}

func (m *MockS3API) PutBucketAccelerateConfigurationRequest(p *s3.PutBucketAccelerateConfigurationInput) (*request.Request, *s3.PutBucketAccelerateConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketAccelerateConfigurationOutput)
}

func (m *MockS3API) PutBucketAccelerateConfiguration(p *s3.PutBucketAccelerateConfigurationInput) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketAccelerateConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketAclRequest(p *s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketAclOutput)
}

func (m *MockS3API) PutBucketAcl(p *s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketAclOutput), args.Error(1)
}

func (m *MockS3API) PutBucketAnalyticsConfigurationRequest(p *s3.PutBucketAnalyticsConfigurationInput) (*request.Request, *s3.PutBucketAnalyticsConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketAnalyticsConfigurationOutput)
}

func (m *MockS3API) PutBucketAnalyticsConfiguration(p *s3.PutBucketAnalyticsConfigurationInput) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketAnalyticsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketCorsRequest(p *s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketCorsOutput)
}

func (m *MockS3API) PutBucketCors(p *s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketCorsOutput), args.Error(1)
}

func (m *MockS3API) PutBucketInventoryConfigurationRequest(p *s3.PutBucketInventoryConfigurationInput) (*request.Request, *s3.PutBucketInventoryConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketInventoryConfigurationOutput)
}

func (m *MockS3API) PutBucketInventoryConfiguration(p *s3.PutBucketInventoryConfigurationInput) (*s3.PutBucketInventoryConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketInventoryConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketLifecycleRequest(p *s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketLifecycleOutput)
}

func (m *MockS3API) PutBucketLifecycle(p *s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketLifecycleOutput), args.Error(1)
}

func (m *MockS3API) PutBucketLifecycleConfigurationRequest(p *s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketLifecycleConfigurationOutput)
}

func (m *MockS3API) PutBucketLifecycleConfiguration(p *s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketLifecycleConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketLoggingRequest(p *s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketLoggingOutput)
}

func (m *MockS3API) PutBucketLogging(p *s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketLoggingOutput), args.Error(1)
}

func (m *MockS3API) PutBucketMetricsConfigurationRequest(p *s3.PutBucketMetricsConfigurationInput) (*request.Request, *s3.PutBucketMetricsConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketMetricsConfigurationOutput)
}

func (m *MockS3API) PutBucketMetricsConfiguration(p *s3.PutBucketMetricsConfigurationInput) (*s3.PutBucketMetricsConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketMetricsConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketNotificationRequest(p *s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketNotificationOutput)
}

func (m *MockS3API) PutBucketNotification(p *s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketNotificationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketNotificationConfigurationRequest(p *s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketNotificationConfigurationOutput)
}

func (m *MockS3API) PutBucketNotificationConfiguration(p *s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketNotificationConfigurationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketPolicyRequest(p *s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketPolicyOutput)
}

func (m *MockS3API) PutBucketPolicy(p *s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketPolicyOutput), args.Error(1)
}

func (m *MockS3API) PutBucketReplicationRequest(p *s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketReplicationOutput)
}

func (m *MockS3API) PutBucketReplication(p *s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketReplicationOutput), args.Error(1)
}

func (m *MockS3API) PutBucketRequestPaymentRequest(p *s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketRequestPaymentOutput)
}

func (m *MockS3API) PutBucketRequestPayment(p *s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketRequestPaymentOutput), args.Error(1)
}

func (m *MockS3API) PutBucketTaggingRequest(p *s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketTaggingOutput)
}

func (m *MockS3API) PutBucketTagging(p *s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketTaggingOutput), args.Error(1)
}

func (m *MockS3API) PutBucketVersioningRequest(p *s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketVersioningOutput)
}

func (m *MockS3API) PutBucketVersioning(p *s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketVersioningOutput), args.Error(1)
}

func (m *MockS3API) PutBucketWebsiteRequest(p *s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutBucketWebsiteOutput)
}

func (m *MockS3API) PutBucketWebsite(p *s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutBucketWebsiteOutput), args.Error(1)
}

func (m *MockS3API) PutObjectRequest(p *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutObjectOutput)
}

func (m *MockS3API) PutObject(p *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func (m *MockS3API) PutObjectAclRequest(p *s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutObjectAclOutput)
}

func (m *MockS3API) PutObjectAcl(p *s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutObjectAclOutput), args.Error(1)
}

func (m *MockS3API) PutObjectTaggingRequest(p *s3.PutObjectTaggingInput) (*request.Request, *s3.PutObjectTaggingOutput) {
	args := m.Called(p)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.PutObjectTaggingOutput)
}

func (m *MockS3API) PutObjectTagging(p *s3.PutObjectTaggingInput) (*s3.PutObjectTaggingOutput, error) {
	args := m.Called(p)
	return args.Get(0).(*s3.PutObjectTaggingOutput), args.Error(1)
}

func (m *MockS3API) RestoreObjectRequest(r *s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput) {
	args := m.Called(r)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.RestoreObjectOutput)
}

func (m *MockS3API) RestoreObject(r *s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error) {
	args := m.Called(r)
	return args.Get(0).(*s3.RestoreObjectOutput), args.Error(1)
}

func (m *MockS3API) UploadPartRequest(u *s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput) {
	args := m.Called(u)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.UploadPartOutput)
}

func (m *MockS3API) UploadPart(u *s3.UploadPartInput) (*s3.UploadPartOutput, error) {
	args := m.Called(u)
	return args.Get(0).(*s3.UploadPartOutput), args.Error(1)
}

func (m *MockS3API) UploadPartCopyRequest(u *s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput) {
	args := m.Called(u)
	return args.Get(0).(*request.Request), args.Get(1).(*s3.UploadPartCopyOutput)
}

func (m *MockS3API) UploadPartCopy(input *s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.UploadPartCopyOutput), args.Error(1)
}

func (m *MockS3API) WaitUntilBucketExists(head *s3.HeadBucketInput) error {
	args := m.Called(head)
	return args.Error(0)
}

func (m *MockS3API) WaitUntilBucketNotExists(head *s3.HeadBucketInput) error {
	args := m.Called(head)
	return args.Error(0)
}

func (m *MockS3API) WaitUntilObjectExists(head *s3.HeadObjectInput) error {
	args := m.Called(head)
	return args.Error(0)
}

func (m *MockS3API) WaitUntilObjectNotExists(head *s3.HeadObjectInput) error {
	args := m.Called(head)
	return args.Error(0)
}
