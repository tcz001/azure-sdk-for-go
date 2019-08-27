// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/tcz001/azure-sdk-for-go/tools/profileBuilder

package botservice

import (
	"context"

	original "github.com/tcz001/azure-sdk-for-go/services/preview/botservice/mgmt/2018-07-12/botservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ChannelName = original.ChannelName

const (
	ChannelNameDirectLineChannel ChannelName = original.ChannelNameDirectLineChannel
	ChannelNameEmailChannel      ChannelName = original.ChannelNameEmailChannel
	ChannelNameFacebookChannel   ChannelName = original.ChannelNameFacebookChannel
	ChannelNameKikChannel        ChannelName = original.ChannelNameKikChannel
	ChannelNameMsTeamsChannel    ChannelName = original.ChannelNameMsTeamsChannel
	ChannelNameSkypeChannel      ChannelName = original.ChannelNameSkypeChannel
	ChannelNameSlackChannel      ChannelName = original.ChannelNameSlackChannel
	ChannelNameSmsChannel        ChannelName = original.ChannelNameSmsChannel
	ChannelNameTelegramChannel   ChannelName = original.ChannelNameTelegramChannel
	ChannelNameWebChatChannel    ChannelName = original.ChannelNameWebChatChannel
)

type ChannelNameBasicChannel = original.ChannelNameBasicChannel

const (
	ChannelNameChannel            ChannelNameBasicChannel = original.ChannelNameChannel
	ChannelNameDirectLineChannel1 ChannelNameBasicChannel = original.ChannelNameDirectLineChannel1
	ChannelNameEmailChannel1      ChannelNameBasicChannel = original.ChannelNameEmailChannel1
	ChannelNameFacebookChannel1   ChannelNameBasicChannel = original.ChannelNameFacebookChannel1
	ChannelNameKikChannel1        ChannelNameBasicChannel = original.ChannelNameKikChannel1
	ChannelNameMsTeamsChannel1    ChannelNameBasicChannel = original.ChannelNameMsTeamsChannel1
	ChannelNameSkypeChannel1      ChannelNameBasicChannel = original.ChannelNameSkypeChannel1
	ChannelNameSlackChannel1      ChannelNameBasicChannel = original.ChannelNameSlackChannel1
	ChannelNameSmsChannel1        ChannelNameBasicChannel = original.ChannelNameSmsChannel1
	ChannelNameTelegramChannel1   ChannelNameBasicChannel = original.ChannelNameTelegramChannel1
	ChannelNameWebChatChannel1    ChannelNameBasicChannel = original.ChannelNameWebChatChannel1
)

type EnterpriseChannelNodeState = original.EnterpriseChannelNodeState

const (
	CreateFailed EnterpriseChannelNodeState = original.CreateFailed
	Creating     EnterpriseChannelNodeState = original.Creating
	DeleteFailed EnterpriseChannelNodeState = original.DeleteFailed
	Deleting     EnterpriseChannelNodeState = original.Deleting
	Started      EnterpriseChannelNodeState = original.Started
	StartFailed  EnterpriseChannelNodeState = original.StartFailed
	Starting     EnterpriseChannelNodeState = original.Starting
	StopFailed   EnterpriseChannelNodeState = original.StopFailed
	Stopped      EnterpriseChannelNodeState = original.Stopped
	Stopping     EnterpriseChannelNodeState = original.Stopping
)

type EnterpriseChannelState = original.EnterpriseChannelState

const (
	EnterpriseChannelStateCreateFailed EnterpriseChannelState = original.EnterpriseChannelStateCreateFailed
	EnterpriseChannelStateCreating     EnterpriseChannelState = original.EnterpriseChannelStateCreating
	EnterpriseChannelStateDeleteFailed EnterpriseChannelState = original.EnterpriseChannelStateDeleteFailed
	EnterpriseChannelStateDeleting     EnterpriseChannelState = original.EnterpriseChannelStateDeleting
	EnterpriseChannelStateStarted      EnterpriseChannelState = original.EnterpriseChannelStateStarted
	EnterpriseChannelStateStartFailed  EnterpriseChannelState = original.EnterpriseChannelStateStartFailed
	EnterpriseChannelStateStarting     EnterpriseChannelState = original.EnterpriseChannelStateStarting
	EnterpriseChannelStateStopFailed   EnterpriseChannelState = original.EnterpriseChannelStateStopFailed
	EnterpriseChannelStateStopped      EnterpriseChannelState = original.EnterpriseChannelStateStopped
	EnterpriseChannelStateStopping     EnterpriseChannelState = original.EnterpriseChannelStateStopping
)

type Kind = original.Kind

const (
	KindBot      Kind = original.KindBot
	KindDesigner Kind = original.KindDesigner
	KindFunction Kind = original.KindFunction
	KindSdk      Kind = original.KindSdk
)

type SkuName = original.SkuName

const (
	F0 SkuName = original.F0
	S1 SkuName = original.S1
)

type SkuTier = original.SkuTier

const (
	Free     SkuTier = original.Free
	Standard SkuTier = original.Standard
)

type BaseClient = original.BaseClient
type BasicChannel = original.BasicChannel
type Bot = original.Bot
type BotChannel = original.BotChannel
type BotConnectionClient = original.BotConnectionClient
type BotProperties = original.BotProperties
type BotResponseList = original.BotResponseList
type BotResponseListIterator = original.BotResponseListIterator
type BotResponseListPage = original.BotResponseListPage
type BotsClient = original.BotsClient
type Channel = original.Channel
type ChannelResponseList = original.ChannelResponseList
type ChannelResponseListIterator = original.ChannelResponseListIterator
type ChannelResponseListPage = original.ChannelResponseListPage
type ChannelsClient = original.ChannelsClient
type CheckNameAvailabilityRequestBody = original.CheckNameAvailabilityRequestBody
type CheckNameAvailabilityResponseBody = original.CheckNameAvailabilityResponseBody
type ConnectionItemName = original.ConnectionItemName
type ConnectionSetting = original.ConnectionSetting
type ConnectionSettingParameter = original.ConnectionSettingParameter
type ConnectionSettingProperties = original.ConnectionSettingProperties
type ConnectionSettingResponseList = original.ConnectionSettingResponseList
type ConnectionSettingResponseListIterator = original.ConnectionSettingResponseListIterator
type ConnectionSettingResponseListPage = original.ConnectionSettingResponseListPage
type DirectLineChannel = original.DirectLineChannel
type DirectLineChannelProperties = original.DirectLineChannelProperties
type DirectLineSite = original.DirectLineSite
type EmailChannel = original.EmailChannel
type EmailChannelProperties = original.EmailChannelProperties
type EnterpriseChannel = original.EnterpriseChannel
type EnterpriseChannelCheckNameAvailabilityRequest = original.EnterpriseChannelCheckNameAvailabilityRequest
type EnterpriseChannelCheckNameAvailabilityResponse = original.EnterpriseChannelCheckNameAvailabilityResponse
type EnterpriseChannelNode = original.EnterpriseChannelNode
type EnterpriseChannelProperties = original.EnterpriseChannelProperties
type EnterpriseChannelResponseList = original.EnterpriseChannelResponseList
type EnterpriseChannelResponseListIterator = original.EnterpriseChannelResponseListIterator
type EnterpriseChannelResponseListPage = original.EnterpriseChannelResponseListPage
type EnterpriseChannelsClient = original.EnterpriseChannelsClient
type EnterpriseChannelsCreateFuture = original.EnterpriseChannelsCreateFuture
type EnterpriseChannelsDeleteFuture = original.EnterpriseChannelsDeleteFuture
type EnterpriseChannelsUpdateFuture = original.EnterpriseChannelsUpdateFuture
type Error = original.Error
type ErrorBody = original.ErrorBody
type FacebookChannel = original.FacebookChannel
type FacebookChannelProperties = original.FacebookChannelProperties
type FacebookPage = original.FacebookPage
type KikChannel = original.KikChannel
type KikChannelProperties = original.KikChannelProperties
type MsTeamsChannel = original.MsTeamsChannel
type MsTeamsChannelProperties = original.MsTeamsChannelProperties
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type ServiceProvider = original.ServiceProvider
type ServiceProviderParameter = original.ServiceProviderParameter
type ServiceProviderProperties = original.ServiceProviderProperties
type ServiceProviderResponseList = original.ServiceProviderResponseList
type Sku = original.Sku
type SkypeChannel = original.SkypeChannel
type SkypeChannelProperties = original.SkypeChannelProperties
type SlackChannel = original.SlackChannel
type SlackChannelProperties = original.SlackChannelProperties
type SmsChannel = original.SmsChannel
type SmsChannelProperties = original.SmsChannelProperties
type TelegramChannel = original.TelegramChannel
type TelegramChannelProperties = original.TelegramChannelProperties
type WebChatChannel = original.WebChatChannel
type WebChatChannelProperties = original.WebChatChannelProperties
type WebChatSite = original.WebChatSite

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewBotConnectionClient(subscriptionID string) BotConnectionClient {
	return original.NewBotConnectionClient(subscriptionID)
}
func NewBotConnectionClientWithBaseURI(baseURI string, subscriptionID string) BotConnectionClient {
	return original.NewBotConnectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewBotResponseListIterator(page BotResponseListPage) BotResponseListIterator {
	return original.NewBotResponseListIterator(page)
}
func NewBotResponseListPage(getNextPage func(context.Context, BotResponseList) (BotResponseList, error)) BotResponseListPage {
	return original.NewBotResponseListPage(getNextPage)
}
func NewBotsClient(subscriptionID string) BotsClient {
	return original.NewBotsClient(subscriptionID)
}
func NewBotsClientWithBaseURI(baseURI string, subscriptionID string) BotsClient {
	return original.NewBotsClientWithBaseURI(baseURI, subscriptionID)
}
func NewChannelResponseListIterator(page ChannelResponseListPage) ChannelResponseListIterator {
	return original.NewChannelResponseListIterator(page)
}
func NewChannelResponseListPage(getNextPage func(context.Context, ChannelResponseList) (ChannelResponseList, error)) ChannelResponseListPage {
	return original.NewChannelResponseListPage(getNextPage)
}
func NewChannelsClient(subscriptionID string) ChannelsClient {
	return original.NewChannelsClient(subscriptionID)
}
func NewChannelsClientWithBaseURI(baseURI string, subscriptionID string) ChannelsClient {
	return original.NewChannelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectionSettingResponseListIterator(page ConnectionSettingResponseListPage) ConnectionSettingResponseListIterator {
	return original.NewConnectionSettingResponseListIterator(page)
}
func NewConnectionSettingResponseListPage(getNextPage func(context.Context, ConnectionSettingResponseList) (ConnectionSettingResponseList, error)) ConnectionSettingResponseListPage {
	return original.NewConnectionSettingResponseListPage(getNextPage)
}
func NewEnterpriseChannelResponseListIterator(page EnterpriseChannelResponseListPage) EnterpriseChannelResponseListIterator {
	return original.NewEnterpriseChannelResponseListIterator(page)
}
func NewEnterpriseChannelResponseListPage(getNextPage func(context.Context, EnterpriseChannelResponseList) (EnterpriseChannelResponseList, error)) EnterpriseChannelResponseListPage {
	return original.NewEnterpriseChannelResponseListPage(getNextPage)
}
func NewEnterpriseChannelsClient(subscriptionID string) EnterpriseChannelsClient {
	return original.NewEnterpriseChannelsClient(subscriptionID)
}
func NewEnterpriseChannelsClientWithBaseURI(baseURI string, subscriptionID string) EnterpriseChannelsClient {
	return original.NewEnterpriseChannelsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleChannelNameBasicChannelValues() []ChannelNameBasicChannel {
	return original.PossibleChannelNameBasicChannelValues()
}
func PossibleChannelNameValues() []ChannelName {
	return original.PossibleChannelNameValues()
}
func PossibleEnterpriseChannelNodeStateValues() []EnterpriseChannelNodeState {
	return original.PossibleEnterpriseChannelNodeStateValues()
}
func PossibleEnterpriseChannelStateValues() []EnterpriseChannelState {
	return original.PossibleEnterpriseChannelStateValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
