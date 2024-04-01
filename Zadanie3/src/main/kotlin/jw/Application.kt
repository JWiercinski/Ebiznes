package jw

import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import jw.plugins.*
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.URI
import java.net.http.HttpResponse

fun main() {
    embeddedServer(Netty, port = 8080, host = "0.0.0.0", module = Application::module)
        .start(wait = true)
}

fun Application.module() {
    configureRouting()
}

fun dishook(message: String)
{
    val hookurl="https://discord.com/api/webhooks/1223667327757975703/WY9lNoK4RQ1AN--z4DpQNQVcrUQikgBXtBdUpZ30_rw0SPnSk2VnOJrC_J9rFSoN1a4e"
    val client=HttpClient.newHttpClient()
    val request=HttpRequest.newBuilder()
    request.uri(URI(hookurl))
    request.header("Content-Type", "application/json")
    val body=HttpRequest.BodyPublishers.ofString("{\"content\": \"$message\"}")
    request.POST(body)
    val postone=request.build()
    println(postone)
    val response=client.send(postone, HttpResponse.BodyHandlers.ofString())
    println(response.body())
}