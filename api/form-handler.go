package apis

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/marcboeker/go-duckdb"
)

type SignUpForm struct {
	Name                                    string `form:"name"`
	Email                                   string `form:"email"`
	Phone                                   string `form:"phone"`
	Age                                     string `form:"age"`
	Religion                                string `form:"religion"`
	Gender                                  string `form:"gender"`
	Degree                                  string `form:"degree"`
	CampusName                              string `form:"campus_name"`
	CityBirthplace                          string `form:"city_birthplace"`
	CityResidency                           string `form:"city_residency"`
	Test                                    string `form:"test"`
	InternalAuditSeniorManager              bool   `form:"internalauditseniormanager"`
	CultureOrganizationalDevelopmentManager bool   `form:"cultureorganizationaldevelopmentmanager"`
	RecruitmentAndSelectionStaff            bool   `form:"recruitmentandselectionstaff"`
	AssessmentStaff                         bool   `form:"assessmentstaff"`
	SSBCorpUniversityStaff                  bool   `form:"ssbcorpuniversitystaff"`
	FinanceSupervisor                       bool   `form:"financesupervisor"`
	SoftwareEngineerStaff                   bool   `form:"softwareengineerstaff"`
	P2OmanagerHeadOffice                    bool   `form:"p2omanagerheadoffice"`
	P2OSupervisorHeadOffice                 bool   `form:"p2osupervisorheadoffice"`
	CateringManagerHeadOffice               bool   `form:"cateringmanagerheadoffice"`
	CateringSupervisorHeadOffice            bool   `form:"cateringsupervisorheadoffice"`
	ProduksiManagerHeadoffice               bool   `form:"produksimanagerheadoffice"`
	ManagerOutletSupervisor                 bool   `form:"manageroutletsupervisor"`
	GeneralAffairAreaStaff                  bool   `form:"generalaffairareastaff"`
	FinanceAreaStaff                        bool   `form:"financeareastaff"`
	AccountingSupervisor                    bool   `form:"accountingsupervisor"`
	AccountPayableCollectorStaff            bool   `form:"accountpayablecollectorstaff"`
	BookKeeperStaff                         bool   `form:"bookkeeperstaff"`
	GeneralAccountingBankStaff              bool   `form:"generalaccountingbankstaff"`
	StaffArsipStaff                         bool   `form:"staffarsipstaff"`
	StaffCostAccountingStaff                bool   `form:"staffcostaccountingstaff"`
	RecruitmentAndAssessmentStaff           bool   `form:"recruitmentandassessmentstaff"`
	LearningAndDevelopmentSupervisor        bool   `form:"learninganddevelopmentsupervisor"`
	LearningAndDevelopmentStaff             bool   `form:"learninganddevelopmentstaff"`
	HCAreaSupervisor                        bool   `form:"hcareasupervisor"`
	HCAreaStaff                             bool   `form:"hcareastaff"`
	MarketingSeniorManager                  bool   `form:"marketingseniormanager"`
	KontenKreatorSupervisor                 bool   `form:"kontenkreatorsupervisor"`
	CustomerRelationMarketingSupervisor     bool   `form:"customerrelationmarketingsupervisor"`
	MarketingBrandingStaff                  bool   `form:"marketingbrandingstaff"`
	MarketingSalesSupervisor                bool   `form:"marketingsalessupervisor"`
	SurveyorStaff                           bool   `form:"surveyorstaff"`
	KemitraanStaff                          bool   `form:"kemitraanstaff"`
	SupplyChainManagementSeniorManager      bool   `form:"supplychainmanagementseniormanager"`
	LogisticAreaSupervisor                  bool   `form:"logisticareasupervisor"`
	PurchasingStaff                         bool   `form:"purchasingstaff"`
	SupplyPlanStaff                         bool   `form:"supplyplanstaff"`
	WarehouseSupervisor                     bool   `form:"warehousesupervisor"`
	MaintenanceAtauMekanikSupervisor        bool   `form:"maintenanceataumekaniksupervisor"`
}

