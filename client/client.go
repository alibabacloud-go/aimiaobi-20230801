// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelAsyncTaskRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskRequest) SetAgentKey(v string) *CancelAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelAsyncTaskRequest) SetTaskId(v string) *CancelAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type CancelAsyncTaskResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponseBody) SetCode(v string) *CancelAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetData(v bool) *CancelAsyncTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetHttpStatusCode(v int32) *CancelAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetMessage(v string) *CancelAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetRequestId(v string) *CancelAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAsyncTaskResponseBody) SetSuccess(v bool) *CancelAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type CancelAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelAsyncTaskResponse) SetHeaders(v map[string]*string) *CancelAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelAsyncTaskResponse) SetStatusCode(v int32) *CancelAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAsyncTaskResponse) SetBody(v *CancelAsyncTaskResponseBody) *CancelAsyncTaskResponse {
	s.Body = v
	return s
}

type ClearIntervenesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ClearIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ClearIntervenesRequest) SetAgentKey(v string) *ClearIntervenesRequest {
	s.AgentKey = &v
	return s
}

type ClearIntervenesResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ClearIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ClearIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBody) SetCode(v string) *ClearIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetData(v *ClearIntervenesResponseBodyData) *ClearIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ClearIntervenesResponseBody) SetHttpStatusCode(v int32) *ClearIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetMessage(v string) *ClearIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetRequestId(v string) *ClearIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClearIntervenesResponseBody) SetSuccess(v bool) *ClearIntervenesResponseBody {
	s.Success = &v
	return s
}

type ClearIntervenesResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	TaskId     *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ClearIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponseBodyData) SetFailIdList(v []*string) *ClearIntervenesResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ClearIntervenesResponseBodyData) SetTaskId(v string) *ClearIntervenesResponseBodyData {
	s.TaskId = &v
	return s
}

type ClearIntervenesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ClearIntervenesResponse) SetHeaders(v map[string]*string) *ClearIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ClearIntervenesResponse) SetStatusCode(v int32) *ClearIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearIntervenesResponse) SetBody(v *ClearIntervenesResponseBody) *ClearIntervenesResponse {
	s.Body = v
	return s
}

type CreateGeneratedContentRequest struct {
	AgentKey      *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content       *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentDomain *string   `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string   `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	Keywords      []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Prompt        *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TaskId        *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title         *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Uuid          *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentRequest) SetAgentKey(v string) *CreateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContent(v string) *CreateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentDomain(v string) *CreateGeneratedContentRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetContentText(v string) *CreateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetKeywords(v []*string) *CreateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *CreateGeneratedContentRequest) SetPrompt(v string) *CreateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTaskId(v string) *CreateGeneratedContentRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetTitle(v string) *CreateGeneratedContentRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentRequest) SetUuid(v string) *CreateGeneratedContentRequest {
	s.Uuid = &v
	return s
}

type CreateGeneratedContentShrinkRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentDomain  *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText    *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Uuid           *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s CreateGeneratedContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentShrinkRequest) SetAgentKey(v string) *CreateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContent(v string) *CreateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentDomain(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentDomain = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetContentText(v string) *CreateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *CreateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetPrompt(v string) *CreateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTaskId(v string) *CreateGeneratedContentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetTitle(v string) *CreateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateGeneratedContentShrinkRequest) SetUuid(v string) *CreateGeneratedContentShrinkRequest {
	s.Uuid = &v
	return s
}

type CreateGeneratedContentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponseBody) SetCode(v string) *CreateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetData(v int64) *CreateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *CreateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetMessage(v string) *CreateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetRequestId(v string) *CreateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGeneratedContentResponseBody) SetSuccess(v bool) *CreateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type CreateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *CreateGeneratedContentResponse) SetHeaders(v map[string]*string) *CreateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *CreateGeneratedContentResponse) SetStatusCode(v int32) *CreateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGeneratedContentResponse) SetBody(v *CreateGeneratedContentResponseBody) *CreateGeneratedContentResponse {
	s.Body = v
	return s
}

type CreateTokenRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s CreateTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateTokenRequest) SetAgentKey(v string) *CreateTokenRequest {
	s.AgentKey = &v
	return s
}

type CreateTokenResponseBody struct {
	Code           *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBody) SetCode(v string) *CreateTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTokenResponseBody) SetData(v *CreateTokenResponseBodyData) *CreateTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateTokenResponseBody) SetHttpStatusCode(v int32) *CreateTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTokenResponseBody) SetMessage(v string) *CreateTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTokenResponseBody) SetRequestId(v string) *CreateTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTokenResponseBody) SetSuccess(v bool) *CreateTokenResponseBody {
	s.Success = &v
	return s
}

type CreateTokenResponseBodyData struct {
	ExpiredTime *int64  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTokenResponseBodyData) SetExpiredTime(v int64) *CreateTokenResponseBodyData {
	s.ExpiredTime = &v
	return s
}

func (s *CreateTokenResponseBodyData) SetToken(v string) *CreateTokenResponseBodyData {
	s.Token = &v
	return s
}

type CreateTokenResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenResponse) SetHeaders(v map[string]*string) *CreateTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenResponse) SetStatusCode(v int32) *CreateTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenResponse) SetBody(v *CreateTokenResponseBody) *CreateTokenResponse {
	s.Body = v
	return s
}

type DeleteGeneratedContentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentRequest) SetAgentKey(v string) *DeleteGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteGeneratedContentRequest) SetId(v int64) *DeleteGeneratedContentRequest {
	s.Id = &v
	return s
}

type DeleteGeneratedContentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponseBody) SetCode(v string) *DeleteGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetData(v bool) *DeleteGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetHttpStatusCode(v int32) *DeleteGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetMessage(v string) *DeleteGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetRequestId(v string) *DeleteGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGeneratedContentResponseBody) SetSuccess(v bool) *DeleteGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type DeleteGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *DeleteGeneratedContentResponse) SetHeaders(v map[string]*string) *DeleteGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *DeleteGeneratedContentResponse) SetStatusCode(v int32) *DeleteGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGeneratedContentResponse) SetBody(v *DeleteGeneratedContentResponseBody) *DeleteGeneratedContentResponse {
	s.Body = v
	return s
}

type DeleteInterveneRuleRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteInterveneRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleRequest) SetAgentKey(v string) *DeleteInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteInterveneRuleRequest) SetRuleId(v int64) *DeleteInterveneRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteInterveneRuleResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *DeleteInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInterveneRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBody) SetCode(v string) *DeleteInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetData(v *DeleteInterveneRuleResponseBodyData) *DeleteInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetHttpStatusCode(v int32) *DeleteInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetMessage(v string) *DeleteInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetRequestId(v string) *DeleteInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInterveneRuleResponseBody) SetSuccess(v bool) *DeleteInterveneRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteInterveneRuleResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	TaskId     *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteInterveneRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponseBodyData) SetFailIdList(v []*string) *DeleteInterveneRuleResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *DeleteInterveneRuleResponseBodyData) SetTaskId(v string) *DeleteInterveneRuleResponseBodyData {
	s.TaskId = &v
	return s
}

type DeleteInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInterveneRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponse) SetHeaders(v map[string]*string) *DeleteInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteInterveneRuleResponse) SetStatusCode(v int32) *DeleteInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponse) SetBody(v *DeleteInterveneRuleResponseBody) *DeleteInterveneRuleResponse {
	s.Body = v
	return s
}

type DeleteMaterialByIdRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteMaterialByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdRequest) SetAgentKey(v string) *DeleteMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *DeleteMaterialByIdRequest) SetId(v int64) *DeleteMaterialByIdRequest {
	s.Id = &v
	return s
}

type DeleteMaterialByIdResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMaterialByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponseBody) SetCode(v string) *DeleteMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetData(v bool) *DeleteMaterialByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetHttpStatusCode(v int32) *DeleteMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetMessage(v string) *DeleteMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetRequestId(v string) *DeleteMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterialByIdResponseBody) SetSuccess(v bool) *DeleteMaterialByIdResponseBody {
	s.Success = &v
	return s
}

type DeleteMaterialByIdResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMaterialByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *DeleteMaterialByIdResponse) SetHeaders(v map[string]*string) *DeleteMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *DeleteMaterialByIdResponse) SetStatusCode(v int32) *DeleteMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMaterialByIdResponse) SetBody(v *DeleteMaterialByIdResponseBody) *DeleteMaterialByIdResponse {
	s.Body = v
	return s
}

type ExportGeneratedContentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentRequest) SetAgentKey(v string) *ExportGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *ExportGeneratedContentRequest) SetId(v int64) *ExportGeneratedContentRequest {
	s.Id = &v
	return s
}

type ExportGeneratedContentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentResponseBody) SetCode(v string) *ExportGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetData(v string) *ExportGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetHttpStatusCode(v int32) *ExportGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetMessage(v string) *ExportGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetRequestId(v string) *ExportGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportGeneratedContentResponseBody) SetSuccess(v bool) *ExportGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type ExportGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *ExportGeneratedContentResponse) SetHeaders(v map[string]*string) *ExportGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *ExportGeneratedContentResponse) SetStatusCode(v int32) *ExportGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportGeneratedContentResponse) SetBody(v *ExportGeneratedContentResponseBody) *ExportGeneratedContentResponse {
	s.Body = v
	return s
}

type ExportIntervenesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ExportIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ExportIntervenesRequest) SetAgentKey(v string) *ExportIntervenesRequest {
	s.AgentKey = &v
	return s
}

type ExportIntervenesResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ExportIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExportIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponseBody) SetCode(v string) *ExportIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetData(v *ExportIntervenesResponseBodyData) *ExportIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ExportIntervenesResponseBody) SetHttpStatusCode(v int32) *ExportIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetMessage(v string) *ExportIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetRequestId(v string) *ExportIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportIntervenesResponseBody) SetSuccess(v bool) *ExportIntervenesResponseBody {
	s.Success = &v
	return s
}

type ExportIntervenesResponseBodyData struct {
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ExportIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponseBodyData) SetFileUrl(v string) *ExportIntervenesResponseBodyData {
	s.FileUrl = &v
	return s
}

type ExportIntervenesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ExportIntervenesResponse) SetHeaders(v map[string]*string) *ExportIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ExportIntervenesResponse) SetStatusCode(v int32) *ExportIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportIntervenesResponse) SetBody(v *ExportIntervenesResponseBody) *ExportIntervenesResponse {
	s.Body = v
	return s
}

type FeedbackDialogueRequest struct {
	AgentKey         *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CustomerResponse *string   `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	GoodText         *string   `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	ModifiedResponse *string   `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	Rating           *string   `json:"Rating,omitempty" xml:"Rating,omitempty"`
	RatingTags       []*string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty" type:"Repeated"`
	SessionId        *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId           *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueRequest) SetAgentKey(v string) *FeedbackDialogueRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueRequest) SetCustomerResponse(v string) *FeedbackDialogueRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetGoodText(v string) *FeedbackDialogueRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueRequest) SetModifiedResponse(v string) *FeedbackDialogueRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRating(v string) *FeedbackDialogueRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueRequest) SetRatingTags(v []*string) *FeedbackDialogueRequest {
	s.RatingTags = v
	return s
}

func (s *FeedbackDialogueRequest) SetSessionId(v string) *FeedbackDialogueRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueRequest) SetTaskId(v string) *FeedbackDialogueRequest {
	s.TaskId = &v
	return s
}

type FeedbackDialogueShrinkRequest struct {
	AgentKey         *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CustomerResponse *string `json:"CustomerResponse,omitempty" xml:"CustomerResponse,omitempty"`
	GoodText         *string `json:"GoodText,omitempty" xml:"GoodText,omitempty"`
	ModifiedResponse *string `json:"ModifiedResponse,omitempty" xml:"ModifiedResponse,omitempty"`
	Rating           *string `json:"Rating,omitempty" xml:"Rating,omitempty"`
	RatingTagsShrink *string `json:"RatingTags,omitempty" xml:"RatingTags,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	TaskId           *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FeedbackDialogueShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueShrinkRequest) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueShrinkRequest) SetAgentKey(v string) *FeedbackDialogueShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetCustomerResponse(v string) *FeedbackDialogueShrinkRequest {
	s.CustomerResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetGoodText(v string) *FeedbackDialogueShrinkRequest {
	s.GoodText = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetModifiedResponse(v string) *FeedbackDialogueShrinkRequest {
	s.ModifiedResponse = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRating(v string) *FeedbackDialogueShrinkRequest {
	s.Rating = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetRatingTagsShrink(v string) *FeedbackDialogueShrinkRequest {
	s.RatingTagsShrink = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetSessionId(v string) *FeedbackDialogueShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *FeedbackDialogueShrinkRequest) SetTaskId(v string) *FeedbackDialogueShrinkRequest {
	s.TaskId = &v
	return s
}

type FeedbackDialogueResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FeedbackDialogueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponseBody) SetCode(v string) *FeedbackDialogueResponseBody {
	s.Code = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetHttpStatusCode(v int32) *FeedbackDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetMessage(v string) *FeedbackDialogueResponseBody {
	s.Message = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetRequestId(v string) *FeedbackDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetSuccess(v bool) *FeedbackDialogueResponseBody {
	s.Success = &v
	return s
}

type FeedbackDialogueResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FeedbackDialogueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackDialogueResponse) String() string {
	return tea.Prettify(s)
}

func (s FeedbackDialogueResponse) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponse) SetHeaders(v map[string]*string) *FeedbackDialogueResponse {
	s.Headers = v
	return s
}

func (s *FeedbackDialogueResponse) SetStatusCode(v int32) *FeedbackDialogueResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackDialogueResponse) SetBody(v *FeedbackDialogueResponseBody) *FeedbackDialogueResponse {
	s.Body = v
	return s
}

type FetchImageTaskRequest struct {
	AgentKey      *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ArticleTaskId *string   `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	TaskIdList    []*string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskRequest) SetAgentKey(v string) *FetchImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskRequest) SetArticleTaskId(v string) *FetchImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskRequest) SetTaskIdList(v []*string) *FetchImageTaskRequest {
	s.TaskIdList = v
	return s
}

type FetchImageTaskShrinkRequest struct {
	AgentKey         *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ArticleTaskId    *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	TaskIdListShrink *string `json:"TaskIdList,omitempty" xml:"TaskIdList,omitempty"`
}

func (s FetchImageTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *FetchImageTaskShrinkRequest) SetAgentKey(v string) *FetchImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetArticleTaskId(v string) *FetchImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *FetchImageTaskShrinkRequest) SetTaskIdListShrink(v string) *FetchImageTaskShrinkRequest {
	s.TaskIdListShrink = &v
	return s
}

type FetchImageTaskResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *FetchImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FetchImageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBody) SetCode(v string) *FetchImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetData(v *FetchImageTaskResponseBodyData) *FetchImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *FetchImageTaskResponseBody) SetHttpStatusCode(v int32) *FetchImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetMessage(v string) *FetchImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetRequestId(v string) *FetchImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchImageTaskResponseBody) SetSuccess(v bool) *FetchImageTaskResponseBody {
	s.Success = &v
	return s
}

type FetchImageTaskResponseBodyData struct {
	TaskInfoList []*FetchImageTaskResponseBodyDataTaskInfoList `json:"TaskInfoList,omitempty" xml:"TaskInfoList,omitempty" type:"Repeated"`
}

func (s FetchImageTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyData) SetTaskInfoList(v []*FetchImageTaskResponseBodyDataTaskInfoList) *FetchImageTaskResponseBodyData {
	s.TaskInfoList = v
	return s
}

