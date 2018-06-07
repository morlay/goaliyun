package org.aliyun.converter

import java.io.File

var sdkModelPath = Regex("aliyun-openapi-java-sdk/([^\\/]+)/src/main/java/com/aliyuncs/([^\\/]+)/model/v([^\\/]+)$")

class Converter {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            val sdks = SDK.collectAll()

            // ignore batchcompute, because of different code struct
            sdks.remove("batchcompute")
            sdks.remove("cloudapi")

            sdks.forEach { s ->
                generateSDK(s.value)
            }

//            val sdk = sdks.getValue("cs")
//            generateSDK(sdk)
        }

        fun generateSDK(sdk: SDK) {
            println("generating sdk ${sdk.name} ...")
            val reForMethod = Regex("MethodType\\.([A-Z]+)")
            val reForUri = Regex("setUriPattern\\(\"([^\"]+)\"\\)")
            val parsedCodes = loadAndParseAllJavaFile(sdk.baseDir)

            parsedCodes.forEach { parsedCode ->
                parsedCode.value.types.forEach { type ->
                    if (type.isTopLevelType) {
                        type.members.forEach { member ->
                            if (member.isConstructorDeclaration) {
                                val constructorDeclaration = member.asConstructorDeclaration()
                                var requestType = ""
                                var product = ""
                                var version = ""
                                var action = ""
                                var serviceCode = ""
                                var method = ""
                                var uriPattern = ""

                                parsedCode.value.imports.forEach { i ->
                                    when (i.name.asString()) {
                                        "com.aliyuncs.RpcAcsRequest" -> {
                                            requestType = "rpc"
                                        }
                                        "com.aliyuncs.RoaAcsRequest" -> {
                                            requestType = "roa"
                                        }
                                    }
                                }

                                constructorDeclaration.body.statements.forEach { stmt ->
                                    if (stmt.isExplicitConstructorInvocationStmt) {
                                        val argsForSuper = stmt.asExplicitConstructorInvocationStmt().arguments
                                        if (argsForSuper.size > 0) {
                                            product = argsForSuper[0].asStringLiteralExpr().asString()
                                        }
                                        if (argsForSuper.size > 1) {
                                            version = argsForSuper[1].asStringLiteralExpr().asString()
                                        }
                                        if (argsForSuper.size > 2) {
                                            val arg2 = argsForSuper[2]
                                            if (arg2.isStringLiteralExpr) {
                                                action = arg2.asStringLiteralExpr().asString()
                                            }
                                        }
                                        if (argsForSuper.size > 3) {
                                            serviceCode = argsForSuper[3].asStringLiteralExpr().asString()
                                        }
                                    }

                                    if (requestType == "roa") {
                                        val stmtString = stmt.toString()
                                        if (reForMethod.containsMatchIn(stmtString)) {
                                            method = reForMethod.find(stmtString)!!.groupValues[1]
                                        }
                                        if (reForUri.containsMatchIn(stmt.toString())) {
                                            uriPattern = reForUri.find(stmtString)!!.groupValues[1]
                                        }
                                    }
                                }

                                // skip all roa
                                if (requestType == "roa") {
                                    println("skip ${sdk.name}")
                                    return
                                }

                                if (action != "") {
                                    sdk.addOperation(Operation(requestType, product, version, action, serviceCode, method, uriPattern))
                                }
                            }
                        }
                    }
                }
            }

            parsedCodes.forEach({ parsedCode ->
                val imports = mutableMapOf<String, List<String>>()

                parsedCode.value.imports.forEach { i ->
                    val importPaths = i.name.asString().split(".")
                    imports.set(importPaths[importPaths.size - 1], importPaths)
                }

                sdk.operations.forEach { op ->
                    parsedCode.value.types.forEach { type ->
                        when (type.name.asString()) {
                            op.key + "Request" -> {
                                val schema = Schema.fromTypeDeclForParameter(type, imports)
                                op.value.setParam(schema)
                            }
                            op.key + "Response" -> {
                                val schema = Schema.fromTypeDecl(type, imports)
                                op.value.setResp(schema)
                            }
                        }
                    }
                }
            })

            val baseDir = "clients/${sdk.name}"
            File(baseDir).deleteRecursively()

            sdk.operations.forEach { op ->
                writeFile("${baseDir}/${toLowerSnakeCase(op.value.action)}.go", sdk.withPkg(op.value.goClientMethod()))
            }

            println(Runtime.getRuntime().exec("goimports -w ./${baseDir}").hashCode())
        }
    }
}
