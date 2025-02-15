// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package table

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/dynamodb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/dynamodb-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.DynamoDB{}
	_ = &svcapitypes.Table{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeTableOutput
	resp, err = rm.sdkapi.DescribeTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeTable", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Table.ArchivalSummary != nil {
		f0 := &svcapitypes.ArchivalSummary{}
		if resp.Table.ArchivalSummary.ArchivalBackupArn != nil {
			f0.ArchivalBackupARN = resp.Table.ArchivalSummary.ArchivalBackupArn
		}
		if resp.Table.ArchivalSummary.ArchivalDateTime != nil {
			f0.ArchivalDateTime = &metav1.Time{*resp.Table.ArchivalSummary.ArchivalDateTime}
		}
		if resp.Table.ArchivalSummary.ArchivalReason != nil {
			f0.ArchivalReason = resp.Table.ArchivalSummary.ArchivalReason
		}
		ko.Status.ArchivalSummary = f0
	} else {
		ko.Status.ArchivalSummary = nil
	}
	if resp.Table.AttributeDefinitions != nil {
		f1 := []*svcapitypes.AttributeDefinition{}
		for _, f1iter := range resp.Table.AttributeDefinitions {
			f1elem := &svcapitypes.AttributeDefinition{}
			if f1iter.AttributeName != nil {
				f1elem.AttributeName = f1iter.AttributeName
			}
			if f1iter.AttributeType != nil {
				f1elem.AttributeType = f1iter.AttributeType
			}
			f1 = append(f1, f1elem)
		}
		ko.Spec.AttributeDefinitions = f1
	} else {
		ko.Spec.AttributeDefinitions = nil
	}
	if resp.Table.BillingModeSummary != nil {
		f2 := &svcapitypes.BillingModeSummary{}
		if resp.Table.BillingModeSummary.BillingMode != nil {
			f2.BillingMode = resp.Table.BillingModeSummary.BillingMode
		}
		if resp.Table.BillingModeSummary.LastUpdateToPayPerRequestDateTime != nil {
			f2.LastUpdateToPayPerRequestDateTime = &metav1.Time{*resp.Table.BillingModeSummary.LastUpdateToPayPerRequestDateTime}
		}
		ko.Status.BillingModeSummary = f2
	} else {
		ko.Status.BillingModeSummary = nil
	}
	if resp.Table.CreationDateTime != nil {
		ko.Status.CreationDateTime = &metav1.Time{*resp.Table.CreationDateTime}
	} else {
		ko.Status.CreationDateTime = nil
	}
	if resp.Table.GlobalSecondaryIndexes != nil {
		f4 := []*svcapitypes.GlobalSecondaryIndex{}
		for _, f4iter := range resp.Table.GlobalSecondaryIndexes {
			f4elem := &svcapitypes.GlobalSecondaryIndex{}
			if f4iter.IndexName != nil {
				f4elem.IndexName = f4iter.IndexName
			}
			if f4iter.KeySchema != nil {
				f4elemf6 := []*svcapitypes.KeySchemaElement{}
				for _, f4elemf6iter := range f4iter.KeySchema {
					f4elemf6elem := &svcapitypes.KeySchemaElement{}
					if f4elemf6iter.AttributeName != nil {
						f4elemf6elem.AttributeName = f4elemf6iter.AttributeName
					}
					if f4elemf6iter.KeyType != nil {
						f4elemf6elem.KeyType = f4elemf6iter.KeyType
					}
					f4elemf6 = append(f4elemf6, f4elemf6elem)
				}
				f4elem.KeySchema = f4elemf6
			}
			if f4iter.Projection != nil {
				f4elemf7 := &svcapitypes.Projection{}
				if f4iter.Projection.NonKeyAttributes != nil {
					f4elemf7f0 := []*string{}
					for _, f4elemf7f0iter := range f4iter.Projection.NonKeyAttributes {
						var f4elemf7f0elem string
						f4elemf7f0elem = *f4elemf7f0iter
						f4elemf7f0 = append(f4elemf7f0, &f4elemf7f0elem)
					}
					f4elemf7.NonKeyAttributes = f4elemf7f0
				}
				if f4iter.Projection.ProjectionType != nil {
					f4elemf7.ProjectionType = f4iter.Projection.ProjectionType
				}
				f4elem.Projection = f4elemf7
			}
			if f4iter.ProvisionedThroughput != nil {
				f4elemf8 := &svcapitypes.ProvisionedThroughput{}
				if f4iter.ProvisionedThroughput.ReadCapacityUnits != nil {
					f4elemf8.ReadCapacityUnits = f4iter.ProvisionedThroughput.ReadCapacityUnits
				}
				if f4iter.ProvisionedThroughput.WriteCapacityUnits != nil {
					f4elemf8.WriteCapacityUnits = f4iter.ProvisionedThroughput.WriteCapacityUnits
				}
				f4elem.ProvisionedThroughput = f4elemf8
			}
			f4 = append(f4, f4elem)
		}
		ko.Spec.GlobalSecondaryIndexes = f4
	} else {
		ko.Spec.GlobalSecondaryIndexes = nil
	}
	if resp.Table.GlobalTableVersion != nil {
		ko.Status.GlobalTableVersion = resp.Table.GlobalTableVersion
	} else {
		ko.Status.GlobalTableVersion = nil
	}
	if resp.Table.ItemCount != nil {
		ko.Status.ItemCount = resp.Table.ItemCount
	} else {
		ko.Status.ItemCount = nil
	}
	if resp.Table.KeySchema != nil {
		f7 := []*svcapitypes.KeySchemaElement{}
		for _, f7iter := range resp.Table.KeySchema {
			f7elem := &svcapitypes.KeySchemaElement{}
			if f7iter.AttributeName != nil {
				f7elem.AttributeName = f7iter.AttributeName
			}
			if f7iter.KeyType != nil {
				f7elem.KeyType = f7iter.KeyType
			}
			f7 = append(f7, f7elem)
		}
		ko.Spec.KeySchema = f7
	} else {
		ko.Spec.KeySchema = nil
	}
	if resp.Table.LatestStreamArn != nil {
		ko.Status.LatestStreamARN = resp.Table.LatestStreamArn
	} else {
		ko.Status.LatestStreamARN = nil
	}
	if resp.Table.LatestStreamLabel != nil {
		ko.Status.LatestStreamLabel = resp.Table.LatestStreamLabel
	} else {
		ko.Status.LatestStreamLabel = nil
	}
	if resp.Table.LocalSecondaryIndexes != nil {
		f10 := []*svcapitypes.LocalSecondaryIndex{}
		for _, f10iter := range resp.Table.LocalSecondaryIndexes {
			f10elem := &svcapitypes.LocalSecondaryIndex{}
			if f10iter.IndexName != nil {
				f10elem.IndexName = f10iter.IndexName
			}
			if f10iter.KeySchema != nil {
				f10elemf4 := []*svcapitypes.KeySchemaElement{}
				for _, f10elemf4iter := range f10iter.KeySchema {
					f10elemf4elem := &svcapitypes.KeySchemaElement{}
					if f10elemf4iter.AttributeName != nil {
						f10elemf4elem.AttributeName = f10elemf4iter.AttributeName
					}
					if f10elemf4iter.KeyType != nil {
						f10elemf4elem.KeyType = f10elemf4iter.KeyType
					}
					f10elemf4 = append(f10elemf4, f10elemf4elem)
				}
				f10elem.KeySchema = f10elemf4
			}
			if f10iter.Projection != nil {
				f10elemf5 := &svcapitypes.Projection{}
				if f10iter.Projection.NonKeyAttributes != nil {
					f10elemf5f0 := []*string{}
					for _, f10elemf5f0iter := range f10iter.Projection.NonKeyAttributes {
						var f10elemf5f0elem string
						f10elemf5f0elem = *f10elemf5f0iter
						f10elemf5f0 = append(f10elemf5f0, &f10elemf5f0elem)
					}
					f10elemf5.NonKeyAttributes = f10elemf5f0
				}
				if f10iter.Projection.ProjectionType != nil {
					f10elemf5.ProjectionType = f10iter.Projection.ProjectionType
				}
				f10elem.Projection = f10elemf5
			}
			f10 = append(f10, f10elem)
		}
		ko.Spec.LocalSecondaryIndexes = f10
	} else {
		ko.Spec.LocalSecondaryIndexes = nil
	}
	if resp.Table.ProvisionedThroughput != nil {
		f11 := &svcapitypes.ProvisionedThroughput{}
		if resp.Table.ProvisionedThroughput.ReadCapacityUnits != nil {
			f11.ReadCapacityUnits = resp.Table.ProvisionedThroughput.ReadCapacityUnits
		}
		if resp.Table.ProvisionedThroughput.WriteCapacityUnits != nil {
			f11.WriteCapacityUnits = resp.Table.ProvisionedThroughput.WriteCapacityUnits
		}
		ko.Spec.ProvisionedThroughput = f11
	} else {
		ko.Spec.ProvisionedThroughput = nil
	}
	if resp.Table.Replicas != nil {
		f12 := []*svcapitypes.ReplicaDescription{}
		for _, f12iter := range resp.Table.Replicas {
			f12elem := &svcapitypes.ReplicaDescription{}
			if f12iter.GlobalSecondaryIndexes != nil {
				f12elemf0 := []*svcapitypes.ReplicaGlobalSecondaryIndexDescription{}
				for _, f12elemf0iter := range f12iter.GlobalSecondaryIndexes {
					f12elemf0elem := &svcapitypes.ReplicaGlobalSecondaryIndexDescription{}
					if f12elemf0iter.IndexName != nil {
						f12elemf0elem.IndexName = f12elemf0iter.IndexName
					}
					if f12elemf0iter.ProvisionedThroughputOverride != nil {
						f12elemf0elemf1 := &svcapitypes.ProvisionedThroughputOverride{}
						if f12elemf0iter.ProvisionedThroughputOverride.ReadCapacityUnits != nil {
							f12elemf0elemf1.ReadCapacityUnits = f12elemf0iter.ProvisionedThroughputOverride.ReadCapacityUnits
						}
						f12elemf0elem.ProvisionedThroughputOverride = f12elemf0elemf1
					}
					f12elemf0 = append(f12elemf0, f12elemf0elem)
				}
				f12elem.GlobalSecondaryIndexes = f12elemf0
			}
			if f12iter.KMSMasterKeyId != nil {
				f12elem.KMSMasterKeyID = f12iter.KMSMasterKeyId
			}
			if f12iter.ProvisionedThroughputOverride != nil {
				f12elemf2 := &svcapitypes.ProvisionedThroughputOverride{}
				if f12iter.ProvisionedThroughputOverride.ReadCapacityUnits != nil {
					f12elemf2.ReadCapacityUnits = f12iter.ProvisionedThroughputOverride.ReadCapacityUnits
				}
				f12elem.ProvisionedThroughputOverride = f12elemf2
			}
			if f12iter.RegionName != nil {
				f12elem.RegionName = f12iter.RegionName
			}
			if f12iter.ReplicaInaccessibleDateTime != nil {
				f12elem.ReplicaInaccessibleDateTime = &metav1.Time{*f12iter.ReplicaInaccessibleDateTime}
			}
			if f12iter.ReplicaStatus != nil {
				f12elem.ReplicaStatus = f12iter.ReplicaStatus
			}
			if f12iter.ReplicaStatusDescription != nil {
				f12elem.ReplicaStatusDescription = f12iter.ReplicaStatusDescription
			}
			if f12iter.ReplicaStatusPercentProgress != nil {
				f12elem.ReplicaStatusPercentProgress = f12iter.ReplicaStatusPercentProgress
			}
			f12 = append(f12, f12elem)
		}
		ko.Status.Replicas = f12
	} else {
		ko.Status.Replicas = nil
	}
	if resp.Table.RestoreSummary != nil {
		f13 := &svcapitypes.RestoreSummary{}
		if resp.Table.RestoreSummary.RestoreDateTime != nil {
			f13.RestoreDateTime = &metav1.Time{*resp.Table.RestoreSummary.RestoreDateTime}
		}
		if resp.Table.RestoreSummary.RestoreInProgress != nil {
			f13.RestoreInProgress = resp.Table.RestoreSummary.RestoreInProgress
		}
		if resp.Table.RestoreSummary.SourceBackupArn != nil {
			f13.SourceBackupARN = resp.Table.RestoreSummary.SourceBackupArn
		}
		if resp.Table.RestoreSummary.SourceTableArn != nil {
			f13.SourceTableARN = resp.Table.RestoreSummary.SourceTableArn
		}
		ko.Status.RestoreSummary = f13
	} else {
		ko.Status.RestoreSummary = nil
	}
	if resp.Table.SSEDescription != nil {
		f14 := &svcapitypes.SSEDescription{}
		if resp.Table.SSEDescription.InaccessibleEncryptionDateTime != nil {
			f14.InaccessibleEncryptionDateTime = &metav1.Time{*resp.Table.SSEDescription.InaccessibleEncryptionDateTime}
		}
		if resp.Table.SSEDescription.KMSMasterKeyArn != nil {
			f14.KMSMasterKeyARN = resp.Table.SSEDescription.KMSMasterKeyArn
		}
		if resp.Table.SSEDescription.SSEType != nil {
			f14.SSEType = resp.Table.SSEDescription.SSEType
		}
		if resp.Table.SSEDescription.Status != nil {
			f14.Status = resp.Table.SSEDescription.Status
		}
		ko.Status.SSEDescription = f14
	} else {
		ko.Status.SSEDescription = nil
	}
	if resp.Table.StreamSpecification != nil {
		f15 := &svcapitypes.StreamSpecification{}
		if resp.Table.StreamSpecification.StreamEnabled != nil {
			f15.StreamEnabled = resp.Table.StreamSpecification.StreamEnabled
		}
		if resp.Table.StreamSpecification.StreamViewType != nil {
			f15.StreamViewType = resp.Table.StreamSpecification.StreamViewType
		}
		ko.Spec.StreamSpecification = f15
	} else {
		ko.Spec.StreamSpecification = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.Table.TableArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.Table.TableArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.Table.TableId != nil {
		ko.Status.TableID = resp.Table.TableId
	} else {
		ko.Status.TableID = nil
	}
	if resp.Table.TableName != nil {
		ko.Spec.TableName = resp.Table.TableName
	} else {
		ko.Spec.TableName = nil
	}
	if resp.Table.TableSizeBytes != nil {
		ko.Status.TableSizeBytes = resp.Table.TableSizeBytes
	} else {
		ko.Status.TableSizeBytes = nil
	}
	if resp.Table.TableStatus != nil {
		ko.Status.TableStatus = resp.Table.TableStatus
	} else {
		ko.Status.TableStatus = nil
	}

	rm.setStatusDefaults(ko)
	if isTableCreating(&resource{ko}) {
		return &resource{ko}, requeueWaitWhileCreating
	}
	if isTableUpdating(&resource{ko}) {
		return &resource{ko}, requeueWaitWhileUpdating
	}
	if err := rm.setResourceAdditionalFields(ctx, ko); err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.TableName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeTableInput, error) {
	res := &svcsdk.DescribeTableInput{}

	if r.ko.Spec.TableName != nil {
		res.SetTableName(*r.ko.Spec.TableName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateTableOutput
	_ = resp
	resp, err = rm.sdkapi.CreateTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateTable", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.TableDescription.ArchivalSummary != nil {
		f0 := &svcapitypes.ArchivalSummary{}
		if resp.TableDescription.ArchivalSummary.ArchivalBackupArn != nil {
			f0.ArchivalBackupARN = resp.TableDescription.ArchivalSummary.ArchivalBackupArn
		}
		if resp.TableDescription.ArchivalSummary.ArchivalDateTime != nil {
			f0.ArchivalDateTime = &metav1.Time{*resp.TableDescription.ArchivalSummary.ArchivalDateTime}
		}
		if resp.TableDescription.ArchivalSummary.ArchivalReason != nil {
			f0.ArchivalReason = resp.TableDescription.ArchivalSummary.ArchivalReason
		}
		ko.Status.ArchivalSummary = f0
	} else {
		ko.Status.ArchivalSummary = nil
	}
	if resp.TableDescription.AttributeDefinitions != nil {
		f1 := []*svcapitypes.AttributeDefinition{}
		for _, f1iter := range resp.TableDescription.AttributeDefinitions {
			f1elem := &svcapitypes.AttributeDefinition{}
			if f1iter.AttributeName != nil {
				f1elem.AttributeName = f1iter.AttributeName
			}
			if f1iter.AttributeType != nil {
				f1elem.AttributeType = f1iter.AttributeType
			}
			f1 = append(f1, f1elem)
		}
		ko.Spec.AttributeDefinitions = f1
	} else {
		ko.Spec.AttributeDefinitions = nil
	}
	if resp.TableDescription.BillingModeSummary != nil {
		f2 := &svcapitypes.BillingModeSummary{}
		if resp.TableDescription.BillingModeSummary.BillingMode != nil {
			f2.BillingMode = resp.TableDescription.BillingModeSummary.BillingMode
		}
		if resp.TableDescription.BillingModeSummary.LastUpdateToPayPerRequestDateTime != nil {
			f2.LastUpdateToPayPerRequestDateTime = &metav1.Time{*resp.TableDescription.BillingModeSummary.LastUpdateToPayPerRequestDateTime}
		}
		ko.Status.BillingModeSummary = f2
	} else {
		ko.Status.BillingModeSummary = nil
	}
	if resp.TableDescription.CreationDateTime != nil {
		ko.Status.CreationDateTime = &metav1.Time{*resp.TableDescription.CreationDateTime}
	} else {
		ko.Status.CreationDateTime = nil
	}
	if resp.TableDescription.GlobalSecondaryIndexes != nil {
		f4 := []*svcapitypes.GlobalSecondaryIndex{}
		for _, f4iter := range resp.TableDescription.GlobalSecondaryIndexes {
			f4elem := &svcapitypes.GlobalSecondaryIndex{}
			if f4iter.IndexName != nil {
				f4elem.IndexName = f4iter.IndexName
			}
			if f4iter.KeySchema != nil {
				f4elemf6 := []*svcapitypes.KeySchemaElement{}
				for _, f4elemf6iter := range f4iter.KeySchema {
					f4elemf6elem := &svcapitypes.KeySchemaElement{}
					if f4elemf6iter.AttributeName != nil {
						f4elemf6elem.AttributeName = f4elemf6iter.AttributeName
					}
					if f4elemf6iter.KeyType != nil {
						f4elemf6elem.KeyType = f4elemf6iter.KeyType
					}
					f4elemf6 = append(f4elemf6, f4elemf6elem)
				}
				f4elem.KeySchema = f4elemf6
			}
			if f4iter.Projection != nil {
				f4elemf7 := &svcapitypes.Projection{}
				if f4iter.Projection.NonKeyAttributes != nil {
					f4elemf7f0 := []*string{}
					for _, f4elemf7f0iter := range f4iter.Projection.NonKeyAttributes {
						var f4elemf7f0elem string
						f4elemf7f0elem = *f4elemf7f0iter
						f4elemf7f0 = append(f4elemf7f0, &f4elemf7f0elem)
					}
					f4elemf7.NonKeyAttributes = f4elemf7f0
				}
				if f4iter.Projection.ProjectionType != nil {
					f4elemf7.ProjectionType = f4iter.Projection.ProjectionType
				}
				f4elem.Projection = f4elemf7
			}
			if f4iter.ProvisionedThroughput != nil {
				f4elemf8 := &svcapitypes.ProvisionedThroughput{}
				if f4iter.ProvisionedThroughput.ReadCapacityUnits != nil {
					f4elemf8.ReadCapacityUnits = f4iter.ProvisionedThroughput.ReadCapacityUnits
				}
				if f4iter.ProvisionedThroughput.WriteCapacityUnits != nil {
					f4elemf8.WriteCapacityUnits = f4iter.ProvisionedThroughput.WriteCapacityUnits
				}
				f4elem.ProvisionedThroughput = f4elemf8
			}
			f4 = append(f4, f4elem)
		}
		ko.Spec.GlobalSecondaryIndexes = f4
	} else {
		ko.Spec.GlobalSecondaryIndexes = nil
	}
	if resp.TableDescription.GlobalTableVersion != nil {
		ko.Status.GlobalTableVersion = resp.TableDescription.GlobalTableVersion
	} else {
		ko.Status.GlobalTableVersion = nil
	}
	if resp.TableDescription.ItemCount != nil {
		ko.Status.ItemCount = resp.TableDescription.ItemCount
	} else {
		ko.Status.ItemCount = nil
	}
	if resp.TableDescription.KeySchema != nil {
		f7 := []*svcapitypes.KeySchemaElement{}
		for _, f7iter := range resp.TableDescription.KeySchema {
			f7elem := &svcapitypes.KeySchemaElement{}
			if f7iter.AttributeName != nil {
				f7elem.AttributeName = f7iter.AttributeName
			}
			if f7iter.KeyType != nil {
				f7elem.KeyType = f7iter.KeyType
			}
			f7 = append(f7, f7elem)
		}
		ko.Spec.KeySchema = f7
	} else {
		ko.Spec.KeySchema = nil
	}
	if resp.TableDescription.LatestStreamArn != nil {
		ko.Status.LatestStreamARN = resp.TableDescription.LatestStreamArn
	} else {
		ko.Status.LatestStreamARN = nil
	}
	if resp.TableDescription.LatestStreamLabel != nil {
		ko.Status.LatestStreamLabel = resp.TableDescription.LatestStreamLabel
	} else {
		ko.Status.LatestStreamLabel = nil
	}
	if resp.TableDescription.LocalSecondaryIndexes != nil {
		f10 := []*svcapitypes.LocalSecondaryIndex{}
		for _, f10iter := range resp.TableDescription.LocalSecondaryIndexes {
			f10elem := &svcapitypes.LocalSecondaryIndex{}
			if f10iter.IndexName != nil {
				f10elem.IndexName = f10iter.IndexName
			}
			if f10iter.KeySchema != nil {
				f10elemf4 := []*svcapitypes.KeySchemaElement{}
				for _, f10elemf4iter := range f10iter.KeySchema {
					f10elemf4elem := &svcapitypes.KeySchemaElement{}
					if f10elemf4iter.AttributeName != nil {
						f10elemf4elem.AttributeName = f10elemf4iter.AttributeName
					}
					if f10elemf4iter.KeyType != nil {
						f10elemf4elem.KeyType = f10elemf4iter.KeyType
					}
					f10elemf4 = append(f10elemf4, f10elemf4elem)
				}
				f10elem.KeySchema = f10elemf4
			}
			if f10iter.Projection != nil {
				f10elemf5 := &svcapitypes.Projection{}
				if f10iter.Projection.NonKeyAttributes != nil {
					f10elemf5f0 := []*string{}
					for _, f10elemf5f0iter := range f10iter.Projection.NonKeyAttributes {
						var f10elemf5f0elem string
						f10elemf5f0elem = *f10elemf5f0iter
						f10elemf5f0 = append(f10elemf5f0, &f10elemf5f0elem)
					}
					f10elemf5.NonKeyAttributes = f10elemf5f0
				}
				if f10iter.Projection.ProjectionType != nil {
					f10elemf5.ProjectionType = f10iter.Projection.ProjectionType
				}
				f10elem.Projection = f10elemf5
			}
			f10 = append(f10, f10elem)
		}
		ko.Spec.LocalSecondaryIndexes = f10
	} else {
		ko.Spec.LocalSecondaryIndexes = nil
	}
	if resp.TableDescription.ProvisionedThroughput != nil {
		f11 := &svcapitypes.ProvisionedThroughput{}
		if resp.TableDescription.ProvisionedThroughput.ReadCapacityUnits != nil {
			f11.ReadCapacityUnits = resp.TableDescription.ProvisionedThroughput.ReadCapacityUnits
		}
		if resp.TableDescription.ProvisionedThroughput.WriteCapacityUnits != nil {
			f11.WriteCapacityUnits = resp.TableDescription.ProvisionedThroughput.WriteCapacityUnits
		}
		ko.Spec.ProvisionedThroughput = f11
	} else {
		ko.Spec.ProvisionedThroughput = nil
	}
	if resp.TableDescription.Replicas != nil {
		f12 := []*svcapitypes.ReplicaDescription{}
		for _, f12iter := range resp.TableDescription.Replicas {
			f12elem := &svcapitypes.ReplicaDescription{}
			if f12iter.GlobalSecondaryIndexes != nil {
				f12elemf0 := []*svcapitypes.ReplicaGlobalSecondaryIndexDescription{}
				for _, f12elemf0iter := range f12iter.GlobalSecondaryIndexes {
					f12elemf0elem := &svcapitypes.ReplicaGlobalSecondaryIndexDescription{}
					if f12elemf0iter.IndexName != nil {
						f12elemf0elem.IndexName = f12elemf0iter.IndexName
					}
					if f12elemf0iter.ProvisionedThroughputOverride != nil {
						f12elemf0elemf1 := &svcapitypes.ProvisionedThroughputOverride{}
						if f12elemf0iter.ProvisionedThroughputOverride.ReadCapacityUnits != nil {
							f12elemf0elemf1.ReadCapacityUnits = f12elemf0iter.ProvisionedThroughputOverride.ReadCapacityUnits
						}
						f12elemf0elem.ProvisionedThroughputOverride = f12elemf0elemf1
					}
					f12elemf0 = append(f12elemf0, f12elemf0elem)
				}
				f12elem.GlobalSecondaryIndexes = f12elemf0
			}
			if f12iter.KMSMasterKeyId != nil {
				f12elem.KMSMasterKeyID = f12iter.KMSMasterKeyId
			}
			if f12iter.ProvisionedThroughputOverride != nil {
				f12elemf2 := &svcapitypes.ProvisionedThroughputOverride{}
				if f12iter.ProvisionedThroughputOverride.ReadCapacityUnits != nil {
					f12elemf2.ReadCapacityUnits = f12iter.ProvisionedThroughputOverride.ReadCapacityUnits
				}
				f12elem.ProvisionedThroughputOverride = f12elemf2
			}
			if f12iter.RegionName != nil {
				f12elem.RegionName = f12iter.RegionName
			}
			if f12iter.ReplicaInaccessibleDateTime != nil {
				f12elem.ReplicaInaccessibleDateTime = &metav1.Time{*f12iter.ReplicaInaccessibleDateTime}
			}
			if f12iter.ReplicaStatus != nil {
				f12elem.ReplicaStatus = f12iter.ReplicaStatus
			}
			if f12iter.ReplicaStatusDescription != nil {
				f12elem.ReplicaStatusDescription = f12iter.ReplicaStatusDescription
			}
			if f12iter.ReplicaStatusPercentProgress != nil {
				f12elem.ReplicaStatusPercentProgress = f12iter.ReplicaStatusPercentProgress
			}
			f12 = append(f12, f12elem)
		}
		ko.Status.Replicas = f12
	} else {
		ko.Status.Replicas = nil
	}
	if resp.TableDescription.RestoreSummary != nil {
		f13 := &svcapitypes.RestoreSummary{}
		if resp.TableDescription.RestoreSummary.RestoreDateTime != nil {
			f13.RestoreDateTime = &metav1.Time{*resp.TableDescription.RestoreSummary.RestoreDateTime}
		}
		if resp.TableDescription.RestoreSummary.RestoreInProgress != nil {
			f13.RestoreInProgress = resp.TableDescription.RestoreSummary.RestoreInProgress
		}
		if resp.TableDescription.RestoreSummary.SourceBackupArn != nil {
			f13.SourceBackupARN = resp.TableDescription.RestoreSummary.SourceBackupArn
		}
		if resp.TableDescription.RestoreSummary.SourceTableArn != nil {
			f13.SourceTableARN = resp.TableDescription.RestoreSummary.SourceTableArn
		}
		ko.Status.RestoreSummary = f13
	} else {
		ko.Status.RestoreSummary = nil
	}
	if resp.TableDescription.SSEDescription != nil {
		f14 := &svcapitypes.SSEDescription{}
		if resp.TableDescription.SSEDescription.InaccessibleEncryptionDateTime != nil {
			f14.InaccessibleEncryptionDateTime = &metav1.Time{*resp.TableDescription.SSEDescription.InaccessibleEncryptionDateTime}
		}
		if resp.TableDescription.SSEDescription.KMSMasterKeyArn != nil {
			f14.KMSMasterKeyARN = resp.TableDescription.SSEDescription.KMSMasterKeyArn
		}
		if resp.TableDescription.SSEDescription.SSEType != nil {
			f14.SSEType = resp.TableDescription.SSEDescription.SSEType
		}
		if resp.TableDescription.SSEDescription.Status != nil {
			f14.Status = resp.TableDescription.SSEDescription.Status
		}
		ko.Status.SSEDescription = f14
	} else {
		ko.Status.SSEDescription = nil
	}
	if resp.TableDescription.StreamSpecification != nil {
		f15 := &svcapitypes.StreamSpecification{}
		if resp.TableDescription.StreamSpecification.StreamEnabled != nil {
			f15.StreamEnabled = resp.TableDescription.StreamSpecification.StreamEnabled
		}
		if resp.TableDescription.StreamSpecification.StreamViewType != nil {
			f15.StreamViewType = resp.TableDescription.StreamSpecification.StreamViewType
		}
		ko.Spec.StreamSpecification = f15
	} else {
		ko.Spec.StreamSpecification = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.TableDescription.TableArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.TableDescription.TableArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.TableDescription.TableId != nil {
		ko.Status.TableID = resp.TableDescription.TableId
	} else {
		ko.Status.TableID = nil
	}
	if resp.TableDescription.TableName != nil {
		ko.Spec.TableName = resp.TableDescription.TableName
	} else {
		ko.Spec.TableName = nil
	}
	if resp.TableDescription.TableSizeBytes != nil {
		ko.Status.TableSizeBytes = resp.TableDescription.TableSizeBytes
	} else {
		ko.Status.TableSizeBytes = nil
	}
	if resp.TableDescription.TableStatus != nil {
		ko.Status.TableStatus = resp.TableDescription.TableStatus
	} else {
		ko.Status.TableStatus = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateTableInput, error) {
	res := &svcsdk.CreateTableInput{}

	if r.ko.Spec.AttributeDefinitions != nil {
		f0 := []*svcsdk.AttributeDefinition{}
		for _, f0iter := range r.ko.Spec.AttributeDefinitions {
			f0elem := &svcsdk.AttributeDefinition{}
			if f0iter.AttributeName != nil {
				f0elem.SetAttributeName(*f0iter.AttributeName)
			}
			if f0iter.AttributeType != nil {
				f0elem.SetAttributeType(*f0iter.AttributeType)
			}
			f0 = append(f0, f0elem)
		}
		res.SetAttributeDefinitions(f0)
	}
	if r.ko.Spec.BillingMode != nil {
		res.SetBillingMode(*r.ko.Spec.BillingMode)
	}
	if r.ko.Spec.GlobalSecondaryIndexes != nil {
		f2 := []*svcsdk.GlobalSecondaryIndex{}
		for _, f2iter := range r.ko.Spec.GlobalSecondaryIndexes {
			f2elem := &svcsdk.GlobalSecondaryIndex{}
			if f2iter.IndexName != nil {
				f2elem.SetIndexName(*f2iter.IndexName)
			}
			if f2iter.KeySchema != nil {
				f2elemf1 := []*svcsdk.KeySchemaElement{}
				for _, f2elemf1iter := range f2iter.KeySchema {
					f2elemf1elem := &svcsdk.KeySchemaElement{}
					if f2elemf1iter.AttributeName != nil {
						f2elemf1elem.SetAttributeName(*f2elemf1iter.AttributeName)
					}
					if f2elemf1iter.KeyType != nil {
						f2elemf1elem.SetKeyType(*f2elemf1iter.KeyType)
					}
					f2elemf1 = append(f2elemf1, f2elemf1elem)
				}
				f2elem.SetKeySchema(f2elemf1)
			}
			if f2iter.Projection != nil {
				f2elemf2 := &svcsdk.Projection{}
				if f2iter.Projection.NonKeyAttributes != nil {
					f2elemf2f0 := []*string{}
					for _, f2elemf2f0iter := range f2iter.Projection.NonKeyAttributes {
						var f2elemf2f0elem string
						f2elemf2f0elem = *f2elemf2f0iter
						f2elemf2f0 = append(f2elemf2f0, &f2elemf2f0elem)
					}
					f2elemf2.SetNonKeyAttributes(f2elemf2f0)
				}
				if f2iter.Projection.ProjectionType != nil {
					f2elemf2.SetProjectionType(*f2iter.Projection.ProjectionType)
				}
				f2elem.SetProjection(f2elemf2)
			}
			if f2iter.ProvisionedThroughput != nil {
				f2elemf3 := &svcsdk.ProvisionedThroughput{}
				if f2iter.ProvisionedThroughput.ReadCapacityUnits != nil {
					f2elemf3.SetReadCapacityUnits(*f2iter.ProvisionedThroughput.ReadCapacityUnits)
				}
				if f2iter.ProvisionedThroughput.WriteCapacityUnits != nil {
					f2elemf3.SetWriteCapacityUnits(*f2iter.ProvisionedThroughput.WriteCapacityUnits)
				}
				f2elem.SetProvisionedThroughput(f2elemf3)
			}
			f2 = append(f2, f2elem)
		}
		res.SetGlobalSecondaryIndexes(f2)
	}
	if r.ko.Spec.KeySchema != nil {
		f3 := []*svcsdk.KeySchemaElement{}
		for _, f3iter := range r.ko.Spec.KeySchema {
			f3elem := &svcsdk.KeySchemaElement{}
			if f3iter.AttributeName != nil {
				f3elem.SetAttributeName(*f3iter.AttributeName)
			}
			if f3iter.KeyType != nil {
				f3elem.SetKeyType(*f3iter.KeyType)
			}
			f3 = append(f3, f3elem)
		}
		res.SetKeySchema(f3)
	}
	if r.ko.Spec.LocalSecondaryIndexes != nil {
		f4 := []*svcsdk.LocalSecondaryIndex{}
		for _, f4iter := range r.ko.Spec.LocalSecondaryIndexes {
			f4elem := &svcsdk.LocalSecondaryIndex{}
			if f4iter.IndexName != nil {
				f4elem.SetIndexName(*f4iter.IndexName)
			}
			if f4iter.KeySchema != nil {
				f4elemf1 := []*svcsdk.KeySchemaElement{}
				for _, f4elemf1iter := range f4iter.KeySchema {
					f4elemf1elem := &svcsdk.KeySchemaElement{}
					if f4elemf1iter.AttributeName != nil {
						f4elemf1elem.SetAttributeName(*f4elemf1iter.AttributeName)
					}
					if f4elemf1iter.KeyType != nil {
						f4elemf1elem.SetKeyType(*f4elemf1iter.KeyType)
					}
					f4elemf1 = append(f4elemf1, f4elemf1elem)
				}
				f4elem.SetKeySchema(f4elemf1)
			}
			if f4iter.Projection != nil {
				f4elemf2 := &svcsdk.Projection{}
				if f4iter.Projection.NonKeyAttributes != nil {
					f4elemf2f0 := []*string{}
					for _, f4elemf2f0iter := range f4iter.Projection.NonKeyAttributes {
						var f4elemf2f0elem string
						f4elemf2f0elem = *f4elemf2f0iter
						f4elemf2f0 = append(f4elemf2f0, &f4elemf2f0elem)
					}
					f4elemf2.SetNonKeyAttributes(f4elemf2f0)
				}
				if f4iter.Projection.ProjectionType != nil {
					f4elemf2.SetProjectionType(*f4iter.Projection.ProjectionType)
				}
				f4elem.SetProjection(f4elemf2)
			}
			f4 = append(f4, f4elem)
		}
		res.SetLocalSecondaryIndexes(f4)
	}
	if r.ko.Spec.ProvisionedThroughput != nil {
		f5 := &svcsdk.ProvisionedThroughput{}
		if r.ko.Spec.ProvisionedThroughput.ReadCapacityUnits != nil {
			f5.SetReadCapacityUnits(*r.ko.Spec.ProvisionedThroughput.ReadCapacityUnits)
		}
		if r.ko.Spec.ProvisionedThroughput.WriteCapacityUnits != nil {
			f5.SetWriteCapacityUnits(*r.ko.Spec.ProvisionedThroughput.WriteCapacityUnits)
		}
		res.SetProvisionedThroughput(f5)
	}
	if r.ko.Spec.SSESpecification != nil {
		f6 := &svcsdk.SSESpecification{}
		if r.ko.Spec.SSESpecification.Enabled != nil {
			f6.SetEnabled(*r.ko.Spec.SSESpecification.Enabled)
		}
		if r.ko.Spec.SSESpecification.KMSMasterKeyID != nil {
			f6.SetKMSMasterKeyId(*r.ko.Spec.SSESpecification.KMSMasterKeyID)
		}
		if r.ko.Spec.SSESpecification.SSEType != nil {
			f6.SetSSEType(*r.ko.Spec.SSESpecification.SSEType)
		}
		res.SetSSESpecification(f6)
	}
	if r.ko.Spec.StreamSpecification != nil {
		f7 := &svcsdk.StreamSpecification{}
		if r.ko.Spec.StreamSpecification.StreamEnabled != nil {
			f7.SetStreamEnabled(*r.ko.Spec.StreamSpecification.StreamEnabled)
		}
		if r.ko.Spec.StreamSpecification.StreamViewType != nil {
			f7.SetStreamViewType(*r.ko.Spec.StreamSpecification.StreamViewType)
		}
		res.SetStreamSpecification(f7)
	}
	if r.ko.Spec.TableName != nil {
		res.SetTableName(*r.ko.Spec.TableName)
	}
	if r.ko.Spec.Tags != nil {
		f9 := []*svcsdk.Tag{}
		for _, f9iter := range r.ko.Spec.Tags {
			f9elem := &svcsdk.Tag{}
			if f9iter.Key != nil {
				f9elem.SetKey(*f9iter.Key)
			}
			if f9iter.Value != nil {
				f9elem.SetValue(*f9iter.Value)
			}
			f9 = append(f9, f9elem)
		}
		res.SetTags(f9)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	return rm.customUpdateTable(ctx, desired, latest, delta)
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	if isTableDeleting(r) {
		return nil, requeueWaitWhileDeleting
	}
	if isTableUpdating(r) {
		return nil, requeueWaitWhileUpdating
	}
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteTableOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteTableWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteTable", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteTableInput, error) {
	res := &svcsdk.DeleteTableInput{}

	if r.ko.Spec.TableName != nil {
		res.SetTableName(*r.ko.Spec.TableName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Table,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InternalServerError",
		"LimitExceededException",
		"ResourceInUseException":
		return true
	default:
		return false
	}
}