type FetchImageTaskResponseBodyDataTaskInfoList struct {
	Id         *int64                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	ImageList  []*FetchImageTaskResponseBodyDataTaskInfoListImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	TaskId     *string                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string                                                `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetId(v int64) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.Id = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetImageList(v []*FetchImageTaskResponseBodyDataTaskInfoListImageList) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.ImageList = v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskId(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskId = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoList) SetTaskStatus(v string) *FetchImageTaskResponseBodyDataTaskInfoList {
	s.TaskStatus = &v
	return s
}

type FetchImageTaskResponseBodyDataTaskInfoListImageList struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponseBodyDataTaskInfoListImageList) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetCode(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Code = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetMessage(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Message = &v
	return s
}

func (s *FetchImageTaskResponseBodyDataTaskInfoListImageList) SetUrl(v string) *FetchImageTaskResponseBodyDataTaskInfoListImageList {
	s.Url = &v
	return s
}

type FetchImageTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FetchImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FetchImageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchImageTaskResponse) GoString() string {
	return s.String()
}

func (s *FetchImageTaskResponse) SetHeaders(v map[string]*string) *FetchImageTaskResponse {
	s.Headers = v
	return s
}

func (s *FetchImageTaskResponse) SetStatusCode(v int32) *FetchImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchImageTaskResponse) SetBody(v *FetchImageTaskResponseBody) *FetchImageTaskResponse {
	s.Body = v
	return s
}

type GenerateFileUrlByKeyRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GenerateFileUrlByKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyRequest) SetAgentKey(v string) *GenerateFileUrlByKeyRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileKey(v string) *GenerateFileUrlByKeyRequest {
	s.FileKey = &v
	return s
}

func (s *GenerateFileUrlByKeyRequest) SetFileName(v string) *GenerateFileUrlByKeyRequest {
	s.FileName = &v
	return s
}

type GenerateFileUrlByKeyResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateFileUrlByKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponseBody) SetCode(v string) *GenerateFileUrlByKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetData(v string) *GenerateFileUrlByKeyResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetHttpStatusCode(v int32) *GenerateFileUrlByKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetMessage(v string) *GenerateFileUrlByKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetRequestId(v string) *GenerateFileUrlByKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFileUrlByKeyResponseBody) SetSuccess(v bool) *GenerateFileUrlByKeyResponseBody {
	s.Success = &v
	return s
}

type GenerateFileUrlByKeyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFileUrlByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFileUrlByKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateFileUrlByKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponse) SetHeaders(v map[string]*string) *GenerateFileUrlByKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetStatusCode(v int32) *GenerateFileUrlByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetBody(v *GenerateFileUrlByKeyResponseBody) *GenerateFileUrlByKeyResponse {
	s.Body = v
	return s
}

type GenerateImageTaskRequest struct {
	AgentKey      *string                                  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ArticleTaskId *string                                  `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	ParagraphList []*GenerateImageTaskRequestParagraphList `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty" type:"Repeated"`
	Size          *string                                  `json:"Size,omitempty" xml:"Size,omitempty"`
	Style         *string                                  `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequest) SetAgentKey(v string) *GenerateImageTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskRequest) SetArticleTaskId(v string) *GenerateImageTaskRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskRequest) SetParagraphList(v []*GenerateImageTaskRequestParagraphList) *GenerateImageTaskRequest {
	s.ParagraphList = v
	return s
}

func (s *GenerateImageTaskRequest) SetSize(v string) *GenerateImageTaskRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskRequest) SetStyle(v string) *GenerateImageTaskRequest {
	s.Style = &v
	return s
}

type GenerateImageTaskRequestParagraphList struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskRequestParagraphList) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskRequestParagraphList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskRequestParagraphList) SetContent(v string) *GenerateImageTaskRequestParagraphList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetId(v int64) *GenerateImageTaskRequestParagraphList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskId(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskRequestParagraphList) SetTaskStatus(v string) *GenerateImageTaskRequestParagraphList {
	s.TaskStatus = &v
	return s
}

type GenerateImageTaskShrinkRequest struct {
	AgentKey            *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ArticleTaskId       *string `json:"ArticleTaskId,omitempty" xml:"ArticleTaskId,omitempty"`
	ParagraphListShrink *string `json:"ParagraphList,omitempty" xml:"ParagraphList,omitempty"`
	Size                *string `json:"Size,omitempty" xml:"Size,omitempty"`
	Style               *string `json:"Style,omitempty" xml:"Style,omitempty"`
}

func (s GenerateImageTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskShrinkRequest) SetAgentKey(v string) *GenerateImageTaskShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetArticleTaskId(v string) *GenerateImageTaskShrinkRequest {
	s.ArticleTaskId = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetParagraphListShrink(v string) *GenerateImageTaskShrinkRequest {
	s.ParagraphListShrink = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetSize(v string) *GenerateImageTaskShrinkRequest {
	s.Size = &v
	return s
}

func (s *GenerateImageTaskShrinkRequest) SetStyle(v string) *GenerateImageTaskShrinkRequest {
	s.Style = &v
	return s
}

type GenerateImageTaskResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GenerateImageTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateImageTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBody) SetCode(v string) *GenerateImageTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetData(v *GenerateImageTaskResponseBodyData) *GenerateImageTaskResponseBody {
	s.Data = v
	return s
}

func (s *GenerateImageTaskResponseBody) SetHttpStatusCode(v int32) *GenerateImageTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetMessage(v string) *GenerateImageTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetRequestId(v string) *GenerateImageTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateImageTaskResponseBody) SetSuccess(v bool) *GenerateImageTaskResponseBody {
	s.Success = &v
	return s
}

type GenerateImageTaskResponseBodyData struct {
	TaskList []*GenerateImageTaskResponseBodyDataTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s GenerateImageTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyData) SetTaskList(v []*GenerateImageTaskResponseBodyDataTaskList) *GenerateImageTaskResponseBodyData {
	s.TaskList = v
	return s
}

type GenerateImageTaskResponseBodyDataTaskList struct {
	Content    *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GenerateImageTaskResponseBodyDataTaskList) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponseBodyDataTaskList) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetContent(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.Content = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetId(v int64) *GenerateImageTaskResponseBodyDataTaskList {
	s.Id = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskId(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskId = &v
	return s
}

func (s *GenerateImageTaskResponseBodyDataTaskList) SetTaskStatus(v string) *GenerateImageTaskResponseBodyDataTaskList {
	s.TaskStatus = &v
	return s
}

type GenerateImageTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateImageTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateImageTaskResponse) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponse) SetHeaders(v map[string]*string) *GenerateImageTaskResponse {
	s.Headers = v
	return s
}

func (s *GenerateImageTaskResponse) SetStatusCode(v int32) *GenerateImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateImageTaskResponse) SetBody(v *GenerateImageTaskResponseBody) *GenerateImageTaskResponse {
	s.Body = v
	return s
}

type GenerateUploadConfigRequest struct {
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	ParentDir *string `json:"ParentDir,omitempty" xml:"ParentDir,omitempty"`
}

func (s GenerateUploadConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigRequest) SetAgentKey(v string) *GenerateUploadConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetFileName(v string) *GenerateUploadConfigRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadConfigRequest) SetParentDir(v string) *GenerateUploadConfigRequest {
	s.ParentDir = &v
	return s
}

type GenerateUploadConfigResponseBody struct {
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GenerateUploadConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateUploadConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBody) SetCode(v string) *GenerateUploadConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetData(v *GenerateUploadConfigResponseBodyData) *GenerateUploadConfigResponseBody {
	s.Data = v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetHttpStatusCode(v int32) *GenerateUploadConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetMessage(v string) *GenerateUploadConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetRequestId(v string) *GenerateUploadConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateUploadConfigResponseBody) SetSuccess(v bool) *GenerateUploadConfigResponseBody {
	s.Success = &v
	return s
}

type GenerateUploadConfigResponseBodyData struct {
	FileKey   *string                `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FormDatas map[string]interface{} `json:"FormDatas,omitempty" xml:"FormDatas,omitempty"`
	PostUrl   *string                `json:"PostUrl,omitempty" xml:"PostUrl,omitempty"`
}

func (s GenerateUploadConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponseBodyData) SetFileKey(v string) *GenerateUploadConfigResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetFormDatas(v map[string]interface{}) *GenerateUploadConfigResponseBodyData {
	s.FormDatas = v
	return s
}

func (s *GenerateUploadConfigResponseBodyData) SetPostUrl(v string) *GenerateUploadConfigResponseBodyData {
	s.PostUrl = &v
	return s
}

type GenerateUploadConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateUploadConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponse) SetHeaders(v map[string]*string) *GenerateUploadConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadConfigResponse) SetStatusCode(v int32) *GenerateUploadConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponse) SetBody(v *GenerateUploadConfigResponseBody) *GenerateUploadConfigResponse {
	s.Body = v
	return s
}

