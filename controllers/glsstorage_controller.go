/*
Copyright 2022.

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

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	glsv1 "github.com/Jade0829/crd/api/v1"
)

// GlsstorageReconciler reconciles a Glsstorage object
type GlsstorageReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=gls.excrd.com,resources=glsstorages,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=gls.excrd.com,resources=glsstorages/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=gls.excrd.com,resources=glsstorages/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Glsstorage object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *GlsstorageReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	Glsstorage := &glsv1.Glsstorage{}

	err := r.Client.Get(ctx, req.NamespacedName, Glsstorage)

	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	spec := Glssotorage.Spec
	scConfig := Sepc.Sc
	pvcConfig := Spec.Pvc

	var provisioner string

	if scConfig.stype == "gluster" {
		provisioner = "kadalu"
	} else if scConfig.stype == "nvme" {
		provisioner = "csi.gluesys.io"
	}
	GlsSC := &storagev1.StorageClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: spec.Name + "sc",
		},
		Provisioner: provisioner,
	}

	_ = r.Client.Create(context.Backgroud(), GlsSC)

	GlsPVC := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: spec.Name,
		},
		Spec: &corev1.PersistentVolumeClaimSpec{
			AccesccModes: []corev1.PersistentVolumeAccessMode{
				"ReadWriteMany",
			},
			storageClassName: spec.Name + "sc",
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: pvcConfig.limit,
				},
			},
		},
	}

	_ = r.Client.Create(context.Backgroud(), GlsPVC)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GlsstorageReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&glsv1.Glsstorage{}).
		Complete(r)
}
