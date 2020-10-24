/*
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
	"k8s.io/apimachinery/pkg/api/errors"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	minecraftv1alpha1 "github.com/robertsmieja/minecraft-operator/api/v1alpha1"
)

// VanillaServerReconciler reconciles a VanillaServer object
type VanillaServerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=minecraft.robertsmieja.com,resources=vanillaservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=minecraft.robertsmieja.com,resources=vanillaservers/status,verbs=get;update;patch

func (r *VanillaServerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("vanillaserver", req.NamespacedName)

	// Fetch the VanillaServer instances
	memcached := &minecraftv1alpha1.VanillaServerList{}
	err := r.Get(ctx, req.NamespacedName, memcached)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			log.Info("VanillaServer resources not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get VanillaServer")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *VanillaServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&minecraftv1alpha1.VanillaServer{}).
		Complete(r)
}
