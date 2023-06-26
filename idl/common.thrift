// idl/common.thrift
namespace go common

include "idl/enum.thrift"

struct Meta {
  1: enum.ResponseCode  status_code
  2: string status_msg
}
