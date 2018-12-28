package restapi

import (
	"fmt"
	"net/url"

	"github.com/creativesoftwarefdn/weaviate/database"
	connutils "github.com/creativesoftwarefdn/weaviate/database/connectors/utils"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/models"
	libnetwork "github.com/creativesoftwarefdn/weaviate/network"
	"github.com/creativesoftwarefdn/weaviate/network/crossrefs"
)

type referenceSchemaUpdater struct {
	schemaManager database.SchemaManager
	network       libnetwork.Network
	fromClass     string
	kind          kind.Kind
}

func newReferenceSchemaUpdater(schemaManager database.SchemaManager, network libnetwork.Network,
	fromClass string, kind kind.Kind) *referenceSchemaUpdater {
	return &referenceSchemaUpdater{schemaManager, network, fromClass, kind}
}

// make sure this only ever called AFTER validtion as it skips
// validation checks itself
func (u *referenceSchemaUpdater) addNetworkDataTypes(schema interface{}) error {
	if schema == nil {
		return nil
	}

	for propName, prop := range schema.(map[string]interface{}) {
		switch propTyped := prop.(type) {
		case *models.SingleRef:
			err := u.singleRef(propTyped, propName)
			if err != nil {
				return err
			}

		case models.MultipleRef:
			for _, single := range propTyped {
				err := u.singleRef(single, propName)
				if err != nil {
					return err
				}
			}

		default:
			// primitive prop, skip
			continue
		}
	}
	return nil
}

func (u *referenceSchemaUpdater) singleRef(prop *models.SingleRef, propName string) error {
	parsed, err := url.Parse(*prop.LocationURL)
	if err != nil {
		return err
	}

	if parsed.Host == "localhost" {
		// local ref, nothign to do
		return nil
	}

	peers, err := u.network.ListPeers()
	if err != nil {
		return fmt.Errorf("could not list network peers: %s", err)
	}

	networkResourceKind, err := kindOfNetworResource(prop, parsed.Host)
	if err != nil {
		return err
	}

	remoteKind, err := peers.RemoteKind(crossrefs.NetworkKind{
		Kind:     networkResourceKind,
		ID:       prop.NrDollarCref,
		PeerName: parsed.Host,
	})
	if err != nil {
		return fmt.Errorf("invalid network reference: %s", err)
	}

	return u.updateSchema(remoteKind, parsed.Host, prop, propName)
}

func (u *referenceSchemaUpdater) updateSchema(remoteKind interface{}, peerName string, prop *models.SingleRef,
	propName string) error {
	switch thingOrAction := remoteKind.(type) {
	case models.Thing:
		remoteClass := fmt.Sprintf("%s/%s", peerName, thingOrAction.AtClass)
		err := u.schemaManager.UpdatePropertyAddDataType(u.kind, u.fromClass, propName, remoteClass)
		if err != nil {
			return fmt.Errorf("could not add network thing class %s to %s.%s: %s",
				remoteClass, u.fromClass, propName, err)
		}
	case models.Action:
		remoteClass := fmt.Sprintf("%s/%s", peerName, thingOrAction.AtClass)
		err := u.schemaManager.UpdatePropertyAddDataType(u.kind, u.fromClass, propName, remoteClass)
		if err != nil {
			return fmt.Errorf("could not add network action class %s to %s.%s: %s",
				remoteClass, u.fromClass, propName, err)
		}
	default:
		return fmt.Errorf("unrecognized kind from remote peer for %s from %s", prop.NrDollarCref, peerName)
	}

	return nil
}

func kindOfNetworResource(prop *models.SingleRef, peerName string) (kind.Kind, error) {
	var k kind.Kind
	switch connutils.RefType(prop.Type) {
	case connutils.RefTypeNetworkThing:
		k = kind.THING_KIND
	case connutils.RefTypeNetworkAction:
		k = kind.ACTION_KIND
	default:
		return kind.Kind(""), fmt.Errorf("unrecognized network reference type '%s' for cref %#v at %s, "+
			"must be one of: %s, %s", prop.Type, prop, peerName, connutils.RefTypeNetworkThing,
			connutils.RefTypeNetworkAction)
	}

	return k, nil
}