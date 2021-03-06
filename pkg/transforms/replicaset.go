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

// ReplicaSetResource ...
type ReplicaSetResource struct {
	node Node
}

// ReplicaSetResourceBuilder ...
func ReplicaSetResourceBuilder(r *v1.ReplicaSet) *ReplicaSetResource {
	node := transformCommon(r)         // Start off with the common properties
	apiGroupVersion(r.TypeMeta, &node) // add kind, apigroup and version
	// Extract the properties specific to this type
	node.Properties["current"] = int64(r.Status.Replicas)
	node.Properties["desired"] = int64(0)
	if r.Spec.Replicas != nil {
		node.Properties["desired"] = int64(*r.Spec.Replicas)
	}

	return &ReplicaSetResource{node: node}
}

// BuildNode construct the node for ReplicaSet Resources
func (r ReplicaSetResource) BuildNode() Node {
	return r.node
}

// BuildEdges construct the edges for ReplicaSet Resources
func (r ReplicaSetResource) BuildEdges(ns NodeStore) []Edge {
	//no op for now to implement interface
	return []Edge{}
}
