package jw.plugins

import io.ktor.server.application.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import jw.dishook
import jw.JDBot
import net.dv8tion.jda.api.entities.Activity

fun Application.configureRouting() {
    routing {
        get("/") {
            call.respondText("Hello World!")
        }
        get("/dismess/{message}")
        {
            val mess= (call.parameters["message"])
            val sendit=mess.toString()
            if (sendit.isEmpty() == false || sendit.isBlank()==false)
            {
                dishook(sendit)
                call.respondText("Dokonano próby wysyłki $sendit")
            }
            else
            {
                call.respondText("Otrzymano pustą wiadomość")
            }
        }
        get("/disbot")
        {
            JDBot.jda.addEventListener(JDBot)
            JDBot.jda.presence.activity = Activity.watching("Czacik na którym jesteś")
            call.respondText { "Disbot uruchomiony" }
        }
    }
}
