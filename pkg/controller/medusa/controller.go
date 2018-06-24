package medusa

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	operatorv1beta1 "github.com/danieloliveira079/kubebuilder-example/pkg/apis/operator/v1beta1"
	operatorv1beta1client "github.com/danieloliveira079/kubebuilder-example/pkg/client/clientset/versioned/typed/operator/v1beta1"
	operatorv1beta1informer "github.com/danieloliveira079/kubebuilder-example/pkg/client/informers/externalversions/operator/v1beta1"
	operatorv1beta1lister "github.com/danieloliveira079/kubebuilder-example/pkg/client/listers/operator/v1beta1"

	"github.com/danieloliveira079/kubebuilder-example/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Medusa resources goes here.

func (bc *MedusaController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on medusa.MedusaController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=operator,version=v1beta1,kind=Medusa,resource=medusas
type MedusaController struct {
	// INSERT ADDITIONAL FIELDS HERE
	medusaLister operatorv1beta1lister.MedusaLister
	medusaclient operatorv1beta1client.OperatorV1beta1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	medusarecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &MedusaController{
		medusaLister: arguments.ControllerManager.GetInformerProvider(&operatorv1beta1.Medusa{}).(operatorv1beta1informer.MedusaInformer).Lister(),

		medusaclient:   arguments.Clientset.OperatorV1beta1(),
		medusarecorder: arguments.CreateRecorder("MedusaController"),
	}

	// Create a new controller that will call MedusaController.Reconcile on changes to Medusas
	gc := &controller.GenericController{
		Name:             "MedusaController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&operatorv1beta1.Medusa{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Medusa Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the MedusaController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
