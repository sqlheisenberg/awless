/* Copyright 2017 WALLIX

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package awsspec

import (
	awssdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/wallix/awless/logger"
)

type CreateTopic struct {
	_      string `action:"create" entity:"topic" awsAPI:"sns" awsCall:"CreateTopic" awsInput:"sns.CreateTopicInput" awsOutput:"sns.CreateTopicOutput"`
	logger *logger.Logger
	api    snsiface.SNSAPI
	Name   *string `awsName:"Name" awsType:"awsstr" templateName:"name" required:""`
}

func (cmd *CreateTopic) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}

func (cmd *CreateTopic) ExtractResult(i interface{}) string {
	return awssdk.StringValue(i.(*sns.CreateTopicOutput).TopicArn)
}

type DeleteTopic struct {
	_      string `action:"delete" entity:"topic" awsAPI:"sns" awsCall:"DeleteTopic" awsInput:"sns.DeleteTopicInput" awsOutput:"sns.DeleteTopicOutput"`
	logger *logger.Logger
	api    snsiface.SNSAPI
	Id     *string `awsName:"TopicArn" awsType:"awsstr" templateName:"id" required:""`
}

func (cmd *DeleteTopic) ValidateParams(params []string) ([]string, error) {
	return validateParams(cmd, params)
}
