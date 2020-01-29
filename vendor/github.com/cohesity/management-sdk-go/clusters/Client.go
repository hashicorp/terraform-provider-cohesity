// Copyright 2019 Cohesity Inc.
package clusters

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/configuration"
	"github.com/cohesity/management-sdk-go/models"
	"github.com/cohesity/management-sdk-go/unirest-go"
)

/*
 * Client structure as interface implementation
 */
type CLUSTERS_IMPL struct {
	config configuration.CONFIGURATION
}

/**
 * Returns the Public Keys for the cluster.
 * @return	Returns the *models.ClusterPublicKeys response from the API call
 */
func (me *CLUSTERS_IMPL) GetClusterKeys() (*models.ClusterPublicKeys, error) {
	//the endpoint path uri
	_pathUrl := "/public/cluster/keys"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ClusterPublicKeys = &models.ClusterPublicKeys{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to destroy a Cohesity Cluster and returns some information
 * about the operation and each Node.
 * @return	Returns the  response from the API call
 */
func (me *CLUSTERS_IMPL) DestroyCluster() error {
	//the endpoint path uri
	_pathUrl := "/public/clusters"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}
	if me.config.AccessToken() == nil {
		return errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Delete(_queryBuilder, headers, nil)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil

}

/**
 * Sends a request to create a new Cloud Edition Cohesity Cluster and returns
 * the IDs, name, and software version of the new cluster. Also returns the
 * status of each node.
 * @param    *models.CreateCloudClusterParameters        body     parameter: Required
 * @return	Returns the *models.CreateClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) CreateCloudCluster(
	body *models.CreateCloudClusterParameters) (*models.CreateClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/cloudEdition"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.CreateClusterResult = &models.CreateClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to expand a Cloud Edition Cohesity Cluster and returns some
 * information about the request and each new Node.
 * @param    *models.ExpandCloudClusterParameters        body     parameter: Required
 * @return	Returns the *models.CreateClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) CreateExpandCloudCluster(
	body *models.ExpandCloudClusterParameters) (*models.CreateClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/cloudEdition/nodes"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.CreateClusterResult = &models.CreateClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to check the progress of the creation of a new Cohesity
 * Cluster and returns some information about the creation process along
 * with an estimated time remaining and completion percentage.
 * @return	Returns the *models.ClusterCreationProgressResult response from the API call
 */
func (me *CLUSTERS_IMPL) GetClusterCreationProgress() (*models.ClusterCreationProgressResult, error) {
	//the endpoint path uri
	_pathUrl := "/public/clusters/creationProgress"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ClusterCreationProgressResult = &models.ClusterCreationProgressResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Get the IO preferential tiers of the cluster.
 * @return	Returns the *models.IoPreferentialTier response from the API call
 */
func (me *CLUSTERS_IMPL) GetIoPreferentialTier() (*models.IoPreferentialTier, error) {
	//the endpoint path uri
	_pathUrl := "/public/clusters/ioPreferentialTier"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.IoPreferentialTier = &models.IoPreferentialTier{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to remove a Node from a Cohesity Cluster.
 * @param    int64        id     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *CLUSTERS_IMPL) RemoveNode(
	id int64) error {
	//the endpoint path uri
	_pathUrl := "/public/clusters/nodes/{id}"

	//variable to hold errors
	var err error = nil
	//process optional template parameters
	_pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		//error in template param handling
		return err
	}

	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}
	if me.config.AccessToken() == nil {
		return errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Delete(_queryBuilder, headers, nil)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil

}

/**
 * Sends a request to create a new Physical Edition Cohesity Cluster and returns
 * the IDs, name, and software version of the new cluster. Also returns the
 * status of each node.
 * @param    *models.CreatePhysicalClusterParameters        body     parameter: Required
 * @return	Returns the *models.CreateClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) CreatePhysicalCluster(
	body *models.CreatePhysicalClusterParameters) (*models.CreateClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/physicalEdition"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.CreateClusterResult = &models.CreateClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to expand a Physical Edition Cohesity Cluster and returns some
 * information about the request and each new Node.
 * @param    *models.ExpandPhysicalClusterParameters        body     parameter: Required
 * @return	Returns the *models.CreateClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) CreateExpandPhysicalCluster(
	body *models.ExpandPhysicalClusterParameters) (*models.CreateClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/physicalEdition/nodes"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.CreateClusterResult = &models.CreateClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to list the states of all of the services on a Cluster.
 * @return	Returns the []*models.ServiceStateResult response from the API call
 */
func (me *CLUSTERS_IMPL) ListServiceStates() ([]*models.ServiceStateResult, error) {
	//the endpoint path uri
	_pathUrl := "/public/clusters/services/states"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal []*models.ServiceStateResult
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to either stop, start, or restart one or more of the services
 * on a Cohesity Cluster and returns a message describing the result.
 * @param    *models.ChangeServiceStateParameters        body     parameter: Required
 * @return	Returns the *models.ChangeServiceStateResult response from the API call
 */
func (me *CLUSTERS_IMPL) ChangeServiceState(
	body *models.ChangeServiceStateParameters) (*models.ChangeServiceStateResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/services/states"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ChangeServiceStateResult = &models.ChangeServiceStateResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to upgrade the software version of a Cohesity Cluster
 * and returns a message specifying the result. Before using this, you need to
 * use the /public/packages endpoint to upload a new package to the Cluster.
 * @param    *models.UpgradeClusterParameters        body     parameter: Required
 * @return	Returns the *models.UpgradeClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) UpdateUpgradeCluster(
	body *models.UpgradeClusterParameters) (*models.UpgradeClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/software"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Put(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.UpgradeClusterResult = &models.UpgradeClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Sends a request to create a new Virtual Edition Cohesity Cluster and returns
 * the IDs, name and software version of the new cluster.
 * @param    *models.CreateVirtualClusterParameters        body     parameter: Required
 * @return	Returns the *models.CreateClusterResult response from the API call
 */
func (me *CLUSTERS_IMPL) CreateVirtualCluster(
	body *models.CreateVirtualClusterParameters) (*models.CreateClusterResult, error) {
	//validating required parameters
	if body == nil {
		return nil, errors.New("The parameter 'body' is a required parameter and cannot be nil.")
	} //the endpoint path uri
	_pathUrl := "/public/clusters/virtualEdition"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.CreateClusterResult = &models.CreateClusterResult{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Returns the external Client Subnets for the cluster.
 * @return	Returns the *models.ExternalClientSubnets response from the API call
 */
func (me *CLUSTERS_IMPL) GetExternalClientSubnets() (*models.ExternalClientSubnets, error) {
	//the endpoint path uri
	_pathUrl := "/public/externalClientSubnets"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ExternalClientSubnets = &models.ExternalClientSubnets{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

/**
 * Returns the updated external Client Subnets of the cluster.
 * @param    *models.ExternalClientSubnets        body     parameter: Optional
 * @return	Returns the *models.ExternalClientSubnets response from the API call
 */
func (me *CLUSTERS_IMPL) UpdateExternalClientSubnets(
	body *models.ExternalClientSubnets) (*models.ExternalClientSubnets, error) {
	//the endpoint path uri
	_pathUrl := "/public/externalClientSubnets"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)

	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}
	if me.config.AccessToken() == nil {
		return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.2",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Put(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.ExternalClientSubnets = &models.ExternalClientSubnets{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil

}

func (me *CLUSTERS_IMPL) ApplyClusterLicence(
	body *models.LicenceClusterParameters) error {
	//the endpoint path uri
	_pathUrl := "/licenseAgreement"

	//variable to hold errors
	var err error = nil
	//the base uri for api requests
	_queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST, me.config)
	//prepare query string for API call
	_queryBuilder = _queryBuilder + _pathUrl

	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}
	if me.config.AccessToken() == nil {
		return errors.New("Access Token not set. Please authorize the client using client.Authorize()")
	}
	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"user-agent":    "cohesity-Go-sdk-1.1.0",
		"accept":        "application/json",
		"content-type":  "application/json; charset=utf-8",
		"Authorization": fmt.Sprintf("%s %s", *me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request, me.config.SkipSSL())
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if _response.Code == 0 {
		err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
	} else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}
	return nil
}
