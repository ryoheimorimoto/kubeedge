/*
Copyright 2018 The KubeEdge Authors.

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

package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang/mock/gomock"
	coreV1 "k8s.io/api/core/v1"

	"github.com/kubeedge/api/apis/common/constants"
	"github.com/kubeedge/kubeedge/edge/mocks/beego"
	"github.com/kubeedge/kubeedge/edge/pkg/common/dbm"
)

// errFailedDBOperation is common DB operation fail error
var errFailedDBOperation = errors.New("Failed DB Operation")

// meta is global variable for passing as test parameter
var meta = Meta{
	Key:   "TestKey",
	Value: "TestValue",
	Type:  "TestType",
}

// TestSaveMeta is function to initialize all global variable and test SaveMeta
func TestSaveMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// returnInt is first return of mock interface ormerMock
		returnInt int64
		// returnErr is second return of mock interface ormerMock which is also expected error
		returnErr error
	}{{
		// Success Case
		name:      "SuccessCase",
		returnInt: int64(1),
		returnErr: nil,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnInt: int64(1),
		returnErr: errFailedDBOperation,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ormerMock.EXPECT().Insert(gomock.Any()).Return(test.returnInt, test.returnErr).Times(1)
			err := SaveMeta(&meta)
			if test.returnErr != err {
				t.Errorf("Save Meta Case failed : wanted error %v and got error %v", test.returnErr, err)
			}
		})
	}
}

// TestDeleteMetaByKey is function to test DeleteMetaByKey
func TestDeleteMetaByKey(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// deleteReturnInt is the first return of mock interface querySeterMock's delete function
		deleteReturnInt int64
		// deleteReturnErr is the second return of mock interface querySeterMocks's delete function also expected error
		deleteReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		deleteReturnInt:  int64(1),
		deleteReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		deleteReturnInt:  int64(0),
		deleteReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Delete().Return(test.deleteReturnInt, test.deleteReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := DeleteMetaByKey("test")
			if test.deleteReturnErr != err {
				t.Errorf("Delete Meta By Key Case failed : wanted %v and got %v", test.deleteReturnErr, err)
			}
		})
	}
}

// TestDeleteMetaByKeyAndPodUID is function to test DeleteMetaByKeyAndPodUID
func TestDeleteMetaByKeyAndPodUID(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock
	rawSeterMock := beego.NewMockRawSeter(mockCtrl)
	deleteRes := beego.NewMockDriverRes(mockCtrl)
	deleteRes.EXPECT().RowsAffected().Return(int64(1), nil).Times(1)

	cases := []struct {
		// name is name of the testcase
		name string
		// deleteReturnRes is first return of mock interface rawSeterMock's Exec function
		deleteReturnRes sql.Result
		// deleteReturnErr is second return of mock interface rawSeterMock's Exec function which is also expected error
		deleteReturnErr error
		// deleteReturnRaw is the return of mock interface ormerMock's Raw function
		deleteReturnRaw orm.RawSeter
	}{{
		// Success Case
		name:            "SuccessCase",
		deleteReturnRes: deleteRes,
		deleteReturnErr: nil,
		deleteReturnRaw: rawSeterMock,
	}, {
		// Failure Case
		name:            "FailureCase",
		deleteReturnRes: nil,
		deleteReturnErr: errFailedDBOperation,
		deleteReturnRaw: rawSeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			rawSeterMock.EXPECT().Exec().Return(test.deleteReturnRes, test.deleteReturnErr).Times(1)
			ormerMock.EXPECT().Raw(gomock.Any(), gomock.Any()).Return(test.deleteReturnRaw).Times(1)
			_, err := DeleteMetaByKeyAndPodUID("test", "testUID")
			if test.deleteReturnErr != err {
				t.Errorf("Delete Meta By Key Case failed : wanted %v and got %v", test.deleteReturnErr, err)
			}
		})
	}
}

// TestUpdateMeta is function to test UpdateMeta
func TestUpdateMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// returnInt is first return of mock interface ormerMock
		returnInt int64
		// returnErr is second return of mock interface ormerMock which is also expected error
		returnErr error
	}{{
		// Success Case
		name:      "SuccessCase",
		returnInt: int64(1),
		returnErr: nil,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnInt: int64(0),
		returnErr: errFailedDBOperation,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ormerMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(test.returnInt, test.returnErr).Times(1)
			err := UpdateMeta(&meta)
			if test.returnErr != err {
				t.Errorf("Update Meta Case failed : wanted %v and got %v", test.returnErr, err)
			}
		})
	}
}

// TestInsertOrUpdate is function to test InsertOrUpdate
func TestInsertOrUpdate(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	rawSeterMock := beego.NewMockRawSeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// returnSQL is first return of mock interface rawSeterMock's Exec function
		returnSQL sql.Result
		// returnErr is second return of mock interface rawSeterMock's Exec function which is also expected error
		returnErr error
		// returnRaw is the return of mock interface ormerMock's Raw function
		returnRaw orm.RawSeter
	}{{
		// Success Case
		name:      "SuccessCase",
		returnSQL: nil,
		returnErr: nil,
		returnRaw: rawSeterMock,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnSQL: nil,
		returnErr: errFailedDBOperation,
		returnRaw: rawSeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			rawSeterMock.EXPECT().Exec().Return(test.returnSQL, test.returnErr).Times(1)
			ormerMock.EXPECT().Raw(gomock.Any(), gomock.Any()).Return(test.returnRaw).Times(1)
			err := InsertOrUpdate(&meta)
			if test.returnErr != err {
				t.Errorf("Insert or Update Meta Case failed : wanted %v and got %v", test.returnErr, err)
			}
		})
	}
}

// TestUpdateMetaField is function to test UpdateMetaField
func TestUpdateMetaField(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// updateReturnInt is the first return of mock interface querySeterMock's update function
		updateReturnInt int64
		// updateReturnErr is the second return of mock interface querySeterMocks's update function also expected error
		updateReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(1),
		updateReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(0),
		updateReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Update(gomock.Any()).Return(test.updateReturnInt, test.updateReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := UpdateMetaField("test", "test", "test")
			if test.updateReturnErr != err {
				t.Errorf("Update Meta Field Case failed : wanted %v and got %v", test.updateReturnErr, err)
			}
		})
	}
}

// TestUpdateMetaFields is function to test UpdateMetaFields
func TestUpdateMetaFields(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// updateReturnInt is the first return of mock interface querySeterMock's update function
		updateReturnInt int64
		// updateReturnErr is the second return of mock interface querySeterMocks's update function also expected error
		updateReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(1),
		updateReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(0),
		updateReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Update(gomock.Any()).Return(test.updateReturnInt, test.updateReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := UpdateMetaFields("test", nil)

			if test.updateReturnErr != err {
				t.Errorf("Update Meta Fields Case failed : wanted %v and got %v", test.updateReturnErr, err)
			}
		})
	}
}

// TestQueryMeta is function to test QueryMeta
func TestQueryMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// allReturnInt is the first return of mock interface querySeterMock's all function
		allReturnInt int64
		// allReturnErr is the second return of mock interface querySeterMocks's all function also expected error
		allReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		allReturnInt:     int64(1),
		allReturnErr:     nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		allReturnInt:     int64(0),
		allReturnErr:     errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// fakeDao is used to set the argument of All function
	fakeDao := new([]Meta)
	fakeDaoArray := make([]Meta, 1)
	fakeDaoArray[0] = Meta{Key: "Test"}
	fakeDao = &fakeDaoArray

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().All(gomock.Any()).SetArg(0, *fakeDao).Return(test.allReturnInt, test.allReturnErr).Times(1)
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			meta, err := QueryMeta("test", "test")
			if test.allReturnErr != err {
				t.Errorf("Query Meta Case Failed : wanted error %v and got error %v", test.allReturnErr, err)
				return
			}

			if err == nil {
				if len(*meta) != 1 {
					t.Errorf("Query Meta Case failed: wanted length 1 and got length %v", len(*meta))
				}
			}
		})
	}
}

// TestQueryAllMeta is function to test QueryAllMeta
func TestQueryAllMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// allReturnInt is the first return of mock interface querySeterMock's all function
		allReturnInt int64
		// allReturnErr is the second return of mock interface querySeterMocks's all function also expected error
		allReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{
		{
			// Success Case
			name:             "SuccessCase",
			filterReturn:     querySeterMock,
			allReturnInt:     int64(1),
			allReturnErr:     nil,
			queryTableReturn: querySeterMock,
		},
		{
			// Failure Case
			name:             "FailureCase",
			filterReturn:     querySeterMock,
			allReturnInt:     int64(0),
			allReturnErr:     errFailedDBOperation,
			queryTableReturn: querySeterMock,
		},
	}

	// fakeDao is used to set the argument of All function
	fakeDao := new([]Meta)
	fakeDaoArray := make([]Meta, 1)
	fakeDaoArray[0] = Meta{Key: "Test", Value: "Test"}
	fakeDao = &fakeDaoArray

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().All(gomock.Any()).SetArg(0, *fakeDao).Return(test.allReturnInt, test.allReturnErr).Times(1)
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			meta, err := QueryAllMeta("test", "test")
			if test.allReturnErr != err {
				t.Errorf("Query All Meta Case Failed : wanted error %v and got error %v", test.allReturnErr, err)
				return
			}

			if err == nil {
				if len(*meta) != 1 {
					t.Errorf("Query All Meta Case failed: wanted length 1 and got length %v", len(*meta))
				}
			}
		})
	}
}

// TestIsNonUniqueNameError is function to test IsNonUniqueNameError().
func TestIsNonUniqueNameError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantBool bool
	}{
		{
			name:     "Suffix-are not unique",
			err:      errors.New("The fields are not unique"),
			wantBool: true,
		},
		{
			name:     "Contains-UNIQUE constraint failed",
			err:      errors.New("Failed-UNIQUE constraint failed"),
			wantBool: true,
		},
		{
			name:     "Contains-constraint failed",
			err:      errors.New("The input constraint failed"),
			wantBool: true,
		},
		{
			name:     "OtherError",
			err:      errors.New("Failed"),
			wantBool: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotBool := IsNonUniqueNameError(test.err)
			if gotBool != test.wantBool {
				t.Errorf("IsNonUniqueError() failed, Got = %v, Want = %v", gotBool, test.wantBool)
			}
		})
	}
}

func TestSaveMQTTMeta(t *testing.T) {
	// Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		name      string
		nodeName  string
		returnInt int64
		returnErr error
	}{
		{
			// Success Case
			name:      "SuccessCase",
			nodeName:  "test-node",
			returnInt: int64(1),
			returnErr: nil,
		},
		{
			// Failure Case
			name:      "FailureCase",
			nodeName:  "test-node",
			returnInt: int64(0),
			returnErr: errFailedDBOperation,
		},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			// Set up expectations
			ormerMock.EXPECT().Insert(gomock.Any()).DoAndReturn(
				func(meta interface{}) (int64, error) {
					// Type assert and verify the meta object
					mqttMeta, ok := meta.(*Meta)
					if !ok {
						t.Error("Expected *Meta type for Insert argument")
						return 0, errFailedDBOperation
					}

					// Verify key format
					expectedKey := fmt.Sprintf("default/pod/%s", constants.DefaultMosquittoContainerName)
					if mqttMeta.Key != expectedKey {
						t.Errorf("Expected key %s but got %s", expectedKey, mqttMeta.Key)
					}

					// Verify type
					if mqttMeta.Type != "pod" {
						t.Errorf("Expected type 'pod' but got %s", mqttMeta.Type)
					}

					// Verify that Value contains valid JSON
					var pod coreV1.Pod
					if err := json.Unmarshal([]byte(mqttMeta.Value), &pod); err != nil {
						t.Errorf("Failed to unmarshal pod JSON: %v", err)
					}

					// Verify essential pod fields
					if pod.Spec.NodeName != test.nodeName {
						t.Errorf("Expected NodeName %s but got %s", test.nodeName, pod.Spec.NodeName)
					}

					return test.returnInt, test.returnErr
				}).Times(1)

			// Execute the function
			err := SaveMQTTMeta(test.nodeName)

			// Verify results
			if test.returnErr != err {
				t.Errorf("SaveMQTTMeta case failed: wanted error %v and got error %v", test.returnErr, err)
			}
		})
	}
}
