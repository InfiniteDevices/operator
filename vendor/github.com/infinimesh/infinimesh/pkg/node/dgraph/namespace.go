//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package dgraph

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ListNamespaces is a method to List details of all Namespaces
func (s *DGraphRepo) ListNamespaces(ctx context.Context) (namespaces []*nodepb.Namespace, err error) {
	const q = `{
                     namespaces(func: eq(type, "namespace")) {
  	               uid
                       name
	             }
                   }`

	res, err := s.Dg.NewReadOnlyTxn().Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, namespace := range resultSet.Namespaces {
		namespaces = append(namespaces, &nodepb.Namespace{
			Id:   namespace.UID,
			Name: namespace.Name,
		})
	}

	return namespaces, nil
}

//DeletePermissionInNamespace is a method to delete permissions for a Namespaces for an account
func (s *DGraphRepo) DeletePermissionInNamespace(ctx context.Context, namespaceID, accountID string) (err error) {
	txn := s.Dg.NewTxn()
	const q = `query deletePermissionInNamespace($namespaceID: string, $accountID: string){
  accounts(func: uid($namespaceID)) @filter(eq(type, "namespace")) @cascade @normalize {
    namespace_uid: uid
    ~access.to.namespace @filter(uid($accountID))  {
      account_uid: uid
    }
  }
}`

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
		"$accountID":   accountID,
	})
	if err != nil {
		return err
	}

	var resultSet struct {
		Accounts []*struct {
			AccountUID   string `json:"account_uid"`
			NamespaceUID string `json:"namespace_uid"`
		} `json:"accounts"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return err
	}

	m := &api.Mutation{CommitNow: true}
	for _, account := range resultSet.Accounts {
		m.Del = append(m.Del, &api.NQuad{
			Subject:   account.AccountUID,
			Predicate: "access.to.namespace",
			ObjectId:  account.NamespaceUID,
		})
	}
	_, err = txn.Mutate(ctx, m)
	return err
}

//ListPermissionsInNamespace is a method to list permissions for all accounts for a Namespace
func (s *DGraphRepo) ListPermissionsInNamespace(ctx context.Context, namespaceid string) (permissions []*nodepb.Permission, err error) {
	const q = `query listPermissionsInNamespace($namespaceid: string) {
		accounts(func: uid($namespaceid)) @filter(eq(type, "namespace")) @normalize @cascade  {
		  ~access.to.namespace {
			uid: uid
			name: name
		  } @facets(permission)
		}
	  }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$namespaceid": namespaceid})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Accounts []*struct {
			UID    string `json:"uid"`
			Name   string `json:"name"`
			Action string `json:"~access.to.namespace|permission"`
		} `json:"accounts"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, account := range resultSet.Accounts {
		permissions = append(permissions, &nodepb.Permission{
			Namespace:   namespaceid,
			AccountId:   account.UID,
			AccountName: account.Name,
			Action:      nodepb.Action(nodepb.Action_value[account.Action]),
		})
	}

	return permissions, nil

}

//ListNamespacesForAccount is a method to list Namespaces for an Account
func (s *DGraphRepo) ListNamespacesForAccount(ctx context.Context, accountID string) (namespaces []*nodepb.Namespace, err error) {
	const q = `query listNamespaces($account: string) {
		namespaces(func: uid($account)) @normalize @cascade  {
		  access.to.namespace @filter(eq(type, "namespace")) @facets(NOT eq(permission,"NONE")) {
			uid : uid
			name : name
		  }
		}
	  }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$account": accountID})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, namespace := range resultSet.Namespaces {
		namespaces = append(namespaces, &nodepb.Namespace{
			Id:   namespace.UID,
			Name: namespace.Name,
		})
	}

	return namespaces, nil
}

//IsAuthorizedNamespace is a method that returns if the access to the namespace for an account is true or false
func (s *DGraphRepo) IsAuthorizedNamespace(ctx context.Context, namespaceid, account string, action nodepb.Action) (decision bool, err error) {
	acc, err := s.GetAccount(ctx, account)
	if err != nil {
		return false, err
	}

	if acc.IsRoot {
		return true, nil
	}

	params := map[string]string{
		"$namespaceid": namespaceid,
		"$user_id":     account,
	}

	txn := s.Dg.NewReadOnlyTxn()

	const q = `query access($namespaceid: string, $user_id: string){
		access(func: uid($user_id)) @normalize @cascade {
		  name
		  uid
		  access.to.namespace @filter(uid($namespaceid)) @facets(permission,inherit) {
			uid
			name
			type
		  }
		}
	  }
`

	res, err := txn.QueryWithVars(ctx, q, params)
	if err != nil {
		return false, err
	}
	var access struct {
		Access []Namespace `json:"access"`
	}

	err = json.Unmarshal(res.Json, &access)
	if err != nil {
		return false, err
	}

	actionValue := strings.Split(action.String(), "_")

	if len(access.Access) > 0 {
		if isPermissionSufficient(actionValue[0], access.Access[0].AccessToPermission) {
			return true, nil
		}
	}

	return false, nil
}

//SoftDeleteNamespace is a method that mark the namespace for deletion
func (s *DGraphRepo) SoftDeleteNamespace(ctx context.Context, namespaceID string) (err error) {
	txn := s.Dg.NewTxn()
	m := &api.Mutation{CommitNow: true}

	const q = `query deleteNodes($namespaceID: string){
        nodes(func: uid($namespaceID)) @filter(eq(type,"namespace")) {
          uid
        }
      }
      `

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
	})
	if err != nil {
		return err
	}

	var result struct {
		Nodes []*Node `json:"nodes"`
	}

	err = json.Unmarshal(res.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Nodes) != 1 {
		return status.Error(codes.NotFound, "The Namespace is not found")
	}

	m.Set = append(m.Set, &api.NQuad{
		Subject:     namespaceID,
		Predicate:   "markfordeletion",
		ObjectId:    namespaceID,
		ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: "true"}},
	})

	m.Set = append(m.Set, &api.NQuad{
		Subject:     namespaceID,
		Predicate:   "deleteinitiationtime",
		ObjectId:    namespaceID,
		ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: time.Now().Format(time.RFC3339)}},
	})

	_, err = txn.Mutate(ctx, m)
	return err
}

//HardDeleteNamespace is a method that deletes a namespace permantly
func (s *DGraphRepo) HardDeleteNamespace(ctx context.Context, namespaceID string) (err error) {
	txn := s.Dg.NewReadOnlyTxn()
	const q = `query deleteNodes($namespaceID: string){
        nodes(func: uid($namespaceID)) @filter(eq(type,"namespace")) @normalize {
          uid
        owns {
          uid
        }
        }
      }
      `

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
	})
	if err != nil {
		return err
	}

	var result struct {
		Nodes []*Node `json:"nodes"`
	}

	err = json.Unmarshal(res.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Nodes) != 1 {
		return status.Error(codes.NotFound, "The Namespace is not found")
	}

	for _, item := range result.Nodes {
		err = s.DeleteObject(ctx, item.UID)
	}

	return err
}
