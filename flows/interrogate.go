//
package flows

import (
	"github.com/golang/protobuf/proto"

	"errors"
	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	"www.velocidex.com/golang/velociraptor/config"
	"www.velocidex.com/golang/velociraptor/constants"
	crypto_proto "www.velocidex.com/golang/velociraptor/crypto/proto"
	"www.velocidex.com/golang/velociraptor/datastore"
	"www.velocidex.com/golang/velociraptor/responder"
	//	utils "www.velocidex.com/golang/velociraptor/testing"
)

const (
	_                        = iota
	processClientInfo uint64 = 1
)

type VInterrogate struct{}

func (self *VInterrogate) Start(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject) (*string, error) {
	if flow_obj.Runner_args.ClientId == "" {
		return nil, errors.New("Client id not provided.")
	}
	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return nil, err
	}

	flow_id := GetNewFlowIdForClient(flow_obj.Runner_args.ClientId)
	queries := []*actions_proto.VQLRequest{
		&actions_proto.VQLRequest{
			VQL:  "select Client_name, Client_build_time, Client_labels from config",
			Name: "Client Info"},
		&actions_proto.VQLRequest{
			VQL: "select OS, Architecture, Platform, PlatformVersion, " +
				"KernelVersion, Fqdn from info()",
			Name: "System Info"},
		&actions_proto.VQLRequest{
			VQL: "select ut_type, ut_id, ut_host, ut_user, " +
				"timestamp(epoch=ut_tv.tv_sec) as login_time from " +
				"users() where ut_type =~ 'USER'",
			Name: "Recent Users"},
	}
	vql_request := &actions_proto.VQLCollectorArgs{
		Query: queries,
	}

	err = db.QueueMessageForClient(
		config_obj, flow_obj.Runner_args.ClientId,
		flow_id,
		"VQLClientAction",
		vql_request, processClientInfo)
	if err != nil {
		return nil, err
	}

	flow_obj.SetState(&actions_proto.ClientInfo{})

	return &flow_id, nil
}

func (self *VInterrogate) ProcessMessage(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject,
	message *crypto_proto.GrrMessage) error {

	switch message.RequestId {
	case processClientInfo:
		err := flow_obj.FailIfError(message)
		if err != nil {
			return err
		}

		if flow_obj.IsRequestComplete(message) {
			defer flow_obj.Complete()

			// The flow is complete - store the client
			// info from our state into the client's AFF4
			// object.
			err := self.StoreClientInfo(config_obj, flow_obj)
			if err != nil {
				return err
			}
			return nil
		}

		// Retrieve the client info from the flow state and
		// modify it by adding the response to it.
		client_info := flow_obj.GetState().(*actions_proto.ClientInfo)
		defer flow_obj.SetState(client_info)

		vql_response, ok := responder.ExtractGrrMessagePayload(
			message).(*actions_proto.VQLResponse)
		if ok {
			client_info.Info = append(client_info.Info, vql_response)
		}
	}

	return nil
}

func (self *VInterrogate) StoreClientInfo(
	config_obj *config.Config,
	flow_obj *AFF4FlowObject) error {

	client_info := flow_obj.GetState().(*actions_proto.ClientInfo)

	client_urn := "aff4:/" + flow_obj.Runner_args.ClientId
	db, err := datastore.GetDB(config_obj)
	if err != nil {
		return err
	}

	data := make(map[string][]byte)
	encoded_client_info, err := proto.Marshal(client_info)
	if err != nil {
		return err
	}

	data[constants.CLIENT_VELOCIRAPTOR_INFO] = encoded_client_info
	db.SetSubjectData(config_obj, client_urn, data)

	return nil
}

func init() {
	impl := VInterrogate{}
	RegisterImplementation("VInterrogate", &impl)
}
