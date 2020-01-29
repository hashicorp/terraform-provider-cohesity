// Copyright 2019 Cohesity Inc.
package statistics


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
type STATISTICS_IMPL struct {
     config configuration.CONFIGURATION
}

/**
 * An entity is an object found on the Cohesity Cluster, such as a disk or a
 * Node.
 * In the Cohesity Dashboard, similar functionality is provided in Advanced
 * Diagnostics.
 * @param    string          schemaName                   parameter: Required
 * @param    *bool           includeAggrMetricSources     parameter: Optional
 * @param    []string        metricNames                  parameter: Optional
 * @return	Returns the []*models.EntityProto response from the API call
 */
func (me *STATISTICS_IMPL) GetEntities (
            schemaName string,
            includeAggrMetricSources *bool,
            metricNames []string) ([]*models.EntityProto, error) {
    //the endpoint path uri
    _pathUrl := "/public/statistics/entities"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "schemaName" : schemaName,
        "includeAggrMetricSources" : includeAggrMetricSources,
        "metricNames" : metricNames,
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
    var retVal []*models.EntityProto
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * An entity schema specifies the meta-data associated with entity such as
 * the list of attributes and a time series of data.
 * For example for a Disk entity, the entity schema specifies the Node that is
 * using this Disk, the type of the Disk, and Metrics about the Disk such as Space
 * Usage, Read IOs and Write IOs. Metrics define data points (time series data)
 * to track over a period of time for a specific interval.
 * If no parameters are specified, all entity schemas found on the Cohesity
 * Cluster are returned.
 * Specifying parameters filters the results that are returned.
 * In the Cohesity Dashboard, similar functionality is provided in Advanced
 * Diagnostics.
 * @param    []string        schemaNames     parameter: Optional
 * @param    []string        metricNames     parameter: Optional
 * @return	Returns the []*models.EntitySchemaProto response from the API call
 */
func (me *STATISTICS_IMPL) GetEntitiesSchema (
            schemaNames []string,
            metricNames []string) ([]*models.EntitySchemaProto, error) {
    //the endpoint path uri
    _pathUrl := "/public/statistics/entitiesSchema"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "schemaNames" : schemaNames,
        "metricNames" : metricNames,
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
    var retVal []*models.EntitySchemaProto
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * An entity schema specifies the meta-data associated with entity such as the
 * list of attributes and a time series of data.
 * For example for a Disk entity, the entity schema specifies the Node that is
 * using this Disk, the type of the Disk, and Metrics about the Disk such as
 * Space Usage, Read IOs and Write IOs. Metrics define data points (time series
 * data) to track over a period of time for a specific interval.
 * In the Cohesity Dashboard, similar functionality is provided in Advanced
 * Diagnostics.
 * @param    string        schemaName     parameter: Required
 * @return	Returns the []*models.EntitySchemaProto response from the API call
 */
func (me *STATISTICS_IMPL) GetEntitySchemaByName (
            schemaName string) ([]*models.EntitySchemaProto, error) {
    //the endpoint path uri
    _pathUrl := "/public/statistics/entitiesSchema/{schemaName}"

    //variable to hold errors
    var err error = nil
    //process optional template parameters
    _pathUrl, err = apihelper.AppendUrlWithTemplateParameters(_pathUrl, map[string]interface{} {
        "schemaName" : schemaName,
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
    var retVal []*models.EntitySchemaProto
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets the Apollo schema information for an entity to list a series of data
 * points.
 * @param    models.EntityTypeEnum        entityType     parameter: Required
 * @param    int64                            entityId       parameter: Required
 * @param    string                           entityName     parameter: Required
 * @return	Returns the *models.TimeSeriesSchemaResponse response from the API call
 */
func (me *STATISTICS_IMPL) GetTimeSeriesSchema (
            entityType models.EntityTypeEnum,
            entityId int64,
            entityName string) (*models.TimeSeriesSchemaResponse, error) {
    //the endpoint path uri
    _pathUrl := "/public/statistics/timeSeriesSchema"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "entityType" : models.EntityTypeEnumToValue(entityType),
        "entityId" : entityId,
        "entityName" : entityName,
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
    var retVal *models.TimeSeriesSchemaResponse = &models.TimeSeriesSchemaResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * A Metric specifies a data point (such as CPU usage and IOPS) to track over a
 * period of time.
 * For example for a disk in the Cluster, you can report on the 'Disk Health'
 * (kDiskAwaitTimeMsecs) Metric of the 'Disk Health Metrics' (kSentryDiskStats)
 * Schema for the last week.
 * You must specify the 'k' names as input and not the descriptive names.
 * You must also specify the id of the entity that you are reporting on such as
 * a Cluster, disk drive, job, etc.
 * Get the entityId by running the GET /public/statistics/entities operation.
 * In the Cohesity Dashboard, similar functionality is provided in Advanced
 * Diagnostics.
 * @param    string          schemaName             parameter: Required
 * @param    string          metricName             parameter: Required
 * @param    int64           startTimeMsecs         parameter: Required
 * @param    *string         entityId               parameter: Optional
 * @param    []string        entityIdList           parameter: Optional
 * @param    *int64          endTimeMsecs           parameter: Optional
 * @param    *string         rollupFunction         parameter: Optional
 * @param    *int64          rollupIntervalSecs     parameter: Optional
 * @return	Returns the *models.MetricDataBlock response from the API call
 */
func (me *STATISTICS_IMPL) GetTimeSeriesStats (
            schemaName string,
            metricName string,
            startTimeMsecs int64,
            entityId *string,
            entityIdList []string,
            endTimeMsecs *int64,
            rollupFunction *string,
            rollupIntervalSecs *int64) (*models.MetricDataBlock, error) {
    //the endpoint path uri
    _pathUrl := "/public/statistics/timeSeriesStats"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "schemaName" : schemaName,
        "metricName" : metricName,
        "startTimeMsecs" : startTimeMsecs,
        "entityId" : entityId,
        "entityIdList" : entityIdList,
        "endTimeMsecs" : endTimeMsecs,
        "rollupFunction" : rollupFunction,
        "rollupIntervalSecs" : rollupIntervalSecs,
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
    var retVal *models.MetricDataBlock = &models.MetricDataBlock{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

/**
 * Gets the progress and status of tasks.
 * @param    []string        taskPaths                parameter: Optional
 * @param    *bool           includeFinishedTasks     parameter: Optional
 * @param    *int64          startTimeSeconds         parameter: Optional
 * @param    *int64          endTimeSeconds           parameter: Optional
 * @param    *int64          maxTasks                 parameter: Optional
 * @param    *bool           excludeSubTasks          parameter: Optional
 * @param    []string        attributes               parameter: Optional
 * @return	Returns the []*models.Task response from the API call
 */
func (me *STATISTICS_IMPL) GetTasks (
            taskPaths []string,
            includeFinishedTasks *bool,
            startTimeSeconds *int64,
            endTimeSeconds *int64,
            maxTasks *int64,
            excludeSubTasks *bool,
            attributes []string) ([]*models.Task, error) {
    //the endpoint path uri
    _pathUrl := "/public/tasks/status"

    //variable to hold errors
    var err error = nil
    //the base uri for api requests
    _queryBuilder := configuration.GetBaseURI(configuration.DEFAULT_HOST,me.config);

    //prepare query string for API call
   _queryBuilder = _queryBuilder + _pathUrl

    //process optional query parameters
    _queryBuilder, err = apihelper.AppendUrlWithQueryParameters(_queryBuilder, map[string]interface{} {
        "taskPaths" : taskPaths,
        "includeFinishedTasks" : includeFinishedTasks,
        "startTimeSeconds" : startTimeSeconds,
        "endTimeSeconds" : endTimeSeconds,
        "maxTasks" : maxTasks,
        "excludeSubTasks" : excludeSubTasks,
        "attributes" : attributes,
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
    var retVal []*models.Task
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil

}

