package host

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	eventgen "github.com/openshift/assisted-service/internal/common/events"
	"github.com/openshift/assisted-service/internal/events"
	"github.com/openshift/assisted-service/internal/events/eventstest"
	"github.com/openshift/assisted-service/internal/hardware"
	"github.com/openshift/assisted-service/internal/host/hostutil"
	"github.com/openshift/assisted-service/internal/metrics"
	"github.com/openshift/assisted-service/internal/operators"
	"github.com/openshift/assisted-service/internal/operators/api"
	"github.com/openshift/assisted-service/internal/provider/registry"
	"github.com/openshift/assisted-service/models"
)

var _ = Describe("Validations test", func() {

	var (
		ctrl            *gomock.Controller
		ctx             = context.Background()
		db              *gorm.DB
		dbName          string
		mockEvents      *events.MockHandler
		mockHwValidator *hardware.MockValidator
		mockMetric      *metrics.MockAPI
		mockOperators   *operators.MockAPI
		m               *Manager

		clusterID, hostID, infraEnvID strfmt.UUID
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		db, dbName = common.PrepareTestDB()
		mockEvents = events.NewMockHandler(ctrl)
		mockHwValidator = hardware.NewMockValidator(ctrl)
		mockMetric = metrics.NewMockAPI(ctrl)
		mockOperators = operators.NewMockAPI(ctrl)
		pr := registry.NewMockProviderRegistry(ctrl)
		pr.EXPECT().IsHostSupported(gomock.Any(), gomock.Any()).Return(true, nil).AnyTimes()
		m = NewManager(common.GetTestLog(), db, mockEvents, mockHwValidator, nil, createValidatorCfg(), mockMetric, defaultConfig, nil, mockOperators, pr)

		clusterID = strfmt.UUID(uuid.New().String())
		hostID = strfmt.UUID(uuid.New().String())
		infraEnvID = strfmt.UUID(uuid.New().String())
	})

	AfterEach(func() {
		ctrl.Finish()
		common.DeleteTestDB(db, dbName)
	})

	mockAndRefreshStatusWithoutEvents := func(h *models.Host) {
		mockDefaultClusterHostRequirements(mockHwValidator)
		mockHwValidator.EXPECT().ListEligibleDisks(gomock.Any()).Return([]*models.Disk{}).AnyTimes()
		mockHwValidator.EXPECT().GetHostInstallationPath(gomock.Any()).Return("/dev/sda").AnyTimes()
		mockOperators.EXPECT().ValidateHost(gomock.Any(), gomock.Any(), gomock.Any()).Return([]api.ValidationResult{
			{Status: api.Success, ValidationId: string(models.HostValidationIDOcsRequirementsSatisfied)},
			{Status: api.Success, ValidationId: string(models.HostValidationIDLsoRequirementsSatisfied)},
			{Status: api.Success, ValidationId: string(models.HostValidationIDCnvRequirementsSatisfied)},
		}, nil).AnyTimes()

		err := m.RefreshStatus(ctx, h, db)
		Expect(err).ToNot(HaveOccurred())
	}

	mockAndRefreshStatus := func(h *models.Host) {
		mockEvents.EXPECT().SendHostEvent(gomock.Any(), eventstest.NewEventMatcher(
			eventstest.WithNameMatcher(eventgen.HostStatusUpdatedEventName),
			eventstest.WithHostIdMatcher(h.ID.String()),
			eventstest.WithInfraEnvIdMatcher(h.InfraEnvID.String()),
		))
		mockAndRefreshStatusWithoutEvents(h)
	}

	Context("Disk encryption validation", func() {

		getDiskEncryptionValidationResult := func(validationsInfo string) (ValidationStatus, string, bool) {

			var validationsRes ValidationsStatus
			err := json.Unmarshal([]byte(validationsInfo), &validationsRes)
			Expect(err).ToNot(HaveOccurred())

			for _, vl := range validationsRes {
				for _, v := range vl {
					if v.ID == validationID(models.HostValidationIDDiskEncryptionRequirementsSatisfied) {
						return v.Status, v.Message, true
					}
				}
			}
			return ValidationStatus(""), "", false
		}

		It("disk-encryption not set", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleMaster)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNr20)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			_, _, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeFalse())
		})

		It("un-affected roles", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnMasters),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleWorker)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNr20)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			_, _, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeFalse())
		})

		It("auto-assigned role", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnMasters),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleAutoAssign)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNr20)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			checkValidation := func(expectedStatus ValidationStatus, expectedMsg string) {
				h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
				validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
				ExpectWithOffset(1, found).To(BeTrue())
				ExpectWithOffset(1, validationStatus).To(Equal(expectedStatus))
				ExpectWithOffset(1, validationMessage).To(Equal(expectedMsg))
			}

			By("effective role is auto-assign (no suggestion yet)")
			mockAndRefreshStatus(&h)
			checkValidation(ValidationPending, "Missing role assignment")

			By("auto-assign node is effectively assigned to master")
			mockEvents.EXPECT().SendHostEvent(gomock.Any(), eventstest.NewEventMatcher(
				eventstest.WithNameMatcher(eventgen.UpdateHostRoleEventName),
				eventstest.WithHostIdMatcher(h.ID.String()),
				eventstest.WithInfraEnvIdMatcher(h.InfraEnvID.String()),
			))
			err := m.RefreshRole(ctx, &h, db)
			Expect(err).ToNot(HaveOccurred())

			mockAndRefreshStatusWithoutEvents(&h)
			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			checkValidation(ValidationSuccess, "Installation disk can be encrypted using tpmv2")
		})

		It("missing inventory", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnMasters),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleMaster)
			h.Inventory = ""
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatusWithoutEvents(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeTrue())
			Expect(validationStatus).To(Equal(ValidationPending))
			Expect(validationMessage).To(Equal("Missing host inventory"))
		})

		It("non TPM mode", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnMasters),
				Mode:     swag.String(models.DiskEncryptionModeTang),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleMaster)
			h.Inventory = common.GenerateTestInventoryWithSetNetwork()
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeTrue())
			Expect(validationStatus).To(Equal(ValidationSuccess))
			Expect(validationMessage).To(Equal(fmt.Sprintf("Installation disk can be encrypted using %s", models.DiskEncryptionModeTang)))
		})

		It("TPM is disabled in host's BIOS", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnWorkers),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleWorker)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNone)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeTrue())
			Expect(validationStatus).To(Equal(ValidationFailure))
			Expect(validationMessage).To(Equal("TPM version could not be found, make sure TPM is enalbed in host's BIOS"))
		})

		It("host's TPM version is unsupported", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnAll),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleMaster)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNr12)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeTrue())
			Expect(validationStatus).To(Equal(ValidationFailure))
			Expect(validationMessage).To(Equal(fmt.Sprintf("The host's TPM version is not supported, expected-version: %s, actual-version: %s",
				models.InventoryTpmVersionNr20, models.InventoryTpmVersionNr12)))
		})

		It("happy flow - explicit role", func() {

			c := hostutil.GenerateTestCluster(clusterID, common.TestIPv4Networking.MachineNetworks)
			c.DiskEncryption = &models.DiskEncryption{
				EnableOn: swag.String(models.DiskEncryptionEnableOnAll),
				Mode:     swag.String(models.DiskEncryptionModeTpmv2),
			}
			Expect(db.Create(&c).Error).ToNot(HaveOccurred())

			h := hostutil.GenerateTestHostByKind(hostID, infraEnvID, &clusterID, models.HostStatusDiscovering, models.HostKindHost, models.HostRoleWorker)
			h.Inventory = common.GenerateTestInventoryWithTpmVersion(models.InventoryTpmVersionNr20)
			Expect(db.Create(&h).Error).ShouldNot(HaveOccurred())

			mockAndRefreshStatus(&h)

			h = hostutil.GetHostFromDB(*h.ID, h.InfraEnvID, db).Host
			validationStatus, validationMessage, found := getDiskEncryptionValidationResult(h.ValidationsInfo)
			Expect(found).To(BeTrue())
			Expect(validationStatus).To(Equal(ValidationSuccess))
			Expect(validationMessage).To(Equal(fmt.Sprintf("Installation disk can be encrypted using %s", models.DiskEncryptionModeTpmv2)))
		})
	})
})
