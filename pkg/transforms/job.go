/*
IBM Confidential
OCO Source Materials
5737-E67
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
*/

package transforms

import (
	v1 "k8s.io/api/batch/v1"
)

type JobResource struct {
	*v1.Job
}

func (j JobResource) BuildNode() Node {
	node := transformCommon(j) // Start off with the common properties

	// Extract the properties specific to this type
	node.Properties["kind"] = "Job"
	node.Properties["apigroup"] = "batch"
	node.Properties["successful"] = int64(j.Status.Succeeded)
	node.Properties["completions"] = int64(0)
	if j.Spec.Completions != nil {
		node.Properties["completions"] = int64(*j.Spec.Completions)
	}
	node.Properties["parallelism"] = int64(0)
	if j.Spec.Completions != nil {
		node.Properties["parallelism"] = int64(*j.Spec.Parallelism)
	}

	return node
}

func (j JobResource) BuildEdges(ns NodeStore) []Edge {
	ret := []Edge{}
	UID := prefixedUID(j.UID)
	//deployer subscriber edges
	nodeInfo := NodeInfo{NameSpace: j.Namespace, UID: UID, Kind: j.Kind, Name: j.Name}
	ret = append(ret, edgesByDeployerSubscriber(nodeInfo, ns)...)
	return ret
}
