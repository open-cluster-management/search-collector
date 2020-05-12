/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package transforms

import (
	app "github.com/open-cluster-management/multicloud-operators-placementrule/pkg/apis/apps/v1"
)

type PlacementRuleResource struct {
	*app.PlacementRule
}

func (p PlacementRuleResource) BuildNode() Node {
	node := transformCommon(p)         // Start off with the common properties
	apiGroupVersion(p.TypeMeta, &node) // add kind, apigroup and version
	//TODO: Add other properties
	return node
}

func (p PlacementRuleResource) BuildEdges(ns NodeStore) []Edge {
	//no op for now to implement interface
	return []Edge{}
}
