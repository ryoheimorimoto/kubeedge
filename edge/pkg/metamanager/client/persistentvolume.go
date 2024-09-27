package client

import (
	"encoding/json"
	"fmt"

	api "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/edge/pkg/common/message"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
	v2 "github.com/kubeedge/kubeedge/edge/pkg/metamanager/dao/v2"
)

// PersistentVolumesGetter is interface to get client PersistentVolumes
type PersistentVolumesGetter interface {
	PersistentVolumes() PersistentVolumesInterface
}

// PersistentVolumesInterface is interface for client PersistentVolumes
type PersistentVolumesInterface interface {
	Create(*api.PersistentVolume) (*api.PersistentVolume, error)
	Update(*api.PersistentVolume) error
	Delete(name string) error
	Get(name string, options metav1.GetOptions) (*api.PersistentVolume, error)
}

type persistentvolumes struct {
	send SendInterface
}

func newPersistentVolumes(s SendInterface) *persistentvolumes {
	return &persistentvolumes{
		send: s,
	}
}

func (c *persistentvolumes) Create(*api.PersistentVolume) (*api.PersistentVolume, error) {
	return nil, nil
}

func (c *persistentvolumes) Update(*api.PersistentVolume) error {
	return nil
}

func (c *persistentvolumes) Delete(string) error {
	return nil
}

func (c *persistentvolumes) Get(name string, _ metav1.GetOptions) (*api.PersistentVolume, error) {
	resource := fmt.Sprintf("%s/%s/%s", v2.NullNamespace, "persistentvolume", name)
	pvMsg := message.BuildMsg(modules.MetaGroup, "", modules.EdgedModuleName, resource, model.QueryOperation, nil)
	msg, err := c.send.SendSync(pvMsg)
	if err != nil {
		return nil, fmt.Errorf("get persistentvolume from metaManager failed, err: %v", err)
	}

	content, err := msg.GetContentData()
	if err != nil {
		return nil, fmt.Errorf("parse message to persistentvolume failed, err: %v", err)
	}

	if msg.GetOperation() == model.ResponseOperation && msg.GetSource() == modules.MetaManagerModuleName {
		return handlePersistentVolumeFromMetaDB(content)
	}
	return handlePersistentVolumeFromMetaManager(content)
}

func handlePersistentVolumeFromMetaDB(content []byte) (*api.PersistentVolume, error) {
	var lists []string
	err := json.Unmarshal(content, &lists)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to persistentvolume list from db failed, err: %v", err)
	}

	if len(lists) != 1 {
		return nil, fmt.Errorf("persistentvolume length from meta db is %d", len(lists))
	}

	var pv *api.PersistentVolume
	err = json.Unmarshal([]byte(lists[0]), &pv)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to persistentvolume from db failed, err: %v", err)
	}
	return pv, nil
}

func handlePersistentVolumeFromMetaManager(content []byte) (*api.PersistentVolume, error) {
	var pv api.PersistentVolume
	err := json.Unmarshal(content, &pv)
	if err != nil {
		return nil, fmt.Errorf("unmarshal message to persistentvolume failed, err: %v", err)
	}
	return &pv, nil
}
