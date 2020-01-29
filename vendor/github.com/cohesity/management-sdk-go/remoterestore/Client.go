// Copyright 2019 Cohesity Inc.
package remoterestore


import(
	"errors"
	"fmt"
	"encoding/json"
	"github.com/cohesity/management-sdk-go/unirest-go"
	"github.com/cohesity/management-sdk-go/apihelper"
	"github.com/cohesity/management-sdk-go/configuration"
	"github.com/cohesity/management-sdk-go/models"
)
/*
 * Client structure as interface implementation
 */
type REMOTERESTORE_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * This request contains multiple files stored as multipart mime data.
 * Each file has a key used to encrypt data between a remote Cluster and the
 * remote Vault.
 * Content of the file should be same as the file downloaded from the remote
 * Cluster.
 * @param    int64                                   id       parameter: Required
 * @param    []*models.VaultEncryptionKey        body     parameter: Optional
 * @return	Returns the  response from the API call
 */
func (me *REMOTERESTORE_IMPL) UploadVaultEncryptionKeys (
            id int64,
            body []*models.VaultEncryptionKey) (error) {
    //the endpoint path uri
    _pathUrl := "/public/remoteVaults/encryptionKeys/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
     if me.config.AccessToken() == nil {
        return errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Put(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

/**
 * A remote Vault restore task can restore archived data from a Vault
 * (External Target) to this local Cluster.
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * @return	Returns the []*models.RemoteVaultRestoreTaskStatus response from the API call
 */
func (me *REMOTERESTORE_IMPL) ListRemoteVaultRestoreTasks () ([]*models.RemoteVaultRestoreTaskStatus, error) {
    //the endpoint path uri
    _pathUrl := "/public/remoteVaults/restoreTasks"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.RemoteVaultRestoreTaskStatus
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Returns the id of the remote Vault restore Task that was created.
 * After a Vault is searched by a search Job, this operation can be
 * called to create a task that restores the indexes and/or the Snapshots
 * of a Protection Job, which were archived on a remote Vault (External Target).
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * @param    *models.CreateRemoteVaultRestoreTaskParameters        body     parameter: Required
 * @return	Returns the *models.UniversalId response from the API call
 */
func (me *REMOTERESTORE_IMPL) CreateRemoteVaultRestoreTask (
            body *models.CreateRemoteVaultRestoreTaskParameters) (*models.UniversalId, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/remoteVaults/restoreTasks"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.UniversalId = &models.UniversalId{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Specify a unique id of the search Job using a combination of the
 * searchJobId, clusterId, and clusterIncarnationId parameters,
 * which are all required.
 * The results can be optionally filtered by the remote Cluster name.
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from a remote Vault.
 * @param    int64          searchJobId              parameter: Required
 * @param    int64          clusterId                parameter: Required
 * @param    int64          clusterIncarnationId     parameter: Required
 * @param    *int64         pageCount                parameter: Optional
 * @param    *string        clusterName              parameter: Optional
 * @param    *string        cookie                   parameter: Optional
 * @return	Returns the *models.RemoteVaultSearchJobResults response from the API call
 */
func (me *REMOTERESTORE_IMPL) GetRemoteVaultSearchJobResults (
            searchJobId int64,
            clusterId int64,
            clusterIncarnationId int64,
            pageCount *int64,
            clusterName *string,
            cookie *string) (*models.RemoteVaultSearchJobResults, error) {
    //the endpoint path uri
    _pathUrl := "/public/remoteVaults/searchJobResults"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "searchJobId" : searchJobId,
        "clusterId" : clusterId,
        "clusterIncarnationId" : clusterIncarnationId,
        "pageCount" : pageCount,
        "clusterName" : clusterName,
        "cookie" : cookie,
    })
    if err != nil {
        //error in query param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RemoteVaultSearchJobResults = &models.RemoteVaultSearchJobResults{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * @param    *models.StopRemoteVaultSearchJobParameters        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *REMOTERESTORE_IMPL) DeleteStopRemoteVaultSearchJob (
            body *models.StopRemoteVaultSearchJobParameters) (error) {
//validating required parameters
    if (body == nil){
        return errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/remoteVaults/searchJobs"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }
     if me.config.AccessToken() == nil {
        return errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Delete(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil

}

/**
 * List all the searches of remote Vaults (External Targets) that
 * have run or are running on this Cohesity Cluster.
 * A search finds Protection Jobs that have archived to a
 * Vault (External Target).
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * NOTE: A Vault is equivalent to an External Target in the Cohesity Dashboard.
 * A search Job is equivalent to a search task in the Cohesity Dashboard.
 * @return	Returns the []*models.RemoteVaultSearchJobInformation response from the API call
 */
func (me *REMOTERESTORE_IMPL) ListRemoteVaultSearchJobs () ([]*models.RemoteVaultSearchJobInformation, error) {
    //the endpoint path uri
    _pathUrl := "/public/remoteVaults/searchJobs"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models.RemoteVaultSearchJobInformation
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * A search Job finds Protection Jobs that archived data to a
 * Vault (External Target) which also match the specified search criteria.
 * The results can be optionally filtered by specifying a Cluster match string,
 * a Protection Job match string, a start time and an end time.
 * This is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * NOTE: A Vault is equivalent to an External Target in the Cohesity Dashboard.
 * A search Job is equivalent to a search task in the Cohesity Dashboard.
 * @param    *models.CreateRemoteVaultSearchJobParameters        body     parameter: Required
 * @return	Returns the *models.CreatedRemoteVaultSearchJobUid response from the API call
 */
func (me *REMOTERESTORE_IMPL) CreateRemoteVaultSearchJob (
            body *models.CreateRemoteVaultSearchJobParameters) (*models.CreatedRemoteVaultSearchJobUid, error) {
//validating required parameters
    if (body == nil){
        return nil,errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/remoteVaults/searchJobs"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Post(_queryBuilder, headers, body)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.CreatedRemoteVaultSearchJobUid = &models.CreatedRemoteVaultSearchJobUid{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Specify an id for a completed or running search Job.
 * A search Job finds data that has been archived to a Vault (External Target).
 * The returned results do not include Job Run (Snapshot) information.
 * It is part of the CloudRetrieve functionality for finding and restoring
 * archived data from remote Vaults to an alternative (non-original) Cluster.
 * @param    int64        id     parameter: Required
 * @return	Returns the *models.RemoteVaultSearchJobInformation response from the API call
 */
func (me *REMOTERESTORE_IMPL) ListRemoteVaultSearchJobById (
            id int64) (*models.RemoteVaultSearchJobInformation, error) {
    //the endpoint path uri
    _pathUrl := "/public/remoteVaults/searchJobs/{id}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "id" : id,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //validate and preprocess url
    _queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }
     if me.config.AccessToken() == nil {
        return nil, errors.New("Access Token not set. Please authorize the client using client.Authorize()");
    }
    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "cohesity-Go-sdk-1.1.2",
        "accept" : "application/json",
        "Authorization" : fmt.Sprintf("%s %s",*me.config.AccessToken().TokenType, *me.config.AccessToken().AccessToken),
    }

    //prepare API request
    _request := unirest.Get(_queryBuilder, headers)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request,me.config.SkipSSL());
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 0) {
        err = apihelper.NewAPIError("Error", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models.RemoteVaultSearchJobInformation = &models.RemoteVaultSearchJobInformation{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

