package ingress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"

	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	extensionsv1beta1informer "k8s.io/client-go/informers/extensions/v1beta1"
	extensionsv1beta1client "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	extensionsv1beta1lister "k8s.io/client-go/listers/extensions/v1beta1"

	"github.com/danieloliveira079/kubebuilder-example/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for Ingress resources goes here.

func (bc *IngressController) Reconcile(k types.ReconcileKey) error {
	ing, err := bc.ingressclient.Ingresses(k.Namespace).Get(k.Name, v1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Println("Object Not found")
			return nil
		}
		fmt.Printf("\nerr: %v", err)
		return err
	}

	annotations := ing.GetAnnotations()
	multiproxy := annotations["octops.io/multiproxy"]
	branch := annotations["octops.io/branch"]

	if multiproxy == "true" && branch != "" {
		rules := ing.Spec.Rules
		newUpstreams := []Upstream{}

		for _, rule := range rules {
			if rule.Host != "" {
				up := Upstream{
					Key:  branch,
					Addr: rule.Host,
				}

				newUpstreams = append(newUpstreams, up)
				log.Printf("Ingress %v has host: %v", ing.GetName(), rule.Host)
			} else {
				log.Printf("No host for ingress: %v", ing.GetName())
			}
		}

		newConfigUpstream := Upstreams{
			Upstreams: newUpstreams,
		}

		outJSON, err := json.Marshal(newConfigUpstream)
		if err != nil {
			fmt.Printf("\nError during Marshal operation: %v", err)
		}

		f, err := os.Open("config.json")
		if err != nil {
			return err
		}
		defer f.Close()

		currentUpstreams := Upstreams{}
		if err := json.NewDecoder(f).Decode(&currentUpstreams); err != nil {
			return err
		}

		log.Printf("current: %v", currentUpstreams)

		//Write final config to disk based on new and current upstreams
		err = ioutil.WriteFile("config.json", outJSON, 0644)

		if err != nil {
			fmt.Printf("\nError writing file to disk: %v", err)
		}

		log.Printf("new: %v", string(outJSON))
	} else {
		log.Printf("Ingress ignored: %v", ing.GetName())
	}

	return nil
}

// +kubebuilder:informers:group=extensions,version=v1beta1,kind=Ingress
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;watch;list
// +kubebuilder:controller:group=extensions,version=v1beta1,kind=Ingress,resource=ingresses
type IngressController struct {
	// INSERT ADDITIONAL FIELDS HERE
	ingressLister extensionsv1beta1lister.IngressLister
	ingressclient extensionsv1beta1client.ExtensionsV1beta1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	ingressrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &IngressController{
		ingressLister:   arguments.ControllerManager.GetInformerProvider(&extensionsv1beta1.Ingress{}).(extensionsv1beta1informer.IngressInformer).Lister(),
		ingressclient:   arguments.KubernetesClientSet.ExtensionsV1beta1(),
		ingressrecorder: arguments.CreateRecorder("IngressController"),
	}

	// Create a new controller that will call IngressController.Reconcile on changes to Ingresss
	gc := &controller.GenericController{
		Name:             "IngressController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&extensionsv1beta1.Ingress{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a Ingress Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the IngressController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
