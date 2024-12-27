package imageScanning

import (
	"github.com/devtron-labs/devtron/pkg/policyGovernance/security/imageScanning/read"
	"github.com/devtron-labs/devtron/pkg/policyGovernance/security/imageScanning/repository"
	"github.com/google/wire"
)

var ImageScanningWireSet = wire.NewSet(
	NewPolicyServiceImpl,
	wire.Bind(new(PolicyService), new(*PolicyServiceImpl)),

	NewImageScanServiceImpl,
	wire.Bind(new(ImageScanService), new(*ImageScanServiceImpl)),

	read.NewImageScanHistoryReadService,
	wire.Bind(new(read.ImageScanHistoryReadService), new(*read.ImageScanHistoryReadServiceImpl)),

	read.NewImageScanDeployInfoReadService,
	wire.Bind(new(read.ImageScanDeployInfoReadService), new(*read.ImageScanDeployInfoReadServiceImpl)),

	NewImageScanDeployInfoService,
	wire.Bind(new(ImageScanDeployInfoService), new(*ImageScanDeployInfoServiceImpl)),

	read.NewImageScanResultReadServiceImpl,
	wire.Bind(new(read.ImageScanResultReadService), new(*read.ImageScanResultReadServiceImpl)),

	NewScanToolMetadataServiceImpl,
	wire.Bind(new(ScanToolMetadataService), new(*ScanToolMetadataServiceImpl)),

	repository.NewImageScanHistoryRepositoryImpl,
	wire.Bind(new(repository.ImageScanHistoryRepository), new(*repository.ImageScanHistoryRepositoryImpl)),
	repository.NewImageScanResultRepositoryImpl,
	wire.Bind(new(repository.ImageScanResultRepository), new(*repository.ImageScanResultRepositoryImpl)),
	repository.NewImageScanObjectMetaRepositoryImpl,
	wire.Bind(new(repository.ImageScanObjectMetaRepository), new(*repository.ImageScanObjectMetaRepositoryImpl)),
	repository.NewCveStoreRepositoryImpl,
	wire.Bind(new(repository.CveStoreRepository), new(*repository.CveStoreRepositoryImpl)),
	repository.NewImageScanDeployInfoRepositoryImpl,
	wire.Bind(new(repository.ImageScanDeployInfoRepository), new(*repository.ImageScanDeployInfoRepositoryImpl)),
	repository.NewScanToolMetadataRepositoryImpl,
	wire.Bind(new(repository.ScanToolMetadataRepository), new(*repository.ScanToolMetadataRepositoryImpl)),

	repository.NewPolicyRepositoryImpl,
	wire.Bind(new(repository.CvePolicyRepository), new(*repository.CvePolicyRepositoryImpl)),
	repository.NewScanToolExecutionHistoryMappingRepositoryImpl,
	wire.Bind(new(repository.ScanToolExecutionHistoryMappingRepository), new(*repository.ScanToolExecutionHistoryMappingRepositoryImpl)),
)