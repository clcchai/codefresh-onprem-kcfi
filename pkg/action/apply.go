/*
Copyright The Codefresh Authors.

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

package action

import (
	"fmt"
	"github.com/stretchr/objx"
	helm "helm.sh/helm/v3/pkg/action"
	c "github.com/codefresh-io/kcfi/pkg/config"
)



// CfApply is an action to create or update Codefresh
type CfApply struct {
	ConfigFile string
	vals map[string]interface{}
	cfg *helm.Configuration
	Helm *helm.Upgrade
}



// NewCfApply creates object
func NewCfApply(cfg *helm.Configuration) *CfApply {
	return &CfApply{
		cfg: cfg,
		Helm: helm.NewUpgrade(cfg),
	}
}



// Run the action
func (o *CfApply) Run(vals map[string]interface{}) error {
	fmt.Printf("Applying Codefresh configuration from %s\n", o.ConfigFile)
	// fmt.Printf("Applying Codefresh configuration from %s\n", o.ConfigFile)
	o.vals = vals
	valsX := objx.New(vals)
	kind := valsX.Get(c.KeyKind).String(); 

	switch kind {
	case kindCodefresh:
		return o.ApplyCodefresh()
		
	default:
		return fmt.Errorf("Wrong installer kind %s", kind)

	}
}


