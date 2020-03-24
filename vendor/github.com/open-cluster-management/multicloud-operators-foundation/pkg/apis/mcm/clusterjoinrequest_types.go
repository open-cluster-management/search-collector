// licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

package mcm

import (
	csrv1beta1 "k8s.io/api/certificates/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterJoinRequest is the request from klusterlet to join manager
type ClusterJoinRequest struct {
	metav1.TypeMeta
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta

	// Spec defines the request to join hcm
	Spec ClusterJoinRequestSpec

	// Status defins the join status
	Status ClusterJoinStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterJoinRequestList is the request list from klusterlet to join manager
type ClusterJoinRequestList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta

	// List of Cluster objects.
	Items []ClusterJoinRequest
}

type ClusterJoinRequestSpec struct {
	// ClusterName is the name of the cluster
	ClusterName string

	// ClusterNamespace is the namespace for cluster
	ClusterNamespace string

	// CSR is the csr request spec
	CSR csrv1beta1.CertificateSigningRequestSpec
}

type JoinRequestPhase string

// These are the possible conditions for a certificate request.
const (
	JoinApproved JoinRequestPhase = "Approved"
	JoinDenied   JoinRequestPhase = "Denied"
)

type ClusterJoinStatus struct {
	// Phase is the pa
	Phase JoinRequestPhase
	// CSRStatus is the status of CSR
	CSRStatus csrv1beta1.CertificateSigningRequestStatus
}
