package org.aliyun.converter


class Operation(
        var requestType: String,
        val product: String,
        val version: String,
        val action: String,
        val serviceCode: String,
        var method: String,
        val uriPattern: String
) {
    private var parameters = Schema("")
    private var response = Schema("")

    fun setParam(s: Schema) {
        s.propOf(Schema("String").named("RegionId").positionIn("Query"))
        this.parameters = s
    }

    fun setResp(s: Schema) {
        this.response = s
    }

    fun goClientMethod(): String {
        val subTypePrefix = this.action
        val sideCodes = mutableMapOf<String, String>()

        var codes = ""

        when (requestType) {
            "roa" -> {
                codes += """
type ${this.action}Request ${this.parameters.goType(subTypePrefix, true, sideCodes)}

func (req *${this.action}Request) Invoke(client goaliyun.Client) (*${this.action}Response, error) {
	resp := &${this.action}Response{}
	err := client.Request("${this.product.toLowerCase()}", "${this.action}", "${this.version}", req.RegionId, "${this.method}", "${this.uriPattern}").Do(req, resp)
	return resp, err
}
"""
            }
            "rpc" -> {
                codes += """
type ${this.action}Request ${this.parameters.goType(subTypePrefix, true, sideCodes)}

func (req *${this.action}Request) Invoke(client goaliyun.Client) (*${this.action}Response, error) {
	resp := &${this.action}Response{}
	err := client.Request("${this.product.toLowerCase()}", "${this.action}", "${this.version}", req.RegionId, "${this.method}", "").Do(req, resp)
	return resp, err
}
"""
            }
        }

        this.parameters.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, true, sideCodes)}
"""
        }

        codes += """
type ${this.action}Response ${this.response.goType(subTypePrefix, false, sideCodes)}
"""
        this.response.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, false, sideCodes)}
"""
        }

        sideCodes.forEach { s ->
            codes += s.value
        }

        return codes
    }
}