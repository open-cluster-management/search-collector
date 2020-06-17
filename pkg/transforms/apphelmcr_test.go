/*
IBM Confidential
OCO Source Materials
(C) Copyright IBM Corporation 2019 All Rights Reserved
The source code for this program is not published or otherwise divested of its trade secrets,
irrespective of what has been deposited with the U.S. Copyright Office.
*/

package transforms

import (
	"testing"

	app "github.com/open-cluster-management/multicloud-operators-subscription-release/pkg/apis/apps/v1"
)

//TODO: Might have to update the json for testing once we have a live example
func TestTransformAppHelmCR(t *testing.T) {
	var a app.HelmRelease

	UnmarshalFile("apphelmcr.json", &a, t)

	node := AppHelmCRResource{&a}.BuildNode()

	// Test only the fields that exist in HelmRelease - the common test will test the other bits
	AssertEqual("name", node.Properties["name"], "testAppHelmCR", t)
	AssertEqual("kind", node.Properties["kind"], "HelmRelease", t)
}
