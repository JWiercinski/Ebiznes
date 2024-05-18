from django.urls import path
from . import views
urlpatterns = [
    path("", views.llama, name="Chat with Llama"),
    path("chat/", views.chatView.as_view(), name="Chat"),
    path("start/", views.chatstart, name="Starter"),
    path("end/", views.chatend, name="Ender")
]
#Jako że Llama 2 Uncensored odpowiada w języku angielskim, interfejs w niniejszym zadaniu zostanie zrealizowany w języku angielskim
#curl -X POST -H "Content-Type: application/json" -d '{"message": "Who won the football world cup in 2002?"}' http://localhost:8000/chat/
