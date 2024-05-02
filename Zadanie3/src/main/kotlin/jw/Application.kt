package jw

import io.ktor.server.application.*
import io.ktor.server.engine.*
import io.ktor.server.netty.*
import jw.plugins.*
import java.net.http.HttpClient
import java.net.http.HttpRequest
import java.net.URI
import java.net.http.HttpResponse
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter

//ISTOTNE: Przejdź do linijki 41 i postąp zgodnie z instrukcją, by uruchomić Disbota
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

object JDBot : ListenerAdapter() {
    //Odkomentuj poniższą linijkę i wprowadź poprawny token bota
    //var jda = JDABuilder.createDefault("poprawnyTokenBota").build()
    override fun onMessageReceived(event: MessageReceivedEvent) {
        if (event.author.isBot) return
        var message = event.message.contentDisplay
        println(event.author.name + ": " + message)
        message = message.lowercase()
        if ("kategorie" in message || "kategorii" in message)
        {
            var responder = "Oto lista kategorii:"
            for (pojedyncza in kategorie){
                responder = responder + pojedyncza.id +"-"+pojedyncza.name + ", "
            }
            responder=responder.substring(0, responder.length-2)
            responder=responder+". Aby zobaczyć listę produktów dla danej kategorii, wyślij do mnie wiadomość zawierającą tylko przypisany do niej numer."
            event.channel.sendMessage(responder).queue()
        }
        else
        {
            val categoy = kategorie.find{it.id.toString()==message}
            if (categoy != null)
            {
                var responder=categoy.name+": "
                val prods= produkty.filter { it.category==categoy.id }
                prods.forEach{responder=responder+it.name+" - Cena "+it.price+" zł, "}
                responder=responder.substring(0, responder.length-2)
                event.channel.sendMessage(responder).queue()
            }
        }
    }
}

data class Kategoria (val id: Int, val name: String)
val kategorie = mutableListOf(
    Kategoria(1, "Wędliny"),
    Kategoria(2, "Chipsy"),
    Kategoria(3, "Herbaty"),
    Kategoria(4, "Krakersy"),
    Kategoria(5, "Yerby")
)
data class Produkt (val name: String, val price: Double, val category: Int)
val produkty = mutableListOf(
    Produkt("Krakersy Sezamowe Rarytas", 2.99, 4),
    Produkt("Herbata Zielona Bonatium", 8.99, 3),
    Produkt("La Rubia Elaborada", 27.99, 5),
    Produkt("Cruz de Malta Elaborada", 29.99, 5),
    Produkt("Herbata Zielona Jaśminowa Dilmah", 8.99, 3),
    Produkt("Kabanosy 100% z Kurczaka Tarczyński", 6.79, 1),
    Produkt("Kiełbaski Śniadaniowe Goodvalley", 13.99, 1),
    Produkt("Szynka bez E Lukullus", 5.40, 1),
    Produkt("Krakersy Lajkonik", 5.50, 4),
    Produkt("Chipsy w Kotle Prażone Śmietana z Cebulą Przysnacki", 6.30, 2),
    Produkt("Chipsy Zielona Cebulka Carrefour", 3.50, 2)
)