type GenerateViewPointRequest struct {
	AgentKey      *string                                `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceData *GenerateViewPointRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
}

func (s GenerateViewPointRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequest) SetAgentKey(v string) *GenerateViewPointRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointRequest) SetReferenceData(v *GenerateViewPointRequestReferenceData) *GenerateViewPointRequest {
	s.ReferenceData = v
	return s
}

type GenerateViewPointRequestReferenceData struct {
	MiniDoc []*string `json:"MiniDoc,omitempty" xml:"MiniDoc,omitempty" type:"Repeated"`
}

func (s GenerateViewPointRequestReferenceData) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointRequestReferenceData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointRequestReferenceData) SetMiniDoc(v []*string) *GenerateViewPointRequestReferenceData {
	s.MiniDoc = v
	return s
}

type GenerateViewPointShrinkRequest struct {
	AgentKey            *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
}

func (s GenerateViewPointShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateViewPointShrinkRequest) SetAgentKey(v string) *GenerateViewPointShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *GenerateViewPointShrinkRequest) SetReferenceDataShrink(v string) *GenerateViewPointShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

type GenerateViewPointResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*GenerateViewPointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateViewPointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBody) SetCode(v string) *GenerateViewPointResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetData(v []*GenerateViewPointResponseBodyData) *GenerateViewPointResponseBody {
	s.Data = v
	return s
}

func (s *GenerateViewPointResponseBody) SetHttpStatusCode(v int32) *GenerateViewPointResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetMessage(v string) *GenerateViewPointResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetRequestId(v string) *GenerateViewPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateViewPointResponseBody) SetSuccess(v bool) *GenerateViewPointResponseBody {
	s.Success = &v
	return s
}

type GenerateViewPointResponseBodyData struct {
	Point *string `json:"Point,omitempty" xml:"Point,omitempty"`
}

func (s GenerateViewPointResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponseBodyData) SetPoint(v string) *GenerateViewPointResponseBodyData {
	s.Point = &v
	return s
}

type GenerateViewPointResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateViewPointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateViewPointResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateViewPointResponse) GoString() string {
	return s.String()
}

func (s *GenerateViewPointResponse) SetHeaders(v map[string]*string) *GenerateViewPointResponse {
	s.Headers = v
	return s
}

func (s *GenerateViewPointResponse) SetStatusCode(v int32) *GenerateViewPointResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateViewPointResponse) SetBody(v *GenerateViewPointResponseBody) *GenerateViewPointResponse {
	s.Body = v
	return s
}

type GetDataSourceOrderConfigRequest struct {
	AgentKey    *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s GetDataSourceOrderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigRequest) SetAgentKey(v string) *GetDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *GetDataSourceOrderConfigRequest) SetProductCode(v string) *GetDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

type GetDataSourceOrderConfigResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetDataSourceOrderConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBody) SetCode(v string) *GetDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetData(v *GetDataSourceOrderConfigResponseBodyData) *GetDataSourceOrderConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *GetDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetMessage(v string) *GetDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetRequestId(v string) *GetDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBody) SetSuccess(v bool) *GetDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

type GetDataSourceOrderConfigResponseBodyData struct {
	UserConfigDataSourceList []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s GetDataSourceOrderConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyData) SetUserConfigDataSourceList(v []*GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) *GetDataSourceOrderConfigResponseBodyData {
	s.UserConfigDataSourceList = v
	return s
}

type GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Number *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetCode(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetName(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetNumber(v int32) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList) SetType(v string) *GetDataSourceOrderConfigResponseBodyDataUserConfigDataSourceList {
	s.Type = &v
	return s
}

type GetDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceOrderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *GetDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetStatusCode(v int32) *GetDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceOrderConfigResponse) SetBody(v *GetDataSourceOrderConfigResponseBody) *GetDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

type GetGeneratedContentRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentRequest) SetAgentKey(v string) *GetGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *GetGeneratedContentRequest) SetId(v int64) *GetGeneratedContentRequest {
	s.Id = &v
	return s
}

type GetGeneratedContentResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetGeneratedContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBody) SetCode(v string) *GetGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetData(v *GetGeneratedContentResponseBodyData) *GetGeneratedContentResponseBody {
	s.Data = v
	return s
}

func (s *GetGeneratedContentResponseBody) SetHttpStatusCode(v int32) *GetGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetMessage(v string) *GetGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetRequestId(v string) *GetGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGeneratedContentResponseBody) SetSuccess(v bool) *GetGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type GetGeneratedContentResponseBodyData struct {
	Content       *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentDomain *string   `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string   `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	CreateTime    *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser    *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DeviceId      *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Id            *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordList   []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	Keywords      *string   `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt        *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TaskId        *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title         *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime    *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser    *string   `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	Uuid          *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetGeneratedContentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponseBodyData) SetContent(v string) *GetGeneratedContentResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentDomain(v string) *GetGeneratedContentResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetContentText(v string) *GetGeneratedContentResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateTime(v string) *GetGeneratedContentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetCreateUser(v string) *GetGeneratedContentResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetDeviceId(v string) *GetGeneratedContentResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetId(v int64) *GetGeneratedContentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywordList(v []*string) *GetGeneratedContentResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetKeywords(v string) *GetGeneratedContentResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetPrompt(v string) *GetGeneratedContentResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTaskId(v string) *GetGeneratedContentResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetTitle(v string) *GetGeneratedContentResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateTime(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUpdateUser(v string) *GetGeneratedContentResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetGeneratedContentResponseBodyData) SetUuid(v string) *GetGeneratedContentResponseBodyData {
	s.Uuid = &v
	return s
}

type GetGeneratedContentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *GetGeneratedContentResponse) SetHeaders(v map[string]*string) *GetGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *GetGeneratedContentResponse) SetStatusCode(v int32) *GetGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGeneratedContentResponse) SetBody(v *GetGeneratedContentResponseBody) *GetGeneratedContentResponse {
	s.Body = v
	return s
}

type GetInterveneGlobalReplyRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneGlobalReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyRequest) SetAgentKey(v string) *GetInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

type GetInterveneGlobalReplyResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBody) SetCode(v string) *GetInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetData(v *GetInterveneGlobalReplyResponseBodyData) *GetInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *GetInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetMessage(v string) *GetInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetRequestId(v string) *GetInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBody) SetSuccess(v bool) *GetInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

type GetInterveneGlobalReplyResponseBodyData struct {
	ReplyMessagList []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s GetInterveneGlobalReplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyData) SetReplyMessagList(v []*GetInterveneGlobalReplyResponseBodyDataReplyMessagList) *GetInterveneGlobalReplyResponseBodyData {
	s.ReplyMessagList = v
	return s
}

type GetInterveneGlobalReplyResponseBodyDataReplyMessagList struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponseBodyDataReplyMessagList) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetMessage(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.Message = &v
	return s
}

func (s *GetInterveneGlobalReplyResponseBodyDataReplyMessagList) SetReplyType(v string) *GetInterveneGlobalReplyResponseBodyDataReplyMessagList {
	s.ReplyType = &v
	return s
}

type GetInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneGlobalReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *GetInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetStatusCode(v int32) *GetInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneGlobalReplyResponse) SetBody(v *GetInterveneGlobalReplyResponseBody) *GetInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

type GetInterveneImportTaskInfoRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetInterveneImportTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoRequest) SetAgentKey(v string) *GetInterveneImportTaskInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneImportTaskInfoRequest) SetTaskId(v string) *GetInterveneImportTaskInfoRequest {
	s.TaskId = &v
	return s
}

type GetInterveneImportTaskInfoResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetInterveneImportTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBody) SetCode(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetData(v *GetInterveneImportTaskInfoResponseBodyData) *GetInterveneImportTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetHttpStatusCode(v int32) *GetInterveneImportTaskInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetMessage(v string) *GetInterveneImportTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetRequestId(v string) *GetInterveneImportTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBody) SetSuccess(v bool) *GetInterveneImportTaskInfoResponseBody {
	s.Success = &v
	return s
}

type GetInterveneImportTaskInfoResponseBodyData struct {
	Status *GetInterveneImportTaskInfoResponseBodyDataStatus `json:"Status,omitempty" xml:"Status,omitempty" type:"Struct"`
}

func (s GetInterveneImportTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyData) SetStatus(v *GetInterveneImportTaskInfoResponseBodyDataStatus) *GetInterveneImportTaskInfoResponseBodyData {
	s.Status = v
	return s
}

type GetInterveneImportTaskInfoResponseBodyDataStatus struct {
	Msg        *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponseBodyDataStatus) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetMsg(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Msg = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetPercentage(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Percentage = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetStatus(v int32) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.Status = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskId(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskId = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponseBodyDataStatus) SetTaskName(v string) *GetInterveneImportTaskInfoResponseBodyDataStatus {
	s.TaskName = &v
	return s
}

type GetInterveneImportTaskInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneImportTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneImportTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneImportTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneImportTaskInfoResponse) SetHeaders(v map[string]*string) *GetInterveneImportTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetStatusCode(v int32) *GetInterveneImportTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneImportTaskInfoResponse) SetBody(v *GetInterveneImportTaskInfoResponseBody) *GetInterveneImportTaskInfoResponse {
	s.Body = v
	return s
}

type GetInterveneRuleDetailRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	RuleId   *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetInterveneRuleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailRequest) SetAgentKey(v string) *GetInterveneRuleDetailRequest {
	s.AgentKey = &v
	return s
}

func (s *GetInterveneRuleDetailRequest) SetRuleId(v int64) *GetInterveneRuleDetailRequest {
	s.RuleId = &v
	return s
}

type GetInterveneRuleDetailResponseBody struct {
	Code           *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetInterveneRuleDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneRuleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBody) SetCode(v string) *GetInterveneRuleDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetData(v *GetInterveneRuleDetailResponseBodyData) *GetInterveneRuleDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetHttpStatusCode(v int32) *GetInterveneRuleDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetMessage(v string) *GetInterveneRuleDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetRequestId(v string) *GetInterveneRuleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBody) SetSuccess(v bool) *GetInterveneRuleDetailResponseBody {
	s.Success = &v
	return s
}

type GetInterveneRuleDetailResponseBodyData struct {
	InterveneRuleDetail *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail `json:"InterveneRuleDetail,omitempty" xml:"InterveneRuleDetail,omitempty" type:"Struct"`
}

func (s GetInterveneRuleDetailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyData) SetInterveneRuleDetail(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) *GetInterveneRuleDetailResponseBodyData {
	s.InterveneRuleDetail = v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail struct {
	AnswerConfig  []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig  *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig   `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	InterveneType *int32                                                                   `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string                                                                `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	RuleId        *int64                                                                   `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName      *string                                                                  `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetAnswerConfig(v []*GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.AnswerConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetEffectConfig(v *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.EffectConfig = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetInterveneType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.InterveneType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetNamespaceList(v []*string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.NamespaceList = v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleId(v int64) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleId = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail) SetRuleName(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetail {
	s.RuleName = &v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig struct {
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetAnswerType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetMessage(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Message = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig) SetNamespace(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailAnswerConfig {
	s.Namespace = &v
	return s
}

type GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig struct {
	EffectType *int32  `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEffectType(v int32) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EffectType = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetEndTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.EndTime = &v
	return s
}

func (s *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig) SetStartTime(v string) *GetInterveneRuleDetailResponseBodyDataInterveneRuleDetailEffectConfig {
	s.StartTime = &v
	return s
}

type GetInterveneRuleDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneRuleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneRuleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneRuleDetailResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneRuleDetailResponse) SetHeaders(v map[string]*string) *GetInterveneRuleDetailResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetStatusCode(v int32) *GetInterveneRuleDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneRuleDetailResponse) SetBody(v *GetInterveneRuleDetailResponseBody) *GetInterveneRuleDetailResponse {
	s.Body = v
	return s
}

type GetInterveneTemplateFileUrlRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetInterveneTemplateFileUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlRequest) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlRequest) SetAgentKey(v string) *GetInterveneTemplateFileUrlRequest {
	s.AgentKey = &v
	return s
}

type GetInterveneTemplateFileUrlResponseBody struct {
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetInterveneTemplateFileUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetCode(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetData(v *GetInterveneTemplateFileUrlResponseBodyData) *GetInterveneTemplateFileUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetHttpStatusCode(v int32) *GetInterveneTemplateFileUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetMessage(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetRequestId(v string) *GetInterveneTemplateFileUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponseBody) SetSuccess(v bool) *GetInterveneTemplateFileUrlResponseBody {
	s.Success = &v
	return s
}

type GetInterveneTemplateFileUrlResponseBodyData struct {
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponseBodyData) SetFileUrl(v string) *GetInterveneTemplateFileUrlResponseBodyData {
	s.FileUrl = &v
	return s
}

type GetInterveneTemplateFileUrlResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInterveneTemplateFileUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInterveneTemplateFileUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInterveneTemplateFileUrlResponse) GoString() string {
	return s.String()
}

func (s *GetInterveneTemplateFileUrlResponse) SetHeaders(v map[string]*string) *GetInterveneTemplateFileUrlResponse {
	s.Headers = v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetStatusCode(v int32) *GetInterveneTemplateFileUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInterveneTemplateFileUrlResponse) SetBody(v *GetInterveneTemplateFileUrlResponseBody) *GetInterveneTemplateFileUrlResponse {
	s.Body = v
	return s
}

type GetMaterialByIdRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Id       *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMaterialByIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdRequest) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdRequest) SetAgentKey(v string) *GetMaterialByIdRequest {
	s.AgentKey = &v
	return s
}

func (s *GetMaterialByIdRequest) SetId(v int64) *GetMaterialByIdRequest {
	s.Id = &v
	return s
}

type GetMaterialByIdResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetMaterialByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMaterialByIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBody) SetCode(v string) *GetMaterialByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetData(v *GetMaterialByIdResponseBodyData) *GetMaterialByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetMaterialByIdResponseBody) SetHttpStatusCode(v int32) *GetMaterialByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetMessage(v string) *GetMaterialByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetRequestId(v string) *GetMaterialByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMaterialByIdResponseBody) SetSuccess(v bool) *GetMaterialByIdResponseBody {
	s.Success = &v
	return s
}

type GetMaterialByIdResponseBodyData struct {
	Author            *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser        *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DocKeywords       []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	DocType           *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl       *string   `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent       *string   `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	Id                *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	PubTime           *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	PublicUrl         *string   `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	ShareAttr         *int32    `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom           *string   `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	ThumbnailInBase64 *string   `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	Title             *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime        *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser        *string   `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	Url               *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMaterialByIdResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponseBodyData) SetAuthor(v string) *GetMaterialByIdResponseBodyData {
	s.Author = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateTime(v string) *GetMaterialByIdResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetCreateUser(v string) *GetMaterialByIdResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocKeywords(v []*string) *GetMaterialByIdResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetDocType(v string) *GetMaterialByIdResponseBodyData {
	s.DocType = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetExternalUrl(v string) *GetMaterialByIdResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetHtmlContent(v string) *GetMaterialByIdResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetId(v int64) *GetMaterialByIdResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPubTime(v string) *GetMaterialByIdResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetPublicUrl(v string) *GetMaterialByIdResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetShareAttr(v int32) *GetMaterialByIdResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSrcFrom(v string) *GetMaterialByIdResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetSummary(v string) *GetMaterialByIdResponseBodyData {
	s.Summary = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTextContent(v string) *GetMaterialByIdResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetThumbnailInBase64(v string) *GetMaterialByIdResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetTitle(v string) *GetMaterialByIdResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateTime(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUpdateUser(v string) *GetMaterialByIdResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *GetMaterialByIdResponseBodyData) SetUrl(v string) *GetMaterialByIdResponseBodyData {
	s.Url = &v
	return s
}

type GetMaterialByIdResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMaterialByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMaterialByIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMaterialByIdResponse) GoString() string {
	return s.String()
}

func (s *GetMaterialByIdResponse) SetHeaders(v map[string]*string) *GetMaterialByIdResponse {
	s.Headers = v
	return s
}

func (s *GetMaterialByIdResponse) SetStatusCode(v int32) *GetMaterialByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMaterialByIdResponse) SetBody(v *GetMaterialByIdResponseBody) *GetMaterialByIdResponse {
	s.Body = v
	return s
}

type GetPropertiesRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s GetPropertiesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesRequest) GoString() string {
	return s.String()
}

func (s *GetPropertiesRequest) SetAgentKey(v string) *GetPropertiesRequest {
	s.AgentKey = &v
	return s
}

type GetPropertiesResponseBody struct {
	Code           *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetPropertiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                         `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPropertiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBody) SetCode(v string) *GetPropertiesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBody) SetData(v *GetPropertiesResponseBodyData) *GetPropertiesResponseBody {
	s.Data = v
	return s
}

func (s *GetPropertiesResponseBody) SetHttpStatusCode(v int32) *GetPropertiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetPropertiesResponseBody) SetMessage(v string) *GetPropertiesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPropertiesResponseBody) SetRequestId(v string) *GetPropertiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPropertiesResponseBody) SetSuccess(v bool) *GetPropertiesResponseBody {
	s.Success = &v
	return s
}

type GetPropertiesResponseBodyData struct {
	ChatConfig               map[string]interface{}                                   `json:"ChatConfig,omitempty" xml:"ChatConfig,omitempty"`
	ConsoleConfig            *GetPropertiesResponseBodyDataConsoleConfig              `json:"ConsoleConfig,omitempty" xml:"ConsoleConfig,omitempty" type:"Struct"`
	GeneralConfigMap         map[string]interface{}                                   `json:"GeneralConfigMap,omitempty" xml:"GeneralConfigMap,omitempty"`
	IntelligentSearchConfig  *GetPropertiesResponseBodyDataIntelligentSearchConfig    `json:"IntelligentSearchConfig,omitempty" xml:"IntelligentSearchConfig,omitempty" type:"Struct"`
	SearchSources            []*GetPropertiesResponseBodyDataSearchSources            `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
	SlrAuthorized            *bool                                                    `json:"SlrAuthorized,omitempty" xml:"SlrAuthorized,omitempty"`
	UserInfo                 *GetPropertiesResponseBodyDataUserInfo                   `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
	WanxiangImageSizeConfig  []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig  `json:"WanxiangImageSizeConfig,omitempty" xml:"WanxiangImageSizeConfig,omitempty" type:"Repeated"`
	WanxiangImageStyleConfig []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig `json:"WanxiangImageStyleConfig,omitempty" xml:"WanxiangImageStyleConfig,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyData) SetChatConfig(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.ChatConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetConsoleConfig(v *GetPropertiesResponseBodyDataConsoleConfig) *GetPropertiesResponseBodyData {
	s.ConsoleConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetGeneralConfigMap(v map[string]interface{}) *GetPropertiesResponseBodyData {
	s.GeneralConfigMap = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetIntelligentSearchConfig(v *GetPropertiesResponseBodyDataIntelligentSearchConfig) *GetPropertiesResponseBodyData {
	s.IntelligentSearchConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSearchSources(v []*GetPropertiesResponseBodyDataSearchSources) *GetPropertiesResponseBodyData {
	s.SearchSources = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetSlrAuthorized(v bool) *GetPropertiesResponseBodyData {
	s.SlrAuthorized = &v
	return s
}

func (s *GetPropertiesResponseBodyData) SetUserInfo(v *GetPropertiesResponseBodyDataUserInfo) *GetPropertiesResponseBodyData {
	s.UserInfo = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageSizeConfig(v []*GetPropertiesResponseBodyDataWanxiangImageSizeConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageSizeConfig = v
	return s
}

func (s *GetPropertiesResponseBodyData) SetWanxiangImageStyleConfig(v []*GetPropertiesResponseBodyDataWanxiangImageStyleConfig) *GetPropertiesResponseBodyData {
	s.WanxiangImageStyleConfig = v
	return s
}

type GetPropertiesResponseBodyDataConsoleConfig struct {
	TipContent *string `json:"TipContent,omitempty" xml:"TipContent,omitempty"`
	Title      *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetPropertiesResponseBodyDataConsoleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataConsoleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTipContent(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.TipContent = &v
	return s
}

func (s *GetPropertiesResponseBodyDataConsoleConfig) SetTitle(v string) *GetPropertiesResponseBodyDataConsoleConfig {
	s.Title = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfig struct {
	ProductDescription *string                                                              `json:"ProductDescription,omitempty" xml:"ProductDescription,omitempty"`
	SearchSamples      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples `json:"SearchSamples,omitempty" xml:"SearchSamples,omitempty" type:"Repeated"`
	SearchSources      []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetProductDescription(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.ProductDescription = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSamples(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSamples = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfig) SetSearchSources(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) *GetPropertiesResponseBodyDataIntelligentSearchConfig {
	s.SearchSources = v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples struct {
	Articles []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles `json:"Articles,omitempty" xml:"Articles,omitempty" type:"Repeated"`
	Prompt   *string                                                                      `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Text     *string                                                                      `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetArticles(v []*GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Articles = v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetPrompt(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Prompt = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples) SetText(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamples {
	s.Text = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles struct {
	Select *bool   `json:"Select,omitempty" xml:"Select,omitempty"`
	Stared *bool   `json:"Stared,omitempty" xml:"Stared,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetSelect(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Select = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetStared(v bool) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Stared = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetTitle(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Title = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles) SetUrl(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSamplesArticles {
	s.Url = &v
	return s
}

type GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetCode(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Code = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetDatasetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.DatasetName = &v
	return s
}

func (s *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources) SetName(v string) *GetPropertiesResponseBodyDataIntelligentSearchConfigSearchSources {
	s.Name = &v
	return s
}

type GetPropertiesResponseBodyDataSearchSources struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataSearchSources) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataSearchSources) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetLabel(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Label = &v
	return s
}

func (s *GetPropertiesResponseBodyDataSearchSources) SetValue(v string) *GetPropertiesResponseBodyDataSearchSources {
	s.Value = &v
	return s
}

type GetPropertiesResponseBodyDataUserInfo struct {
	AgentId  *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetPropertiesResponseBodyDataUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataUserInfo) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetAgentId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.AgentId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetTenantId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.TenantId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUserId(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.UserId = &v
	return s
}

func (s *GetPropertiesResponseBodyDataUserInfo) SetUsername(v string) *GetPropertiesResponseBodyDataUserInfo {
	s.Username = &v
	return s
}

type GetPropertiesResponseBodyDataWanxiangImageSizeConfig struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageSizeConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageSizeConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageSizeConfig {
	s.Value = &v
	return s
}

type GetPropertiesResponseBodyDataWanxiangImageStyleConfig struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pic   *string `json:"Pic,omitempty" xml:"Pic,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponseBodyDataWanxiangImageStyleConfig) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetName(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Name = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetPic(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Pic = &v
	return s
}

func (s *GetPropertiesResponseBodyDataWanxiangImageStyleConfig) SetValue(v string) *GetPropertiesResponseBodyDataWanxiangImageStyleConfig {
	s.Value = &v
	return s
}

type GetPropertiesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPropertiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPropertiesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPropertiesResponse) GoString() string {
	return s.String()
}

func (s *GetPropertiesResponse) SetHeaders(v map[string]*string) *GetPropertiesResponse {
	s.Headers = v
	return s
}

func (s *GetPropertiesResponse) SetStatusCode(v int32) *GetPropertiesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPropertiesResponse) SetBody(v *GetPropertiesResponseBody) *GetPropertiesResponse {
	s.Body = v
	return s
}

type ImportInterveneFileRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocName  *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileRequest) SetAgentKey(v string) *ImportInterveneFileRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetDocName(v string) *ImportInterveneFileRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileKey(v string) *ImportInterveneFileRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileRequest) SetFileUrl(v string) *ImportInterveneFileRequest {
	s.FileUrl = &v
	return s
}

type ImportInterveneFileResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ImportInterveneFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportInterveneFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBody) SetCode(v string) *ImportInterveneFileResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetData(v *ImportInterveneFileResponseBodyData) *ImportInterveneFileResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetMessage(v string) *ImportInterveneFileResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetRequestId(v string) *ImportInterveneFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileResponseBody) SetSuccess(v bool) *ImportInterveneFileResponseBody {
	s.Success = &v
	return s
}

type ImportInterveneFileResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	TaskId     *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileResponseBodyData) SetTaskId(v string) *ImportInterveneFileResponseBodyData {
	s.TaskId = &v
	return s
}

type ImportInterveneFileResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileResponse) SetHeaders(v map[string]*string) *ImportInterveneFileResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileResponse) SetStatusCode(v int32) *ImportInterveneFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileResponse) SetBody(v *ImportInterveneFileResponseBody) *ImportInterveneFileResponse {
	s.Body = v
	return s
}

type ImportInterveneFileAsyncRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	DocName  *string `json:"DocName,omitempty" xml:"DocName,omitempty"`
	FileKey  *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	FileUrl  *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
}

func (s ImportInterveneFileAsyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncRequest) SetAgentKey(v string) *ImportInterveneFileAsyncRequest {
	s.AgentKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetDocName(v string) *ImportInterveneFileAsyncRequest {
	s.DocName = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileKey(v string) *ImportInterveneFileAsyncRequest {
	s.FileKey = &v
	return s
}

func (s *ImportInterveneFileAsyncRequest) SetFileUrl(v string) *ImportInterveneFileAsyncRequest {
	s.FileUrl = &v
	return s
}

type ImportInterveneFileAsyncResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ImportInterveneFileAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBody) SetCode(v string) *ImportInterveneFileAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetData(v *ImportInterveneFileAsyncResponseBodyData) *ImportInterveneFileAsyncResponseBody {
	s.Data = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetHttpStatusCode(v int32) *ImportInterveneFileAsyncResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetMessage(v string) *ImportInterveneFileAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetRequestId(v string) *ImportInterveneFileAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportInterveneFileAsyncResponseBody) SetSuccess(v bool) *ImportInterveneFileAsyncResponseBody {
	s.Success = &v
	return s
}

type ImportInterveneFileAsyncResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	TaskId     *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ImportInterveneFileAsyncResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetFailIdList(v []*string) *ImportInterveneFileAsyncResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *ImportInterveneFileAsyncResponseBodyData) SetTaskId(v string) *ImportInterveneFileAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

type ImportInterveneFileAsyncResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportInterveneFileAsyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportInterveneFileAsyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportInterveneFileAsyncResponse) GoString() string {
	return s.String()
}

func (s *ImportInterveneFileAsyncResponse) SetHeaders(v map[string]*string) *ImportInterveneFileAsyncResponse {
	s.Headers = v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetStatusCode(v int32) *ImportInterveneFileAsyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportInterveneFileAsyncResponse) SetBody(v *ImportInterveneFileAsyncResponseBody) *ImportInterveneFileAsyncResponse {
	s.Body = v
	return s
}

type InsertInterveneGlobalReplyRequest struct {
	AgentKey        *string                                             `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReplyMessagList []*InsertInterveneGlobalReplyRequestReplyMessagList `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty" type:"Repeated"`
}

func (s InsertInterveneGlobalReplyRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequest) SetReplyMessagList(v []*InsertInterveneGlobalReplyRequestReplyMessagList) *InsertInterveneGlobalReplyRequest {
	s.ReplyMessagList = v
	return s
}

type InsertInterveneGlobalReplyRequestReplyMessagList struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ReplyType *string `json:"ReplyType,omitempty" xml:"ReplyType,omitempty"`
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyRequestReplyMessagList) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetMessage(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyRequestReplyMessagList) SetReplyType(v string) *InsertInterveneGlobalReplyRequestReplyMessagList {
	s.ReplyType = &v
	return s
}

type InsertInterveneGlobalReplyShrinkRequest struct {
	AgentKey              *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReplyMessagListShrink *string `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty"`
}

func (s InsertInterveneGlobalReplyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetReplyMessagListShrink(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.ReplyMessagListShrink = &v
	return s
}

type InsertInterveneGlobalReplyResponseBody struct {
	Code           *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *InsertInterveneGlobalReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBody) SetCode(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetData(v *InsertInterveneGlobalReplyResponseBodyData) *InsertInterveneGlobalReplyResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetHttpStatusCode(v int32) *InsertInterveneGlobalReplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetMessage(v string) *InsertInterveneGlobalReplyResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetRequestId(v string) *InsertInterveneGlobalReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBody) SetSuccess(v bool) *InsertInterveneGlobalReplyResponseBody {
	s.Success = &v
	return s
}

type InsertInterveneGlobalReplyResponseBodyData struct {
	FailIdList []*string `json:"FailIdList,omitempty" xml:"FailIdList,omitempty" type:"Repeated"`
	TaskId     *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InsertInterveneGlobalReplyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetFailIdList(v []*string) *InsertInterveneGlobalReplyResponseBodyData {
	s.FailIdList = v
	return s
}

func (s *InsertInterveneGlobalReplyResponseBodyData) SetTaskId(v string) *InsertInterveneGlobalReplyResponseBodyData {
	s.TaskId = &v
	return s
}

type InsertInterveneGlobalReplyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneGlobalReplyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneGlobalReplyResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneGlobalReplyResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyResponse) SetHeaders(v map[string]*string) *InsertInterveneGlobalReplyResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetStatusCode(v int32) *InsertInterveneGlobalReplyResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneGlobalReplyResponse) SetBody(v *InsertInterveneGlobalReplyResponseBody) *InsertInterveneGlobalReplyResponse {
	s.Body = v
	return s
}

type InsertInterveneRuleRequest struct {
	AgentKey            *string                                        `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfig *InsertInterveneRuleRequestInterveneRuleConfig `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty" type:"Struct"`
}

func (s InsertInterveneRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequest) SetAgentKey(v string) *InsertInterveneRuleRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleRequest) SetInterveneRuleConfig(v *InsertInterveneRuleRequestInterveneRuleConfig) *InsertInterveneRuleRequest {
	s.InterveneRuleConfig = v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfig struct {
	AnswerConfig        []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig        `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	EffectConfig        *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig          `json:"EffectConfig,omitempty" xml:"EffectConfig,omitempty" type:"Struct"`
	InterveneConfigList []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList `json:"InterveneConfigList,omitempty" xml:"InterveneConfigList,omitempty" type:"Repeated"`
	InterveneType       *int32                                                              `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList       []*string                                                           `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	RuleId              *int64                                                              `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName            *string                                                             `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetAnswerConfig(v []*InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.AnswerConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetEffectConfig(v *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.EffectConfig = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneConfigList(v []*InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneConfigList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetInterveneType(v int32) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.InterveneType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetNamespaceList(v []*string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.NamespaceList = v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleId(v int64) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleId = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfig) SetRuleName(v string) *InsertInterveneRuleRequestInterveneRuleConfig {
	s.RuleName = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig struct {
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetAnswerType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetMessage(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig) SetNamespace(v string) *InsertInterveneRuleRequestInterveneRuleConfigAnswerConfig {
	s.Namespace = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigEffectConfig struct {
	EffectType *int32  `json:"EffectType,omitempty" xml:"EffectType,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEffectType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EffectType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetEndTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.EndTime = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig) SetStartTime(v string) *InsertInterveneRuleRequestInterveneRuleConfigEffectConfig {
	s.StartTime = &v
	return s
}

type InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList struct {
	// id
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	OperationType *int32  `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetId(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Id = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetOperationType(v int32) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.OperationType = &v
	return s
}

func (s *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList) SetQuery(v string) *InsertInterveneRuleRequestInterveneRuleConfigInterveneConfigList {
	s.Query = &v
	return s
}

type InsertInterveneRuleShrinkRequest struct {
	AgentKey                  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneRuleConfigShrink *string `json:"InterveneRuleConfig,omitempty" xml:"InterveneRuleConfig,omitempty"`
}

func (s InsertInterveneRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleShrinkRequest) SetAgentKey(v string) *InsertInterveneRuleShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneRuleShrinkRequest) SetInterveneRuleConfigShrink(v string) *InsertInterveneRuleShrinkRequest {
	s.InterveneRuleConfigShrink = &v
	return s
}

type InsertInterveneRuleResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *InsertInterveneRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InsertInterveneRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBody) SetCode(v string) *InsertInterveneRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetData(v *InsertInterveneRuleResponseBodyData) *InsertInterveneRuleResponseBody {
	s.Data = v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetHttpStatusCode(v int32) *InsertInterveneRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetMessage(v string) *InsertInterveneRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetRequestId(v string) *InsertInterveneRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertInterveneRuleResponseBody) SetSuccess(v bool) *InsertInterveneRuleResponseBody {
	s.Success = &v
	return s
}

type InsertInterveneRuleResponseBodyData struct {
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s InsertInterveneRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponseBodyData) SetRuleId(v int64) *InsertInterveneRuleResponseBodyData {
	s.RuleId = &v
	return s
}

type InsertInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertInterveneRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s InsertInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *InsertInterveneRuleResponse) SetHeaders(v map[string]*string) *InsertInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *InsertInterveneRuleResponse) SetStatusCode(v int32) *InsertInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertInterveneRuleResponse) SetBody(v *InsertInterveneRuleResponseBody) *InsertInterveneRuleResponse {
	s.Body = v
	return s
}

type ListAsyncTasksRequest struct {
	AgentKey        *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CreateTimeEnd   *string   `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart *string   `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Current         *int32    `json:"Current,omitempty" xml:"Current,omitempty"`
	Size            *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	TaskCode        *string   `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskName        *string   `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus      *int32    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusList  []*int32  `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty" type:"Repeated"`
	TaskType        *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeList    []*string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty" type:"Repeated"`
}

func (s ListAsyncTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) SetAgentKey(v string) *ListAsyncTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeEnd(v string) *ListAsyncTasksRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeStart(v string) *ListAsyncTasksRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCurrent(v int32) *ListAsyncTasksRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksRequest) SetSize(v int32) *ListAsyncTasksRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskCode(v string) *ListAsyncTasksRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskName(v string) *ListAsyncTasksRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatus(v int32) *ListAsyncTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatusList(v []*int32) *ListAsyncTasksRequest {
	s.TaskStatusList = v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskType(v string) *ListAsyncTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskTypeList(v []*string) *ListAsyncTasksRequest {
	s.TaskTypeList = v
	return s
}

type ListAsyncTasksShrinkRequest struct {
	AgentKey             *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	CreateTimeEnd        *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart      *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Current              *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	Size                 *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	TaskCode             *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskName             *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus           *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusListShrink *string `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeListShrink   *string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty"`
}

func (s ListAsyncTasksShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksShrinkRequest) SetAgentKey(v string) *ListAsyncTasksShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeEnd(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeStart(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCurrent(v int32) *ListAsyncTasksShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetSize(v int32) *ListAsyncTasksShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskCode(v string) *ListAsyncTasksShrinkRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskName(v string) *ListAsyncTasksShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatus(v int32) *ListAsyncTasksShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatusListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskStatusListShrink = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskType(v string) *ListAsyncTasksShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskTypeListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskTypeListShrink = &v
	return s
}

type ListAsyncTasksResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Current        *int32                            `json:"Current,omitempty" xml:"Current,omitempty"`
	Data           []*ListAsyncTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int32                            `json:"Size,omitempty" xml:"Size,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                            `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAsyncTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBody) SetCode(v string) *ListAsyncTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetCurrent(v int32) *ListAsyncTasksResponseBody {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetData(v []*ListAsyncTasksResponseBodyData) *ListAsyncTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListAsyncTasksResponseBody) SetHttpStatusCode(v int32) *ListAsyncTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetMessage(v string) *ListAsyncTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetRequestId(v string) *ListAsyncTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSize(v int32) *ListAsyncTasksResponseBody {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetSuccess(v bool) *ListAsyncTasksResponseBody {
	s.Success = &v
	return s
}

func (s *ListAsyncTasksResponseBody) SetTotal(v int32) *ListAsyncTasksResponseBody {
	s.Total = &v
	return s
}

type ListAsyncTasksResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser             *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Id                     *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TaskCode               *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskDefinition         *string `json:"TaskDefinition,omitempty" xml:"TaskDefinition,omitempty"`
	TaskEndTime            *string `json:"TaskEndTime,omitempty" xml:"TaskEndTime,omitempty"`
	TaskErrorMessage       *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	TaskExecuteTime        *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskInnerErrorMessage  *string `json:"TaskInnerErrorMessage,omitempty" xml:"TaskInnerErrorMessage,omitempty"`
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskParam              *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	TaskProgressMessage    *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	TaskResult             *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskRetryCount         *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	TaskStartTime          *string `json:"TaskStartTime,omitempty" xml:"TaskStartTime,omitempty"`
	TaskStatus             *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType               *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	UpdateTime             *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser             *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListAsyncTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponseBodyData) SetCreateTime(v string) *ListAsyncTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetCreateUser(v string) *ListAsyncTasksResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetId(v int64) *ListAsyncTasksResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskCode(v string) *ListAsyncTasksResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskDefinition(v string) *ListAsyncTasksResponseBodyData {
	s.TaskDefinition = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskEndTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskEndTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskExecuteTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskExecuteTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskId(v string) *ListAsyncTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskInnerErrorMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskInnerErrorMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskIntermediateResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskName(v string) *ListAsyncTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskParam(v string) *ListAsyncTasksResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskProgressMessage(v string) *ListAsyncTasksResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskResult(v string) *ListAsyncTasksResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskRetryCount(v string) *ListAsyncTasksResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStartTime(v string) *ListAsyncTasksResponseBodyData {
	s.TaskStartTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskStatus(v int32) *ListAsyncTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetTaskType(v string) *ListAsyncTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateTime(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListAsyncTasksResponseBodyData) SetUpdateUser(v string) *ListAsyncTasksResponseBodyData {
	s.UpdateUser = &v
	return s
}

type ListAsyncTasksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAsyncTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAsyncTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAsyncTasksResponse) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksResponse) SetHeaders(v map[string]*string) *ListAsyncTasksResponse {
	s.Headers = v
	return s
}

func (s *ListAsyncTasksResponse) SetStatusCode(v int32) *ListAsyncTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAsyncTasksResponse) SetBody(v *ListAsyncTasksResponseBody) *ListAsyncTasksResponse {
	s.Body = v
	return s
}

type ListBuildConfigsRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Type     *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBuildConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsRequest) SetAgentKey(v string) *ListBuildConfigsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListBuildConfigsRequest) SetType(v string) *ListBuildConfigsRequest {
	s.Type = &v
	return s
}

type ListBuildConfigsResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListBuildConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBuildConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBody) SetCode(v string) *ListBuildConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetData(v []*ListBuildConfigsResponseBodyData) *ListBuildConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListBuildConfigsResponseBody) SetHttpStatusCode(v int32) *ListBuildConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetMessage(v string) *ListBuildConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetRequestId(v string) *ListBuildConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBuildConfigsResponseBody) SetSuccess(v bool) *ListBuildConfigsResponseBody {
	s.Success = &v
	return s
}

type ListBuildConfigsResponseBodyData struct {
	BuildIn        *bool                                       `json:"BuildIn,omitempty" xml:"BuildIn,omitempty"`
	CreateTime     *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser     *string                                     `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Id             *int64                                      `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords       []*ListBuildConfigsResponseBodyDataKeywords `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Tag            *string                                     `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TagDescription *string                                     `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	Type           *string                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime     *string                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser     *string                                     `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s ListBuildConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyData) SetBuildIn(v bool) *ListBuildConfigsResponseBodyData {
	s.BuildIn = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateTime(v string) *ListBuildConfigsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetCreateUser(v string) *ListBuildConfigsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetId(v int64) *ListBuildConfigsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetKeywords(v []*ListBuildConfigsResponseBodyDataKeywords) *ListBuildConfigsResponseBodyData {
	s.Keywords = v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTag(v string) *ListBuildConfigsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetTagDescription(v string) *ListBuildConfigsResponseBodyData {
	s.TagDescription = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetType(v string) *ListBuildConfigsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateTime(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListBuildConfigsResponseBodyData) SetUpdateUser(v string) *ListBuildConfigsResponseBodyData {
	s.UpdateUser = &v
	return s
}

type ListBuildConfigsResponseBodyDataKeywords struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Key         *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListBuildConfigsResponseBodyDataKeywords) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponseBodyDataKeywords) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetDescription(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Description = &v
	return s
}

func (s *ListBuildConfigsResponseBodyDataKeywords) SetKey(v string) *ListBuildConfigsResponseBodyDataKeywords {
	s.Key = &v
	return s
}

type ListBuildConfigsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBuildConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBuildConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBuildConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListBuildConfigsResponse) SetHeaders(v map[string]*string) *ListBuildConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListBuildConfigsResponse) SetStatusCode(v int32) *ListBuildConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBuildConfigsResponse) SetBody(v *ListBuildConfigsResponseBody) *ListBuildConfigsResponse {
	s.Body = v
	return s
}

