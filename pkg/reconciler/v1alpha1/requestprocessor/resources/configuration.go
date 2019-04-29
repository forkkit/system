/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"fmt"

	"github.com/knative/pkg/kmeta"
	knservingv1alpha1 "github.com/knative/serving/pkg/apis/serving/v1alpha1"
	runv1alpha1 "github.com/projectriff/system/pkg/apis/run/v1alpha1"
	"github.com/projectriff/system/pkg/reconciler/v1alpha1/requestprocessor/resources/names"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MakeConfiguration creates a Configuration from an RequestProcessor object.
func MakeConfiguration(rp *runv1alpha1.RequestProcessor, i int) (*knservingv1alpha1.Configuration, error) {
	rpsi := rp.Spec[i]
	var name string
	if rpsi.Tag != "" {
		name = fmt.Sprintf("%s-%s", names.Configuration(rp), rpsi.Tag)
	} else {
		name = fmt.Sprintf("%s-%d", names.Configuration(rp), i)
	}
	configuration := &knservingv1alpha1.Configuration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: rp.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*kmeta.NewControllerRef(rp),
			},
			Labels: makeLabels(rp),
		},
		Spec: knservingv1alpha1.ConfigurationSpec{
			RevisionTemplate: knservingv1alpha1.RevisionTemplateSpec{
				Spec: knservingv1alpha1.RevisionSpec{
					ServiceAccountName: rpsi.ServiceAccountName,
					Container:          rpsi.Containers[0],
					Volumes:            rpsi.Volumes,
				},
			},
		},
	}

	return configuration, nil
}