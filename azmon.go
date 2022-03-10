//def build_signature(customer_id, shared_key, date, content_length, method, content_type, resource):
//#    x_headers = 'x-ms-date:' + date
//    string_to_hash = method + "\n" + str(content_length) + "\n" + content_type + "\n" + x_headers + "\n" + resource
//    bytes_to_hash = bytes(string_to_hash, encoding="utf-8")
//    decoded_key = base64.b64decode(shared_key)
//    encoded_hash = base64.b64encode(hmac.new(decoded_key, bytes_to_hash, digestmod=hashlib.sha256).digest()).decode()
//    authorization = "SharedKey {}:{}".format(customer_id,encoded_hash)
//    return authorization


package main

import (
	"fmt"
     "strconv"
)

func build_signature(customer_id, shared_key, date, content_length, method, content_type, resource string) string {
  var x_headers string
  var string_to_hash string
  fmt.Println(customer_id, shared_key, date, content_length, method, content_type, resource)
  x_headers = "x-ms-date" + date 
  string_to_hash = method + "\n" + strconv.Itoa(len(content_length)) + "\n" + content_type + "\n" + x_headers + "\n" + resource
  fmt.Println(string_to_hash)
 
  return method
}



func main() {
	fmt.Println("Test")
    build_signature("ovidiu", "borlean", "dumitru", "test", "metoda", "tipcontinut", "resursa" )
}