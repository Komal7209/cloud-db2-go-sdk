//go:build examples

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
	 "encoding/json"
	 "fmt"
	 "os"
 
	 "github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	 "github.com/IBM/go-sdk-core/v5/core"
	 . "github.com/onsi/ginkgo"
	 . "github.com/onsi/gomega"
 )
 
 //
 // This file provides an example of how to use the db2saas service.
 //
 // The following configuration properties are assumed to be defined:
 // DB2SAAS_URL=<service base url>
 // DB2SAAS_AUTH_TYPE=iam
 // DB2SAAS_APIKEY=<IAM apikey>
 // DB2SAAS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
 //
 // These configuration properties can be exported as environment variables, or stored
 // in a configuration file and then:
 // export IBM_CREDENTIALS_FILE=<name of configuration file>
 //
 var _ = Describe(`Db2saasV1 Examples Tests`, func() {
 
	 const externalConfigFile = "../db2saas_v1.env"
 
	 var (
		 db2saasService *db2saasv1.Db2saasV1
		 config       map[string]string
	 )
 
	 var shouldSkipTest = func() {
		 Skip("External configuration is not available, skipping examples...")
	 }
 
	 Describe(`External configuration`, func() {
		 It("Successfully load the configuration", func() {
			 var err error
			 _, err = os.Stat(externalConfigFile)
			 if err != nil {
				 Skip("External configuration file not found, skipping examples: " + err.Error())
			 }
 
			 os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			 config, err = core.GetServiceProperties(db2saasv1.DefaultServiceName)
			 if err != nil {
				 Skip("Error loading service properties, skipping examples: " + err.Error())
			 } else if len(config) == 0 {
				 Skip("Unable to load service properties, skipping examples")
			 }
 
			 shouldSkipTest = func() {}
		 })
	 })
 
	 Describe(`Client initialization`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It("Successfully construct the service client instance", func() {
			 var err error
 
			 // begin-common
 
			 db2saasServiceOptions := &db2saasv1.Db2saasV1Options{}
 
			 db2saasService, err = db2saasv1.NewDb2saasV1UsingExternalConfig(db2saasServiceOptions)
 
			 if err != nil {
				 panic(err)
			 }
 
			 // end-common
 
			 Expect(db2saasService).ToNot(BeNil())
		 })
	 })
 
	 Describe(`Db2saasV1 request examples`, func() {
		 BeforeEach(func() {
			 shouldSkipTest()
		 })
		 It(`GetDb2SaasConnectionInfo request example`, func() {
			 fmt.Println("\nGetDb2SaasConnectionInfo() result:")
			 // begin-get_db2_saas_connection_info
 
			 getDb2SaasConnectionInfoOptions := db2saasService.NewGetDb2SaasConnectionInfoOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A69db420f-33d5-4953-8bd8-1950abd356f6%3A%3A",
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successConnectionInfo, response, err := db2saasService.GetDb2SaasConnectionInfo(getDb2SaasConnectionInfoOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successConnectionInfo, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_connection_info
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successConnectionInfo).ToNot(BeNil())
		 })
		 It(`GetDb2SaasAllowlist request example`, func() {
			 fmt.Println("\nGetDb2SaasAllowlist() result:")
			 // begin-get_db2_saas_allowlist
 
			 getDb2SaasAllowlistOptions := db2saasService.NewGetDb2SaasAllowlistOptions(
				 "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::",
			 )
 
			 successGetAllowlistIPs, response, err := db2saasService.GetDb2SaasAllowlist(getDb2SaasAllowlistOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successGetAllowlistIPs, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_allowlist
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetAllowlistIPs).ToNot(BeNil())
		 })
		 It(`PutDb2SaasAutoscale request example`, func() {
			 fmt.Println("\nPutDb2SaasAutoscale() result:")
			 // begin-put_db2_saas_autoscale
 
			 putDb2SaasAutoscaleOptions := db2saasService.NewPutDb2SaasAutoscaleOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A",
			 )
 
			 successUpdateAutoScale, response, err := db2saasService.PutDb2SaasAutoscale(putDb2SaasAutoscaleOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successUpdateAutoScale, "", "  ")
			 fmt.Println(string(b))
 
			 // end-put_db2_saas_autoscale
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successUpdateAutoScale).ToNot(BeNil())
		 })
		 It(`GetDb2SaasAutoscale request example`, func() {
			 fmt.Println("\nGetDb2SaasAutoscale() result:")
			 // begin-get_db2_saas_autoscale
 
			 getDb2SaasAutoscaleOptions := db2saasService.NewGetDb2SaasAutoscaleOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A",
			 )
 
			 successAutoScaling, response, err := db2saasService.GetDb2SaasAutoscale(getDb2SaasAutoscaleOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successAutoScaling, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_autoscale
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successAutoScaling).ToNot(BeNil())
		 })
		 It(`PostDb2SaasDbConfiguration request example`, func() {
			 fmt.Println("\nPostDb2SaasDbConfiguration() result:")
			 // begin-post_db2_saas_db_configuration
 
			 postDb2SaasDbConfigurationOptions := db2saasService.NewPostDb2SaasDbConfigurationOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A",
			 )
 
			 successPostCustomSettings, response, err := db2saasService.PostDb2SaasDbConfiguration(postDb2SaasDbConfigurationOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successPostCustomSettings, "", "  ")
			 fmt.Println(string(b))
 
			 // end-post_db2_saas_db_configuration
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successPostCustomSettings).ToNot(BeNil())
		 })
		 It(`GetDb2SaasTuneableParam request example`, func() {
			 fmt.Println("\nGetDb2SaasTuneableParam() result:")
			 // begin-get_db2_saas_tuneable_param
 
			 getDb2SaasTuneableParamOptions := db2saasService.NewGetDb2SaasTuneableParamOptions()
 
			 successTuneableParams, response, err := db2saasService.GetDb2SaasTuneableParam(getDb2SaasTuneableParamOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successTuneableParams, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_tuneable_param
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successTuneableParams).ToNot(BeNil())
		 })
		 It(`GetDb2SaasBackup request example`, func() {
			 fmt.Println("\nGetDb2SaasBackup() result:")
			 // begin-get_db2_saas_backup
 
			 getDb2SaasBackupOptions := db2saasService.NewGetDb2SaasBackupOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A",
			 )
 
			 successGetBackups, response, err := db2saasService.GetDb2SaasBackup(getDb2SaasBackupOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successGetBackups, "", "  ")
			 fmt.Println(string(b))
 
			 // end-get_db2_saas_backup
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successGetBackups).ToNot(BeNil())
		 })
		 It(`PostDb2SaasBackup request example`, func() {
			 fmt.Println("\nPostDb2SaasBackup() result:")
			 // begin-post_db2_saas_backup
 
			 postDb2SaasBackupOptions := db2saasService.NewPostDb2SaasBackupOptions(
				 "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A",
			 )
 
			 successCreateBackup, response, err := db2saasService.PostDb2SaasBackup(postDb2SaasBackupOptions)
			 if err != nil {
				 panic(err)
			 }
			 b, _ := json.MarshalIndent(successCreateBackup, "", "  ")
			 fmt.Println(string(b))
 
			 // end-post_db2_saas_backup
 
			 Expect(err).To(BeNil())
			 Expect(response.StatusCode).To(Equal(200))
			 Expect(successCreateBackup).ToNot(BeNil())
		 })
	 })
 })
 