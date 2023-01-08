// idl/api.thrift
namespace go api

include "idl/common.thrift"

struct MetadataGeneralPageableListRequest {
}

struct MetadataGeneralPageableListResponse {
  1: common.Meta meta
  2: list<common.MetadataGeneral> data
}

service ApiService {
    MetadataGeneralPageableListResponse MetadataGeneralPageableList(1: MetadataGeneralPageableListRequest request) (api.get="/v1/resource/general/pageable_list");
}