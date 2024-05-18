from django.urls import path
from . import views
urlpatterns = [
    path("", views.llama, name="Chat with Llama"),
    path("chat/", views.chatView.as_view(), name="Chat"),
    path("start/", views.chatstart, name="Starter"),
    path("end/", views.chatend, name="Ender")
]
#As this local Llama2 instance answers in English, the whole interface will be in English as well
#curl -X POST -H "Content-Type: application/json" -d '{"message": "Who won the football world cup in 2002?"}' http://localhost:8000/chat/
