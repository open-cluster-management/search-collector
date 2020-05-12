/*
Copyright 2019 The Kubernetes Authors.

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

package v1

import (
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//ChartsDir env variable name which contains the directory where the charts are installed
const ChartsDir = "CHARTS_DIR"

//SourceTypeEnum types of sources
type SourceTypeEnum string

const (
	// HelmRepoSourceType helmrepo source type
	HelmRepoSourceType SourceTypeEnum = "helmrepo"
	// GitHubSourceType github source type
	GitHubSourceType SourceTypeEnum = "github"
	// GitSourceType git source type
	GitSourceType SourceTypeEnum = "git"
)

//GitHub provides the parameters to access the helm-chart located in a github repo
type GitHub struct {
	Urls      []string `json:"urls,omitempty"`
	ChartPath string   `json:"chartPath,omitempty"`
	Branch    string   `json:"branch,omitempty"`
}

//Git provides the parameters to access the helm-chart located in a git repo
type Git struct {
	Urls      []string `json:"urls,omitempty"`
	ChartPath string   `json:"chartPath,omitempty"`
	Branch    string   `json:"branch,omitempty"`
}

//HelmRepo provides the urls to retrieve the helm-chart
type HelmRepo struct {
	Urls []string `json:"urls,omitempty"`
}

//Source holds the different types of repository
type Source struct {
	SourceType SourceTypeEnum `json:"type,omitempty"`
	GitHub     *GitHub        `json:"github,omitempty"`
	Git        *Git           `json:"git,omitempty"`
	HelmRepo   *HelmRepo      `json:"helmRepo,omitempty"`
}

func (s Source) String() string {
	switch strings.ToLower(string(s.SourceType)) {
	case string(HelmRepoSourceType):
		return fmt.Sprintf("%v", s.HelmRepo.Urls)
	case string(GitHubSourceType):
		return fmt.Sprintf("%v|%s|%s", s.GitHub.Urls, s.GitHub.Branch, s.GitHub.ChartPath)
	case string(GitSourceType):
		return fmt.Sprintf("%v|%s|%s", s.Git.Urls, s.Git.Branch, s.Git.ChartPath)
	default:
		return fmt.Sprintf("SourceType %s not supported", s.SourceType)
	}
}

// HelmReleaseRepo defines the repository of HelmRelease
// +k8s:openapi-gen=true
type HelmReleaseRepo struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// Source holds the url toward the helm-chart
	Source *Source `json:"source,omitempty"`
	// ChartName is the name of the chart within the repo
	ChartName string `json:"chartName,omitempty"`
	// Version is the chart version
	Version string `json:"version,omitempty"`
	// Secret to use to access the helm-repo defined in the CatalogSource.
	SecretRef *corev1.ObjectReference `json:"secretRef,omitempty"`
	// Configuration parameters to access the helm-repo defined in the CatalogSource
	ConfigMapRef *corev1.ObjectReference `json:"configMapRef,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRelease is the Schema for the subscriptionreleases API
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:subresource:status
type HelmRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Repo HelmReleaseRepo `json:"repo,omitempty"`

	Spec   HelmAppSpec   `json:"spec,omitempty"`
	Status HelmAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmReleaseList contains a list of HelmRelease
type HelmReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelmRelease{}, &HelmReleaseList{})
}

// Below are mostly copied from operator sdk's internal package. See: github.com/operator-framework/operator-sdk

type HelmAppSpec interface{} // modified

type HelmAppConditionType string
type ConditionStatus string
type HelmAppConditionReason string

type HelmAppCondition struct {
	Type    HelmAppConditionType   `json:"type"`
	Status  ConditionStatus        `json:"status"`
	Reason  HelmAppConditionReason `json:"reason,omitempty"`
	Message string                 `json:"message,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type HelmAppRelease struct {
	Name     string `json:"name,omitempty"`
	Manifest string `json:"manifest,omitempty"`
}

const (
	ConditionInitialized    HelmAppConditionType = "Initialized"
	ConditionDeployed       HelmAppConditionType = "Deployed"
	ConditionReleaseFailed  HelmAppConditionType = "ReleaseFailed"
	ConditionIrreconcilable HelmAppConditionType = "Irreconcilable"
	ConditionTimedout       HelmAppConditionType = "Timedout"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   HelmAppConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    HelmAppConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful HelmAppConditionReason = "UninstallSuccessful"
	ReasonInstallError        HelmAppConditionReason = "InstallError"
	ReasonUpdateError         HelmAppConditionReason = "UpdateError"
	ReasonReconcileError      HelmAppConditionReason = "ReconcileError"
	ReasonUninstallError      HelmAppConditionReason = "UninstallError"
)

type HelmAppStatus struct {
	Conditions      []HelmAppCondition `json:"conditions"`
	DeployedRelease *HelmAppRelease    `json:"deployedRelease,omitempty"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *HelmAppStatus) SetCondition(condition HelmAppCondition) *HelmAppStatus {
	now := metav1.Now()

	for i := range s.Conditions {
		if s.Conditions[i].Type == condition.Type {
			if s.Conditions[i].Status != condition.Status {
				condition.LastTransitionTime = now
			} else {
				condition.LastTransitionTime = s.Conditions[i].LastTransitionTime
			}

			s.Conditions[i] = condition

			return s
		}
	}

	// If the condition does not exist,
	// initialize the lastTransitionTime
	condition.LastTransitionTime = now
	s.Conditions = append(s.Conditions, condition)

	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *HelmAppStatus) RemoveCondition(conditionType HelmAppConditionType) *HelmAppStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)

			return s
		}
	}

	return s
}
