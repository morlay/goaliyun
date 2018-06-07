package org.aliyun.converter

import java.io.File

class SDK(val name: String, val version: String, val baseDir: String) {
    companion object {
        fun collectAll(): MutableMap<String, SDK> {
            val dirs = glob(sdkModelPath, File("aliyun-openapi-java-sdk"))

            val sdks = mutableMapOf<String, SDK>()

            for (dir in dirs) {
                val names = dir.split("/")
                val group = names[1].replace(Regex("^aliyun-java-sdk-"), "")
                val versionDates = Regex("^v([0-9]{4})([0-9]{2})([0-9]{2})").matchEntire(names[9])!!.groupValues
                val version = "${versionDates[1]}-${versionDates[2]}-${versionDates[3]}"

                sdks[group] = SDK(group, version, dir)
            }

            return sdks
        }
    }

    val operations = mutableMapOf<String, Operation>()

    fun addOperation(op: Operation) {
        this.operations[op.action] = op
    }

    fun safeName(): String {
        return toLowerSnakeCase(this.name)
    }

    fun withPkg(code: String): String {
        return """
package ${this.safeName()}

import (
	"encoding/json"

	"github.com/morlay/goaliyun"
)

${code}
"""
    }
}
