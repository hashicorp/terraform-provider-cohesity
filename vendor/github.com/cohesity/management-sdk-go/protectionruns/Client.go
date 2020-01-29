// Copyright 2019 Cohesity Inc.
package protectionruns


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
type PROTECTIONRUNS_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * If no parameters are specified, Job Runs currently
 * on the Cohesity Cluster are returned. Both running and completed Job Runs
 * are reported.
 * Specifying parameters filters the results that are returned.
 * @param    *int64          jobId                         parameter: Optional
 * @param    *bool           includeRpoSnapshots           parameter: Optional
 * @param    *int64          startedTimeUsecs              parameter: Optional
 * @param    *int64          startTimeUsecs                parameter: Optional
 * @param    *int64          endTimeUsecs                  parameter: Optional
 * @param    *int64          numRuns                       parameter: Optional
 * @param    *bool           excludeTasks                  parameter: Optional
 * @param    *int64          sourceId                      parameter: Optional
 * @param    []string        runTypes                      parameter: Optional
 * @param    *bool           excludeErrorRuns              parameter: Optional
 * @param    *bool           excludeNonRestoreableRuns     parameter: Optional
 * @return	Returns the []*models.ProtectionRunInstance response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) GetProtectionRuns (
            jobId *int64,
            includeRpoSnapshots *bool,
            startedTimeUsecs *int64,
            startTimeUsecs *int64,
            endTimeUsecs *int64,
            numRuns *int64,
            excludeTasks *bool,
            sourceId *int64,
            runTypes []string,
            excludeErrorRuns *bool,
            excludeNonRestoreableRuns *bool) ([]*models.ProtectionRunInstance, error) {
    //the endpoint path uri
    _pathUrl := "/public/protectionRuns"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "jobId" : jobId,
        "includeRpoSnapshots" : includeRpoSnapshots,
        "startedTimeUsecs" : startedTimeUsecs,
        "startTimeUsecs" : startTimeUsecs,
        "endTimeUsecs" : endTimeUsecs,
        "numRuns" : numRuns,
        "excludeTasks" : excludeTasks,
        "sourceId" : sourceId,
        "runTypes" : runTypes,
        "excludeErrorRuns" : excludeErrorRuns,
        "excludeNonRestoreableRuns" : excludeNonRestoreableRuns,
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
    var retVal []*models.ProtectionRunInstance
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Update the expiration date (retention period) for the specified Protection
 * Job Runs and their snapshots.
 * After an expiration time is reached, the Job Run and its snapshots are deleted.
 * If an expiration time of 0 is specified, a Job Run and its snapshots
 * are immediately deleted.
 * @param    *models.UpdateProtectionJobRunsParam        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) UpdateProtectionRuns (
            body *models.UpdateProtectionJobRunsParam) (error) {
//validating required parameters
    if (body == nil){
        return errors.New("The parameter 'body' is a required parameter and cannot be nil.")
}     //the endpoint path uri
    _pathUrl := "/public/protectionRuns"

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
 * Cancel a Protection Job run.
 * @param    int64                                          id       parameter: Required
 * @param    *models.CancelProtectionJobRunParam        body     parameter: Optional
 * @return	Returns the  response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) CreateCancelProtectionJobRun (
            id int64,
            body *models.CancelProtectionJobRunParam) (error) {
    //the endpoint path uri
    _pathUrl := "/public/protectionRuns/cancel/{id}"

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
    _request := unirest.Post(_queryBuilder, headers, body)
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
 * jobId, startTimeUsecs and taskId have to be provided to get the a list of
 * errors for a job run task.
 * @param    int64          jobId                   parameter: Required
 * @param    int64          startTimeUsecs          parameter: Required
 * @param    int64          taskId                  parameter: Required
 * @param    *int64         limitNumberOfErrors     parameter: Optional
 * @param    *string        paginationCookie        parameter: Optional
 * @return	Returns the *models.ProtectionRunErrors response from the API call
 */
func (me *PROTECTIONRUNS_IMPL) GetProtectionRunErrors (
            jobId int64,
            startTimeUsecs int64,
            taskId int64,
            limitNumberOfErrors *int64,
            paginationCookie *string) (*models.ProtectionRunErrors, error) {
    //the endpoint path uri
    _pathUrl := "/public/protectionRuns/errors"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "jobId" : jobId,
        "startTimeUsecs" : startTimeUsecs,
        "taskId" : taskId,
        "limitNumberOfErrors" : limitNumberOfErrors,
        "paginationCookie" : paginationCookie,
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
    var retVal *models.ProtectionRunErrors = &models.ProtectionRunErrors{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

