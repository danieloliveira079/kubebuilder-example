package inject

import (
	operatorv1beta1 "github.com/danieloliveira079/kubebuilder-example/pkg/apis/operator/v1beta1"
	rscheme "github.com/danieloliveira079/kubebuilder-example/pkg/client/clientset/versioned/scheme"
	"github.com/danieloliveira079/kubebuilder-example/pkg/controller/ingress"
	"github.com/danieloliveira079/kubebuilder-example/pkg/controller/medusa"
	"github.com/danieloliveira079/kubebuilder-example/pkg/inject/args"
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&operatorv1beta1.Medusa{}, arguments.Informers.Operator().V1beta1().Medusas()); err != nil {
			return err
		}

		// Add Kubernetes informers
		if err := arguments.ControllerManager.AddInformerProvider(&extensionsv1beta1.Ingress{}, arguments.KubernetesInformers.Extensions().V1beta1().Ingresses()); err != nil {
			return err
		}

		if c, err := ingress.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		if c, err := medusa.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &operatorv1beta1.MedusaCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"operator.octops.io"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{
			"extensions",
		},
		Resources: []string{
			"ingresses",
		},
		Verbs: []string{
			"get", "list", "watch",
		},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "operator.octops.io",
		Version: "v1beta1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
