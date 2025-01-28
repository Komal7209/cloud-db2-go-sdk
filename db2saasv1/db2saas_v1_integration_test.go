//go:build integration

/**
 * (C) Copyright IBM Corp. 2025.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

 package db2saasv1_test

 import (
	 "fmt"
	 "log"
	 "os"
	 "time"
 
	 "github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	 "github.com/IBM/go-sdk-core/v5/core"
	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
 )
 
 /**
  * This file contains an integration test for the db2saasv1 package.
  *
  * Notes:
  *
  * The integration test will automatically skip tests if the required config file is not available.
  */
 
 var _ = Describe(`Db2saasV1 Integration Tests`, func() {
	 const externalConfigFile = "../db2saas_v1.env"
 
	 var (
		 err          error
		 db2saasService *db2saasv1.Db2saasV1
		 serviceURL   string
		 config       map[string]string
	 )
 
	 var shouldSkipTest = func() {
		 Skip("External configuration is not available, skipping tests...")
	 }
 
	 Describe(`External configuration`, func() {
		 It("Successfully load the configuration", func() {
			 _, err = os.Stat(externalConfigFile)
			 if err != nil {
				 Skip("External configuration file not found, skipping tests: " + err.Error())
			 }
 
			 os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			 config, err = core.GetServiceProperties(db2saasv1.DefaultServiceName)
			 if err != nil {
				 Skip("Error loading service properties, skipping tests: " + err.Error())
			 }
			 serviceURL = config["URL"]
			 if serviceURL == "" {
				 Skip("Unable to load service URL configuration property, skipping tests")
			 }
 
			 fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			 shouldSkipTest = func() {}
		 })
	 })
 
	 Describe(`Client initialization`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It("Successfully construct the service client instance", func() {
			 db2saasServiceOptions := &db2saasv1.Db2saasV1Options{}
 
			 db2saasService, err = db2saasv1.NewDb2saasV1UsingExternalConfig(db2saasServiceOptions)
			 Expect(err).To(BeNil())
			 Expect(db2saasService).ToNot(BeNil())
			 Expect(db2saasService.Service.Options.URL).To(Equal(serviceURL))
 
			 core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			 db2saasService.EnableRetries(4, 30*time.Second)
		 })
	 })
 
	 Describe(`GetDb2SaasConnectionInfo - Get Db2 connection information`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions *GetDb2SaasConnectionInfoOptions)`, func() {
			 getDb2SaasConnectionInfoOptions := &db2saasv1.GetDb2SaasConnectionInfoOptions{
				 DeploymentID: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A"),
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successConnectionInfo, response, err := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successConnectionInfo).ToNot(BeNil())
		 })
	 })
 
	 Describe(`PostDb2SaasAllowlist - Allow listing of new IPs`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`PostDb2SaasAllowlist(postDb2SaasAllowlistOptions *PostDb2SaasAllowlistOptions)`, func() {
			 ipAddressModel := &db2saasv1.IpAddress{
				 Address: core.StringPtr("127.0.0.1"),
				 Description: core.StringPtr("A sample IP address"),
			 }
 
			 postDb2SaasAllowlistOptions := &db2saasv1.PostDb2SaasAllowlistOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
				 IpAddresses: []db2saasv1.IpAddress{*ipAddressModel},
			 }
 
			 successPostAllowedlistIPs, response, err := db2saasService.PostDb2SaasAllowlist(postDb2SaasAllowlistOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successPostAllowedlistIPs).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetDb2SaasAllowlist - Get allowed list of IPs`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasAllowlist(getDb2SaasAllowlistOptions *GetDb2SaasAllowlistOptions)`, func() {
			 getDb2SaasAllowlistOptions := &db2saasv1.GetDb2SaasAllowlistOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successGetAllowlistIPs, response, err := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetAllowlistIPs).ToNot(BeNil())
		 })
	 })
 
	 Describe(`PostDb2SaasUser - Create new user ( available only for platform users)`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`PostDb2SaasUser(postDb2SaasUserOptions *PostDb2SaasUserOptions)`, func() {
			 createUserAuthenticationModel := &db2saasv1.CreateUserAuthentication{
				 Method: core.StringPtr("internal"),
				 PolicyID: core.StringPtr("Default"),
			 }
 
			 postDb2SaasUserOptions := &db2saasv1.PostDb2SaasUserOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
				 ID: core.StringPtr("test-user"),
				 Iam: core.BoolPtr(false),
				 Ibmid: core.StringPtr("test-ibm-id"),
				 Name: core.StringPtr("test_user"),
				 Password: core.StringPtr("dEkMc43@gfAPl!867^dSbu"),
				 Role: core.StringPtr("bluuser"),
				 Email: core.StringPtr("test_user@mycompany.com"),
				 Locked: core.StringPtr("no"),
				 Authentication: createUserAuthenticationModel,
			 }
 
			 successUserResponse, response, err := db2saasService.PostDb2SaasUser(postDb2SaasUserOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successUserResponse).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetDb2SaasUser - Get the list of Users`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasUser(getDb2SaasUserOptions *GetDb2SaasUserOptions)`, func() {
			 getDb2SaasUserOptions := &db2saasv1.GetDb2SaasUserOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successGetUserInfo, response, err := db2saasService.GetDb2SaasUser(getDb2SaasUserOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetUserInfo).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetbyidDb2SaasUser - Get specific user by Id`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetbyidDb2SaasUser(getbyidDb2SaasUserOptions *GetbyidDb2SaasUserOptions)`, func() {
			 getbyidDb2SaasUserOptions := &db2saasv1.GetbyidDb2SaasUserOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successGetUserByID, response, err := db2saasService.GetbyidDb2SaasUser(getbyidDb2SaasUserOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetUserByID).ToNot(BeNil())
		 })
	 })
 
	 Describe(`PutDb2SaasAutoscale - Update auto scaling configuration`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions *PutDb2SaasAutoscaleOptions)`, func() {
			 putDb2SaasAutoscaleOptions := &db2saasv1.PutDb2SaasAutoscaleOptions{
				 XDbProfile: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"),
				 AutoScalingThreshold: core.Int64Ptr(int64(90)),
				 AutoScalingPauseLimit: core.Int64Ptr(int64(70)),
			 }
 
			 successUpdateAutoScale, response, err := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successUpdateAutoScale).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetDb2SaasAutoscale - Get auto scaling info`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions *GetDb2SaasAutoscaleOptions)`, func() {
			 getDb2SaasAutoscaleOptions := &db2saasv1.GetDb2SaasAutoscaleOptions{
				 XDbProfile: core.StringPtr("crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"),
			 }
 
			 successAutoScaling, response, err := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successAutoScaling).ToNot(BeNil())
		 })
	 })
 
	 Describe(`PostDb2SaasDbConfiguration - Set database and database manager configuration`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions *PostDb2SaasDbConfigurationOptions)`, func() {
			 createCustomSettingsRegistryModel := &db2saasv1.CreateCustomSettingsRegistry{
				 DB2BIDI: core.StringPtr("YES"),
				 DB2COMPOPT: core.StringPtr("-"),
				 DB2LOCKTORB: core.StringPtr("STATEMENT"),
				 DB2STMM: core.StringPtr("YES"),
				 DB2ALTERNATEAUTHZBEHAVIOUR: core.StringPtr("EXTERNAL_ROUTINE_DBADM"),
				 DB2ANTIJOIN: core.StringPtr("EXTEND"),
				 DB2ATSENABLE: core.StringPtr("YES"),
				 DB2DEFERREDPREPARESEMANTICS: core.StringPtr("YES"),
				 DB2EVALUNCOMMITTED: core.StringPtr("NO"),
				 DB2EXTENDEDOPTIMIZATION: core.StringPtr("-"),
				 DB2INDEXPCTFREEDEFAULT: core.StringPtr("10"),
				 DB2INLISTTONLJN: core.StringPtr("YES"),
				 DB2MINIMIZELISTPREFETCH: core.StringPtr("NO"),
				 DB2OBJECTTABLEENTRIES: core.StringPtr("5000"),
				 DB2OPTPROFILE: core.StringPtr("NO"),
				 DB2OPTSTATSLOG: core.StringPtr("-"),
				 DB2OPTMAXTEMPSIZE: core.StringPtr("-"),
				 DB2PARALLELIO: core.StringPtr("-"),
				 DB2REDUCEDOPTIMIZATION: core.StringPtr("-"),
				 DB2SELECTIVITY: core.StringPtr("YES"),
				 DB2SKIPDELETED: core.StringPtr("NO"),
				 DB2SKIPINSERTED: core.StringPtr("YES"),
				 DB2SYNCRELEASELOCKATTRIBUTES: core.StringPtr("YES"),
				 DB2TRUNCATEREUSESTORAGE: core.StringPtr("IMPORT"),
				 DB2USEALTERNATEPAGECLEANING: core.StringPtr("ON"),
				 DB2VIEWREOPTVALUES: core.StringPtr("NO"),
				 DB2WLMSETTINGS: core.StringPtr("-"),
				 DB2WORKLOAD: core.StringPtr("SAP"),
			 }
 
			 createCustomSettingsDbModel := &db2saasv1.CreateCustomSettingsDb{
				 ACTSORTMEMLIMIT: core.StringPtr("NONE"),
				 ALTCOLLATE: core.StringPtr("NULL"),
				 APPGROUPMEMSZ: core.StringPtr("10"),
				 APPLHEAPSZ: core.StringPtr("AUTOMATIC"),
				 APPLMEMORY: core.StringPtr("AUTOMATIC"),
				 APPCTLHEAPSZ: core.StringPtr("64000"),
				 ARCHRETRYDELAY: core.StringPtr("65535"),
				 AUTHNCACHEDURATION: core.StringPtr("10000"),
				 AUTORESTART: core.StringPtr("ON"),
				 AUTOCGSTATS: core.StringPtr("ON"),
				 AUTOMAINT: core.StringPtr("OFF"),
				 AUTOREORG: core.StringPtr("ON"),
				 AUTOREVAL: core.StringPtr("IMMEDIATE"),
				 AUTORUNSTATS: core.StringPtr("ON"),
				 AUTOSAMPLING: core.StringPtr("OFF"),
				 AUTOSTATSVIEWS: core.StringPtr("ON"),
				 AUTOSTMTSTATS: core.StringPtr("OFF"),
				 AUTOTBLMAINT: core.StringPtr("ON"),
				 AVGAPPLS: core.StringPtr("-"),
				 CATALOGCACHESZ: core.StringPtr("-"),
				 CHNGPGSTHRESH: core.StringPtr("50"),
				 CURCOMMIT: core.StringPtr("AVAILABLE"),
				 DATABASEMEMORY: core.StringPtr("AUTOMATIC"),
				 DBHEAP: core.StringPtr("AUTOMATIC"),
				 DBCOLLNAME: core.StringPtr("-"),
				 DBMEMTHRESH: core.StringPtr("75"),
				 DDLCOMPRESSIONDEF: core.StringPtr("YES"),
				 DDLCONSTRAINTDEF: core.StringPtr("NO"),
				 DECFLTROUNDING: core.StringPtr("ROUND_HALF_UP"),
				 DECARITHMETIC: core.StringPtr("-"),
				 DECTOCHARFMT: core.StringPtr("NEW"),
				 DFTDEGREE: core.StringPtr("-1"),
				 DFTEXTENTSZ: core.StringPtr("32"),
				 DFTLOADRECSES: core.StringPtr("1000"),
				 DFTMTTBTYPES: core.StringPtr("-"),
				 DFTPREFETCHSZ: core.StringPtr("AUTOMATIC"),
				 DFTQUERYOPT: core.StringPtr("3"),
				 DFTREFRESHAGE: core.StringPtr("-"),
				 DFTSCHEMASDCC: core.StringPtr("YES"),
				 DFTSQLMATHWARN: core.StringPtr("YES"),
				 DFTTABLEORG: core.StringPtr("COLUMN"),
				 DLCHKTIME: core.StringPtr("10000"),
				 ENABLEXMLCHAR: core.StringPtr("YES"),
				 EXTENDEDROWSZ: core.StringPtr("ENABLE"),
				 GROUPHEAPRATIO: core.StringPtr("50"),
				 INDEXREC: core.StringPtr("SYSTEM"),
				 LARGEAGGREGATION: core.StringPtr("YES"),
				 LOCKLIST: core.StringPtr("AUTOMATIC"),
				 LOCKTIMEOUT: core.StringPtr("-1"),
				 LOGINDEXBUILD: core.StringPtr("ON"),
				 LOGAPPLINFO: core.StringPtr("YES"),
				 LOGDDLSTMTS: core.StringPtr("NO"),
				 LOGDISKCAP: core.StringPtr("0"),
				 MAXAPPLS: core.StringPtr("5000"),
				 MAXFILOP: core.StringPtr("1024"),
				 MAXLOCKS: core.StringPtr("AUTOMATIC"),
				 MINDECDIV3: core.StringPtr("NO"),
				 MONACTMETRICS: core.StringPtr("EXTENDED"),
				 MONDEADLOCK: core.StringPtr("HISTORY"),
				 MONLCKMSGLVL: core.StringPtr("2"),
				 MONLOCKTIMEOUT: core.StringPtr("HISTORY"),
				 MONLOCKWAIT: core.StringPtr("WITHOUT_HIST"),
				 MONLWTHRESH: core.StringPtr("10000"),
				 MONOBJMETRICS: core.StringPtr("BASE"),
				 MONPKGLISTSZ: core.StringPtr("512"),
				 MONREQMETRICS: core.StringPtr("NONE"),
				 MONRTNDATA: core.StringPtr("BASE"),
				 MONRTNEXECLIST: core.StringPtr("ON"),
				 MONUOWDATA: core.StringPtr("NONE"),
				 MONUOWEXECLIST: core.StringPtr("ON"),
				 MONUOWPKGLIST: core.StringPtr("OFF"),
				 NCHARMAPPING: core.StringPtr("CHAR_CU32"),
				 NUMFREQVALUES: core.StringPtr("50"),
				 NUMIOCLEANERS: core.StringPtr("AUTOMATIC"),
				 NUMIOSERVERS: core.StringPtr("AUTOMATIC"),
				 NUMLOGSPAN: core.StringPtr("10"),
				 NUMQUANTILES: core.StringPtr("100"),
				 OPTBUFFPAGE: core.StringPtr("-"),
				 OPTDIRECTWRKLD: core.StringPtr("ON"),
				 OPTLOCKLIST: core.StringPtr("-"),
				 OPTMAXLOCKS: core.StringPtr("-"),
				 OPTSORTHEAP: core.StringPtr("-"),
				 PAGEAGETRGTGCR: core.StringPtr("5000"),
				 PAGEAGETRGTMCR: core.StringPtr("3000"),
				 PCKCACHESZ: core.StringPtr("AUTOMATIC"),
				 PLSTACKTRACE: core.StringPtr("UNHANDLED"),
				 SELFTUNINGMEM: core.StringPtr("ON"),
				 SEQDETECT: core.StringPtr("YES"),
				 SHEAPTHRESSHR: core.StringPtr("AUTOMATIC"),
				 SOFTMAX: core.StringPtr("-"),
				 SORTHEAP: core.StringPtr("AUTOMATIC"),
				 SQLCCFLAGS: core.StringPtr("-"),
				 STATHEAPSZ: core.StringPtr("AUTOMATIC"),
				 STMTHEAP: core.StringPtr("AUTOMATIC"),
				 STMTCONC: core.StringPtr("LITERALS"),
				 STRINGUNITS: core.StringPtr("SYSTEM"),
				 SYSTIMEPERIODADJ: core.StringPtr("NO"),
				 TRACKMOD: core.StringPtr("YES"),
				 UTILHEAPSZ: core.StringPtr("AUTOMATIC"),
				 WLMADMISSIONCTRL: core.StringPtr("YES"),
				 WLMAGENTLOADTRGT: core.StringPtr("1000"),
				 WLMCPULIMIT: core.StringPtr("80"),
				 WLMCPUSHARES: core.StringPtr("1000"),
				 WLMCPUSHAREMODE: core.StringPtr("SOFT"),
			 }
 
			 createCustomSettingsDbmModel := &db2saasv1.CreateCustomSettingsDbm{
				 COMMBANDWIDTH: core.StringPtr("1000"),
				 CPUSPEED: core.StringPtr("0.5"),
				 DFTMONBUFPOOL: core.StringPtr("ON"),
				 DFTMONLOCK: core.StringPtr("OFF"),
				 DFTMONSORT: core.StringPtr("ON"),
				 DFTMONSTMT: core.StringPtr("ON"),
				 DFTMONTABLE: core.StringPtr("OFF"),
				 DFTMONTIMESTAMP: core.StringPtr("ON"),
				 DFTMONUOW: core.StringPtr("ON"),
				 DIAGLEVEL: core.StringPtr("2"),
				 FEDERATEDASYNC: core.StringPtr("32767"),
				 INDEXREC: core.StringPtr("RESTART"),
				 INTRAPARALLEL: core.StringPtr("YES"),
				 KEEPFENCED: core.StringPtr("YES"),
				 MAXCONNRETRIES: core.StringPtr("5"),
				 MAXQUERYDEGREE: core.StringPtr("4"),
				 MONHEAPSZ: core.StringPtr("AUTOMATIC"),
				 MULTIPARTSIZEMB: core.StringPtr("100"),
				 NOTIFYLEVEL: core.StringPtr("2"),
				 NUMINITAGENTS: core.StringPtr("100"),
				 NUMINITFENCED: core.StringPtr("20"),
				 NUMPOOLAGENTS: core.StringPtr("10"),
				 RESYNCINTERVAL: core.StringPtr("1000"),
				 RQRIOBLK: core.StringPtr("8192"),
				 STARTSTOPTIME: core.StringPtr("10"),
				 UTILIMPACTLIM: core.StringPtr("50"),
				 WLMDISPATCHER: core.StringPtr("YES"),
				 WLMDISPCONCUR: core.StringPtr("16"),
				 WLMDISPCPUSHARES: core.StringPtr("YES"),
				 WLMDISPMINUTIL: core.StringPtr("10"),
			 }
 
			 postDb2SaasDbConfigurationOptions := &db2saasv1.PostDb2SaasDbConfigurationOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successPostCustomSettings, response, err := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successPostCustomSettings).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetDb2SaasTuneableParam - Retrieves the values of tunable parameters of the DB2 instance`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions *GetDb2SaasTuneableParamOptions)`, func() {
			 getDb2SaasTuneableParamOptions := &db2saasv1.GetDb2SaasTuneableParamOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successTuneableParams, response, err := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successTuneableParams).ToNot(BeNil())
		 })
	 })
 
	 Describe(`GetDb2SaasBackup - Get Db2 instance backup information`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasBackup(getDb2SaasBackupOptions *GetDb2SaasBackupOptions)`, func() {
			 getDb2SaasBackupOptions := &db2saasv1.GetDb2SaasBackupOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successGetBackups, response, err := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetBackups).ToNot(BeNil())
		 })
	 })
 
	 Describe(`PostDb2SaasBackup - Create backup of an instance`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`PostDb2SaasBackup(postDb2SaasBackupOptions *PostDb2SaasBackupOptions)`, func() {
			 postDb2SaasBackupOptions := &db2saasv1.PostDb2SaasBackupOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
			 }
 
			 successGetBackup, response, err := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetBackup).ToNot(BeNil())
		 })
	 })
 
	 Describe(`DeleteDb2SaasUser - Delete a user (only platform admin)`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`DeleteDb2SaasUser(deleteDb2SaasUserOptions *DeleteDb2SaasUserOptions)`, func() {
			 deleteDb2SaasUserOptions := &db2saasv1.DeleteDb2SaasUserOptions{
				 XDeploymentID: core.StringPtr("crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"),
				 ID: core.StringPtr("test-user"),
			 }
 
			 result, response, err := db2saasService.DeleteDb2SaasUser(deleteDb2SaasUserOptions)
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(result).ToNot(BeNil())
		 })
	 })
 })
 
 //
 // Utility functions are declared in the unit test file
 //
 