/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets,
irrespective of what has been deposited with the U.S. Copyright Office.
Copyright (c) 2020 Red Hat, Inc.
*/
// Copyright Contributors to the Open Cluster Management project

package transforms

import (
	v1 "k8s.io/api/apps/v1"
)

// StatefulSetResource ...
type StatefulSetResource struct {
	node Node
}

// StatefulSetResourceBuilder ...
func StatefulSetResourceBuilder(s *v1.StatefulSet) *StatefulSetResource {
	node := transformCommon(s)         // Start off with the common properties
	apiGroupVersion(s.TypeMeta, &node) // add kind, apigroup and version
	// Extract the properties specific to this type
	node.Properties["current"] = int64(s.Status.Replicas)
	node.Properties["desired"] = int64(0)
	if s.Spec.Replicas != nil {
		node.Properties["desired"] = int64(*s.Spec.Replicas)
	}

	return &StatefulSetResource{node: node}
}

// BuildNode construct the node for the StatefulSet Resources
func (s StatefulSetResource) BuildNode() Node {
	return s.node
}

// BuildEdges construct the edges for the StatefulSet Resources
func (s StatefulSetResource) BuildEdges(ns NodeStore) []Edge {
	//no op for now to implement interface
	return []Edge{}
}
