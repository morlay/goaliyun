package org.aliyun.converter

import com.github.javaparser.ast.body.TypeDeclaration
import com.github.javaparser.ast.type.Type
import converter.Inflection

class Schema(val type: String) {
    companion object {
        var positionMethod = Regex("put([A-Za-z]+)Parameter")

        fun fromTypeDeclForParameter(typeDeclaration: TypeDeclaration<*>, imports: Map<String, List<String>>): Schema {
            val schema = Schema("object").named(typeDeclaration.name.asString())

            typeDeclaration.fields.forEach { field ->
                val name = toUpperCamelCase(field.asFieldDeclaration().getVariable(0).nameAsString)
                schema.propOf(Schema.fromType(field.commonType, name, imports))
            }

            typeDeclaration.methods.forEach { method ->
                if (method.name.asString().startsWith("set")) {
                    val body = method.body.toString()
                    if (positionMethod.containsMatchIn(body)) {
                        val fieldName = toUpperCamelCase(method.name.asString().replace(Regex("^set"), ""))
                        if (schema.props != null) {
                            if (schema.props!![fieldName] != null) {
                                schema.props!![fieldName]!!.positionIn(positionMethod.find(body)!!.groupValues[1])
                            }
                        }
                    }
                }
            }
            typeDeclaration.members.forEach { member ->
                if (member.isTypeDeclaration) {
                    schema.def(Schema.fromTypeDecl(member.asTypeDeclaration(), imports))
                }
            }
            return schema
        }


        fun fromTypeDecl(typeDeclaration: TypeDeclaration<*>, imports: Map<String, List<String>>): Schema {
            val schema = Schema("object").named(typeDeclaration.name.asString())
            typeDeclaration.fields.forEach { field ->
                val name = toUpperCamelCase(field.asFieldDeclaration().getVariable(0).nameAsString)
                schema.propOf(Schema.fromType(field.commonType, name, imports))
            }
            typeDeclaration.members.forEach { member ->
                if (member.isTypeDeclaration) {
                    schema.def(Schema.fromTypeDecl(member.asTypeDeclaration(), imports))
                }
            }
            return schema
        }

        fun getRefName(importPaths: List<String>): String {
            val action = importPaths[importPaths.size - 2].replace(Regex("Request|Response"), "")
            return action + importPaths[importPaths.size - 1]
        }

        fun fromType(type: Type, name: String, imports: Map<String, List<String>>): Schema {
            if (type.childNodes.size == 2) {
                val generalTypeName = type.childNodes.get(0).toString()
                val itemName = toUpperCamelCase(type.childNodes.get(1).toString())
                val itemSchema = Schema(itemName).named(itemName)

                if (imports[itemName] != null) {
                    itemSchema.withRef(getRefName(imports[itemName]!!))
                }

                return Schema(generalTypeName)
                        .named(name)
                        .itemOf(itemSchema)
            }

            val t = type.asString()
            if (imports[t] != null) {
                return Schema("").named(name).withRef(getRefName(imports[t]!!))
            }
            return Schema(type.asString()).named(name)
        }

        fun assignDefinitions(definitions: MutableMap<String, Schema>, schema: Schema) {
            schema.definitions.forEach { def ->
                definitions[def.key] = def.value
                assignDefinitions(definitions, def.value)
            }

            if (schema.isArray()) {
                assignDefinitions(definitions, schema.items!!)
            }

            if (schema.isObject()) {
                if (schema.props != null) {
                    schema.props!!.forEach { s ->
                        assignDefinitions(definitions, s.value)
                    }
                }
            }
        }
    }

    var name = ""
    var ref = ""
    var position = ""
    var items: Schema? = null
    var props: MutableMap<String, Schema>? = null
    val definitions = mutableMapOf<String, Schema>()

    fun positionIn(position: String): Schema {
        this.position = position
        return this
    }

    fun getAllDefinitions(): Map<String, Schema> {
        val definitions = mutableMapOf<String, Schema>()

        Schema.assignDefinitions(definitions, this)

        return definitions
    }

    fun isArray(): Boolean {
        return this.items != null
    }

    fun isObject(): Boolean {
        return this.type == "object" || this.props != null
    }

    fun withRef(ref: String): Schema {
        this.ref = ref
        return this
    }

    fun named(name: String): Schema {
        this.name = name
        return this
    }

    fun itemOf(s: Schema): Schema {
        this.items = s
        return this
    }

    fun propOf(s: Schema): Schema {
        if (this.props == null) {
            this.props = mutableMapOf()
        }
        this.props!![s.name] = s
        return this
    }

    fun def(s: Schema): Schema {
        this.definitions[s.name] = s
        return this
    }

    fun goType(prefix: String, asParameter: Boolean, sideCodes: MutableMap<String, String>): String {
        if (this.ref != "") {
            return this.ref
        }

        if (this.isObject()) {
            var t = """struct {
"""
            if (this.props != null) {
                this.props!!.forEach { prop ->
                    t += "${toUpperCamelCase(prop.key)} "
                    if (asParameter && prop.value.isArray()) {
                        t += "*"
                    }
                    t += prop.value.goType(prefix, asParameter, sideCodes)
                    if (asParameter) {
                        t += " `"
                        if (prop.value.position != "") {
                            t += "position:\"${prop.value.position}\" "
                        }
                        name = prop.value.name
                        if (prop.value.isArray()) {
                            t += "type:\"Repeated\" "
                            name = Inflection.singularize(name)
                        }
                        name = name.replace(Regex("([0-9]+)"), ".$1.").trimEnd { s -> s.toString() == "." }
                        t += "name:\"${name}\""
                        t += "`"
                    }
                    t += "\n"
                }
            }
            t += "}"
            return t
        }

        if (this.isArray()) {
            when (this.type) {
                "List" -> {
                    var itemName = this.items!!.name
                    if (itemName == "" || listOf("String", "Long", "Integer", "Float", "Boolean", "boolean").contains(itemName)) {
                        itemName = Inflection.singularize(this.name)
                    }

                    val listTypeName = prefix + toUpperCamelCase(itemName) + "List"
                    val subType = Schema("array").named(itemName).itemOf(this.items!!).goType(prefix, asParameter, sideCodes)

                    sideCodes.set(listTypeName, """
                    type ${listTypeName} ${subType}

                    func (list *${listTypeName}) UnmarshalJSON(data []byte) error {
                        m := make(map[string]${subType})
                        err := json.Unmarshal(data, &m)
                        if err != nil {
                            return err
                        }
                        for _, v := range m {
                            *list = v
                            break
                        }
                        return nil
                    }
                    """)

                    return listTypeName
                }
                "array" -> {
                    return "[]${this.items!!.goType(prefix, asParameter, sideCodes)}"
                }
                else -> {
                    println("unsupported type ${this.type} ${this.items!!.type}")
                    return ""
                }
            }
        }

        when (this.type) {
            "String" -> {
                if (asParameter) {
                    return "string"
                }
                return "goaliyun.String"
            }
            "Long" -> {
                if (asParameter) {
                    return "int64"
                }
                return "goaliyun.Integer"
            }
            "Float" -> {
                if (asParameter) {
                    return "float64"
                }
                return "goaliyun.Float"
            }
            "Integer" -> {
                if (asParameter) {
                    return "int64"
                }
                return "goaliyun.Integer"
            }
            "boolean", "Boolean" -> {
                if (asParameter) {
                    return "string"
                }
                return "bool"
            }
            else -> {
                return prefix + this.type
            }
        }
    }
}