type ListDialoguesRequest struct {
	AgentKey     *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Current      *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	DialogueType *int32  `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Size         *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListDialoguesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesRequest) GoString() string {
	return s.String()
}

func (s *ListDialoguesRequest) SetAgentKey(v string) *ListDialoguesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListDialoguesRequest) SetCurrent(v int32) *ListDialoguesRequest {
	s.Current = &v
	return s
}

func (s *ListDialoguesRequest) SetDialogueType(v int32) *ListDialoguesRequest {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesRequest) SetEndTime(v string) *ListDialoguesRequest {
	s.EndTime = &v
	return s
}

func (s *ListDialoguesRequest) SetSize(v int32) *ListDialoguesRequest {
	s.Size = &v
	return s
}

func (s *ListDialoguesRequest) SetStartTime(v string) *ListDialoguesRequest {
	s.StartTime = &v
	return s
}

func (s *ListDialoguesRequest) SetTaskId(v string) *ListDialoguesRequest {
	s.TaskId = &v
	return s
}

type ListDialoguesResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Current        *int32                           `json:"Current,omitempty" xml:"Current,omitempty"`
	Data           []*ListDialoguesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int32                           `json:"Size,omitempty" xml:"Size,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                           `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDialoguesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBody) SetCode(v string) *ListDialoguesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDialoguesResponseBody) SetCurrent(v int32) *ListDialoguesResponseBody {
	s.Current = &v
	return s
}

func (s *ListDialoguesResponseBody) SetData(v []*ListDialoguesResponseBodyData) *ListDialoguesResponseBody {
	s.Data = v
	return s
}

func (s *ListDialoguesResponseBody) SetHttpStatusCode(v int32) *ListDialoguesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDialoguesResponseBody) SetMessage(v string) *ListDialoguesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDialoguesResponseBody) SetRequestId(v string) *ListDialoguesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSize(v int32) *ListDialoguesResponseBody {
	s.Size = &v
	return s
}

func (s *ListDialoguesResponseBody) SetSuccess(v bool) *ListDialoguesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDialoguesResponseBody) SetTotal(v int32) *ListDialoguesResponseBody {
	s.Total = &v
	return s
}