func FormHandler(c *gin.Context) {
	log.Print("INFO: Parsing Payload")
	var req SignUpForm

	if err := c.Bind(&req); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"msg":    "The request could not be understood by the server due to incorrect syntax."})
		return
	}

	db, err := sql.Open("duckdb", "")
	if err != nil {
		fmt.Printf("db unable to connect: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}
	defer db.Close()

	_, err = db.Exec(`
		SET autoinstall_known_extensions=1;
		SET autoload_known_extensions=1;
		SET s3_endpoint = 'storage.googleapis.com';
		SET s3_access_key_id = '` + os.Getenv("KEYID") + `';
		SET s3_secret_access_key = '` + os.Getenv("KEYSECRET") + `';`)
	if err != nil {
		fmt.Printf("failed setting credentials: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	_, err = db.Exec(`
		CREATE TABLE datum(
			name varchar(100),
			email varchar(100),
			phone varchar(100),
			age varchar(100),
			religion varchar(100),
			gender varchar(100),
			degree varchar(100),
			campus_name varchar(100),
			city_birthplace varchar(100),
			city_residency varchar(100)
			internalauditseniormanager boolean,
			cultureorganizationaldevelopmentmanager boolean,
			recruitmentandselectionstaff boolean,
			assessmentstaff boolean,
			ssbcorpuniversitystaff boolean,
			financesupervisor boolean,
			softwareengineerstaff boolean,
			p2omanagerheadoffice boolean,
			p2osupervisorheadoffice boolean,
			cateringmanagerheadoffice boolean,
			cateringsupervisorheadoffice boolean,
			produksimanagerheadoffice boolean,
			manageroutletsupervisor boolean,
			generalaffairareastaff boolean,
			financeareastaff boolean,
			accountingsupervisor boolean,
			accountpayablecollectorstaff boolean,
			bookkeeperstaff boolean,
			generalaccountingbankstaff boolean,
			staffarsipstaff boolean,
			staffcostaccountingstaff boolean,
			recruitmentandassessmentstaff boolean,
			learninganddevelopmentsupervisor boolean,
			learninganddevelopmentstaff boolean,
			hcareasupervisor boolean,
			hcareastaff boolean,
			marketingseniormanager boolean,
			kontenkreatorsupervisor boolean,
			customerrelationmarketingsupervisor boolean,
			marketingbrandingstaff boolean,
			marketingsalessupervisor boolean,
			surveyorstaff boolean,
			kemitraanstaff boolean,
			supplychainmanagementseniormanager boolean,
			logisticareasupervisor boolean,
			purchasingstaff boolean,
			supplyplanstaff boolean,
			warehousesupervisor boolean,
			maintenanceataumekaniksupervisor boolean
		)`)
	if err != nil {
		fmt.Printf("failed creating table: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	_, err = db.Exec(`
		INSERT INTO datum(
			name,
			email,
			phone,
			age,
			religion,
			gender,
			degree,
			campus_name,
			city_birthplace,
			city_residency,
			internalauditseniormanager, cultureorganizationaldevelopmentmanager, recruitmentandselectionstaff, assessmentstaff, ssbcorpuniversitystaff, financesupervisor, softwareengineerstaff, p2omanagerheadoffice, p2osupervisorheadoffice, cateringmanagerheadoffice, cateringsupervisorheadoffice, produksimanagerheadoffice, manageroutletsupervisor, generalaffairareastaff, financeareastaff, accountingsupervisor, accountpayablecollectorstaff, bookkeeperstaff, generalaccountingbankstaff, staffarsipstaff, staffcostaccountingstaff, recruitmentandassessmentstaff, learninganddevelopmentsupervisor, learninganddevelopmentstaff, hcareasupervisor, hcareastaff, marketingseniormanager, kontenkreatorsupervisor, customerrelationmarketingsupervisor, marketingbrandingstaff, marketingsalessupervisor, surveyorstaff, kemitraanstaff, supplychainmanagementseniormanager, logisticareasupervisor, purchasingstaff, supplyplanstaff, warehousesupervisor, maintenanceataumekaniksupervisor
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49)`, req.Name, req.Email, req.Phone, req.Age, req.Religion, req.Gender, req.Degree, req.CampusName, req.CityBirthplace, req.CityResidency, req.InternalAuditSeniorManager, req.CultureOrganizationalDevelopmentManager, req.RecruitmentAndSelectionStaff, req.AssessmentStaff, req.SSBCorpUniversityStaff, req.FinanceSupervisor, req.SoftwareEngineerStaff, req.P2OmanagerHeadOffice, req.P2OSupervisorHeadOffice, req.CateringManagerHeadOffice, req.CateringSupervisorHeadOffice, req.ProduksiManagerHeadoffice, req.ManagerOutletSupervisor, req.GeneralAffairAreaStaff, req.FinanceAreaStaff, req.AccountingSupervisor, req.AccountPayableCollectorStaff, req.BookKeeperStaff, req.GeneralAccountingBankStaff, req.StaffArsipStaff, req.StaffCostAccountingStaff, req.RecruitmentAndAssessmentStaff, req.LearningAndDevelopmentSupervisor, req.LearningAndDevelopmentStaff, req.HCAreaSupervisor, req.HCAreaStaff, req.MarketingSeniorManager, req.KontenKreatorSupervisor, req.CustomerRelationMarketingSupervisor, req.MarketingBrandingStaff, req.MarketingSalesSupervisor, req.SurveyorStaff, req.KemitraanStaff, req.SupplyChainManagementSeniorManager, req.LogisticAreaSupervisor, req.PurchasingStaff, req.SupplyPlanStaff, req.WarehouseSupervisor, req.MaintenanceAtauMekanikSupervisor)
	if err != nil {
		fmt.Printf("failed inserting data to table: %+v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	id := uuid.New()
	_, err = db.Exec(`
		COPY datum TO 's3://` + os.Getenv("BUCKET_NAME") + `/forms/` + id.String() + `.parquet'`)
	if err != nil {
		fmt.Printf("failed uploading to S3: %+v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "failure",
		})
		return
	}

	f, err := c.FormFile("cv")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	blobFile, err := f.Open()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "Error Opening File",
		})
		return
	}

	obj := fmt.Sprintf("cv-%s-%s", id.String(), time.Now().Format("20060102-150405"))

	err = fileUploader("cvs/"+obj, &blobFile)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"msg":    "Failed Uploading File",
		})
		return
	}
}

func fileUploader(objName string, file *multipart.File) error {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return err
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := client.Bucket(os.Getenv("BUCKET_NAME")).Object(objName).NewWriter(ctx)
	if _, err := io.Copy(wc, *file); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
