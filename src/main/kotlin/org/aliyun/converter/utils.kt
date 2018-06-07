package org.aliyun.converter

import com.github.javaparser.JavaParser
import com.github.javaparser.ast.CompilationUnit
import java.io.File
import java.io.FileOutputStream
import java.io.InputStream

fun toLowerSnakeCase(s: String): String {
    return toUpperCamelCase(s).replace(Regex("([A-Z])"), "_$1").toLowerCase().trimStart { c -> c.toString() == "_" }
}

fun toUpperCamelCase(s: String): String {
    val chars = s.split("-").toMutableList()
    return chars.map { word -> word.capitalize() }.joinToString("")
}

fun loadAndParseAllJavaFile(baseDir: String): Map<String, CompilationUnit> {
    val files = glob(Regex("(.+)\\.java"), File(baseDir))
    val parsedCodes = mutableMapOf<String, CompilationUnit>()

    for (file in files) {
        parsedCodes[file] = JavaParser.parse(readFileContent(file))
    }

    return parsedCodes
}

fun readFileContent(pathname: String): String {
    val inputStream: InputStream = File(pathname).inputStream()
    return inputStream.bufferedReader().use { it.readText() }
}

fun glob(tester: Regex, folder: File): List<String> {
    val paths = mutableListOf<String>()
    for (fileEntry in folder.listFiles()) {
        if (fileEntry.isDirectory()) {
            paths.addAll(glob(tester, fileEntry))
        }
        if (tester.matches(fileEntry.path)) {
            paths.add(fileEntry.path)
        }
    }
    return paths
}


fun writeFile(filename: String, content: String) {
    val file = File(filename)
    file.getParentFile().mkdirs()
    file.createNewFile()
    FileOutputStream(file, false).write(content.toByteArray())
}