type ListDialoguesResponseBodyData struct {
	Bot          *string `json:"Bot,omitempty" xml:"Bot,omitempty"`
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser   *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DialogueType *int32  `json:"DialogueType,omitempty" xml:"DialogueType,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	User         *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s ListDialoguesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponseBodyData) SetBot(v string) *ListDialoguesResponseBodyData {
	s.Bot = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateTime(v string) *ListDialoguesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetCreateUser(v string) *ListDialoguesResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetDialogueType(v int32) *ListDialoguesResponseBodyData {
	s.DialogueType = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetTaskId(v string) *ListDialoguesResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListDialoguesResponseBodyData) SetUser(v string) *ListDialoguesResponseBodyData {
	s.User = &v
	return s
}

type ListDialoguesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDialoguesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDialoguesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDialoguesResponse) GoString() string {
	return s.String()
}

func (s *ListDialoguesResponse) SetHeaders(v map[string]*string) *ListDialoguesResponse {
	s.Headers = v
	return s
}

func (s *ListDialoguesResponse) SetStatusCode(v int32) *ListDialoguesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDialoguesResponse) SetBody(v *ListDialoguesResponseBody) *ListDialoguesResponse {
	s.Body = v
	return s
}

type ListGeneratedContentsRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ContentDomain *string `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	Current       *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Size          *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListGeneratedContentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsRequest) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsRequest) SetAgentKey(v string) *ListGeneratedContentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetContentDomain(v string) *ListGeneratedContentsRequest {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetCurrent(v int32) *ListGeneratedContentsRequest {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetEndTime(v string) *ListGeneratedContentsRequest {
	s.EndTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetSize(v int32) *ListGeneratedContentsRequest {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetStartTime(v string) *ListGeneratedContentsRequest {
	s.StartTime = &v
	return s
}

func (s *ListGeneratedContentsRequest) SetTitle(v string) *ListGeneratedContentsRequest {
	s.Title = &v
	return s
}

type ListGeneratedContentsResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Current        *int32                                   `json:"Current,omitempty" xml:"Current,omitempty"`
	Data           []*ListGeneratedContentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int32                                   `json:"Size,omitempty" xml:"Size,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListGeneratedContentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBody) SetCode(v string) *ListGeneratedContentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetCurrent(v int32) *ListGeneratedContentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetData(v []*ListGeneratedContentsResponseBodyData) *ListGeneratedContentsResponseBody {
	s.Data = v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetHttpStatusCode(v int32) *ListGeneratedContentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetMessage(v string) *ListGeneratedContentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetRequestId(v string) *ListGeneratedContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSize(v int32) *ListGeneratedContentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetSuccess(v bool) *ListGeneratedContentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGeneratedContentsResponseBody) SetTotal(v int32) *ListGeneratedContentsResponseBody {
	s.Total = &v
	return s
}

type ListGeneratedContentsResponseBodyData struct {
	Content       *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentDomain *string   `json:"ContentDomain,omitempty" xml:"ContentDomain,omitempty"`
	ContentText   *string   `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	CreateTime    *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser    *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	DeviceId      *string   `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	Id            *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordList   []*string `json:"KeywordList,omitempty" xml:"KeywordList,omitempty" type:"Repeated"`
	Keywords      *string   `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt        *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TaskId        *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Title         *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime    *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser    *string   `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	Uuid          *string   `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListGeneratedContentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponseBodyData) SetContent(v string) *ListGeneratedContentsResponseBodyData {
	s.Content = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentDomain(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentDomain = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetContentText(v string) *ListGeneratedContentsResponseBodyData {
	s.ContentText = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetCreateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetDeviceId(v string) *ListGeneratedContentsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetId(v int64) *ListGeneratedContentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywordList(v []*string) *ListGeneratedContentsResponseBodyData {
	s.KeywordList = v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetKeywords(v string) *ListGeneratedContentsResponseBodyData {
	s.Keywords = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetPrompt(v string) *ListGeneratedContentsResponseBodyData {
	s.Prompt = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTaskId(v string) *ListGeneratedContentsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetTitle(v string) *ListGeneratedContentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateTime(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUpdateUser(v string) *ListGeneratedContentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListGeneratedContentsResponseBodyData) SetUuid(v string) *ListGeneratedContentsResponseBodyData {
	s.Uuid = &v
	return s
}

type ListGeneratedContentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGeneratedContentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGeneratedContentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGeneratedContentsResponse) GoString() string {
	return s.String()
}

func (s *ListGeneratedContentsResponse) SetHeaders(v map[string]*string) *ListGeneratedContentsResponse {
	s.Headers = v
	return s
}

func (s *ListGeneratedContentsResponse) SetStatusCode(v int32) *ListGeneratedContentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGeneratedContentsResponse) SetBody(v *ListGeneratedContentsResponseBody) *ListGeneratedContentsResponse {
	s.Body = v
	return s
}

type ListHotNewsWithTypeRequest struct {
	AgentKey  *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Current   *int32    `json:"Current,omitempty" xml:"Current,omitempty"`
	NewsType  *string   `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypes []*string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty" type:"Repeated"`
	Size      *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeRequest) SetAgentKey(v string) *ListHotNewsWithTypeRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetCurrent(v int32) *ListHotNewsWithTypeRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsType(v string) *ListHotNewsWithTypeRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetNewsTypes(v []*string) *ListHotNewsWithTypeRequest {
	s.NewsTypes = v
	return s
}

func (s *ListHotNewsWithTypeRequest) SetSize(v int32) *ListHotNewsWithTypeRequest {
	s.Size = &v
	return s
}

type ListHotNewsWithTypeShrinkRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Current         *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	NewsType        *string `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypesShrink *string `json:"NewsTypes,omitempty" xml:"NewsTypes,omitempty"`
	Size            *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListHotNewsWithTypeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeShrinkRequest) SetAgentKey(v string) *ListHotNewsWithTypeShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetCurrent(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsType(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetNewsTypesShrink(v string) *ListHotNewsWithTypeShrinkRequest {
	s.NewsTypesShrink = &v
	return s
}

func (s *ListHotNewsWithTypeShrinkRequest) SetSize(v int32) *ListHotNewsWithTypeShrinkRequest {
	s.Size = &v
	return s
}

type ListHotNewsWithTypeResponseBody struct {
	Code           *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListHotNewsWithTypeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListHotNewsWithTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBody) SetCode(v string) *ListHotNewsWithTypeResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetData(v []*ListHotNewsWithTypeResponseBodyData) *ListHotNewsWithTypeResponseBody {
	s.Data = v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetHttpStatusCode(v int32) *ListHotNewsWithTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetMessage(v string) *ListHotNewsWithTypeResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetRequestId(v string) *ListHotNewsWithTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBody) SetSuccess(v bool) *ListHotNewsWithTypeResponseBody {
	s.Success = &v
	return s
}

type ListHotNewsWithTypeResponseBodyData struct {
	News         []*ListHotNewsWithTypeResponseBodyDataNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
	NewsType     *string                                    `json:"NewsType,omitempty" xml:"NewsType,omitempty"`
	NewsTypeName *string                                    `json:"NewsTypeName,omitempty" xml:"NewsTypeName,omitempty"`
	TotalPages   *int32                                     `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNews(v []*ListHotNewsWithTypeResponseBodyDataNews) *ListHotNewsWithTypeResponseBodyData {
	s.News = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsType(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsType = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetNewsTypeName(v string) *ListHotNewsWithTypeResponseBodyData {
	s.NewsTypeName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyData) SetTotalPages(v int32) *ListHotNewsWithTypeResponseBodyData {
	s.TotalPages = &v
	return s
}

type ListHotNewsWithTypeResponseBodyDataNews struct {
	Author           *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	Content          *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	DocUuid          *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls        []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	PubTime          *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string   `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string   `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Source           *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary          *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag              *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime       *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Url              *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListHotNewsWithTypeResponseBodyDataNews) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetAuthor(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Author = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetContent(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetDocUuid(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.DocUuid = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetImageUrls(v []*string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetPubTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSource = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSearchSourceName(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.SearchSourceName = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSource(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetSummary(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Summary = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTag(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Tag = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetTitle(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUpdateTime(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.UpdateTime = &v
	return s
}

func (s *ListHotNewsWithTypeResponseBodyDataNews) SetUrl(v string) *ListHotNewsWithTypeResponseBodyDataNews {
	s.Url = &v
	return s
}

type ListHotNewsWithTypeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotNewsWithTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotNewsWithTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListHotNewsWithTypeResponse) GoString() string {
	return s.String()
}

func (s *ListHotNewsWithTypeResponse) SetHeaders(v map[string]*string) *ListHotNewsWithTypeResponse {
	s.Headers = v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetStatusCode(v int32) *ListHotNewsWithTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotNewsWithTypeResponse) SetBody(v *ListHotNewsWithTypeResponseBody) *ListHotNewsWithTypeResponse {
	s.Body = v
	return s
}

type ListInterveneCntRequest struct {
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneCntRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneCntRequest) SetAgentKey(v string) *ListInterveneCntRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageIndex(v int32) *ListInterveneCntRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntRequest) SetPageSize(v int32) *ListInterveneCntRequest {
	s.PageSize = &v
	return s
}

type ListInterveneCntResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListInterveneCntResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneCntResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBody) SetCode(v string) *ListInterveneCntResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetData(v *ListInterveneCntResponseBodyData) *ListInterveneCntResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneCntResponseBody) SetHttpStatusCode(v int32) *ListInterveneCntResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetMessage(v string) *ListInterveneCntResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetRequestId(v string) *ListInterveneCntResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneCntResponseBody) SetSuccess(v bool) *ListInterveneCntResponseBody {
	s.Success = &v
	return s
}

type ListInterveneCntResponseBodyData struct {
	CntList   []interface{} `json:"CntList,omitempty" xml:"CntList,omitempty" type:"Repeated"`
	PageCnt   *int32        `json:"PageCnt,omitempty" xml:"PageCnt,omitempty"`
	PageIndex *int32        `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneCntResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponseBodyData) SetCntList(v []interface{}) *ListInterveneCntResponseBodyData {
	s.CntList = v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageCnt(v int32) *ListInterveneCntResponseBodyData {
	s.PageCnt = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageIndex(v int32) *ListInterveneCntResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneCntResponseBodyData) SetPageSize(v int32) *ListInterveneCntResponseBodyData {
	s.PageSize = &v
	return s
}

type ListInterveneCntResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneCntResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneCntResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneCntResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneCntResponse) SetHeaders(v map[string]*string) *ListInterveneCntResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneCntResponse) SetStatusCode(v int32) *ListInterveneCntResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneCntResponse) SetBody(v *ListInterveneCntResponseBody) *ListInterveneCntResponse {
	s.Body = v
	return s
}

type ListInterveneImportTasksRequest struct {
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneImportTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksRequest) SetAgentKey(v string) *ListInterveneImportTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageIndex(v int32) *ListInterveneImportTasksRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksRequest) SetPageSize(v int32) *ListInterveneImportTasksRequest {
	s.PageSize = &v
	return s
}

type ListInterveneImportTasksResponseBody struct {
	Code           *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListInterveneImportTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneImportTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBody) SetCode(v string) *ListInterveneImportTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetData(v *ListInterveneImportTasksResponseBodyData) *ListInterveneImportTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetHttpStatusCode(v int32) *ListInterveneImportTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetMessage(v string) *ListInterveneImportTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetRequestId(v string) *ListInterveneImportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBody) SetSuccess(v bool) *ListInterveneImportTasksResponseBody {
	s.Success = &v
	return s
}

type ListInterveneImportTasksResponseBodyData struct {
	PageIndex  *int32                                                `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize   *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StatusList []*ListInterveneImportTasksResponseBodyDataStatusList `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	TotalSize  *int32                                                `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageIndex(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetPageSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetStatusList(v []*ListInterveneImportTasksResponseBodyDataStatusList) *ListInterveneImportTasksResponseBodyData {
	s.StatusList = v
	return s
}

func (s *ListInterveneImportTasksResponseBodyData) SetTotalSize(v int32) *ListInterveneImportTasksResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListInterveneImportTasksResponseBodyDataStatusList struct {
	Msg        *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Percentage *int32  `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponseBodyDataStatusList) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetMsg(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Msg = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetPercentage(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Percentage = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetStatus(v int32) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.Status = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskId(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskId = &v
	return s
}

func (s *ListInterveneImportTasksResponseBodyDataStatusList) SetTaskName(v string) *ListInterveneImportTasksResponseBodyDataStatusList {
	s.TaskName = &v
	return s
}

type ListInterveneImportTasksResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneImportTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneImportTasksResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneImportTasksResponse) SetHeaders(v map[string]*string) *ListInterveneImportTasksResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneImportTasksResponse) SetStatusCode(v int32) *ListInterveneImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneImportTasksResponse) SetBody(v *ListInterveneImportTasksResponseBody) *ListInterveneImportTasksResponse {
	s.Body = v
	return s
}

type ListInterveneRulesRequest struct {
	AgentKey  *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	PageIndex *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesRequest) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesRequest) SetAgentKey(v string) *ListInterveneRulesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageIndex(v int32) *ListInterveneRulesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesRequest) SetPageSize(v int32) *ListInterveneRulesRequest {
	s.PageSize = &v
	return s
}

type ListInterveneRulesResponseBody struct {
	Code           *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListInterveneRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInterveneRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBody) SetCode(v string) *ListInterveneRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetData(v *ListInterveneRulesResponseBodyData) *ListInterveneRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListInterveneRulesResponseBody) SetHttpStatusCode(v int32) *ListInterveneRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetMessage(v string) *ListInterveneRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetRequestId(v string) *ListInterveneRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterveneRulesResponseBody) SetSuccess(v bool) *ListInterveneRulesResponseBody {
	s.Success = &v
	return s
}

type ListInterveneRulesResponseBodyData struct {
	Count             *int64                                                 `json:"Count,omitempty" xml:"Count,omitempty"`
	InterveneRuleList []*ListInterveneRulesResponseBodyDataInterveneRuleList `json:"InterveneRuleList,omitempty" xml:"InterveneRuleList,omitempty" type:"Repeated"`
	PageIndex         *int32                                                 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize          *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListInterveneRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyData) SetCount(v int64) *ListInterveneRulesResponseBodyData {
	s.Count = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetInterveneRuleList(v []*ListInterveneRulesResponseBodyDataInterveneRuleList) *ListInterveneRulesResponseBodyData {
	s.InterveneRuleList = v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageIndex(v int32) *ListInterveneRulesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListInterveneRulesResponseBodyData) SetPageSize(v int32) *ListInterveneRulesResponseBodyData {
	s.PageSize = &v
	return s
}

type ListInterveneRulesResponseBodyDataInterveneRuleList struct {
	AnswerConfig  []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig `json:"AnswerConfig,omitempty" xml:"AnswerConfig,omitempty" type:"Repeated"`
	CreateTime    *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EffectTime    *string                                                            `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	InterveneType *int32                                                             `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	NamespaceList []*string                                                          `json:"NamespaceList,omitempty" xml:"NamespaceList,omitempty" type:"Repeated"`
	RuleId        *int64                                                             `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName      *string                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleList) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetAnswerConfig(v []*ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.AnswerConfig = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetCreateTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.CreateTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetEffectTime(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.EffectTime = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetInterveneType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.InterveneType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetNamespaceList(v []*string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.NamespaceList = v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleId(v int64) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleId = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleList) SetRuleName(v string) *ListInterveneRulesResponseBodyDataInterveneRuleList {
	s.RuleName = &v
	return s
}

type ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig struct {
	AnswerType *int32  `json:"AnswerType,omitempty" xml:"AnswerType,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Namespace  *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetAnswerType(v int32) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.AnswerType = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetMessage(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Message = &v
	return s
}

func (s *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig) SetNamespace(v string) *ListInterveneRulesResponseBodyDataInterveneRuleListAnswerConfig {
	s.Namespace = &v
	return s
}

type ListInterveneRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterveneRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterveneRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterveneRulesResponse) GoString() string {
	return s.String()
}

func (s *ListInterveneRulesResponse) SetHeaders(v map[string]*string) *ListInterveneRulesResponse {
	s.Headers = v
	return s
}

func (s *ListInterveneRulesResponse) SetStatusCode(v int32) *ListInterveneRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterveneRulesResponse) SetBody(v *ListInterveneRulesResponseBody) *ListInterveneRulesResponse {
	s.Body = v
	return s
}

type ListIntervenesRequest struct {
	AgentKey      *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	InterveneType *int32  `json:"InterveneType,omitempty" xml:"InterveneType,omitempty"`
	PageIndex     *int32  `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query         *string `json:"Query,omitempty" xml:"Query,omitempty"`
	RuleId        *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListIntervenesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesRequest) GoString() string {
	return s.String()
}

func (s *ListIntervenesRequest) SetAgentKey(v string) *ListIntervenesRequest {
	s.AgentKey = &v
	return s
}

func (s *ListIntervenesRequest) SetInterveneType(v int32) *ListIntervenesRequest {
	s.InterveneType = &v
	return s
}

func (s *ListIntervenesRequest) SetPageIndex(v int32) *ListIntervenesRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesRequest) SetPageSize(v int32) *ListIntervenesRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesRequest) SetQuery(v string) *ListIntervenesRequest {
	s.Query = &v
	return s
}

func (s *ListIntervenesRequest) SetRuleId(v int64) *ListIntervenesRequest {
	s.RuleId = &v
	return s
}

type ListIntervenesResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ListIntervenesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntervenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBody) SetCode(v string) *ListIntervenesResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntervenesResponseBody) SetData(v *ListIntervenesResponseBodyData) *ListIntervenesResponseBody {
	s.Data = v
	return s
}

func (s *ListIntervenesResponseBody) SetHttpStatusCode(v int32) *ListIntervenesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntervenesResponseBody) SetMessage(v string) *ListIntervenesResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntervenesResponseBody) SetRequestId(v string) *ListIntervenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntervenesResponseBody) SetSuccess(v bool) *ListIntervenesResponseBody {
	s.Success = &v
	return s
}

type ListIntervenesResponseBodyData struct {
	InterveneList []*ListIntervenesResponseBodyDataInterveneList `json:"InterveneList,omitempty" xml:"InterveneList,omitempty" type:"Repeated"`
	PageIndex     *int32                                         `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	PageSize      *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalSize     *int64                                         `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListIntervenesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyData) SetInterveneList(v []*ListIntervenesResponseBodyDataInterveneList) *ListIntervenesResponseBodyData {
	s.InterveneList = v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageIndex(v int32) *ListIntervenesResponseBodyData {
	s.PageIndex = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetPageSize(v int32) *ListIntervenesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListIntervenesResponseBodyData) SetTotalSize(v int64) *ListIntervenesResponseBodyData {
	s.TotalSize = &v
	return s
}

type ListIntervenesResponseBodyDataInterveneList struct {
	// id
	Id    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
}

func (s ListIntervenesResponseBodyDataInterveneList) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponseBodyDataInterveneList) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetId(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Id = &v
	return s
}

func (s *ListIntervenesResponseBodyDataInterveneList) SetQuery(v string) *ListIntervenesResponseBodyDataInterveneList {
	s.Query = &v
	return s
}

type ListIntervenesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIntervenesResponse) GoString() string {
	return s.String()
}

func (s *ListIntervenesResponse) SetHeaders(v map[string]*string) *ListIntervenesResponse {
	s.Headers = v
	return s
}

func (s *ListIntervenesResponse) SetStatusCode(v int32) *ListIntervenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervenesResponse) SetBody(v *ListIntervenesResponseBody) *ListIntervenesResponse {
	s.Body = v
	return s
}

type ListMaterialDocumentsRequest struct {
	AgentKey          *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content           *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTimeEnd     *string   `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart   *string   `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Current           *int32    `json:"Current,omitempty" xml:"Current,omitempty"`
	DocType           *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocTypeList       []*string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty" type:"Repeated"`
	GeneratePublicUrl *bool     `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	Id                *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords          []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Query             *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	ShareAttr         *int32    `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	Size              *int32    `json:"Size,omitempty" xml:"Size,omitempty"`
	Title             *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTimeEnd     *string   `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	UpdateTimeStart   *string   `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsRequest) SetAgentKey(v string) *ListMaterialDocumentsRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetContent(v string) *ListMaterialDocumentsRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetCurrent(v int32) *ListMaterialDocumentsRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocType(v string) *ListMaterialDocumentsRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetDocTypeList(v []*string) *ListMaterialDocumentsRequest {
	s.DocTypeList = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetId(v int64) *ListMaterialDocumentsRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetKeywords(v []*string) *ListMaterialDocumentsRequest {
	s.Keywords = v
	return s
}

func (s *ListMaterialDocumentsRequest) SetQuery(v string) *ListMaterialDocumentsRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetShareAttr(v int32) *ListMaterialDocumentsRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetSize(v int32) *ListMaterialDocumentsRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetTitle(v string) *ListMaterialDocumentsRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsRequest {
	s.UpdateTimeStart = &v
	return s
}

type ListMaterialDocumentsShrinkRequest struct {
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTimeEnd     *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	CreateTimeStart   *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	Current           *int32  `json:"Current,omitempty" xml:"Current,omitempty"`
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	DocTypeListShrink *string `json:"DocTypeList,omitempty" xml:"DocTypeList,omitempty"`
	GeneratePublicUrl *bool   `json:"GeneratePublicUrl,omitempty" xml:"GeneratePublicUrl,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordsShrink    *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Query             *string `json:"Query,omitempty" xml:"Query,omitempty"`
	ShareAttr         *int32  `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	Size              *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTimeEnd     *string `json:"UpdateTimeEnd,omitempty" xml:"UpdateTimeEnd,omitempty"`
	UpdateTimeStart   *string `json:"UpdateTimeStart,omitempty" xml:"UpdateTimeStart,omitempty"`
}

func (s ListMaterialDocumentsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsShrinkRequest) SetAgentKey(v string) *ListMaterialDocumentsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetContent(v string) *ListMaterialDocumentsShrinkRequest {
	s.Content = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCreateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetCurrent(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocType(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetDocTypeListShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.DocTypeListShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetGeneratePublicUrl(v bool) *ListMaterialDocumentsShrinkRequest {
	s.GeneratePublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetId(v int64) *ListMaterialDocumentsShrinkRequest {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetKeywordsShrink(v string) *ListMaterialDocumentsShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetQuery(v string) *ListMaterialDocumentsShrinkRequest {
	s.Query = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetShareAttr(v int32) *ListMaterialDocumentsShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetSize(v int32) *ListMaterialDocumentsShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetTitle(v string) *ListMaterialDocumentsShrinkRequest {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeEnd(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeEnd = &v
	return s
}

func (s *ListMaterialDocumentsShrinkRequest) SetUpdateTimeStart(v string) *ListMaterialDocumentsShrinkRequest {
	s.UpdateTimeStart = &v
	return s
}

type ListMaterialDocumentsResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Current        *int32                                   `json:"Current,omitempty" xml:"Current,omitempty"`
	Data           []*ListMaterialDocumentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int32                                   `json:"Size,omitempty" xml:"Size,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                                   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMaterialDocumentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBody) SetCode(v string) *ListMaterialDocumentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetCurrent(v int32) *ListMaterialDocumentsResponseBody {
	s.Current = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetData(v []*ListMaterialDocumentsResponseBodyData) *ListMaterialDocumentsResponseBody {
	s.Data = v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetHttpStatusCode(v int32) *ListMaterialDocumentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetMessage(v string) *ListMaterialDocumentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetRequestId(v string) *ListMaterialDocumentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSize(v int32) *ListMaterialDocumentsResponseBody {
	s.Size = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetSuccess(v bool) *ListMaterialDocumentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMaterialDocumentsResponseBody) SetTotal(v int32) *ListMaterialDocumentsResponseBody {
	s.Total = &v
	return s
}

type ListMaterialDocumentsResponseBodyData struct {
	Author            *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser        *string   `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	CreateUserName    *string   `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	DocKeywords       []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	DocType           *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl       *string   `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent       *string   `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	Id                *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	PubTime           *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	PublicUrl         *string   `json:"PublicUrl,omitempty" xml:"PublicUrl,omitempty"`
	ShareAttr         *int32    `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom           *string   `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	ThumbnailInBase64 *string   `json:"ThumbnailInBase64,omitempty" xml:"ThumbnailInBase64,omitempty"`
	Title             *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime        *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser        *string   `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
	UpdateUserName    *string   `json:"UpdateUserName,omitempty" xml:"UpdateUserName,omitempty"`
	Url               *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListMaterialDocumentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponseBodyData) SetAuthor(v string) *ListMaterialDocumentsResponseBodyData {
	s.Author = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetCreateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocKeywords(v []*string) *ListMaterialDocumentsResponseBodyData {
	s.DocKeywords = v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetDocType(v string) *ListMaterialDocumentsResponseBodyData {
	s.DocType = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetExternalUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.ExternalUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetHtmlContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.HtmlContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetId(v int64) *ListMaterialDocumentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPubTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetPublicUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.PublicUrl = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetShareAttr(v int32) *ListMaterialDocumentsResponseBodyData {
	s.ShareAttr = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSrcFrom(v string) *ListMaterialDocumentsResponseBodyData {
	s.SrcFrom = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetSummary(v string) *ListMaterialDocumentsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTextContent(v string) *ListMaterialDocumentsResponseBodyData {
	s.TextContent = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetThumbnailInBase64(v string) *ListMaterialDocumentsResponseBodyData {
	s.ThumbnailInBase64 = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetTitle(v string) *ListMaterialDocumentsResponseBodyData {
	s.Title = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateTime(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUser(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUpdateUserName(v string) *ListMaterialDocumentsResponseBodyData {
	s.UpdateUserName = &v
	return s
}

func (s *ListMaterialDocumentsResponseBodyData) SetUrl(v string) *ListMaterialDocumentsResponseBodyData {
	s.Url = &v
	return s
}

type ListMaterialDocumentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMaterialDocumentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMaterialDocumentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMaterialDocumentsResponse) GoString() string {
	return s.String()
}

func (s *ListMaterialDocumentsResponse) SetHeaders(v map[string]*string) *ListMaterialDocumentsResponse {
	s.Headers = v
	return s
}

func (s *ListMaterialDocumentsResponse) SetStatusCode(v int32) *ListMaterialDocumentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMaterialDocumentsResponse) SetBody(v *ListMaterialDocumentsResponseBody) *ListMaterialDocumentsResponse {
	s.Body = v
	return s
}

type ListVersionsRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ListVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListVersionsRequest) SetAgentKey(v string) *ListVersionsRequest {
	s.AgentKey = &v
	return s
}

type ListVersionsResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBody) SetCode(v string) *ListVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListVersionsResponseBody) SetData(v []*ListVersionsResponseBodyData) *ListVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListVersionsResponseBody) SetHttpStatusCode(v int32) *ListVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVersionsResponseBody) SetMessage(v string) *ListVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListVersionsResponseBody) SetRequestId(v string) *ListVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVersionsResponseBody) SetSuccess(v bool) *ListVersionsResponseBody {
	s.Success = &v
	return s
}

type ListVersionsResponseBodyData struct {
	ConcurrentCount *int32  `json:"ConcurrentCount,omitempty" xml:"ConcurrentCount,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceCount   *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId         *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ProductType     *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	Quota           *int32  `json:"Quota,omitempty" xml:"Quota,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UseQuota        *int32  `json:"UseQuota,omitempty" xml:"UseQuota,omitempty"`
	VersionDetail   *string `json:"VersionDetail,omitempty" xml:"VersionDetail,omitempty"`
	VersionName     *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
	VersionStatus   *int32  `json:"VersionStatus,omitempty" xml:"VersionStatus,omitempty"`
}

func (s ListVersionsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVersionsResponseBodyData) SetConcurrentCount(v int32) *ListVersionsResponseBodyData {
	s.ConcurrentCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetEndTime(v string) *ListVersionsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceCount(v int32) *ListVersionsResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetInstanceId(v string) *ListVersionsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetOrderId(v int64) *ListVersionsResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetProductType(v string) *ListVersionsResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetQuota(v int32) *ListVersionsResponseBodyData {
	s.Quota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetStartTime(v string) *ListVersionsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetUseQuota(v int32) *ListVersionsResponseBodyData {
	s.UseQuota = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionDetail(v string) *ListVersionsResponseBodyData {
	s.VersionDetail = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionName(v string) *ListVersionsResponseBodyData {
	s.VersionName = &v
	return s
}

func (s *ListVersionsResponseBodyData) SetVersionStatus(v int32) *ListVersionsResponseBodyData {
	s.VersionStatus = &v
	return s
}

type ListVersionsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListVersionsResponse) SetHeaders(v map[string]*string) *ListVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListVersionsResponse) SetStatusCode(v int32) *ListVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVersionsResponse) SetBody(v *ListVersionsResponseBody) *ListVersionsResponse {
	s.Body = v
	return s
}

type QueryAsyncTaskRequest struct {
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s QueryAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskRequest) SetAgentKey(v string) *QueryAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *QueryAsyncTaskRequest) SetTaskId(v string) *QueryAsyncTaskRequest {
	s.TaskId = &v
	return s
}

type QueryAsyncTaskResponseBody struct {
	Code           *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *QueryAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                           `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBody) SetCode(v string) *QueryAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetData(v *QueryAsyncTaskResponseBodyData) *QueryAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetHttpStatusCode(v int32) *QueryAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetMessage(v string) *QueryAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetRequestId(v string) *QueryAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAsyncTaskResponseBody) SetSuccess(v bool) *QueryAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type QueryAsyncTaskResponseBodyData struct {
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser             *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	TaskCode               *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskErrorMessage       *string `json:"TaskErrorMessage,omitempty" xml:"TaskErrorMessage,omitempty"`
	TaskId                 *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskIntermediateResult *string `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskParam              *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
	TaskProgressMessage    *string `json:"TaskProgressMessage,omitempty" xml:"TaskProgressMessage,omitempty"`
	TaskResult             *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	TaskRetryCount         *string `json:"TaskRetryCount,omitempty" xml:"TaskRetryCount,omitempty"`
	TaskStatus             *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UpdateTime             *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateUser             *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s QueryAsyncTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetCreateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskCode(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskCode = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskErrorMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskErrorMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskId(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskIntermediateResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskName(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskParam(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskParam = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskProgressMessage(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskProgressMessage = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskResult(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskResult = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskRetryCount(v string) *QueryAsyncTaskResponseBodyData {
	s.TaskRetryCount = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetTaskStatus(v int32) *QueryAsyncTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateTime(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *QueryAsyncTaskResponseBodyData) SetUpdateUser(v string) *QueryAsyncTaskResponseBodyData {
	s.UpdateUser = &v
	return s
}

type QueryAsyncTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *QueryAsyncTaskResponse) SetHeaders(v map[string]*string) *QueryAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *QueryAsyncTaskResponse) SetStatusCode(v int32) *QueryAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAsyncTaskResponse) SetBody(v *QueryAsyncTaskResponseBody) *QueryAsyncTaskResponse {
	s.Body = v
	return s
}

type SaveDataSourceOrderConfigRequest struct {
	AgentKey                 *string                                                     `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ProductCode              *string                                                     `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserConfigDataSourceList []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty" type:"Repeated"`
}

func (s SaveDataSourceOrderConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetProductCode(v string) *SaveDataSourceOrderConfigRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequest) SetUserConfigDataSourceList(v []*SaveDataSourceOrderConfigRequestUserConfigDataSourceList) *SaveDataSourceOrderConfigRequest {
	s.UserConfigDataSourceList = v
	return s
}

type SaveDataSourceOrderConfigRequestUserConfigDataSourceList struct {
	Code   *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Number *int32  `json:"Number,omitempty" xml:"Number,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigRequestUserConfigDataSourceList) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetCode(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetName(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Name = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetNumber(v int32) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Number = &v
	return s
}

func (s *SaveDataSourceOrderConfigRequestUserConfigDataSourceList) SetType(v string) *SaveDataSourceOrderConfigRequestUserConfigDataSourceList {
	s.Type = &v
	return s
}

type SaveDataSourceOrderConfigShrinkRequest struct {
	AgentKey                       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ProductCode                    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	UserConfigDataSourceListShrink *string `json:"UserConfigDataSourceList,omitempty" xml:"UserConfigDataSourceList,omitempty"`
}

func (s SaveDataSourceOrderConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetAgentKey(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetProductCode(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigShrinkRequest) SetUserConfigDataSourceListShrink(v string) *SaveDataSourceOrderConfigShrinkRequest {
	s.UserConfigDataSourceListShrink = &v
	return s
}

type SaveDataSourceOrderConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveDataSourceOrderConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponseBody) SetCode(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetData(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Data = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetHttpStatusCode(v int32) *SaveDataSourceOrderConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetMessage(v string) *SaveDataSourceOrderConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetRequestId(v string) *SaveDataSourceOrderConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponseBody) SetSuccess(v bool) *SaveDataSourceOrderConfigResponseBody {
	s.Success = &v
	return s
}

type SaveDataSourceOrderConfigResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveDataSourceOrderConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveDataSourceOrderConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveDataSourceOrderConfigResponse) GoString() string {
	return s.String()
}

func (s *SaveDataSourceOrderConfigResponse) SetHeaders(v map[string]*string) *SaveDataSourceOrderConfigResponse {
	s.Headers = v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetStatusCode(v int32) *SaveDataSourceOrderConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveDataSourceOrderConfigResponse) SetBody(v *SaveDataSourceOrderConfigResponseBody) *SaveDataSourceOrderConfigResponse {
	s.Body = v
	return s
}

type SaveMaterialDocumentRequest struct {
	AgentKey                *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author                  *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	BothSavePrivateAndShare *bool     `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywords             []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	DocType                 *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl             *string   `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent             *string   `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	PubTime                 *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ShareAttr               *int32    `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom                 *string   `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary                 *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent             *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title                   *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Url                     *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SaveMaterialDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentRequest) SetAgentKey(v string) *SaveMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetAuthor(v string) *SaveMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocKeywords(v []*string) *SaveMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *SaveMaterialDocumentRequest) SetDocType(v string) *SaveMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetExternalUrl(v string) *SaveMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetHtmlContent(v string) *SaveMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetPubTime(v string) *SaveMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetShareAttr(v int32) *SaveMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSrcFrom(v string) *SaveMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetSummary(v string) *SaveMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTextContent(v string) *SaveMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetTitle(v string) *SaveMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentRequest) SetUrl(v string) *SaveMaterialDocumentRequest {
	s.Url = &v
	return s
}

type SaveMaterialDocumentShrinkRequest struct {
	AgentKey                *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author                  *string `json:"Author,omitempty" xml:"Author,omitempty"`
	BothSavePrivateAndShare *bool   `json:"BothSavePrivateAndShare,omitempty" xml:"BothSavePrivateAndShare,omitempty"`
	DocKeywordsShrink       *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	DocType                 *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl             *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent             *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	PubTime                 *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ShareAttr               *int32  `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom                 *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary                 *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent             *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title                   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url                     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SaveMaterialDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentShrinkRequest) SetAgentKey(v string) *SaveMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetAuthor(v string) *SaveMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetBothSavePrivateAndShare(v bool) *SaveMaterialDocumentShrinkRequest {
	s.BothSavePrivateAndShare = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetDocType(v string) *SaveMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetExternalUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetHtmlContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetPubTime(v string) *SaveMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetShareAttr(v int32) *SaveMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSrcFrom(v string) *SaveMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetSummary(v string) *SaveMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTextContent(v string) *SaveMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetTitle(v string) *SaveMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *SaveMaterialDocumentShrinkRequest) SetUrl(v string) *SaveMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

type SaveMaterialDocumentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveMaterialDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponseBody) SetCode(v string) *SaveMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetData(v int64) *SaveMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *SaveMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetMessage(v string) *SaveMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetRequestId(v string) *SaveMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveMaterialDocumentResponseBody) SetSuccess(v bool) *SaveMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

type SaveMaterialDocumentResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveMaterialDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *SaveMaterialDocumentResponse) SetHeaders(v map[string]*string) *SaveMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *SaveMaterialDocumentResponse) SetStatusCode(v int32) *SaveMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveMaterialDocumentResponse) SetBody(v *SaveMaterialDocumentResponseBody) *SaveMaterialDocumentResponse {
	s.Body = v
	return s
}

type SearchNewsRequest struct {
	AgentKey       *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FilterNotNull  *bool     `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	IncludeContent *bool     `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	Page           *int32    `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize       *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query          *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchSources  []*string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty" type:"Repeated"`
}

func (s SearchNewsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsRequest) SetAgentKey(v string) *SearchNewsRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsRequest) SetFilterNotNull(v bool) *SearchNewsRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsRequest) SetIncludeContent(v bool) *SearchNewsRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsRequest) SetPage(v int32) *SearchNewsRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsRequest) SetPageSize(v int32) *SearchNewsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsRequest) SetQuery(v string) *SearchNewsRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsRequest) SetSearchSources(v []*string) *SearchNewsRequest {
	s.SearchSources = v
	return s
}

type SearchNewsShrinkRequest struct {
	AgentKey            *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	FilterNotNull       *bool   `json:"FilterNotNull,omitempty" xml:"FilterNotNull,omitempty"`
	IncludeContent      *bool   `json:"IncludeContent,omitempty" xml:"IncludeContent,omitempty"`
	Page                *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Query               *string `json:"Query,omitempty" xml:"Query,omitempty"`
	SearchSourcesShrink *string `json:"SearchSources,omitempty" xml:"SearchSources,omitempty"`
}

func (s SearchNewsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchNewsShrinkRequest) SetAgentKey(v string) *SearchNewsShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetFilterNotNull(v bool) *SearchNewsShrinkRequest {
	s.FilterNotNull = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetIncludeContent(v bool) *SearchNewsShrinkRequest {
	s.IncludeContent = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPage(v int32) *SearchNewsShrinkRequest {
	s.Page = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetPageSize(v int32) *SearchNewsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetQuery(v string) *SearchNewsShrinkRequest {
	s.Query = &v
	return s
}

func (s *SearchNewsShrinkRequest) SetSearchSourcesShrink(v string) *SearchNewsShrinkRequest {
	s.SearchSourcesShrink = &v
	return s
}

type SearchNewsResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Current        *int32                        `json:"Current,omitempty" xml:"Current,omitempty"`
	Data           []*SearchNewsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Size           *int32                        `json:"Size,omitempty" xml:"Size,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
	Total          *int32                        `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchNewsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBody) SetCode(v string) *SearchNewsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchNewsResponseBody) SetCurrent(v int32) *SearchNewsResponseBody {
	s.Current = &v
	return s
}

func (s *SearchNewsResponseBody) SetData(v []*SearchNewsResponseBodyData) *SearchNewsResponseBody {
	s.Data = v
	return s
}

func (s *SearchNewsResponseBody) SetHttpStatusCode(v int32) *SearchNewsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchNewsResponseBody) SetMessage(v string) *SearchNewsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchNewsResponseBody) SetRequestId(v string) *SearchNewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchNewsResponseBody) SetSize(v int32) *SearchNewsResponseBody {
	s.Size = &v
	return s
}

func (s *SearchNewsResponseBody) SetSuccess(v bool) *SearchNewsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchNewsResponseBody) SetTotal(v int32) *SearchNewsResponseBody {
	s.Total = &v
	return s
}

type SearchNewsResponseBodyData struct {
	Author           *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	Content          *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	DocUuid          *string   `json:"DocUuid,omitempty" xml:"DocUuid,omitempty"`
	ImageUrls        []*string `json:"ImageUrls,omitempty" xml:"ImageUrls,omitempty" type:"Repeated"`
	PubTime          *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string   `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string   `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Source           *string   `json:"Source,omitempty" xml:"Source,omitempty"`
	Summary          *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	Tag              *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Title            *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdateTime       *string   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Url              *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SearchNewsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchNewsResponseBodyData) SetAuthor(v string) *SearchNewsResponseBodyData {
	s.Author = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetContent(v string) *SearchNewsResponseBodyData {
	s.Content = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetDocUuid(v string) *SearchNewsResponseBodyData {
	s.DocUuid = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetImageUrls(v []*string) *SearchNewsResponseBodyData {
	s.ImageUrls = v
	return s
}

func (s *SearchNewsResponseBodyData) SetPubTime(v string) *SearchNewsResponseBodyData {
	s.PubTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSource(v string) *SearchNewsResponseBodyData {
	s.SearchSource = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSearchSourceName(v string) *SearchNewsResponseBodyData {
	s.SearchSourceName = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSource(v string) *SearchNewsResponseBodyData {
	s.Source = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetSummary(v string) *SearchNewsResponseBodyData {
	s.Summary = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTag(v string) *SearchNewsResponseBodyData {
	s.Tag = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetTitle(v string) *SearchNewsResponseBodyData {
	s.Title = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUpdateTime(v string) *SearchNewsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *SearchNewsResponseBodyData) SetUrl(v string) *SearchNewsResponseBodyData {
	s.Url = &v
	return s
}

type SearchNewsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchNewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchNewsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchNewsResponse) GoString() string {
	return s.String()
}

func (s *SearchNewsResponse) SetHeaders(v map[string]*string) *SearchNewsResponse {
	s.Headers = v
	return s
}

func (s *SearchNewsResponse) SetStatusCode(v int32) *SearchNewsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchNewsResponse) SetBody(v *SearchNewsResponseBody) *SearchNewsResponse {
	s.Body = v
	return s
}

type SubmitAsyncTaskRequest struct {
	AgentKey        *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	TaskCode        *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskExecuteTime *string `json:"TaskExecuteTime,omitempty" xml:"TaskExecuteTime,omitempty"`
	TaskName        *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskParam       *string `json:"TaskParam,omitempty" xml:"TaskParam,omitempty"`
}

func (s SubmitAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskRequest) SetAgentKey(v string) *SubmitAsyncTaskRequest {
	s.AgentKey = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskCode(v string) *SubmitAsyncTaskRequest {
	s.TaskCode = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskExecuteTime(v string) *SubmitAsyncTaskRequest {
	s.TaskExecuteTime = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskName(v string) *SubmitAsyncTaskRequest {
	s.TaskName = &v
	return s
}

func (s *SubmitAsyncTaskRequest) SetTaskParam(v string) *SubmitAsyncTaskRequest {
	s.TaskParam = &v
	return s
}

type SubmitAsyncTaskResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *SubmitAsyncTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBody) SetCode(v string) *SubmitAsyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetData(v *SubmitAsyncTaskResponseBodyData) *SubmitAsyncTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetHttpStatusCode(v int32) *SubmitAsyncTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetMessage(v string) *SubmitAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetRequestId(v string) *SubmitAsyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBody) SetSuccess(v bool) *SubmitAsyncTaskResponseBody {
	s.Success = &v
	return s
}

type SubmitAsyncTaskResponseBodyData struct {
	TaskId                 *string     `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskIntermediateResult interface{} `json:"TaskIntermediateResult,omitempty" xml:"TaskIntermediateResult,omitempty"`
	TaskName               *string     `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s SubmitAsyncTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskId(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskIntermediateResult(v interface{}) *SubmitAsyncTaskResponseBodyData {
	s.TaskIntermediateResult = v
	return s
}

func (s *SubmitAsyncTaskResponseBodyData) SetTaskName(v string) *SubmitAsyncTaskResponseBodyData {
	s.TaskName = &v
	return s
}

type SubmitAsyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *SubmitAsyncTaskResponse) SetHeaders(v map[string]*string) *SubmitAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *SubmitAsyncTaskResponse) SetStatusCode(v int32) *SubmitAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAsyncTaskResponse) SetBody(v *SubmitAsyncTaskResponseBody) *SubmitAsyncTaskResponse {
	s.Body = v
	return s
}

type UpdateGeneratedContentRequest struct {
	AgentKey    *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content     *string   `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentText *string   `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	Keywords    []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	Prompt      *string   `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Title       *string   `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentRequest) SetAgentKey(v string) *UpdateGeneratedContentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContent(v string) *UpdateGeneratedContentRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetContentText(v string) *UpdateGeneratedContentRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetId(v int64) *UpdateGeneratedContentRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetKeywords(v []*string) *UpdateGeneratedContentRequest {
	s.Keywords = v
	return s
}

func (s *UpdateGeneratedContentRequest) SetPrompt(v string) *UpdateGeneratedContentRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentRequest) SetTitle(v string) *UpdateGeneratedContentRequest {
	s.Title = &v
	return s
}

type UpdateGeneratedContentShrinkRequest struct {
	AgentKey       *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Content        *string `json:"Content,omitempty" xml:"Content,omitempty"`
	ContentText    *string `json:"ContentText,omitempty" xml:"ContentText,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	KeywordsShrink *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	Title          *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateGeneratedContentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentShrinkRequest) SetAgentKey(v string) *UpdateGeneratedContentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContent(v string) *UpdateGeneratedContentShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetContentText(v string) *UpdateGeneratedContentShrinkRequest {
	s.ContentText = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetId(v int64) *UpdateGeneratedContentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetKeywordsShrink(v string) *UpdateGeneratedContentShrinkRequest {
	s.KeywordsShrink = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetPrompt(v string) *UpdateGeneratedContentShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *UpdateGeneratedContentShrinkRequest) SetTitle(v string) *UpdateGeneratedContentShrinkRequest {
	s.Title = &v
	return s
}

type UpdateGeneratedContentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateGeneratedContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponseBody) SetCode(v string) *UpdateGeneratedContentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetData(v bool) *UpdateGeneratedContentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetHttpStatusCode(v int32) *UpdateGeneratedContentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetMessage(v string) *UpdateGeneratedContentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetRequestId(v string) *UpdateGeneratedContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGeneratedContentResponseBody) SetSuccess(v bool) *UpdateGeneratedContentResponseBody {
	s.Success = &v
	return s
}

type UpdateGeneratedContentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGeneratedContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGeneratedContentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGeneratedContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateGeneratedContentResponse) SetHeaders(v map[string]*string) *UpdateGeneratedContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateGeneratedContentResponse) SetStatusCode(v int32) *UpdateGeneratedContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGeneratedContentResponse) SetBody(v *UpdateGeneratedContentResponseBody) *UpdateGeneratedContentResponse {
	s.Body = v
	return s
}

type UpdateMaterialDocumentRequest struct {
	AgentKey    *string   `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author      *string   `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywords []*string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty" type:"Repeated"`
	DocType     *string   `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl *string   `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent *string   `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	Id          *int64    `json:"Id,omitempty" xml:"Id,omitempty"`
	PubTime     *string   `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ShareAttr   *int32    `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom     *string   `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary     *string   `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent *string   `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title       *string   `json:"Title,omitempty" xml:"Title,omitempty"`
	Url         *string   `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateMaterialDocumentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentRequest) SetAgentKey(v string) *UpdateMaterialDocumentRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetAuthor(v string) *UpdateMaterialDocumentRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocKeywords(v []*string) *UpdateMaterialDocumentRequest {
	s.DocKeywords = v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetDocType(v string) *UpdateMaterialDocumentRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetExternalUrl(v string) *UpdateMaterialDocumentRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetHtmlContent(v string) *UpdateMaterialDocumentRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetId(v int64) *UpdateMaterialDocumentRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetPubTime(v string) *UpdateMaterialDocumentRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetShareAttr(v int32) *UpdateMaterialDocumentRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSrcFrom(v string) *UpdateMaterialDocumentRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetSummary(v string) *UpdateMaterialDocumentRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTextContent(v string) *UpdateMaterialDocumentRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetTitle(v string) *UpdateMaterialDocumentRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentRequest) SetUrl(v string) *UpdateMaterialDocumentRequest {
	s.Url = &v
	return s
}

type UpdateMaterialDocumentShrinkRequest struct {
	AgentKey          *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	Author            *string `json:"Author,omitempty" xml:"Author,omitempty"`
	DocKeywordsShrink *string `json:"DocKeywords,omitempty" xml:"DocKeywords,omitempty"`
	DocType           *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	ExternalUrl       *string `json:"ExternalUrl,omitempty" xml:"ExternalUrl,omitempty"`
	HtmlContent       *string `json:"HtmlContent,omitempty" xml:"HtmlContent,omitempty"`
	Id                *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	PubTime           *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	ShareAttr         *int32  `json:"ShareAttr,omitempty" xml:"ShareAttr,omitempty"`
	SrcFrom           *string `json:"SrcFrom,omitempty" xml:"SrcFrom,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TextContent       *string `json:"TextContent,omitempty" xml:"TextContent,omitempty"`
	Title             *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url               *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateMaterialDocumentShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAgentKey(v string) *UpdateMaterialDocumentShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetAuthor(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Author = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocKeywordsShrink(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocKeywordsShrink = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetDocType(v string) *UpdateMaterialDocumentShrinkRequest {
	s.DocType = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetExternalUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.ExternalUrl = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetHtmlContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.HtmlContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetId(v int64) *UpdateMaterialDocumentShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetPubTime(v string) *UpdateMaterialDocumentShrinkRequest {
	s.PubTime = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetShareAttr(v int32) *UpdateMaterialDocumentShrinkRequest {
	s.ShareAttr = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSrcFrom(v string) *UpdateMaterialDocumentShrinkRequest {
	s.SrcFrom = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetSummary(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Summary = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTextContent(v string) *UpdateMaterialDocumentShrinkRequest {
	s.TextContent = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetTitle(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Title = &v
	return s
}

func (s *UpdateMaterialDocumentShrinkRequest) SetUrl(v string) *UpdateMaterialDocumentShrinkRequest {
	s.Url = &v
	return s
}

type UpdateMaterialDocumentResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateMaterialDocumentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponseBody) SetCode(v string) *UpdateMaterialDocumentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetData(v int64) *UpdateMaterialDocumentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetHttpStatusCode(v int32) *UpdateMaterialDocumentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetMessage(v string) *UpdateMaterialDocumentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetRequestId(v string) *UpdateMaterialDocumentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMaterialDocumentResponseBody) SetSuccess(v bool) *UpdateMaterialDocumentResponseBody {
	s.Success = &v
	return s
}

type UpdateMaterialDocumentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMaterialDocumentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMaterialDocumentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMaterialDocumentResponse) GoString() string {
	return s.String()
}

func (s *UpdateMaterialDocumentResponse) SetHeaders(v map[string]*string) *UpdateMaterialDocumentResponse {
	s.Headers = v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetStatusCode(v int32) *UpdateMaterialDocumentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMaterialDocumentResponse) SetBody(v *UpdateMaterialDocumentResponseBody) *UpdateMaterialDocumentResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("aimiaobi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelAsyncTaskWithOptions(request *CancelAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *CancelAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelAsyncTask(request *CancelAsyncTaskRequest) (_result *CancelAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelAsyncTaskResponse{}
	_body, _err := client.CancelAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ClearIntervenesWithOptions(request *ClearIntervenesRequest, runtime *util.RuntimeOptions) (_result *ClearIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ClearIntervenes(request *ClearIntervenesRequest) (_result *ClearIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ClearIntervenesResponse{}
	_body, _err := client.ClearIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGeneratedContentWithOptions(tmpReq *CreateGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *CreateGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentDomain)) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !tea.BoolValue(util.IsUnset(request.ContentText)) {
		body["ContentText"] = request.ContentText
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Uuid)) {
		body["Uuid"] = request.Uuid
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGeneratedContent(request *CreateGeneratedContentRequest) (_result *CreateGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGeneratedContentResponse{}
	_body, _err := client.CreateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTokenWithOptions(request *CreateTokenRequest, runtime *util.RuntimeOptions) (_result *CreateTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateToken"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateToken(request *CreateTokenRequest) (_result *CreateTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTokenResponse{}
	_body, _err := client.CreateTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGeneratedContentWithOptions(request *DeleteGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *DeleteGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGeneratedContent(request *DeleteGeneratedContentRequest) (_result *DeleteGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGeneratedContentResponse{}
	_body, _err := client.DeleteGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInterveneRuleWithOptions(request *DeleteInterveneRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteInterveneRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInterveneRule"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInterveneRule(request *DeleteInterveneRuleRequest) (_result *DeleteInterveneRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInterveneRuleResponse{}
	_body, _err := client.DeleteInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMaterialByIdWithOptions(request *DeleteMaterialByIdRequest, runtime *util.RuntimeOptions) (_result *DeleteMaterialByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMaterialById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMaterialById(request *DeleteMaterialByIdRequest) (_result *DeleteMaterialByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMaterialByIdResponse{}
	_body, _err := client.DeleteMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportGeneratedContentWithOptions(request *ExportGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *ExportGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportGeneratedContent(request *ExportGeneratedContentRequest) (_result *ExportGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportGeneratedContentResponse{}
	_body, _err := client.ExportGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExportIntervenesWithOptions(request *ExportIntervenesRequest, runtime *util.RuntimeOptions) (_result *ExportIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExportIntervenes(request *ExportIntervenesRequest) (_result *ExportIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportIntervenesResponse{}
	_body, _err := client.ExportIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FeedbackDialogueWithOptions(tmpReq *FeedbackDialogueRequest, runtime *util.RuntimeOptions) (_result *FeedbackDialogueResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FeedbackDialogueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.RatingTags)) {
		request.RatingTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RatingTags, tea.String("RatingTags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomerResponse)) {
		body["CustomerResponse"] = request.CustomerResponse
	}

	if !tea.BoolValue(util.IsUnset(request.GoodText)) {
		body["GoodText"] = request.GoodText
	}

	if !tea.BoolValue(util.IsUnset(request.ModifiedResponse)) {
		body["ModifiedResponse"] = request.ModifiedResponse
	}

	if !tea.BoolValue(util.IsUnset(request.Rating)) {
		body["Rating"] = request.Rating
	}

	if !tea.BoolValue(util.IsUnset(request.RatingTagsShrink)) {
		body["RatingTags"] = request.RatingTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FeedbackDialogue"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FeedbackDialogue(request *FeedbackDialogueRequest) (_result *FeedbackDialogueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FeedbackDialogueResponse{}
	_body, _err := client.FeedbackDialogueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchImageTaskWithOptions(tmpReq *FetchImageTaskRequest, runtime *util.RuntimeOptions) (_result *FetchImageTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FetchImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskIdList)) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, tea.String("TaskIdList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleTaskId)) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIdListShrink)) {
		body["TaskIdList"] = request.TaskIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchImageTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchImageTask(request *FetchImageTaskRequest) (_result *FetchImageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchImageTaskResponse{}
	_body, _err := client.FetchImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateFileUrlByKeyWithOptions(request *GenerateFileUrlByKeyRequest, runtime *util.RuntimeOptions) (_result *GenerateFileUrlByKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateFileUrlByKey"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateFileUrlByKey(request *GenerateFileUrlByKeyRequest) (_result *GenerateFileUrlByKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateFileUrlByKeyResponse{}
	_body, _err := client.GenerateFileUrlByKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateImageTaskWithOptions(tmpReq *GenerateImageTaskRequest, runtime *util.RuntimeOptions) (_result *GenerateImageTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateImageTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParagraphList)) {
		request.ParagraphListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParagraphList, tea.String("ParagraphList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArticleTaskId)) {
		body["ArticleTaskId"] = request.ArticleTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ParagraphListShrink)) {
		body["ParagraphList"] = request.ParagraphListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Style)) {
		body["Style"] = request.Style
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateImageTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateImageTask(request *GenerateImageTaskRequest) (_result *GenerateImageTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateImageTaskResponse{}
	_body, _err := client.GenerateImageTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateUploadConfigWithOptions(request *GenerateUploadConfigRequest, runtime *util.RuntimeOptions) (_result *GenerateUploadConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["FileName"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDir)) {
		body["ParentDir"] = request.ParentDir
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateUploadConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateUploadConfig(request *GenerateUploadConfigRequest) (_result *GenerateUploadConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateUploadConfigResponse{}
	_body, _err := client.GenerateUploadConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateViewPointWithOptions(tmpReq *GenerateViewPointRequest, runtime *util.RuntimeOptions) (_result *GenerateViewPointResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GenerateViewPointShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReferenceData)) {
		request.ReferenceDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReferenceData, tea.String("ReferenceData"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReferenceDataShrink)) {
		body["ReferenceData"] = request.ReferenceDataShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateViewPoint"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateViewPoint(request *GenerateViewPointRequest) (_result *GenerateViewPointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateViewPointResponse{}
	_body, _err := client.GenerateViewPointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSourceOrderConfigWithOptions(request *GetDataSourceOrderConfigRequest, runtime *util.RuntimeOptions) (_result *GetDataSourceOrderConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDataSourceOrderConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSourceOrderConfig(request *GetDataSourceOrderConfigRequest) (_result *GetDataSourceOrderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDataSourceOrderConfigResponse{}
	_body, _err := client.GetDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGeneratedContentWithOptions(request *GetGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *GetGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGeneratedContent(request *GetGeneratedContentRequest) (_result *GetGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGeneratedContentResponse{}
	_body, _err := client.GetGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInterveneGlobalReplyWithOptions(request *GetInterveneGlobalReplyRequest, runtime *util.RuntimeOptions) (_result *GetInterveneGlobalReplyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneGlobalReply"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInterveneGlobalReply(request *GetInterveneGlobalReplyRequest) (_result *GetInterveneGlobalReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneGlobalReplyResponse{}
	_body, _err := client.GetInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInterveneImportTaskInfoWithOptions(request *GetInterveneImportTaskInfoRequest, runtime *util.RuntimeOptions) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneImportTaskInfo"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInterveneImportTaskInfo(request *GetInterveneImportTaskInfoRequest) (_result *GetInterveneImportTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneImportTaskInfoResponse{}
	_body, _err := client.GetInterveneImportTaskInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInterveneRuleDetailWithOptions(request *GetInterveneRuleDetailRequest, runtime *util.RuntimeOptions) (_result *GetInterveneRuleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneRuleDetail"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInterveneRuleDetail(request *GetInterveneRuleDetailRequest) (_result *GetInterveneRuleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneRuleDetailResponse{}
	_body, _err := client.GetInterveneRuleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInterveneTemplateFileUrlWithOptions(request *GetInterveneTemplateFileUrlRequest, runtime *util.RuntimeOptions) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInterveneTemplateFileUrl"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInterveneTemplateFileUrl(request *GetInterveneTemplateFileUrlRequest) (_result *GetInterveneTemplateFileUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInterveneTemplateFileUrlResponse{}
	_body, _err := client.GetInterveneTemplateFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMaterialByIdWithOptions(request *GetMaterialByIdRequest, runtime *util.RuntimeOptions) (_result *GetMaterialByIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMaterialById"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMaterialById(request *GetMaterialByIdRequest) (_result *GetMaterialByIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMaterialByIdResponse{}
	_body, _err := client.GetMaterialByIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPropertiesWithOptions(request *GetPropertiesRequest, runtime *util.RuntimeOptions) (_result *GetPropertiesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetProperties"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPropertiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProperties(request *GetPropertiesRequest) (_result *GetPropertiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPropertiesResponse{}
	_body, _err := client.GetPropertiesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportInterveneFileWithOptions(request *ImportInterveneFileRequest, runtime *util.RuntimeOptions) (_result *ImportInterveneFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		body["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportInterveneFile"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportInterveneFile(request *ImportInterveneFileRequest) (_result *ImportInterveneFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportInterveneFileResponse{}
	_body, _err := client.ImportInterveneFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportInterveneFileAsyncWithOptions(request *ImportInterveneFileAsyncRequest, runtime *util.RuntimeOptions) (_result *ImportInterveneFileAsyncResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocName)) {
		body["DocName"] = request.DocName
	}

	if !tea.BoolValue(util.IsUnset(request.FileKey)) {
		body["FileKey"] = request.FileKey
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportInterveneFileAsync"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportInterveneFileAsync(request *ImportInterveneFileAsyncRequest) (_result *ImportInterveneFileAsyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportInterveneFileAsyncResponse{}
	_body, _err := client.ImportInterveneFileAsyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertInterveneGlobalReplyWithOptions(tmpReq *InsertInterveneGlobalReplyRequest, runtime *util.RuntimeOptions) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertInterveneGlobalReplyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ReplyMessagList)) {
		request.ReplyMessagListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ReplyMessagList, tea.String("ReplyMessagList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReplyMessagListShrink)) {
		body["ReplyMessagList"] = request.ReplyMessagListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertInterveneGlobalReply"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertInterveneGlobalReply(request *InsertInterveneGlobalReplyRequest) (_result *InsertInterveneGlobalReplyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertInterveneGlobalReplyResponse{}
	_body, _err := client.InsertInterveneGlobalReplyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InsertInterveneRuleWithOptions(tmpReq *InsertInterveneRuleRequest, runtime *util.RuntimeOptions) (_result *InsertInterveneRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &InsertInterveneRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InterveneRuleConfig)) {
		request.InterveneRuleConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InterveneRuleConfig, tea.String("InterveneRuleConfig"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InterveneRuleConfigShrink)) {
		body["InterveneRuleConfig"] = request.InterveneRuleConfigShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InsertInterveneRule"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InsertInterveneRule(request *InsertInterveneRuleRequest) (_result *InsertInterveneRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InsertInterveneRuleResponse{}
	_body, _err := client.InsertInterveneRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAsyncTasksWithOptions(tmpReq *ListAsyncTasksRequest, runtime *util.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAsyncTasksShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TaskStatusList)) {
		request.TaskStatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskStatusList, tea.String("TaskStatusList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TaskTypeList)) {
		request.TaskTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTypeList, tea.String("TaskTypeList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.TaskCode)) {
		body["TaskCode"] = request.TaskCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		body["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatusListShrink)) {
		body["TaskStatusList"] = request.TaskStatusListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		body["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TaskTypeListShrink)) {
		body["TaskTypeList"] = request.TaskTypeListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAsyncTasks"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAsyncTasks(request *ListAsyncTasksRequest) (_result *ListAsyncTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.ListAsyncTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBuildConfigsWithOptions(request *ListBuildConfigsRequest, runtime *util.RuntimeOptions) (_result *ListBuildConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBuildConfigs"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBuildConfigs(request *ListBuildConfigsRequest) (_result *ListBuildConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListBuildConfigsResponse{}
	_body, _err := client.ListBuildConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDialoguesWithOptions(request *ListDialoguesRequest, runtime *util.RuntimeOptions) (_result *ListDialoguesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.DialogueType)) {
		body["DialogueType"] = request.DialogueType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDialogues"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDialoguesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDialogues(request *ListDialoguesRequest) (_result *ListDialoguesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDialoguesResponse{}
	_body, _err := client.ListDialoguesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGeneratedContentsWithOptions(request *ListGeneratedContentsRequest, runtime *util.RuntimeOptions) (_result *ListGeneratedContentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentDomain)) {
		body["ContentDomain"] = request.ContentDomain
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGeneratedContents"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGeneratedContents(request *ListGeneratedContentsRequest) (_result *ListGeneratedContentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGeneratedContentsResponse{}
	_body, _err := client.ListGeneratedContentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListHotNewsWithTypeWithOptions(tmpReq *ListHotNewsWithTypeRequest, runtime *util.RuntimeOptions) (_result *ListHotNewsWithTypeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListHotNewsWithTypeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NewsTypes)) {
		request.NewsTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NewsTypes, tea.String("NewsTypes"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.NewsType)) {
		body["NewsType"] = request.NewsType
	}

	if !tea.BoolValue(util.IsUnset(request.NewsTypesShrink)) {
		body["NewsTypes"] = request.NewsTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListHotNewsWithType"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListHotNewsWithType(request *ListHotNewsWithTypeRequest) (_result *ListHotNewsWithTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListHotNewsWithTypeResponse{}
	_body, _err := client.ListHotNewsWithTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterveneCntWithOptions(request *ListInterveneCntRequest, runtime *util.RuntimeOptions) (_result *ListInterveneCntResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneCnt"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterveneCnt(request *ListInterveneCntRequest) (_result *ListInterveneCntResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneCntResponse{}
	_body, _err := client.ListInterveneCntWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterveneImportTasksWithOptions(request *ListInterveneImportTasksRequest, runtime *util.RuntimeOptions) (_result *ListInterveneImportTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneImportTasks"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterveneImportTasks(request *ListInterveneImportTasksRequest) (_result *ListInterveneImportTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneImportTasksResponse{}
	_body, _err := client.ListInterveneImportTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterveneRulesWithOptions(request *ListInterveneRulesRequest, runtime *util.RuntimeOptions) (_result *ListInterveneRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterveneRules"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterveneRules(request *ListInterveneRulesRequest) (_result *ListInterveneRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInterveneRulesResponse{}
	_body, _err := client.ListInterveneRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListIntervenesWithOptions(request *ListIntervenesRequest, runtime *util.RuntimeOptions) (_result *ListIntervenesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InterveneType)) {
		body["InterveneType"] = request.InterveneType
	}

	if !tea.BoolValue(util.IsUnset(request.PageIndex)) {
		body["PageIndex"] = request.PageIndex
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIntervenes"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIntervenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListIntervenes(request *ListIntervenesRequest) (_result *ListIntervenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIntervenesResponse{}
	_body, _err := client.ListIntervenesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMaterialDocumentsWithOptions(tmpReq *ListMaterialDocumentsRequest, runtime *util.RuntimeOptions) (_result *ListMaterialDocumentsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMaterialDocumentsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocTypeList)) {
		request.DocTypeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocTypeList, tea.String("DocTypeList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeEnd)) {
		body["CreateTimeEnd"] = request.CreateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeStart)) {
		body["CreateTimeStart"] = request.CreateTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Current)) {
		body["Current"] = request.Current
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.DocTypeListShrink)) {
		body["DocTypeList"] = request.DocTypeListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GeneratePublicUrl)) {
		body["GeneratePublicUrl"] = request.GeneratePublicUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeEnd)) {
		body["UpdateTimeEnd"] = request.UpdateTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateTimeStart)) {
		body["UpdateTimeStart"] = request.UpdateTimeStart
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMaterialDocuments"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMaterialDocuments(request *ListMaterialDocumentsRequest) (_result *ListMaterialDocumentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMaterialDocumentsResponse{}
	_body, _err := client.ListMaterialDocumentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVersionsWithOptions(request *ListVersionsRequest, runtime *util.RuntimeOptions) (_result *ListVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVersions"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVersions(request *ListVersionsRequest) (_result *ListVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVersionsResponse{}
	_body, _err := client.ListVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryAsyncTaskWithOptions(request *QueryAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *QueryAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryAsyncTask(request *QueryAsyncTaskRequest) (_result *QueryAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryAsyncTaskResponse{}
	_body, _err := client.QueryAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveDataSourceOrderConfigWithOptions(tmpReq *SaveDataSourceOrderConfigRequest, runtime *util.RuntimeOptions) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveDataSourceOrderConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserConfigDataSourceList)) {
		request.UserConfigDataSourceListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserConfigDataSourceList, tea.String("UserConfigDataSourceList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductCode)) {
		body["ProductCode"] = request.ProductCode
	}

	if !tea.BoolValue(util.IsUnset(request.UserConfigDataSourceListShrink)) {
		body["UserConfigDataSourceList"] = request.UserConfigDataSourceListShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveDataSourceOrderConfig"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveDataSourceOrderConfig(request *SaveDataSourceOrderConfigRequest) (_result *SaveDataSourceOrderConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveDataSourceOrderConfigResponse{}
	_body, _err := client.SaveDataSourceOrderConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveMaterialDocumentWithOptions(tmpReq *SaveMaterialDocumentRequest, runtime *util.RuntimeOptions) (_result *SaveMaterialDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SaveMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocKeywords)) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, tea.String("DocKeywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Author)) {
		body["Author"] = request.Author
	}

	if !tea.BoolValue(util.IsUnset(request.BothSavePrivateAndShare)) {
		body["BothSavePrivateAndShare"] = request.BothSavePrivateAndShare
	}

	if !tea.BoolValue(util.IsUnset(request.DocKeywordsShrink)) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUrl)) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HtmlContent)) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !tea.BoolValue(util.IsUnset(request.PubTime)) {
		body["PubTime"] = request.PubTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.SrcFrom)) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveMaterialDocument"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveMaterialDocument(request *SaveMaterialDocumentRequest) (_result *SaveMaterialDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SaveMaterialDocumentResponse{}
	_body, _err := client.SaveMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchNewsWithOptions(tmpReq *SearchNewsRequest, runtime *util.RuntimeOptions) (_result *SearchNewsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchNewsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SearchSources)) {
		request.SearchSourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SearchSources, tea.String("SearchSources"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterNotNull)) {
		body["FilterNotNull"] = request.FilterNotNull
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeContent)) {
		body["IncludeContent"] = request.IncludeContent
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		body["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["Query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SearchSourcesShrink)) {
		body["SearchSources"] = request.SearchSourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchNews"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchNewsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchNews(request *SearchNewsRequest) (_result *SearchNewsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchNewsResponse{}
	_body, _err := client.SearchNewsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitAsyncTaskWithOptions(request *SubmitAsyncTaskRequest, runtime *util.RuntimeOptions) (_result *SubmitAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskCode)) {
		body["TaskCode"] = request.TaskCode
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecuteTime)) {
		body["TaskExecuteTime"] = request.TaskExecuteTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		body["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.TaskParam)) {
		body["TaskParam"] = request.TaskParam
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitAsyncTask"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitAsyncTask(request *SubmitAsyncTaskRequest) (_result *SubmitAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitAsyncTaskResponse{}
	_body, _err := client.SubmitAsyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGeneratedContentWithOptions(tmpReq *UpdateGeneratedContentRequest, runtime *util.RuntimeOptions) (_result *UpdateGeneratedContentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateGeneratedContentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Keywords)) {
		request.KeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Keywords, tea.String("Keywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentText)) {
		body["ContentText"] = request.ContentText
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.KeywordsShrink)) {
		body["Keywords"] = request.KeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Prompt)) {
		body["Prompt"] = request.Prompt
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGeneratedContent"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGeneratedContent(request *UpdateGeneratedContentRequest) (_result *UpdateGeneratedContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGeneratedContentResponse{}
	_body, _err := client.UpdateGeneratedContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMaterialDocumentWithOptions(tmpReq *UpdateMaterialDocumentRequest, runtime *util.RuntimeOptions) (_result *UpdateMaterialDocumentResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateMaterialDocumentShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DocKeywords)) {
		request.DocKeywordsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocKeywords, tea.String("DocKeywords"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AgentKey)) {
		query["AgentKey"] = request.AgentKey
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Author)) {
		body["Author"] = request.Author
	}

	if !tea.BoolValue(util.IsUnset(request.DocKeywordsShrink)) {
		body["DocKeywords"] = request.DocKeywordsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DocType)) {
		body["DocType"] = request.DocType
	}

	if !tea.BoolValue(util.IsUnset(request.ExternalUrl)) {
		body["ExternalUrl"] = request.ExternalUrl
	}

	if !tea.BoolValue(util.IsUnset(request.HtmlContent)) {
		body["HtmlContent"] = request.HtmlContent
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.PubTime)) {
		body["PubTime"] = request.PubTime
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAttr)) {
		body["ShareAttr"] = request.ShareAttr
	}

	if !tea.BoolValue(util.IsUnset(request.SrcFrom)) {
		body["SrcFrom"] = request.SrcFrom
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		body["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TextContent)) {
		body["TextContent"] = request.TextContent
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		body["Title"] = request.Title
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		body["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMaterialDocument"),
		Version:     tea.String("2023-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMaterialDocument(request *UpdateMaterialDocumentRequest) (_result *UpdateMaterialDocumentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMaterialDocumentResponse{}
	_body, _err := client.UpdateMaterialDocumentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
